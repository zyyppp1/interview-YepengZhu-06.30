package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
	"gorm.io/gorm"
)

// GameLog 游戏日志
type GameLog struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	PlayerID   *uint          `json:"player_id,omitempty"`
	ActionType string         `gorm:"not null" json:"action_type"` // register, login, logout, enter_room, exit_room, join_challenge, challenge_result
	Details    JSONB          `gorm:"type:jsonb" json:"details"`
	IP         string         `json:"ip,omitempty"`
	UserAgent  string         `json:"user_agent,omitempty"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	Player *Player `gorm:"foreignKey:PlayerID" json:"player,omitempty"`
}

// JSONB 自定义类型用于处理 PostgreSQL 的 JSONB
type JSONB map[string]interface{}

// Value 实现 driver.Valuer 接口
func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

// Scan 实现 sql.Scanner 接口
func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, j)
}

// TableName 指定表名
func (GameLog) TableName() string {
	return "game_logs"
}