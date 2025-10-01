package services

import (
	"errors"
	"go-study-blog/models"
	"go-study-blog/repositories"
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
	return s.postRepo.FindByID(id)
}

func (s *PostService) GetPost(post models.Post, page models.Pagination) (models.PaginationResult, error) {

	return s.postRepo.Find(post, page)

}

func (s *PostService) UpdatePost(paramPost models.Post, userid uint) error {

	post1, _ := s.postRepo.FindByID(paramPost.ID)

	if post1.UserID != userid {
		return errors.New("auth error")
	}

	return s.postRepo.Update(paramPost)
}

func (s *PostService) DeletePost(id int, userid uint) error {

	post, err := s.postRepo.FindByID(uint(id))

	if err != nil {
		return err
	}

	if post.UserID != userid {
		return errors.New("auth error")
	}

	return s.postRepo.Delete(post.ID)
}

func (s *PostService) CreatePost(post models.Post, userid uint) error {

	user, err := s.userRepo.FindByID(userid)
	if err != nil {
		return err
	}

	if !user.IsAuthen {
		return errors.New("user has not auchenticate")
	}

	return s.postRepo.Create(post)
}
