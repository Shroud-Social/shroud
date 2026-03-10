package role

import (
	"services/internal/domain/realm/guild"

	"github.com/google/uuid"
)

type Role struct {
	ID       uuid.UUID `json:"id"`
	GuildID  uuid.UUID `json:"guild_id"`
	Name     string    `json:"name"`
	Color    string    `json:"color"`
	Position int       `json:"position"`

	Permissions guild.PermissionMask `json:"permissions"`
}

func (r Role) HasPermission(p guild.Permission) bool {
	if r.Permissions&guild.PermissionMask(p) != 0 {
		return true
	}
	return false
}
