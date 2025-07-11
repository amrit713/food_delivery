package repositories

import (
	"github.com/amrit713/food-delivery/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RestaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) *RestaurantRepository {
	return &RestaurantRepository{db: db}

}

func (r *RestaurantRepository) Create(restaurant *models.Restaurant) error {
	return r.db.Create(restaurant).Error
}

func (r *RestaurantRepository) GetResturantById(id string) (*models.Restaurant, error) {

	uid, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	var restaurant models.Restaurant

	result := r.db.Where("id =?", uid).First(&restaurant)
	return &restaurant, result.Error
}

func (r *RestaurantRepository) GetAllResturant() ([]models.Restaurant, error) {
	var resturants []models.Restaurant

	err := r.db.Find(&resturants).Error

	if err != nil {
		return nil, err
	}

	return resturants, nil
}

func (r *RestaurantRepository) Update(resturant *models.Restaurant) error {
	return r.db.Save(resturant).Error
}
func (r *RestaurantRepository) Delete(id string) error {
	uid, err := uuid.Parse(id)

	if err != nil {
		return err
	}
	return r.db.Delete(&models.Restaurant{}, "id= ?", uid).Error
}
