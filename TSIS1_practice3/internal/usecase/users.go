package usecase

import (
	"fmt"

	"golang/internal/repository"
	"golang/pkg/modules"
)

type UsersUsecase struct {
	repo repository.UserRepository
}

func NewUsersUsecase(r repository.UserRepository) *UsersUsecase {
	return &UsersUsecase{repo: r}
}

func (u *UsersUsecase) GetUsers() ([]modules.User, error) {
	return u.repo.GetUsers()
}

func (u *UsersUsecase) GetUserByID(id int) (*modules.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *UsersUsecase) CreateUser(user modules.User) (int, error) {
	id, err := u.repo.CreateUser(user)
	if err != nil {
		return 0, fmt.Errorf("create user: %w", err)
	}
	return id, nil
}

func (u *UsersUsecase) UpdateUser(id int, user modules.User) error {
	return u.repo.UpdateUser(id, user)
}

func (u *UsersUsecase) DeleteUser(id int) error {
	return u.repo.DeleteUser(id)
}
