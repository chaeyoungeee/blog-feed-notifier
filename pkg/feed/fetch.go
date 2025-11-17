package feed

import (
	"github.com/chaeyoungeee/blog-feed-notifier/dto"
	"github.com/mmcdole/gofeed"
)

func FetchFeedItems(rssURL string) (feedItems []*dto.FeedItem, err error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(rssURL)
	if err != nil {
		return nil, err
	}
	if len(feed.Items) == 0 {
		return nil, nil
	}

	items := make([]*dto.FeedItem, 0, len(feed.Items))
	for _, item := range feed.Items {
		var thumbnail string
		if item.Image != nil {
			thumbnail = item.Image.URL
		}

		items = append(items, &dto.FeedItem{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			Published:   item.Published,
			Thumbnail:   thumbnail,
			GUID:        item.GUID,
		})
	}
	return items, nil
}
