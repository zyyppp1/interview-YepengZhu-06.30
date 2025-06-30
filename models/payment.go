package models

import (
	"fmt"
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