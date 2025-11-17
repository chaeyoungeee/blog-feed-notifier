package scheduler

import (
	"log"

	"github.com/chaeyoungeee/blog-feed-notifier/dto"
	"github.com/chaeyoungeee/blog-feed-notifier/service"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	SubscriptionService *service.SubscriptionService
	FeedService         *service.FeedService
	BlogService         *service.BlogService
}

func NewScheduler(subService *service.SubscriptionService, feedService *service.FeedService, blogService *service.BlogService) *Scheduler {
	return &Scheduler{
		SubscriptionService: subService,
		FeedService:         feedService,
		BlogService:         blogService,
	}
}

func (s *Scheduler) Start() {
	c := cron.New()

	c.AddFunc("@every 5s", func() {
		log.Println("[Scheduler] Checking for new blog posts...")
		s.CheckNewPosts()
	})

	c.Start()
}

func (s *Scheduler) CheckNewPosts() {
	blogs, err := s.BlogService.GetBlogs()
	if err != nil {
		return
	}

	for _, blog := range blogs {
		items, err := s.FeedService.GetFeedItems(blog.RSSURL)
		if err != nil || len(items) == 0 {
			continue
		}

		lastIdx := -1
		for i, item := range items {
			guid := item.GUID
			if guid == "" {
				guid = item.Link
			}

			if guid == blog.LastID {
				lastIdx = i
				break
			}
		}
		var newItems []*dto.FeedItem

		if lastIdx == 0 {
			continue
		} else if lastIdx > 0 {
			newItems = items[:lastIdx]
		} else {
			newItems = items[:10]
		}

		if len(newItems) == 0 {
			continue
		}

		s.SubscriptionService.NotifySubscribers(blog, newItems)

		newLastID := newItems[0].GUID
		if newLastID == "" {
			newLastID = newItems[0].Link
		}
		s.BlogService.UpdateLastID(blog.ID, newLastID)
	}
}
