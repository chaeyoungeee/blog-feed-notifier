package domain

import "time"

type User struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	Username          string    `gorm:"size:20;not null;unique" json:"username"`
	Nickname          string    `gorm:"size:20;not null" json:"nickname"`
	Password          string    `gorm:"size:100;not null" json:"password"`
	DiscordWebhookURL string    `gorm:"size:200" json:"discord_webhook_url"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
