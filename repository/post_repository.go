package repository

import (
	"gorm_blogging_application/model"
	"gorm_blogging_application/utility"

	"gorm.io/gorm"
)

// THis layer is interacting between service layer and data layer
type PostRepository interface {
	GetAllPosts() ([]model.Post, error)
	GetPostsPaginated(pageNo int, pageSize int) ([]model.Post, error)
	GetPostsPaginatedByTitle(pageNo int, pageSize int) ([]model.Post, error)
	GetPostsPaginatedByAuthorID(pageNo int, pageSize int) ([]model.Post, error)
}

type PostRepoImpl struct{}

func (postRepo PostRepoImpl) GetAllPosts() ([]model.Post, error) {

	posts := []model.Post{}

	var err error
	if err := utility.Db.Find(&posts).Error; err != nil {

		return posts, err
	}

	return posts, err

}
func (postRepo PostRepoImpl) GetPostsPaginated(pageNo int, pageSize int) ([]model.Post, error) {

	posts := []model.Post{}

	var err error

	if err := utility.Db.Offset(pageNo).Limit(pageSize).Find(&posts).Error; err != nil {
		return posts, err
	}

	return posts, err
}

// Custom scope for pagination ordered by Title in ascending order
func PaginateByTitle(db *gorm.DB) *gorm.DB {
	return db.Order("title ASC")
}

func (postRepo PostRepoImpl) GetPostsPaginatedByTitle(pageNo int, pageSize int) ([]model.Post, error) {

	posts := []model.Post{}
	var err error

	if err := utility.Db.Scopes(PaginateByTitle).Offset((pageNo - 1) * pageSize).Limit(pageSize).Error; err != nil {
		return posts, err
	}

	return posts, err
}

// Custom scope for pagination ordered by Author in descending order
func PaginateByAuthor(db *gorm.DB) *gorm.DB {
	return db.Order("id DESC")
}

func (postRepo PostRepoImpl) GetPostsPaginatedByAuthorID(pageNo int, pageSize int) ([]model.Post, error) {

	posts := []model.Post{}

	var err error

	if err := utility.Db.Scopes(PaginateByAuthor).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&posts).Error; err != nil {
		return posts, err
	}

	return posts, err
}
