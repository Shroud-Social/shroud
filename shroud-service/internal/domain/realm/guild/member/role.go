package member

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	UserID     uuid.UUID `json:"user_id"`
	GuildID    uuid.UUID `json:"guild_id"`
	RoleID     uuid.UUID `json:"role_id"`
	AssignedAt time.Time `json:"assigned_at"`
}
