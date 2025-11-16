package repository

import (
	"gorm.io/gorm"

	"github.com/chaeyoungeee/blog-feed-notifier/domain"
)

var ErrBlogNotFound = gorm.ErrRecordNotFound

type BlogRepo struct {
	DB *gorm.DB
}

func NewBlogRepo(db *gorm.DB) *BlogRepo {
	return &BlogRepo{DB: db}
}

func (r *BlogRepo) Create(blog *domain.Blog) error {
	return r.DB.Create(blog).Error
}

func (r *BlogRepo) GetAll() ([]*domain.Blog, error) {
	var blogs []*domain.Blog
	err := r.DB.Find(&blogs).Error
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (r *BlogRepo) GetByName(name string) (*domain.Blog, error) {
	var blog domain.Blog
	err := r.DB.Where("name = ?", name).First(&blog).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrBlogNotFound
		}
		return nil, err
	}
	return &blog, nil
}

func (r *BlogRepo) UpdateLastID(blogID uint, lastID string) error {
	return r.DB.Model(&domain.Blog{}).Where("id = ?", blogID).Update("last_id", lastID).Error
}
