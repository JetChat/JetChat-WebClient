package api

import "time"

type User struct {
	UserID        uint64    `json:"user_id"`
	AvatarUrl     string    `json:"avatar_url"`
	CreatedAt     time.Time `json:"created_at"`
	Description   string    `json:"description"`
	Discriminator int       `json:"discriminator"`
	Username      string    `json:"username"`
}
