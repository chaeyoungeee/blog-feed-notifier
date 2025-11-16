package repository

import (
	"github.com/chaeyoungeee/blog-feed-notifier/domain"
	"gorm.io/gorm"
)

var ErrUserNotFound = gorm.ErrRecordNotFound

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(user *domain.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepo) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) ExistsByUsername(username string) (bool, error) {
	var count int64
	err := r.DB.Model(&domain.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
