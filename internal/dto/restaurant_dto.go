package dto

type ResturantInput struct {
	Name        string `gorm:"not null" json:"name"`
	Image       string `json:"image,omitempty"`
	Address     string `json:"address,omitempty"`
	Description string `json:"description,omitempty"`
}
