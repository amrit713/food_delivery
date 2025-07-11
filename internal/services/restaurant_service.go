package services

import (
	"errors"

	"github.com/amrit713/food-delivery/internal/dto"
	"github.com/amrit713/food-delivery/internal/models"
	"github.com/amrit713/food-delivery/internal/repositories"
	"github.com/google/uuid"
)

type RestaurantService struct {
	repo *repositories.RestaurantRepository
}

func NewRestaurantService(repo *repositories.RestaurantRepository) *RestaurantService {
	return &RestaurantService{repo: repo}
}

func (s *RestaurantService) Create(input *dto.ResturantInput, user *models.User) (*models.Restaurant, error) {

	restaurant := &models.Restaurant{

		Name:        input.Name,
		Email:       user.Email,
		Phone:       user.Phone,
		OwnerID:     user.ID,
		Description: input.Description,
		Address:     input.Address,
		Image:       input.Image,
	}

	err := s.repo.Create(restaurant)
	return restaurant, err

}

func (s *RestaurantService) GetAllResturant() ([]models.Restaurant, error) {
	return s.repo.GetAllResturant()
}

func (s *RestaurantService) UpdateRestaurant(id string, input *dto.ResturantInput, userID uuid.UUID) (*models.Restaurant, error) {

	restaurant, err := s.repo.GetResturantById(id)

	if err != nil {
		return nil, errors.New("restaurant not found")
	}

	if restaurant.OwnerID != userID {
		return nil, errors.New("unauthorized to edit this restaurant")
	}

	if input.Name != "" {
		restaurant.Name = input.Name
	}
	if input.Description != "" {
		restaurant.Description = input.Description
	}

	if input.Address != "" {
		restaurant.Address = input.Address
	}

	if input.Image != "" {
		restaurant.Image = input.Image
	}

	return restaurant, nil

}

func (s *RestaurantService) DeleteRestaurant(id string, userID uuid.UUID) error {
	restaurant, err := s.repo.GetResturantById(id)
	if err != nil {
		return errors.New("restaurant not found")
	}

	if restaurant.OwnerID != userID {
		return errors.New("unauthorized to delete this restaurant")
	}

	return s.repo.Delete(id)
}
