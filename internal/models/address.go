package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Address model
type Address struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID      uuid.UUID `gorm:"type:uuid;index" json:"user_id"`
	Label       string    `json:"label"`
	FullAddress string    `gorm:"not null" json:"full_address"`
	Latitude    float64   `json:"latitude,omitempty"`
	Longitude   float64   `json:"longitude,omitempty"`
	IsDefault   bool      `gorm:"default:false" json:"is_default"`
}

func (a *Address) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
