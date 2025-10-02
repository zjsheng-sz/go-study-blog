package services

import (
	"go-study-blog/common"
	"go-study-blog/models"
	"go-study-blog/repositories"

	"gorm.io/gorm"
)

type CommentService struct {
	// Define comment-related business logic methods here
	commentRepo *repositories.CommentRepository
	userRepo    *repositories.UserRepository
	postRepo    *repositories.PostRepository
}

func NewCommentService(commentRepository *repositories.CommentRepository, userRepository *repositories.UserRepository, postRepository *repositories.PostRepository) *CommentService {
	return &CommentService{commentRepo: commentRepository, userRepo: userRepository, postRepo: postRepository}
}

func (s *CommentService) GetCommentByID(id uint) (models.Comment, error) {
	comment, err := s.commentRepo.FindByID(id)
	if err == gorm.ErrRecordNotFound {
		return comment, common.ErrCommentNotFound
	}
	return comment, err
}

func (s *CommentService) GetComment(comment models.Comment, page common.Pagination) (common.PaginationResult, error) {

	return s.commentRepo.Find(comment, page)

}

func (s *CommentService) UpdateComment(paramComment models.Comment, userid uint) error {
	comment, err := s.commentRepo.FindByID(paramComment.ID)
	if err == gorm.ErrRecordNotFound {
		return common.ErrCommentNotFound
	} else if err != nil {
		return common.ErrDBOperation
	}

	if comment.UserID != userid {
		return common.ErrForbidden
	}

	if err := s.commentRepo.Update(paramComment); err != nil {
		return common.ErrDBOperation
	}

	return nil
}

func (s *CommentService) DeleteComment(id int, userid uint) error {
	comment, err := s.commentRepo.FindByID(uint(id))
	if err == gorm.ErrRecordNotFound {
		return common.ErrCommentNotFound
	} else if err != nil {
		return common.ErrDBOperation
	}

	if comment.UserID != userid {
		return common.ErrForbidden
	}

	if err := s.commentRepo.Delete(comment.ID); err != nil {
		return common.ErrDBOperation
	}

	return nil
}

func (s *CommentService) CreateComment(comment models.Comment) error {
	_, err := s.userRepo.FindByID(comment.UserID)
	if err == gorm.ErrRecordNotFound {
		return common.ErrUserNotFound
	} else if err != nil {
		return common.ErrDBOperation
	}

	_, err = s.postRepo.FindByID(comment.PostID)
	if err == gorm.ErrRecordNotFound {
		return common.ErrPostNotFound
	} else if err != nil {
		return common.ErrDBOperation
	}

	if err := s.commentRepo.Create(comment); err != nil {
		return common.ErrCommentCreateFailed
	}

	return nil
}
