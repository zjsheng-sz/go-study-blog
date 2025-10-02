package services

import (
	"go-study-blog/common"
	"go-study-blog/models"
	"go-study-blog/repositories"

	"gorm.io/gorm"
)

type UserService struct {
	// Define user-related business logic methods here
	repo *repositories.UserRepository
}

func NewUserCtrl(userRepository *repositories.UserRepository) *UserService {
	return &UserService{repo: userRepository}
}

func (s *UserService) GetUserByID(id uint) (models.User, error) {
	user, err := s.repo.FindByID(id)
	if err == gorm.ErrRecordNotFound {
		return user, common.ErrUserNotFound
	}
	return user, err
}

func (s *UserService) GetUserByName(username string) (models.User, error) {
	user, err := s.repo.FindByName(username)
	if err == gorm.ErrRecordNotFound {
		return user, common.ErrUserNotFound
	}
	return user, err
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) UpdateUser(user models.User) error {
	// 检查用户是否存在
	_, err := s.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	// 更新用户信息
	if err := s.repo.Update(user); err != nil {
		return common.ErrDBOperation
	}

	return nil
}

func (s *UserService) DeleteUser(id uint) error {
	// 检查用户是否存在
	_, err := s.GetUserByID(id)
	if err != nil {
		return err
	}

	// 删除用户
	if err := s.repo.Delete(id); err != nil {
		return common.ErrDBOperation
	}

	return nil
}

func (s *UserService) Register(user models.User) error {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return common.ErrInvalidInput
	}

	// 检查用户名是否已存在
	_, err := s.GetUserByName(user.Username)
	if err == nil {
		return common.ErrUserExists
	} else if err != common.ErrUserNotFound {
		return err
	}

	// 加密密码
	if err := user.HashPassword(user.Password); err != nil {
		return common.ErrInternal
	}

	// 创建用户
	if err := s.repo.Create(user); err != nil {
		return common.ErrDBOperation
	}

	return nil
}
