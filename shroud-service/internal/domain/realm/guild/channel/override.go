package channel

import (
	"services/internal/domain/realm/guild"

	"github.com/google/uuid"
)

type Scope string

const (
	ScopeRole Scope = "ROLE"
	ScopeUser Scope = "USER"
)

type Override struct {
	GuildID   uuid.UUID `json:"guild_id"`
	ChannelID uuid.UUID `json:"channel_id"`

	Scope            Scope                `json:"scope"`
	ObjectID         uuid.UUID            `json:"object_id"`
	PermissionsAllow guild.PermissionMask `json:"permissions_allow"`
	PermissionsDeny  guild.PermissionMask `json:"permissions_deny"`
}

func (o *Override) HasAllow(p guild.Permission) bool {
	if o.PermissionsAllow&guild.PermissionMask(p) != 0 {
		return true
	}
	return false
}

func (o *Override) HasDeny(p guild.Permission) bool {
	if o.PermissionsDeny&guild.PermissionMask(p) != 0 {
		return true
	}
	return false
}
