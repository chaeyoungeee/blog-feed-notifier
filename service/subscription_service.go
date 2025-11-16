package service

import (
	"github.com/chaeyoungeee/blog-feed-notifier/domain"
	"github.com/chaeyoungeee/blog-feed-notifier/repository"
)

type SubscriptionService struct {
	Repo *repository.SubscriptionRepo
}

func NewSubscriptionService(repo *repository.SubscriptionRepo) *SubscriptionService {
	return &SubscriptionService{Repo: repo}
}

func (s *SubscriptionService) GetUserSubscriptions(userID uint) ([]*domain.Subscription, error) {
	return s.Repo.GetAllByUserID(userID)
}
