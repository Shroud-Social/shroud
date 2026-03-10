package interaction

import "github.com/google/uuid"

type Emote struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Image string    `json:"image"`
}
