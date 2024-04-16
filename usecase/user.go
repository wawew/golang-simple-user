package usecase

import (
	"github.com/wawew/golang-simple-user/entity"
	"github.com/wawew/golang-simple-user/repository"
)

type IUserUseCase interface {
	GetUserByID(id string) (*entity.User, error)
	GetAllUsers() ([]*entity.User, error)
}

type UserUseCase struct {
	userRepository repository.IUserRepository
}

func NewUserUseCase(userRepository repository.IUserRepository) IUserUseCase {
	return &UserUseCase{userRepository: userRepository}
}

func (u *UserUseCase) GetUserByID(id string) (*entity.User, error) {
	user, err := u.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, entity.ErrNotFound
	}
	return user, nil
}

func (u *UserUseCase) GetAllUsers() ([]*entity.User, error) {
	return u.userRepository.GetAllUsers()
}
