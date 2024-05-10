package service

import (
	"errors"
	"gorm_blogging_application/model"
	"gorm_blogging_application/repository"
)

type PostService interface {
	GetAllPosts() ([]model.Post, error)
	GetPostsPaginated(pageNo int, pageSize int) ([]model.Post, error)
	GetPostsPaginatedByTitle(pageNo int, pageSize int) ([]model.Post, error)
	GetPostsPaginatedByAuthorID(pageNo int, pageSize int) ([]model.Post, error)
}

func NewPostServiceImpl(postRepo repository.PostRepository) PostService {
	return PostServiceImpl{postRepository: postRepo}
}

type PostServiceImpl struct {
	postRepository repository.PostRepository
}

func (postService PostServiceImpl) GetAllPosts() ([]model.Post, error) {
	posts := []model.Post{}
	var err error
	postModel, err := postService.postRepository.GetAllPosts()

	if err != nil {
		err1 := errors.New("posts not found")
		return posts, err1
	}

	return postModel, err
}
func (postService PostServiceImpl) GetPostsPaginated(pageNo int, pageSize int) ([]model.Post, error) {

	posts := []model.Post{}
	var err error

	postModel, err := postService.postRepository.GetPostsPaginated(pageNo, pageSize)

	if err != nil {
		err1 := errors.New("posts not found")
		return posts, err1
	}

	return postModel, err
}

func (postService PostServiceImpl) GetPostsPaginatedByTitle(pageNo int, pageSize int) ([]model.Post, error) {

	posts := []model.Post{}
	var err error

	postModel, err := postService.postRepository.GetPostsPaginatedByTitle(pageNo, pageSize)
	if err != nil {
		err1 := errors.New("posts not found")
		return posts, err1
	}

	return postModel, err

}

func (postService PostServiceImpl) GetPostsPaginatedByAuthorID(pageNo int, pageSize int) ([]model.Post, error) {

	posts := []model.Post{}
	var err error

	postModel, err := postService.postRepository.GetPostsPaginatedByAuthorID(pageNo, pageSize)
	if err != nil {
		err1 := errors.New("posts not found")
		return posts, err1
	}

	return postModel, err
}
