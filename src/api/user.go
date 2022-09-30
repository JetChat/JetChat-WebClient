package api

import (
	"JetChatClientGo/utils"
	"encoding/json"
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

type UserSession struct {
	Username      string `json:"username"`
	Discriminator int    `json:"discriminator"`
	Id            uint64 `json:"userId"`
}

func GetSelf(r *http.Request) (*User, error) {
	request := NewRequest[User]("/users/me").Json().Session(r)
	user, _, err := request.SendWithResponse()

	return user, err
}

func GetSelfId(r *http.Request) (uint64, error) {
	session, err := GetSession(r)
	return session.Id, err
}

func GetSession(r *http.Request) (UserSession, error) {
	sessionCookie, err := r.Cookie(utils.ConnectionCookie)
	if err != nil {
		return UserSession{}, err
	}

	session := UserSession{}
	err = json.Unmarshal([]byte(sessionCookie.Value), &session)
	return session, nil
}
