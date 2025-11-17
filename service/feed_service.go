package service

import (
	"github.com/chaeyoungeee/blog-feed-notifier/dto"
	"github.com/chaeyoungeee/blog-feed-notifier/pkg/feed"
)

type FeedService struct {
}

func NewFeedService() *FeedService {
	return &FeedService{}
}

func (s *FeedService) GetFeedItems(rssURL string) ([]*dto.FeedItem, error) {
	items, err := feed.FetchFeedItems(rssURL)
	if err != nil {
		return nil, err
	}
	return items, nil
}
