package twist

// Anime represents an anime feed object returned by the twist.moe API.
type Anime struct {
	Episodes []*Episode `json:"items"`
}
