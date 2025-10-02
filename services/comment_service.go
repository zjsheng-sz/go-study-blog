package services

import (
	"errors"
	"go-study-blog/common"
	"go-study-blog/models"
	"go-study-blog/repositories"
	"log"
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
	return s.commentRepo.FindByID(id)
}

func (s *CommentService) GetComment(comment models.Comment, page common.Pagination) (common.PaginationResult, error) {

	return s.commentRepo.Find(comment, page)

}

func (s *CommentService) UpdateComment(paramComment models.Comment, userid uint) error {

	comment1, _ := s.commentRepo.FindByID(paramComment.ID)

	if comment1.UserID != userid {
		return errors.New("auth error")
	}

	return s.commentRepo.Update(paramComment)
}

func (s *CommentService) DeleteComment(id int, userid uint) error {

	comment, err := s.commentRepo.FindByID(uint(id))

	if err != nil {
		return err
	}

	if comment.UserID != userid {
		return errors.New("auth error")
	}

	return s.commentRepo.Delete(comment.ID)
}

func (s *CommentService) CreateComment(comment models.Comment) error {

	_, err := s.userRepo.FindByID(comment.UserID)

	if err != nil {
		log.Println("user not exist")
		return err
	}

	_, err = s.postRepo.FindByID(comment.PostID)

	if err != nil {
		log.Println("post not exist")
		return err
	}

	return s.commentRepo.Create(comment)
}
