package notification

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/chaeyoungeee/blog-feed-notifier/dto"
)

func SendDiscordWebhook(webhookURL string, payload *dto.DiscordWebhookPayload) error {
	body, _ := json.Marshal(payload)

	_, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(body))
	return err
}
