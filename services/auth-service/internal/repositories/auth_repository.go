package repositories

import (
	"errors"

	"github.com/wingobank/auth-service/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

// Save implements UserRepository.
func (u *userRepo) Save(user models.User) (models.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// FindByEmail implements UserRepository.
func (u *userRepo) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, err
	}
	return user, nil
}
