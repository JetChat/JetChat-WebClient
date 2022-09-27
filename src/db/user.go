package db

import "time"

type SQLUser struct {
	UserID        uint64    `json:"user_id" gorm:"column:user_id"`
	AvatarUrl     string    `json:"avatar_url" gorm:"column:avatar_url"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`
	Description   string    `json:"description" gorm:"column:description"`
	Discriminator int       `json:"discriminator" gorm:"column:discriminator"`
	Username      string    `json:"username" gorm:"column:username"`
}

func (m *SQLUser) TableName() string {
	return "user"
}
