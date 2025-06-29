package services

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/zyyppp1/interview-YepengZhu-06.30/db"
	"github.com/zyyppp1/interview-YepengZhu-06.30/models"
	"gorm.io/gorm"
)

// PlayerService 玩家服务接口
type PlayerService interface {
	Create(player *models.Player) error
	GetByID(id uuid.UUID) (*models.Player, error)
	GetAll(page, pageSize int) ([]*models.Player, int64, error)
	Update(id uuid.UUID, updates map[string]interface{}) error
	Delete(id uuid.UUID) error
}

type playerService struct {
	db *gorm.DB
}

// NewPlayerService 创建玩家服务
func NewPlayerService() PlayerService {
	return &playerService{
		db: db.GetDB(),
	}
}

// Create 创建玩家
func (s *playerService) Create(player *models.Player) error {
	// 验证等级是否存在
	var level models.Level
	if err := s.db.First(&level, "id = ?", player.LevelID).Error; err != nil {
		return errors.New("invalid level ID")
	}

	// 创建玩家
	if err := s.db.Create(player).Error; err != nil {
		return fmt.Errorf("failed to create player: %w", err)
	}

	// 加载关联数据
	s.db.Preload("Level").First(player, player.ID)
	return nil
}

// GetByID 根据ID获取玩家
func (s *playerService) GetByID(id uuid.UUID) (*models.Player, error) {
	var player models.Player
	err := s.db.Preload("Level").First(&player, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("player not found")
		}
		return nil, err
	}
	return &player, nil
}

// GetAll 获取所有玩家（分页）
func (s *playerService) GetAll(page, pageSize int) ([]*models.Player, int64, error) {
	var players []*models.Player
	var total int64

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 获取总数
	if err := s.db.Model(&models.Player{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	err := s.db.Preload("Level").
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&players).Error

	if err != nil {
		return nil, 0, err
	}

	return players, total, nil
}

// Update 更新玩家信息
func (s *playerService) Update(id uuid.UUID, updates map[string]interface{}) error {
	// 检查玩家是否存在
	var player models.Player
	if err := s.db.First(&player, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("player not found")
		}
		return err
	}

	// 如果更新等级，验证等级是否存在
	if levelID, ok := updates["level_id"]; ok {
		var level models.Level
		if err := s.db.First(&level, "id = ?", levelID).Error; err != nil {
			return errors.New("invalid level ID")
		}
	}

	// 更新玩家信息
	if err := s.db.Model(&player).Updates(updates).Error; err != nil {
		return fmt.Errorf("failed to update player: %w", err)
	}

	return nil
}

// Delete 删除玩家（软删除）
func (s *playerService) Delete(id uuid.UUID) error {
	result := s.db.Delete(&models.Player{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("player not found")
	}
	return nil
}

// LevelService 等级服务接口
type LevelService interface {
	GetAll() ([]*models.Level, error)
	Create(level *models.Level) error
}

type levelService struct {
	db *gorm.DB
}

// NewLevelService 创建等级服务
func NewLevelService() LevelService {
	return &levelService{
		db: db.GetDB(),
	}
}

// GetAll 获取所有等级
func (s *levelService) GetAll() ([]*models.Level, error) {
	var levels []*models.Level
	err := s.db.Order("created_at ASC").Find(&levels).Error
	return levels, err
}

// Create 创建等级
func (s *levelService) Create(level *models.Level) error {
	return s.db.Create(level).Error
}