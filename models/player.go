// models/player.go
package models

import (
	"time"
	"gorm.io/gorm"
)

// Player 玩家模型
type Player struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"not null;unique" json:"name"`
	LevelID   uint           `gorm:"not null" json:"level_id"`
	Balance   float64        `gorm:"default:0" json:"balance"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	Level *Level `gorm:"foreignKey:LevelID" json:"level,omitempty"`
}

// Level 玩家等级
type Level struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName 指定表名
func (Player) TableName() string {
	return "players"
}

// TableName 指定表名
func (Level) TableName() string {
	return "levels"
}