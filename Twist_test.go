package twist

import (
	"testing"

	"github.com/akyoto/assert"
)

func TestGetFeed(t *testing.T) {
	anime, err := GetFeed("872")

	assert.Nil(t, err)
	assert.NotNil(t, anime)
	assert.NotNil(t, anime.Episodes)
}

func TestGetFeedByKitsuID(t *testing.T) {
	anime, err := GetFeedByKitsuID("10902")

	assert.Nil(t, err)
	assert.NotNil(t, anime)
	assert.NotNil(t, anime.Episodes)
}

func TestGetAnimeIndex(t *testing.T) {
	list, err := GetAnimeIndex()

	assert.Nil(t, err)
	assert.NotNil(t, list)
	assert.NotNil(t, list.Items)
	assert.NotEqual(t, len(list.KitsuIDs()), 0)
}
