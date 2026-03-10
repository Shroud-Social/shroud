package channel

import (
	"time"

	"github.com/google/uuid"
)

type Type string

const (
	TypeText   Type = "TEXT"
	TypeVoice  Type = "VOICE"
	TypeForum  Type = "FORUM"
	TypeThread Type = "THREAD"
	TypeDM     Type = "DM"
	TypeGroup  Type = "GROUP"
)

type Channel struct {
	ChannelId uuid.UUID `json:"channel_id"`
	ParentId  uuid.UUID `json:"parent_id"`
	GuildId   uuid.UUID `json:"guild_id"`

	CreationTime time.Time `json:"creation_time"`

	Name     string `json:"name"`
	Type     string `json:"type"`
	NSFW     bool   `json:"nsfw"`
	Slowmode uint32 `json:"slowmode"`

	// Type Specific

	// Type: Text
	Topic string `json:"topic"`
	// Type: Voice
	Bitrate   uint32 `json:"bitrate"`
	UserLimit uint32 `json:"user_limit"`
	// Type: Forum
	Tags          []ForumTag `json:"tags"`
	DefaultLayout string     `json:"default_layout"`
	DefaultOrder  string     `json:"default_order"`
	// Type: Thread
	Locked  bool `json:"locked"`
	Private bool `json:"private"`
	// Type: DM

	// Type: Group
	Image string `json:"image"`
}

type ForumTag struct {
	Icon          string `json:"icon"`
	Name          string `json:"name"`
	ModeratorOnly bool   `json:"moderator_only"`
}
