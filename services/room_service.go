package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zyyppp1/interview-YepengZhu-06.30/db"
	"github.com/zyyppp1/interview-YepengZhu-06.30/models"
	"gorm.io/gorm"
)

// RoomService 房间服务接口
type RoomService interface {
	Create(room *models.Room) error
	GetByID(id uuid.UUID) (*models.Room, error)
	GetAll(status string) ([]*models.Room, error)
	Update(id uuid.UUID, updates map[string]interface{}) error
	Delete(id uuid.UUID) error
}

type roomService struct {
	db *gorm.DB
}

// NewRoomService 创建房间服务
func NewRoomService() RoomService {
	return &roomService{
		db: db.GetDB(),
	}
}

// Create 创建房间
func (s *roomService) Create(room *models.Room) error {
	if room.MaxPlayers <= 0 {
		room.MaxPlayers = 4 // 默认值
	}
	return s.db.Create(room).Error
}

// GetByID 根据ID获取房间
func (s *roomService) GetByID(id uuid.UUID) (*models.Room, error) {
	var room models.Room
	err := s.db.First(&room, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("room not found")
		}
		return nil, err
	}
	return &room, nil
}

// GetAll 获取所有房间
func (s *roomService) GetAll(status string) ([]*models.Room, error) {
	var rooms []*models.Room
	query := s.db.Order("created_at DESC")
	
	// 如果指定了状态，添加过滤条件
	if status != "" {
		query = query.Where("status = ?", status)
	}
	
	err := query.Find(&rooms).Error
	return rooms, err
}

// Update 更新房间信息
func (s *roomService) Update(id uuid.UUID, updates map[string]interface{}) error {
	// 验证状态值
	if status, ok := updates["status"].(string); ok {
		validStatuses := []string{"available", "occupied", "maintenance"}
		isValid := false
		for _, v := range validStatuses {
			if v == status {
				isValid = true
				break
			}
		}
		if !isValid {
			return errors.New("invalid status value")
		}
	}

	result := s.db.Model(&models.Room{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("room not found")
	}
	return nil
}

// Delete 删除房间
func (s *roomService) Delete(id uuid.UUID) error {
	result := s.db.Delete(&models.Room{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("room not found")
	}
	return nil
}

// ReservationService 预约服务接口
type ReservationService interface {
	Create(reservation *models.Reservation) error
	GetAll(roomID *uuid.UUID, date *time.Time, limit int) ([]*models.Reservation, error)
	CheckConflict(roomID uuid.UUID, date time.Time, startTime, endTime string) (bool, error)
}

type reservationService struct {
	db *gorm.DB
}

// NewReservationService 创建预约服务
func NewReservationService() ReservationService {
	return &reservationService{
		db: db.GetDB(),
	}
}

// Create 创建预约
func (s *reservationService) Create(reservation *models.Reservation) error {
	// 检查房间是否存在
	var room models.Room
	if err := s.db.First(&room, "id = ?", reservation.RoomID).Error; err != nil {
		return errors.New("room not found")
	}

	// 检查玩家是否存在
	var player models.Player
	if err := s.db.First(&player, "id = ?", reservation.PlayerID).Error; err != nil {
		return errors.New("player not found")
	}

	// 检查时间冲突
	hasConflict, err := s.CheckConflict(
		reservation.RoomID,
		reservation.ReservationDate,
		reservation.StartTime,
		reservation.EndTime,
	)
	if err != nil {
		return err
	}
	if hasConflict {
		return errors.New("time slot already reserved")
	}

	// 创建预约
	if err := s.db.Create(reservation).Error; err != nil {
		return fmt.Errorf("failed to create reservation: %w", err)
	}

	// 加载关联数据
	s.db.Preload("Room").Preload("Player").First(reservation, reservation.ID)
	return nil
}

// GetAll 获取预约列表
func (s *reservationService) GetAll(roomID *uuid.UUID, date *time.Time, limit int) ([]*models.Reservation, error) {
	var reservations []*models.Reservation
	query := s.db.Preload("Room").Preload("Player")

	// 添加过滤条件
	if roomID != nil {
		query = query.Where("room_id = ?", *roomID)
	}
	if date != nil {
		query = query.Where("reservation_date = ?", date.Format("2006-01-02"))
	}
	
	// 只显示活跃的预约
	query = query.Where("status = ?", "active")
	
	// 添加排序和限制
	query = query.Order("reservation_date DESC, start_time ASC")
	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Find(&reservations).Error
	return reservations, err
}

// CheckConflict 检查时间冲突
func (s *reservationService) CheckConflict(roomID uuid.UUID, date time.Time, startTime, endTime string) (bool, error) {
	var count int64
	err := s.db.Model(&models.Reservation{}).
		Where("room_id = ? AND reservation_date = ? AND status = ?", roomID, date.Format("2006-01-02"), "active").
		Where("(start_time < ? AND end_time > ?) OR (start_time < ? AND end_time > ?)",
			endTime, startTime, endTime, startTime).
		Count(&count).Error
	
	if err != nil {
		return false, err
	}
	
	return count > 0, nil
}