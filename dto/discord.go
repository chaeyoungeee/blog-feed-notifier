package dto

type DiscordEmbed struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	URL         string            `json:"url"`
	Color       int               `json:"color,omitempty"`
	Author      *DiscordAuthor    `json:"author,omitempty"`
	Thumbnail   *DiscordThumbnail `json:"thumbnail,omitempty"`
}

type DiscordAuthor struct {
	Name    string `json:"name"`
	IconURL string `json:"icon_url,omitempty"`
}

type DiscordThumbnail struct {
	URL string `json:"url"`
}

type DiscordWebhookPayload struct {
	Content  string         `json:"content"`
	Embeds   []DiscordEmbed `json:"embeds"`
	Username string         `json:"username"`
}
