package tier

import (
	"time"

	"github.com/google/uuid"
)

type Tier struct {
	ID uuid.UUID `json:"id"`

	Name  string `json:"name"`
	Badge string `json:"badge"`

	TierExpiry time.Time `json:"tier_expiry"`

	UploadSize uint32 `json:"upload_size"`
}
