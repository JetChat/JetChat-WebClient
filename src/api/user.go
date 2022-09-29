package api

import (
	"net/http"
	"time"
)

type User struct {
	UserID        uint64    `json:"user_id"`
	AvatarUrl     string    `json:"avatar_url"`
	CreatedAt     time.Time `json:"created_at"`
	Description   string    `json:"description"`
	Discriminator int       `json:"discriminator"`
	Username      string    `json:"username"`
}

func GetSelf(r *http.Request) (*User, error) {
	request := NewRequest[User]("/users/me")
	request.Json()
	request.Cookies = r.Cookies()
	user, _, err := request.SendWithResponse()

	return user, err
}
