package service

import (
	"errors"
	"online-ticketing/internal/model"
	"online-ticketing/internal/repository"
	"online-ticketing/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("authentication failed")

type UserService struct {
	up        *repository.UserRepository
	secretKey []byte
}

func NewUserService(up *repository.UserRepository, secret []byte) *UserService {
	return &UserService{
		up:        up,
		secretKey: secret,
	}
}

func (u *UserService) Register(user model.RegisterRequest) (model.UserResponse, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.UserResponse{}, err
	}
	user.Password = string(bytes)
	result, err := u.up.CreateUser(user)
	if err != nil {
		return model.UserResponse{}, err
	}
	return result, nil
}

func (u *UserService) FetchAllUsers() ([]model.UserResponse, error) {
	result, err := u.up.GetUsers()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *UserService) Login(data model.LoginRequest) (string, error) {
	result, err := u.up.GetUserByEmail(data.Email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return "", ErrInvalidCredentials
		}
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(data.Password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	tokenString, err := utils.CreateToken(result.Id, u.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
