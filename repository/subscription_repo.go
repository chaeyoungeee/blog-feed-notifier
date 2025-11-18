package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/chaeyoungeee/blog-feed-notifier/domain"
)

var ErrSubscriptionNotFound = gorm.ErrRecordNotFound

type SubscriptionRepo struct {
	DB *gorm.DB
}

func NewSubscriptionRepo(db *gorm.DB) *SubscriptionRepo {
	return &SubscriptionRepo{DB: db}
}

func (r *SubscriptionRepo) Create(subscription *domain.Subscription) error {
	return r.DB.Create(subscription).Error
}

func (r *SubscriptionRepo) CreateBatch(subscriptions []*domain.Subscription) error {
	return r.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&subscriptions).Error
}
func (r *SubscriptionRepo) GetAllByUserID(userID uint) ([]*domain.Subscription, error) {
	var subscriptions []*domain.Subscription
	err := r.DB.
		Preload("Blog").
		Where("user_id = ?", userID).
		Find(&subscriptions).Error
	if err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (r *SubscriptionRepo) GetAllByBlogID(blogID uint) ([]*domain.Subscription, error) {
	var subscriptions []*domain.Subscription
	err := r.DB.
		Preload("User").
		Where("blog_id = ?", blogID).
		Find(&subscriptions).Error
	if err != nil {
		return nil, err
	}
	return subscriptions, nil
}
