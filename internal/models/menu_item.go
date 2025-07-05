package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MenuItem model
type MenuItem struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	RestaurantID uuid.UUID `gorm:"type:uuid;index" json:"restaurant_id"`
	Name         string    `gorm:"not null" json:"name"`
	Description  string    `json:"description,omitempty"`
	Price        float64   `gorm:"not null" json:"price"`
	Image        string    `json:"image,omitempty"`
	IsAvailable  bool      `gorm:"default:true" json:"is_available"`
	CreatedAt    time.Time `json:"created_at"`
}

func (m *MenuItem) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}
