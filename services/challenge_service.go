// services/challenge_service.go
package services

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/zyyppp1/interview-YepengZhu-06.30/db"
	"github.com/zyyppp1/interview-YepengZhu-06.30/models"
	"gorm.io/gorm"
)

// ChallengeService 挑战服务接口
type ChallengeService interface {
	JoinChallenge(playerID uint, amount float64) (*models.Challenge, error)
	ProcessChallengeResult(challengeID uint) error
	GetRecentResults(limit int) ([]*models.Challenge, error)
	CheckCooldown(playerID uint) (bool, time.Time, error)
}

type challengeService struct {
	db *gorm.DB
}

// NewChallengeService 创建挑战服务
func NewChallengeService() ChallengeService {
	return &challengeService{
		db: db.GetDB(),
	}
}

// JoinChallenge 参加挑战
func (s *challengeService) JoinChallenge(playerID uint, amount float64) (*models.Challenge, error) {
	// 验证金额
	if amount != 20.01 {
		return nil, errors.New("challenge amount must be 20.01")
	}

	// 检查玩家是否存在
	var player models.Player
	if err := s.db.First(&player, "id = ?", playerID).Error; err != nil {
		return nil, errors.New("player not found")
	}

	// 检查冷却时间（每分钟只能参加一次）
	canJoin, nextTime, err := s.CheckCooldown(playerID)
	if err != nil {
		return nil, err
	}
	if !canJoin {
		return nil, fmt.Errorf("please wait until %s to join next challenge", nextTime.Format("15:04:05"))
	}

	// 检查玩家余额
	if player.Balance < amount {
		return nil, errors.New("insufficient balance")
	}

	// 开始事务
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 扣除玩家余额
	if err := tx.Model(&player).Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to deduct balance: %w", err)
	}

	// 创建挑战记录
	challenge := &models.Challenge{
		PlayerID:  playerID,
		Amount:    amount,
		StartedAt: time.Now(),
	}

	if err := tx.Create(challenge).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create challenge: %w", err)
	}

	// 更新奖池
	var prizePool models.PrizePool
	if err := tx.First(&prizePool).Error; err != nil {
		// 如果奖池不存在，创建一个
		prizePool = models.PrizePool{CurrentAmount: 0}
		tx.Create(&prizePool)
	}

	if err := tx.Model(&prizePool).Update("current_amount", gorm.Expr("current_amount + ?", amount)).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to update prize pool: %w", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// 30秒后处理结果（实际生产环境应该使用消息队列或定时任务）
	go func() {
		time.Sleep(30 * time.Second)
		s.ProcessChallengeResult(challenge.ID)
	}()

	return challenge, nil
}

// ProcessChallengeResult 处理挑战结果
func (s *challengeService) ProcessChallengeResult(challengeID uint) error {
	// 获取挑战记录
	var challenge models.Challenge
	if err := s.db.First(&challenge, "id = ?", challengeID).Error; err != nil {
		return err
	}

	// 如果已经处理过，直接返回
	if challenge.EndedAt != nil {
		return nil
	}

	// 开始事务
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 计算是否中奖（1%概率）
	rand.Seed(time.Now().UnixNano())
	isWinner := rand.Float64() < 0.01

	endTime := time.Now()
	challenge.EndedAt = &endTime
	challenge.IsWinner = isWinner

	if isWinner {
		// 获取奖池
		var prizePool models.PrizePool
		if err := tx.First(&prizePool).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 设置奖金
		challenge.PrizeAmount = prizePool.CurrentAmount

		// 更新玩家余额
		if err := tx.Model(&models.Player{}).
			Where("id = ?", challenge.PlayerID).
			Update("balance", gorm.Expr("balance + ?", prizePool.CurrentAmount)).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 重置奖池
		prizePool.CurrentAmount = 0
		prizePool.LastWinnerID = &challenge.PlayerID
		prizePool.LastWinAmount = challenge.PrizeAmount
		prizePool.LastWinTime = &endTime

		if err := tx.Save(&prizePool).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 记录日志
		log := &models.GameLog{
			PlayerID:   &challenge.PlayerID,
			ActionType: "challenge_result",
			Details: models.JSONB{
				"challenge_id": challenge.ID,
				"is_winner":    true,
				"prize_amount": challenge.PrizeAmount,
			},
		}
		tx.Create(log)
	}

	// 更新挑战记录
	if err := tx.Save(&challenge).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	return tx.Commit().Error
}

// GetRecentResults 获取最近的挑战结果
func (s *challengeService) GetRecentResults(limit int) ([]*models.Challenge, error) {
	var challenges []*models.Challenge
	query := s.db.Preload("Player").
		Where("ended_at IS NOT NULL").
		Order("ended_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&challenges).Error
	return challenges, err
}

// CheckCooldown 检查冷却时间
func (s *challengeService) CheckCooldown(playerID uint) (bool, time.Time, error) {
	var lastChallenge models.Challenge
	err := s.db.Where("player_id = ?", playerID).
		Order("created_at DESC").
		First(&lastChallenge).Error

	// 如果没有记录，可以参加
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true, time.Now(), nil
	}

	if err != nil {
		return false, time.Now(), err
	}

	// 检查是否过了1分钟
	nextAvailable := lastChallenge.CreatedAt.Add(1 * time.Minute)
	if time.Now().Before(nextAvailable) {
		return false, nextAvailable, nil
	}

	return true, time.Now(), nil
}

