// models/room.go
package models

import (
	"time"
	"gorm.io/gorm"
)

// Room 游戏房间
type Room struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string         `gorm:"not null;unique" json:"name"`
	Description string         `json:"description"`
	Status      string         `gorm:"default:'available'" json:"status"` // available, occupied, maintenance
	MaxPlayers  int            `gorm:"default:4" json:"max_players"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// Reservation 房间预约
type Reservation struct {
	ID              uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	RoomID          uint           `gorm:"not null" json:"room_id"`
	PlayerID        uint           `gorm:"not null" json:"player_id"`
	ReservationDate time.Time      `gorm:"type:date;not null" json:"reservation_date"`
	StartTime       string         `gorm:"not null" json:"start_time"` // HH:MM
	EndTime         string         `gorm:"not null" json:"end_time"`   // HH:MM
	Status          string         `gorm:"default:'active'" json:"status"` // active, cancelled, completed
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	Room   *Room   `gorm:"foreignKey:RoomID" json:"room,omitempty"`
	Player *Player `gorm:"foreignKey:PlayerID" json:"player,omitempty"`
}

// TableName 指定表名
func (Room) TableName() string {
	return "rooms"
}

// TableName 指定表名
func (Reservation) TableName() string {
	return "reservations"
}