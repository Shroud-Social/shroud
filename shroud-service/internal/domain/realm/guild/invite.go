package guild

import (
	"time"

	"github.com/google/uuid"
)

type Invite struct {
	Code              string    `json:"invite_code"`
	GeneratedBy       uuid.UUID `json:"generated_by"`
	GeneratedAt       time.Time `json:"generated_at"`
	ExpiresAt         time.Time `json:"expires_at"`
	InviteLimit       int       `json:"invite_limit"`
	BypassApplication bool      `json:"bypass_application"`
}
