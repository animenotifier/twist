package twist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFeed(t *testing.T) {
	anime, err := GetFeed("872")

	assert.NoError(t, err)
	assert.NotNil(t, anime)
	assert.NotNil(t, anime.Episodes)
}

func TestGetFeedByKitsuID(t *testing.T) {
	anime, err := GetFeedByKitsuID("10902")

	assert.NoError(t, err)
	assert.NotNil(t, anime)
	assert.NotNil(t, anime.Episodes)
}
