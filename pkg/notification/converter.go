package notification

import (
	"fmt"

	"github.com/chaeyoungeee/blog-feed-notifier/domain"
	"github.com/chaeyoungeee/blog-feed-notifier/dto"
)

func ConvertFeedItemToWebhookPayload(items []*dto.FeedItem, blog *domain.Blog) *dto.DiscordWebhookPayload {
	content := fmt.Sprintf("📢 %s 블로그에 새 글이 올라왔습니다!\n", blog.Name)
	username := "테크 블로그 알림봇"

	embeds := make([]dto.DiscordEmbed, 0, len(items))
	for _, item := range items {
		embeds = append(embeds, dto.DiscordEmbed{
			Title:       item.Title,
			Description: item.Description,
			URL:         item.Link,
			Author: &dto.DiscordAuthor{
				Name:    blog.Name,
				IconURL: blog.IconURL,
			},
			Thumbnail: &dto.DiscordThumbnail{
				URL: item.Thumbnail,
			},
		})
	}

	return &dto.DiscordWebhookPayload{
		Content:  content,
		Username: username,
		Embeds:   embeds,
	}
}
