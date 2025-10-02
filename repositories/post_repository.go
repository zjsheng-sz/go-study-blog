package repositories

import (
	"go-study-blog/common"
	"go-study-blog/models"

	"gorm.io/gorm"
)

type PostRepository struct {
	// db *gorm.DB // Assuming you're using GORM for database operations
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) FindByID(id uint) (models.Post, error) {

	post := models.Post{}

	err := r.db.First(&post, id).Error

	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (r *PostRepository) Find(post models.Post, page common.Pagination) (common.PaginationResult, error) {

	var total int64

	posts := []models.Post{}

	query := r.db.Model(&models.Post{}).Where(post)

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return common.PaginationResult{}, nil
	}

	// 分页查询
	if err := query.Scopes(page.Paginate()).Find(&posts).Error; err != nil {
		return common.PaginationResult{}, nil
	}

	// 计算总页数
	totalpage := int(total) / page.PageSize
	if int(total)%page.PageSize != 0 {
		totalpage++
	}

	return common.PaginationResult{
		List:      posts,
		Total:     total,
		Page:      page.Page,
		PageSize:  page.PageSize,
		TotalPage: totalpage,
	}, nil

}

func (r *PostRepository) Create(post models.Post) error {

	return r.db.Create(&post).Error
}

func (r *PostRepository) Update(post models.Post) error {

	return r.db.Model(&models.Post{}).Where("id=?", post.ID).Updates(post).Error
}

func (r *PostRepository) Delete(id uint) error {

	return r.db.Delete(&models.Post{}, id).Error

}
