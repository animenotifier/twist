package twist

import (
	"errors"
	"net/http"

	"github.com/aerogo/http/client"
)

var endpointEpisodes = "https://twist.moe/feed/episodes"
var endpointAnime = "https://twist.moe/feed/anime"

// GetFeed queries a twist.moe feed by the twist.moe internal ID.
func GetFeed(twistID string) (*Feed, error) {
	var anime *Feed
	resp, err := client.Get(endpointEpisodes + "?format=json&animeId=" + twistID).EndStruct(&anime)

	if resp.StatusCode() == http.StatusNotFound {
		return nil, errors.New("Doesn't exist on twist.moe")
	}

	return anime, err
}

// GetFeedByKitsuID queries a twist.moe feed by the Kitsu ID.
func GetFeedByKitsuID(kitsuID string) (*Feed, error) {
	var anime *Feed
	resp, err := client.Get(endpointEpisodes + "?format=json&kitsuId=" + kitsuID).EndStruct(&anime)

	if resp.StatusCode() == http.StatusNotFound {
		return nil, errors.New("Doesn't exist on twist.moe")
	}

	return anime, err
}

// GetAnimeIndex queries information about all anime on twist.moe.
func GetAnimeIndex() (*AnimeIndex, error) {
	var animeList *AnimeIndex
	resp, err := client.Get(endpointAnime + "?format=json").EndStruct(&animeList)

	if resp.StatusCode() == http.StatusNotFound {
		return nil, errors.New("Doesn't exist on twist.moe")
	}

	return animeList, err
}
