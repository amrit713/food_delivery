package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Restaurant struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Email       string    `json:"email,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	Image       string    `json:"image,omitempty"`
	Address     string    `json:"address,omitempty"`
	Description string    `json:"description,omitempty"`
	IsOpen      bool      `gorm:"default:true" json:"is_open"`
	OwnerID     uuid.UUID `gorm:"type:uuid" json:"owner_id"`
	CreatedAt   time.Time `json:"created_at"`

	MenuItems []MenuItem `json:"menu_items,omitempty"`
}

func (r *Restaurant) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return
}
