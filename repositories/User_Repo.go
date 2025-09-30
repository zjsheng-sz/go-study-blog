package repositories

import (
	"go-study-blog/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	// Define user-related data access methods here
	FindByID(id uint) (models.User, error)
	FindAll() ([]models.User, error)
	Create(user models.User) error
	Update(user models.User) error
	Delete(id uint) error
}

type userRepo struct {
	// db *gorm.DB // Assuming you're using GORM for database operations
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}
