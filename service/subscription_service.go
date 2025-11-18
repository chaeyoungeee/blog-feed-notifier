package service

import (
	"github.com/chaeyoungeee/blog-feed-notifier/domain"
	"github.com/chaeyoungeee/blog-feed-notifier/dto"
	"github.com/chaeyoungeee/blog-feed-notifier/pkg/notification"
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

func (s *SubscriptionService) CreateSubscription(subscription *domain.Subscription) error {
	return s.Repo.Create(subscription)
}

func (s *SubscriptionService) CreateSubscriptions(userID uint, subscriptions []*domain.Subscription) error {
	return s.Repo.CreateBatch(subscriptions)
}

func (s *SubscriptionService) DeleteSubscription(userID uint, subscriptionID uint) error {
	return s.Repo.Delete(subscriptionID)
}

func (s *SubscriptionService) NotifySubscribers(blog *domain.Blog, items []*dto.FeedItem) error {
	subscriptions, err := s.Repo.GetAllByBlogID(blog.ID)
	if err != nil {
		return err
	}

	payload := notification.ConvertFeedItemToWebhookPayload(items, blog)
	for _, sub := range subscriptions {
		notification.SendDiscordWebhook(sub.User.DiscordWebhookURL, payload)
	}

	return nil
}
