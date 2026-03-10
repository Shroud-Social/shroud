package realm

type Realm struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	FirstSeen   time.Time `json:"first_seen"`
	LastSeen    time.Time `json:"last_seen"`
	allows_nsfw bool
}
