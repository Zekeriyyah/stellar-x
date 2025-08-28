package repositories

import (
	"fmt"

	"github.com/Zekeriyyah/stellar-x/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	if user == nil {
		return fmt.Errorf("user details not provided")
	}

	result := r.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {

	var user models.User

	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
			return &models.User{}, err
		}
	return &user, nil
}

func(r *UserRepository) FindUserByID(id uint) (*models.User, error) {

	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
			return &models.User{}, err
		}

	return &user, nil
}