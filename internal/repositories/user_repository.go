package repositories

import (
	"log"

	"github.com/amrit713/food-delivery/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {

	log.Println("result", email)
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)

	return &user, result.Error

}

func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	var user models.User

	result := r.db.Where("id =?", uid).First(&user)

	return &user, result.Error
}

func (r *UserRepository) GetAllUser() ([]models.User, error) {
	var users []models.User

	err := r.db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) EditUser(user *models.User) error {

	return r.db.Model(user).Where("id = ?", user.ID).Updates(user).Error
}
