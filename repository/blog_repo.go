package repository

import (
	"gorm.io/gorm"

	"github.com/chaeyoungeee/blog-feed-notifier/domain"
)

type BlogRepo struct {
	DB *gorm.DB
}

func NewBlogRepo(db *gorm.DB) *BlogRepo {
	return &BlogRepo{DB: db}
}

func (r *BlogRepo) Create(blog *domain.Blog) error {
	return r.DB.Create(blog).Error
}

func (r *BlogRepo) GetAll() ([]domain.Blog, error) {
	var blogs []domain.Blog
	err := r.DB.Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	return blogs, nil
}
