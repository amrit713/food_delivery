package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Order model
type Order struct {
	ID               uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID           uuid.UUID  `gorm:"type:uuid;index" json:"user_id"`
	RestaurantID     uuid.UUID  `gorm:"type:uuid;index" json:"restaurant_id"`
	DeliveryPersonID *uuid.UUID `gorm:"type:uuid" json:"delivery_person_id,omitempty"`
	TotalAmount      float64    `json:"total_amount"`
	Status           string     `gorm:"default:'pending'" json:"status"`
	PaymentStatus    string     `gorm:"default:'pending'" json:"payment_status"`
	DeliveryAddress  string     `json:"delivery_address"`
	OrderedAt        time.Time  `json:"ordered_at"`
	DeliveredAt      *time.Time `json:"delivered_at,omitempty"`

	OrderItems []OrderItem `json:"order_items,omitempty"`
	Payment    Payment     `json:"payment,omitempty"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New()
	return
}
