package twist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAnime(t *testing.T) {
	anime, err := GetAnime("123")

	assert.NoError(t, err)
	assert.NotNil(t, anime)
	assert.NotNil(t, anime.Episodes)
}

func TestGetAnimeByKitsuID(t *testing.T) {
	anime, err := GetAnimeByKitsuID("123")

	assert.NoError(t, err)
	assert.NotNil(t, anime)
	assert.NotNil(t, anime.Episodes)
}
