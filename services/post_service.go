package services

import (
	"go-study-blog/common"
	"go-study-blog/models"
	"go-study-blog/repositories"

	"gorm.io/gorm"
)

type PostService struct {
	// Define post-related business logic methods here
	postRepo *repositories.PostRepository
	userRepo *repositories.UserRepository
}

func NewPostService(postRepository *repositories.PostRepository, userRepository *repositories.UserRepository) *PostService {
	return &PostService{postRepo: postRepository, userRepo: userRepository}
}

func (s *PostService) GetPostByID(id uint) (models.Post, error) {
	post, err := s.postRepo.FindByID(id)
	if err == gorm.ErrRecordNotFound {
		return post, common.ErrPostNotFound
	}
	return post, err
}

func (s *PostService) GetPost(post models.Post, page common.Pagination) (common.PaginationResult, error) {

	return s.postRepo.Find(post, page)

}

func (s *PostService) UpdatePost(paramPost models.Post, userid uint) error {
	post, err := s.postRepo.FindByID(paramPost.ID)
	if err == gorm.ErrRecordNotFound {
		return common.ErrPostNotFound
	} else if err != nil {
		return common.ErrDBOperation
	}

	if post.UserID != userid {
		return common.ErrForbidden
	}

	if err := s.postRepo.Update(paramPost); err != nil {
		return common.ErrDBOperation
	}

	return nil
}

func (s *PostService) DeletePost(id int, userid uint) error {
	post, err := s.postRepo.FindByID(uint(id))
	if err == gorm.ErrRecordNotFound {
		return common.ErrPostNotFound
	} else if err != nil {
		return common.ErrDBOperation
	}

	if post.UserID != userid {
		return common.ErrForbidden
	}

	if err := s.postRepo.Delete(post.ID); err != nil {
		return common.ErrDBOperation
	}

	return nil
}

func (s *PostService) CreatePost(post models.Post, userid uint) error {
	_, err := s.userRepo.FindByID(userid)
	if err == gorm.ErrRecordNotFound {
		return common.ErrUserNotFound
	} else if err != nil {
		return common.ErrDBOperation
	}

	if err := s.postRepo.Create(post); err != nil {
		return common.ErrPostCreateFailed
	}

	return nil
}
