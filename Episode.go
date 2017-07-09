package twist

// Episode represents a single episode on twist.moe.
type Episode struct {
	Link         string `json:"link"`
	Number       int    `json:"episode:number"`
	AnimeTwistID int    `json:"animetwist:id"`
}
