package twist

import "github.com/parnurzeal/gorequest"

// GetFeed queries a twist.moe feed by the twist.moe internal ID.
func GetFeed(twistID string) (*Feed, error) {
	var anime *Feed
	_, _, errs := gorequest.New().Get("https://twist.moe/upload-feed?format=json&animeId=" + twistID).EndStruct(&anime)

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return anime, nil
}

// GetFeedByKitsuID queries a twist.moe feed by the Kitsu ID.
func GetFeedByKitsuID(kitsuID string) (*Feed, error) {
	var anime *Feed
	_, _, errs := gorequest.New().Get("https://twist.moe/upload-feed?format=json&kitsuId=" + kitsuID).EndStruct(&anime)

	if len(errs) > 0 {
		return nil, errs[0]
	}

	return anime, nil
}
