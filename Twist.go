package twist

import (
	"errors"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

var endpointEpisodes = "https://twist.moe/feed/episodes"
var endpointAnime = "https://twist.moe/feed/anime"

// GetFeed queries a twist.moe feed by the twist.moe internal ID.
func GetFeed(twistID string) (*Feed, error) {
	var anime *Feed
	resp, _, errs := gorequest.New().Get(endpointEpisodes + "?format=json&animeId=" + twistID).EndStruct(&anime)

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("Doesn't exist on twist.moe")
	}

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return anime, nil
}

// GetFeedByKitsuID queries a twist.moe feed by the Kitsu ID.
func GetFeedByKitsuID(kitsuID string) (*Feed, error) {
	var anime *Feed
	resp, _, errs := gorequest.New().Get(endpointEpisodes + "?format=json&kitsuId=" + kitsuID).EndStruct(&anime)

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("Doesn't exist on twist.moe")
	}

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return anime, nil
}

// GetAnimeIndex queries information about all anime on twist.moe.
func GetAnimeIndex() (*AnimeIndex, error) {
	var animeList *AnimeIndex
	resp, _, errs := gorequest.New().Get(endpointAnime + "?format=json").EndStruct(&animeList)

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("Doesn't exist on twist.moe")
	}

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return animeList, nil
}
