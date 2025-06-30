// models/log.go
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

// ============================================

// models/payment.go
package models

import (
	"time"
	"gorm.io/gorm"
)

// Payment 支付记录
type Payment struct {
	ID             uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	PlayerID       uint           `gorm:"not null" json:"player_id"`
	PaymentMethod  string         `gorm:"not null" json:"payment_method"` // credit_card, bank_transfer, third_party, blockchain
	Amount         float64        `gorm:"not null" json:"amount"`
	Currency       string         `gorm:"default:'CNY'" json:"currency"`
	Status         string         `gorm:"default:'pending'" json:"status"` // pending, processing, success, failed, refunded
	TransactionID  string         `gorm:"unique" json:"transaction_id"`
	PaymentDetails JSONB          `gorm:"type:jsonb" json:"payment_details,omitempty"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	Player *Player `gorm:"foreignKey:PlayerID" json:"player,omitempty"`
}

// BeforeCreate 创建前的钩子
func (p *Payment) BeforeCreate(tx *gorm.DB) error {
	// 生成交易ID
	if p.TransactionID == "" {
		p.TransactionID = fmt.Sprintf("TXN%d%d", time.Now().Unix(), p.PlayerID)
	}
	return nil
}

// TableName 指定表名
func (Payment) TableName() string {
	return "payments"
}