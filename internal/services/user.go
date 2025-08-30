package services

import (
	"errors"
	"net/http"
	"time"

	"github.com/Zekeriyyah/stellar-x/internal/database"
	"github.com/Zekeriyyah/stellar-x/internal/models"
	"github.com/Zekeriyyah/stellar-x/internal/repositories"
	"github.com/Zekeriyyah/stellar-x/pkg"
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

func (s *UserService) Login(email, password string) (string, int, string) {

	// get user from database using email
	// verify that user password is correct
	// generate token for the user

	user, err :=s.userRepo.FindUserByEmail(email) 
	if err != nil {
		return "", http.StatusNotFound, "user not found"
	}

	if !user.VerifyPassword(password) {
		return "", http.StatusUnauthorized, "password verification failed"
	}

	tokenStr, err := pkg.GeneratJWT(user.ID, time.Now().Add(24*time.Minute))
	if err != nil {
		return "", http.StatusInternalServerError, "invalid input"
	}

	return tokenStr, 200, "login successful"
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