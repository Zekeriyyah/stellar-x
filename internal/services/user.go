package services

import (
	"errors"
	"net/http"

	"github.com/Zekeriyyah/stellar-x/internal/database"
	"github.com/Zekeriyyah/stellar-x/internal/models"
	"github.com/Zekeriyyah/stellar-x/internal/repositories"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

func(s *UserService) CreateUser(user *models.User) (int, string) {

	// // hash password before save
	// if err := user.SetPassword(user.Password); err != nil {
	// 	return http.StatusInternalServerError, "failed to hash password"
	// }

	// save user to database
	if database.DB == nil {
		return http.StatusInternalServerError, "database connection is not initialized"
	}

	err := s.userRepo.CreateUser(user)
	if err != nil {
		return http.StatusInternalServerError, "failed to create user"
	}

	return 200, "user created successfully"
}

func(u *UserService) GetUserByID(id uint) (*models.User, int, string) {

	user, err := u.userRepo.FindUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, "user not found"
		}

		return nil, http.StatusInternalServerError, "failed to fetch user"

	}
	return user, 200, "success"
}

func(u *UserService) GetUserByEmail(email string) (*models.User, int, string) {
	user, err := u.userRepo.FindUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, "user not found"
		}

		return nil, http.StatusInternalServerError, "failed to fetch user"
	}

	return user, 200, "success"
}