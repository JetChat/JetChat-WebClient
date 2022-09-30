package api

import "time"

type Member struct {
	UserId   uint64    `json:"userId"`
	JoinedAt time.Time `json:"joinedAt"`
	GuildId  uint64    `json:"guildId"`
}
