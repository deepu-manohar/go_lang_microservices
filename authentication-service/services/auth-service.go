package services

import (
	"authentication/data"
	"errors"
)

type AuthService struct {
	userRepo data.User
}

func NewAuthService(models data.Models) AuthService {
	var authService = AuthService{
		userRepo: models.User,
	}
	return authService
}

func (a *AuthService) GetUser(email string, password string) (*UserDTO, error) {
	user, err := a.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	valid, err := user.PasswordMatches(password)
	if err != nil || !valid {
		return nil, errors.New("invalid credentials")
	}
	userDTO := UserDTO{
		Name:     user.FirstName + " " + user.LastName,
		Email:    user.Email,
		Password: user.Password,
		IsActive: user.Active,
	}
	return &userDTO, nil
}
