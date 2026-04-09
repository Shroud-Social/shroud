package storage

type UploadType string

const (
	TypeUserAvatar      UploadType = "USER_AVATAR"
	TypeUserBanner      UploadType = "USER_BANNER"
	TypeGuildIcon       UploadType = "GUILD_ICON"
	TypeGuildBanner     UploadType = "GUILD_BANNER"
	TypeGuildBackground UploadType = "INVITE_BACKGROUND"
	TypeEmote           UploadType = "EMOTE"
	TypeSticker         UploadType = "STICKER"
	TypeSound           UploadType = "SOUND"
	TypeAttachment      UploadType = "ATTACHMENT"
)

var UploadConfigs = map[UploadType]struct {
	MaxSize      int64
	AllowedTypes []string
	UseVolume    bool
	Path         string
}{
	TypeUserAvatar: {
		MaxSize:      2 * 1024 * 1024,
		AllowedTypes: []string{"image/webp"},
		UseVolume:    false,
		Path:         "/avatars/{user_id}/{file_name}",
	},
	TypeUserBanner: {
		MaxSize:      5 * 1024 * 1024,
		AllowedTypes: []string{"image/webp"},
		UseVolume:    false,
		Path:         "/banners/{user_id}/{file_name}",
	},
	TypeGuildIcon: {
		MaxSize:      2 * 1024 * 1024,
		AllowedTypes: []string{"image/webp"},
		UseVolume:    false,
		Path:         "/guild-icons/{guild_id}/{file_name}",
	},
	TypeGuildBanner: {
		MaxSize:      5 * 1024 * 1024,
		AllowedTypes: []string{"image/webp"},
		UseVolume:    false,
		Path:         "/guild-banners/{guild_id}/{file_name}",
	},
	TypeGuildBackground: {
		MaxSize:      10 * 1024 * 1024,
		AllowedTypes: []string{"image/webp"},
		UseVolume:    false,
		Path:         "/guild-backgrounds/{guild_id}/{file_name}",
	},
	TypeEmote: {
		MaxSize:      2 * 1024 * 1024,
		AllowedTypes: []string{"image/webp"},
		UseVolume:    false,
		Path:         "/emotes/{emote_id}/{file_name}",
	},
	TypeSticker: {
		MaxSize:      5 * 1024 * 1024,
		AllowedTypes: []string{"image/webp"},
		UseVolume:    false,
		Path:         "/stickers/{sticker_id}/{file_name}",
	},
	TypeSound: {
		MaxSize:      5 * 1024 * 1024,
		AllowedTypes: []string{"audio/webm"},
		UseVolume:    false,
		Path:         "/sounds/{sound_id}/{file_name}",
	},
	TypeAttachment: {
		MaxSize:      1000 * 1024 * 1024,
		AllowedTypes: []string{"*/*"},
		UseVolume:    true,
		Path:         "/attachments/{channel_id}/{file_name}",
	},
}
