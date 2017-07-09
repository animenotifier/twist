package twist

// Episode represents a single episode on twist.moe.
type Episode struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Link          string `json:"link"`
	PubDate       string `json:"pubDate"`
	AnimeTitle    string `json:"anime:title"`
	EpisodeNumber string `json:"episode:number"`
	AnimeTwistID  string `json:"animetwist:id"`
	KitsuID       string `json:"kitsu:id"`
	MyAnimeListID string `json:"myanimelist:id"`
}
