package services

import (
	"errors"

	"github.com/wingobank/auth-service/internal/models"
	"github.com/wingobank/auth-service/internal/repositories"
	"github.com/wingobank/auth-service/utils"
)

type AuthService interface {
	CreateUser(name, email, password string) (models.User, error)
	Authenticate(email, password string) (models.User, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

// CreateUser implements AuthService.
func (s *authService) CreateUser(name string, email string, password string) (models.User, error) {
	switch _, err := s.userRepo.FindByEmail(email); {
	case err == nil:
		return models.User{}, errors.New("email already registered")
	case !errors.Is(err, repositories.ErrUserNotFound):
		return models.User{}, err
	}

	// Hashes password
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return models.User{}, err
	}

	// Persists new user
	return s.userRepo.Save(models.User{
		Name:     name,
		Email:    email,
		Password: hashed,
	})
}

// Authenticate implements AuthService.
func (a *authService) Authenticate(email string, password string) (models.User, error) {
	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return models.User{}, errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return models.User{}, errors.New("invalid credentials")
	}

	return user, nil
}
