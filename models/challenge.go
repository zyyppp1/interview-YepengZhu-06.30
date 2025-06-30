// models/challenge.go
package models

import (
	"time"
	"gorm.io/gorm"
)

// Challenge 挑战记录
type Challenge struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	PlayerID    uint           `gorm:"not null" json:"player_id"`
	Amount      float64        `gorm:"default:20.01" json:"amount"`
	IsWinner    bool           `gorm:"default:false" json:"is_winner"`
	PrizeAmount float64        `gorm:"default:0" json:"prize_amount"`
	StartedAt   time.Time      `json:"started_at"`
	EndedAt     *time.Time     `json:"ended_at,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	Player *Player `gorm:"foreignKey:PlayerID" json:"player,omitempty"`
}

// PrizePool 奖池
type PrizePool struct {
	ID            uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	CurrentAmount float64    `json:"current_amount"`
	LastWinnerID  *uint      `json:"last_winner_id,omitempty"`
	LastWinAmount float64    `json:"last_win_amount"`
	LastWinTime   *time.Time `json:"last_win_time,omitempty"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

// TableName 指定表名
func (Challenge) TableName() string {
	return "challenges"
}

// TableName 指定表名
func (PrizePool) TableName() string {
	return "prize_pools"
}