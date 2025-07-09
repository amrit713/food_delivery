package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
	RoleRider Role = "rider"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `json:"-"`
	Phone     string    `gorm:"unique;not null" json:"phone"`
	Role      Role      `gorm:"type:user_role; default:'user'" json:"role"`
	Image     string    `json:"image,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Addresses []Address `json:"addresses,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
