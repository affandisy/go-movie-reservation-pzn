package services

import (
	"errors"
	"go-movie-reservation/internal/model"
	"go-movie-reservation/internal/utils"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		DB: db,
	}
}

func (auth *AuthService) SignUp(user model.User) (*model.User, error) {
	user.ID = uuid.New()
	user.Role = model.RegularUserRole

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	if err := auth.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (auth *AuthService) Login(credential model.Credentials) (*model.User, string, error) {
	var user model.User
	if err := auth.DB.Where("email = ?", credential.Email).First(&user).Error; err != nil {
		return nil, "", err
	}

	if !utils.CheckPassword(credential.Password, user.Password) {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWTToken(user.ID.String(), strconv.Itoa(int(user.Role)))
	if err != nil {
		return nil, "", err
	}

	return &user, token, nil
}

func (auth *AuthService) PromoteToAdmin(userID string) error {
	var user model.User
	if err := auth.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return err
	}

	user.Role = model.AdminRole
	if err := auth.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
