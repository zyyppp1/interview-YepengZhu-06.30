package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Challenge 挑战记录
type Challenge struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	PlayerID    uuid.UUID      `gorm:"type:uuid;not null" json:"player_id"`
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
	ID            uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CurrentAmount float64    `json:"current_amount"`
	LastWinnerID  *uuid.UUID `json:"last_winner_id,omitempty"`
	LastWinAmount float64    `json:"last_win_amount"`
	LastWinTime   *time.Time `json:"last_win_time,omitempty"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

// BeforeCreate 创建前的钩子
func (c *Challenge) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	c.StartedAt = time.Now()
	return nil
}

// BeforeCreate 创建前的钩子
func (p *PrizePool) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

// TableName 指定表名
func (Challenge) TableName() string {
	return "challenges"
}

// TableName 指定表名
func (PrizePool) TableName() string {
	return "prize_pools"
}