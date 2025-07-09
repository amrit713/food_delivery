package services

import (
	"errors"
	"fmt"

	"github.com/amrit713/food-delivery/internal/dto"
	"github.com/amrit713/food-delivery/internal/models"
	"github.com/amrit713/food-delivery/internal/repositories"
	"github.com/amrit713/food-delivery/internal/utils"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// register Service
func (s *AuthService) Register(user *models.User) (*dto.AuthResponse, error) {
	_, err := s.userRepo.GetByEmail(user.Email)

	if err == nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	fmt.Println("user", user)

	err = s.userRepo.Create(user)

	if err != nil {
		return nil, errors.New("unable to create user")
	}

	token, err := utils.GenerateJWT(user.ID.String(), user.Email)

	if err != nil {
		return nil, errors.New("unable to sign jwt token")
	}

	return &dto.AuthResponse{User: user, Token: token}, nil

}

func (s *AuthService) Login(input *dto.LoginInput) (*dto.AuthResponse, error) {
	user, err := s.userRepo.GetByEmail(input.Email)

	if err != nil {
		return nil, errors.New("invalid email or passwords")
	}

	if !utils.ComparePasswordHash(input.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID.String(), string(user.Role))

	if err != nil {
		return nil, errors.New("unable to sign jwt token")
	}

	return &dto.AuthResponse{User: user, Token: token}, nil
}
