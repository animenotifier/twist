package twist

// Feed represents an episode feed object returned by the twist.moe API.
type Feed struct {
	Episodes []*Episode `json:"items"`
}
