package domain

import "time"

type Blog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:20;not null" json:"name"`
	LastID    string    `gorm:"size:200" json:"last_id"`
	MainURL   string    `gorm:"size:100;not null" json:"main_url"`
	RSSURL    string    `gorm:"size:200;not null" json:"rss_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
