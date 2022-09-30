package api

import (
	"net/http"
	"strconv"
	"time"
)

type Guild struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	OwnerId   uint64    `json:"ownerId"`
	CreatedAt time.Time `json:"createdAt"`
	Members   []Member  `json:"members"`
	Channels  []Channel `json:"channels"`
	IconUrl   string    `json:"iconUrl,omitempty"`
}

func GetGuild(r *http.Request, id uint64) (*Guild, error) {
	request := NewRequest[Guild]("/guilds/" + strconv.FormatUint(id, 10)).Json().Session(r)
	guild, _, err := request.SendWithResponse()
	return guild, err
}

func GetGuilds(r *http.Request) ([]Guild, error) {
	request := NewRequest[[]Guild]("/guilds").Json().Session(r)
	guilds, _, err := request.SendWithResponse()
	return *guilds, err
}

func (g *Guild) Icon() string {
	if g.IconUrl == "" {
		return "/static/images/guild.png"
	}

	return g.IconUrl
}
