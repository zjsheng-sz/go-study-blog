package services

import (
	"errors"
	"go-study-blog/models"
	"go-study-blog/repositories"
)

type UserService struct {
	// Define user-related business logic methods here
	repo *repositories.UserRepository
}

func NewUserCtrl(userRepository *repositories.UserRepository) *UserService {
	return &UserService{repo: userRepository}
}

func (s *UserService) GetUserByID(id uint) (models.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetUserByName(username string) (models.User, error) {
	return s.repo.FindByName(username)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) UpdateUser(user models.User) error {

	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {

	return s.repo.Delete(id)
}

func (s *UserService) Register(user models.User) error {

	if user.Username == "" || user.Email == "" || user.Password == "" {
		return errors.New("username, email, and password cannot be empty")
	}

	user.HashPassword(user.Password)

	return s.repo.Create(user)
}
