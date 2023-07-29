package repository

import (
	"github.com/mrefawaladam/server/pkg/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUserByID(userID uint) (*entity.User, error) {
	var user entity.User
	if err := r.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	return r.DB.Create(user).Error
}
