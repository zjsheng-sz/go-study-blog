package repositories

import (
	"go-study-blog/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	// db *gorm.DB // Assuming you're using GORM for database operations
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(id uint) (models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) FindByName(username string) (models.User, error) {
	var user models.User
	if err := r.db.Where("username", username).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Create(user models.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) Update(user models.User) error {

	return r.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
