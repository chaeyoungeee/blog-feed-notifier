package repository

import (
	"github.com/chaeyoungeee/blog-feed-notifier/domain"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(user *domain.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepo) ExistsByUsername(username string) (bool, error) {
	var count int64
	err := r.DB.Model(&domain.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
