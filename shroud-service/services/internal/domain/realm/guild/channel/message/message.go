package message

import (
	"time"

	"github.com/google/uuid"
)

type AuthorType string

const (
	AuthorTypeUser    AuthorType = "USER"
	AuthorTypeBot     AuthorType = "BOT"
	AuthorTypeWebhook AuthorType = "WEBHOOK"
	AuthorTypeSystem  AuthorType = "SYSTEM"
)

type MessageType string

const (
	MessageTypeText    MessageType = "TEXT"
	MessageTypeSticker MessageType = "STICKER"
)

type Message struct {
	ID         uuid.UUID  `json:"id"`
	Bucket     int        `json:"bucket"`
	Sequence   int64      `json:"sequence"`
	GuildID    uuid.UUID  `json:"guild_id"`
	ChannelID  uuid.UUID  `json:"channel_id"`
	AuthorID   uuid.UUID  `json:"author_id"`
	AuthorType AuthorType `json:"author_type"`
	IsPinned   bool       `json:"is_pinned"`

	MessageType MessageType `json:"message_type"`

	// Type: TEXT
	Content     string    `json:"content"`
	CreatedTime time.Time `json:"created_time"`
	Attachments []string  `json:"attachments"`
	Embeds      []string  `json:"embeds"`

	// Type: STICKER
	StickerID uuid.UUID `json:"sticker_id"`

	// Metadata
	MentionsRoles []uuid.UUID `json:"mentions_roles"`
	MentionsUsers []uuid.UUID `json:"mentions"`
	HasImage      bool        `json:"has_image"`
	HasVideo      bool        `json:"has_video"`
	HasLink       bool        `json:"has_links"`
	HasSound      bool        `json:"has_sound"`
	IsForwarded   bool        `json:"is_forwarded"`

	// Reply
	RepliesTo uuid.UUID `json:"replies_to"`
}

func (m *Message) Dispatch() error {
	return nil
}

func (m *Message) Persist() error {
	return nil
}
