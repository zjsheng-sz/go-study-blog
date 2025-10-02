package repositories

import (
	"go-study-blog/common"
	"go-study-blog/models"

	"gorm.io/gorm"
)

type CommentRepository struct {
	// db *gorm.DB // Assuming you're using GORM for database operations
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) FindByID(id uint) (models.Comment, error) {

	comment := models.Comment{}

	err := r.db.First(&comment, id).Error

	if err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

func (r *CommentRepository) Find(comment models.Comment, page common.Pagination) (common.PaginationResult, error) {

	var total int64

	comments := []models.Comment{}

	query := r.db.Model(&models.Comment{}).Where(comment)

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return common.PaginationResult{}, nil
	}

	// 分页查询
	if err := query.Scopes(page.Paginate()).Find(&comments).Error; err != nil {
		return common.PaginationResult{}, nil
	}

	// 计算总页数
	totalpage := int(total) / page.PageSize
	if int(total)%page.PageSize != 0 {
		totalpage++
	}

	return common.PaginationResult{
		List:      comments,
		Total:     total,
		Page:      page.Page,
		PageSize:  page.PageSize,
		TotalPage: totalpage,
	}, nil

}

func (r *CommentRepository) Create(comment models.Comment) error {

	return r.db.Create(&comment).Error
}

func (r *CommentRepository) Update(comment models.Comment) error {

	return r.db.Model(&models.Comment{}).Where("id=?", comment.ID).Updates(comment).Error
}

func (r *CommentRepository) Delete(id uint) error {

	return r.db.Delete(&models.Comment{}, id).Error

}
