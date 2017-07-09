package twist

import "github.com/parnurzeal/gorequest"

// GetAnime queries a twist.moe anime by the twist.moe internal ID.
// API endpoint: https://twist.moe/api/anime/:id
func GetAnime(id string) (*Anime, error) {
	var anime *Anime
	_, _, errs := gorequest.New().Get("https://twist.moe/api/anime/" + id).EndStruct(&anime)

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return anime, nil
}

// GetAnimeByKitsuID queries a twist.moe anime by the Kitsu ID.
// API endpoint: https://twist.moe/api/animebykitsu/:id
func GetAnimeByKitsuID(id string) (*Anime, error) {
	var anime *Anime
	_, _, errs := gorequest.New().Get("https://twist.moe/api/animebykitsu/" + id).EndStruct(&anime)

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return anime, nil
}
