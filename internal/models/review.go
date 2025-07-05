package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Review model
type Review struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID       uuid.UUID `gorm:"type:uuid;index" json:"user_id"`
	RestaurantID uuid.UUID `gorm:"type:uuid;index" json:"restaurant_id"`
	Rating       int       `json:"rating"`
	Comment      string    `json:"comment"`
	CreatedAt    time.Time `json:"created_at"`
}

func (r *Review) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return
}
