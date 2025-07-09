package services

import (
	"errors"

	"github.com/amrit713/food-delivery/internal/dto"
	"github.com/amrit713/food-delivery/internal/models"
	"github.com/amrit713/food-delivery/internal/repositories"
	"github.com/amrit713/food-delivery/internal/utils"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAllUser()
}

func (s *UserService) EditUser(user *models.User, input *dto.UpdateUserInput) error {

	// TODO: TRIM SPACE IN FIRST AND LAST IN SENTENCE
	// Update name/email
	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}

	// Change password if provided
	if input.OldPassword != "" && input.NewPassword != "" {
		// Match old password
		if !utils.ComparePasswordHash(input.OldPassword, user.Password) {
			return errors.New("old password is incorrect")
		}

		if input.NewPassword != input.ConfirmPassword {
			return errors.New("new password and confirm password do not match")
		}

		hashed, err := utils.HashPassword(input.NewPassword)

		if err != nil {
			return errors.New("failed to hashed password")
		}

		user.Password = hashed
	}

	return s.userRepo.EditUser(user)
}
