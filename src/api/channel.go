package api

import "time"

type Channel struct {
	Id              uint64    `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	ParentId        uint64    `json:"parentId"`
	GuildId         int       `json:"guildId"`
	CreatedAt       time.Time `json:"createdAt"`
	ChannelType     int       `json:"channelType"`
	ChannelPosition int       `json:"channelPosition"`
}
