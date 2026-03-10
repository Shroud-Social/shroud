package interaction

import "github.com/google/uuid"

type Sticker struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Image string    `json:"image"`
}
