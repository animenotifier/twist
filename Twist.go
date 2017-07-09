package twist

import "github.com/parnurzeal/gorequest"

// GetAnime queries a twist.moe anime by the twist.moe internal ID.
func GetAnime(id string) (*Feed, error) {
	var anime *Feed
	_, _, errs := gorequest.New().Get("https://twist.moe/upload-feed?format=json&animeId=" + id).EndStruct(&anime)

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return anime, nil
}

// GetAnimeByKitsuID queries a twist.moe anime by the Kitsu ID.
func GetAnimeByKitsuID(id string) (*Feed, error) {
	var anime *Feed
	_, _, errs := gorequest.New().Get("https://twist.moe/upload-feed?format=json&kitsuId=" + id).EndStruct(&anime)

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return anime, nil
}
