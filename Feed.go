package twist

// Feed represents an anime feed object returned by the twist.moe API.
type Feed struct {
	Episodes []*Episode `json:"items"`
}
