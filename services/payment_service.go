package services

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/zyyppp1/interview-YepengZhu-06.30/db"
	"github.com/zyyppp1/interview-YepengZhu-06.30/models"
	"gorm.io/gorm"
)

// PaymentService 支付服务接口
type PaymentService interface {
	ProcessPayment(payment *models.Payment) error
	GetByID(id uuid.UUID) (*models.Payment, error)
	SimulatePaymentGateway(method string, amount float64) (bool, string)
}

type paymentService struct {
	db *gorm.DB
}

// NewPaymentService 创建支付服务
func NewPaymentService() PaymentService {
	return &paymentService{
		db: db.GetDB(),
	}
}

// ProcessPayment 处理支付
func (s *paymentService) ProcessPayment(payment *models.Payment) error {
	// 验证玩家是否存在
	var player models.Player
	if err := s.db.First(&player, "id = ?", payment.PlayerID).Error; err != nil {
		return errors.New("player not found")
	}

	// 验证支付方式
	validMethods := []string{"credit_card", "bank_transfer", "third_party", "blockchain"}
	isValidMethod := false
	for _, method := range validMethods {
		if method == payment.PaymentMethod {
			isValidMethod = true
			break
		}
	}
	if !isValidMethod {
		return errors.New("invalid payment method")
	}

	// 验证金额
	if payment.Amount <= 0 {
		return errors.New("invalid payment amount")
	}

	// 开始事务
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 创建支付记录
	payment.Status = "processing"
	if err := tx.Create(payment).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create payment: %w", err)
	}

	// 模拟调用支付网关
	success, message := s.SimulatePaymentGateway(payment.PaymentMethod, payment.Amount)

	if success {
		// 支付成功，更新玩家余额
		if err := tx.Model(&player).Update("balance", gorm.Expr("balance + ?", payment.Amount)).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update balance: %w", err)
		}

		payment.Status = "success"
		
		// 记录日志
		log := &models.GameLog{
			PlayerID:   &payment.PlayerID,
			ActionType: "payment",
			Details: models.JSONB{
				"payment_id":     payment.ID.String(),
				"amount":         payment.Amount,
				"payment_method": payment.PaymentMethod,
				"status":         "success",
			},
		}
		tx.Create(log)
	} else {
		payment.Status = "failed"
		if payment.PaymentDetails == nil {
			payment.PaymentDetails = make(models.JSONB)
		}
		payment.PaymentDetails["error"] = message
	}

	// 更新支付记录
	if err := tx.Save(payment).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}

	// 重新加载数据
	s.db.Preload("Player").First(payment, payment.ID)
	return nil
}

// GetByID 根据ID获取支付记录
func (s *paymentService) GetByID(id uuid.UUID) (*models.Payment, error) {
	var payment models.Payment
	err := s.db.Preload("Player").First(&payment, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}
	return &payment, nil
}

// SimulatePaymentGateway 模拟支付网关
func (s *paymentService) SimulatePaymentGateway(method string, amount float64) (bool, string) {
	// 模拟处理延迟
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	// 模拟不同支付方式的成功率
	successRates := map[string]float64{
		"credit_card":   0.95, // 95% 成功率
		"bank_transfer": 0.98, // 98% 成功率
		"third_party":   0.90, // 90% 成功率
		"blockchain":    0.85, // 85% 成功率
	}

	rate, ok := successRates[method]
	if !ok {
		rate = 0.80 // 默认80%成功率
	}

	// 随机决定是否成功
	rand.Seed(time.Now().UnixNano())
	if rand.Float64() < rate {
		return true, "Payment processed successfully"
	}

	// 返回随机错误消息
	errors := []string{
		"Insufficient funds",
		"Payment gateway timeout",
		"Invalid card information",
		"Transaction declined by bank",
		"Network error",
	}
	
	return false, errors[rand.Intn(len(errors))]
}

// GameLogService 游戏日志服务
type GameLogService interface {
	Create(log *models.GameLog) error
	GetAll(playerID *uuid.UUID, actionType string, startTime, endTime *time.Time, limit int) ([]*models.GameLog, error)
}

type gameLogService struct {
	db *gorm.DB
}

// NewGameLogService 创建游戏日志服务
func NewGameLogService() GameLogService {
	return &gameLogService{
		db: db.GetDB(),
	}
}

// Create 创建日志
func (s *gameLogService) Create(log *models.GameLog) error {
	// 验证操作类型
	validActions := []string{"register", "login", "logout", "enter_room", "exit_room", "join_challenge", "challenge_result"}
	isValid := false
	for _, action := range validActions {
		if action == log.ActionType {
			isValid = true
			break
		}
	}
	if !isValid {
		return errors.New("invalid action type")
	}

	// 如果提供了玩家ID，验证玩家是否存在
	if log.PlayerID != nil {
		var player models.Player
		if err := s.db.First(&player, "id = ?", *log.PlayerID).Error; err != nil {
			return errors.New("player not found")
		}
	}

	return s.db.Create(log).Error
}

// GetAll 获取日志列表
func (s *gameLogService) GetAll(playerID *uuid.UUID, actionType string, startTime, endTime *time.Time, limit int) ([]*models.GameLog, error) {
	var logs []*models.GameLog
	query := s.db.Preload("Player").Order("created_at DESC")

	// 添加过滤条件
	if playerID != nil {
		query = query.Where("player_id = ?", *playerID)
	}
	if actionType != "" {
		query = query.Where("action_type = ?", actionType)
	}
	if startTime != nil {
		query = query.Where("created_at >= ?", *startTime)
	}
	if endTime != nil {
		query = query.Where("created_at <= ?", *endTime)
	}

	// 添加限制
	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&logs).Error
	return logs, err
}