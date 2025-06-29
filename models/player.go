package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Player 玩家模型
type Player struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	LevelID   uuid.UUID      `gorm:"type:uuid" json:"level_id"`
	Balance   float64        `gorm:"default:0" json:"balance"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	Level *Level `gorm:"foreignKey:LevelID" json:"level,omitempty"`
}

// Level 玩家等级
type Level struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// BeforeCreate 创建前的钩子
func (p *Player) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

// BeforeCreate 创建前的钩子
func (l *Level) BeforeCreate(tx *gorm.DB) error {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	return nil
}

// TableName 指定表名
func (Player) TableName() string {
	return "players"
}

// TableName 指定表名
func (Level) TableName() string {
	return "levels"
}