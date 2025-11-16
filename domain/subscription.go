package domain

import "time"

type Subscription struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	BlogID    uint      `gorm:"not null" json:"blog_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User User `gorm:"foreignKey:UserID"`
	Blog Blog `gorm:"foreignKey:BlogID"`
}
