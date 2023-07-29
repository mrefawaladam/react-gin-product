package usecase

import (
	"github.com/mrefawaladam/server/pkg/entity"
	"github.com/mrefawaladam/server/pkg/repository"
)

type UserUsecase interface {
	GetUserByID(userID uint) (*entity.User, error)
	CreateUser(user *entity.User) error
}

type userUsecase struct {
	UserRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{UserRepo: userRepo}
}

func (u *userUsecase) GetUserByID(userID uint) (*entity.User, error) {
	return u.UserRepo.GetUserByID(userID)
}

func (u *userUsecase) CreateUser(user *entity.User) error {
	return u.UserRepo.CreateUser(user)
}
