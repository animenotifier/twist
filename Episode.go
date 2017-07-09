package twist

// Episode represents a single episode on twist.moe.
type Episode struct {
	Link         string `json:"link"`
	Number       string `json:"episode:number"`
	AnimeTwistID string `json:"animetwist:id"`
	KitsuID      string `json:"kitsu:id"`
}
