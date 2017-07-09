package twist

// Anime represents an anime API object returned by the twist.moe API.
type Anime struct {
	ID       string     `json:"id"`
	KitsuID  string     `json:"kitsuId"`
	Episodes []*Episode `json:"episodes"`
}
