package db

import "time"

type SQLUser struct {
	UserID        uint64    `json:"user_id" db:"column:user_id"`
	AvatarUrl     string    `json:"avatar_url" db:"column:avatar_url"`
	CreatedAt     time.Time `json:"created_at" db:"column:created_at"`
	Description   string    `json:"description" db:"column:description"`
	Discriminator int       `json:"discriminator" db:"column:discriminator"`
	Username      string    `json:"username" db:"column:username"`
}
