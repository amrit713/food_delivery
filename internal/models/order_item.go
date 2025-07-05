package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// OrderItem model
type OrderItem struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	OrderID    uuid.UUID `gorm:"type:uuid;index" json:"order_id"`
	MenuItemID uuid.UUID `gorm:"type:uuid;index" json:"menu_item_id"`
	Quantity   int       `json:"quantity"`
	Price      float64   `json:"price"`
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	oi.ID = uuid.New()
	return
}
