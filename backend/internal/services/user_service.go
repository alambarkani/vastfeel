package services

import (
	"vastfeel-backend/internal/models"
	"vastfeel-backend/internal/repositories"
)

type UserService interface {
	Create(user *models.User) error
	GetByID(id int) (*models.User, error)
	GetAllUser() ([]models.User, error)
	Update(user *models.User) error
	Delete(id int) error
}

type UserServiceImpl struct {
	Repo repositories.UserRepository
}

// Create implements UserService.
func (u *UserServiceImpl) Create(user *models.User) error {
	return u.Repo.Create(user)
}

// Delete implements UserService.
func (u *UserServiceImpl) Delete(id int) error {
	return u.Repo.Delete(id)
}

// GetAllUser implements UserService.
func (u *UserServiceImpl) GetAllUser() ([]models.User, error) {
	return u.Repo.GetAllUser()
}

// GetByID implements UserService.
func (u *UserServiceImpl) GetByID(id int) (*models.User, error) {
	return u.Repo.GetByID(id)
}

// Update implements UserService.
func (u *UserServiceImpl) Update(user *models.User) error {
	return u.Repo.Update(user)
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &UserServiceImpl{
		Repo: repo,
	}
}
