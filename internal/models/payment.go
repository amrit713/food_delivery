package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Payment model
type Payment struct {
	ID      uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	OrderID uuid.UUID  `gorm:"type:uuid;uniqueIndex" json:"order_id"`
	Amount  float64    `gorm:"not null" json:"amount"`
	Method  string     `json:"method"`
	Status  string     `json:"status"`
	PaidAt  *time.Time `json:"paid_at,omitempty"`
}

func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
