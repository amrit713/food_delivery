package dto

import (
	"github.com/amrit713/food-delivery/internal/models"
	"github.com/google/uuid"
)

// TODO: need store this user in local rather than model user for more secure
type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Role  string    `json:"role"`
}

func ToUserResponse(u *models.User) UserResponse {
	return UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Role:  string(u.Role),
	}
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	User  *models.User `json:"user"`
	Token string       `json:"token"`
}

type UpdateUserInput struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	OldPassword     string `json:"old_password,omitempty"`
	NewPassword     string `json:"new_password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
}
