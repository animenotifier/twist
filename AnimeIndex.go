package twist

import "strconv"

// AnimeIndex represents all anime on twist.moe.
type AnimeIndex struct {
	Items []*Anime `json:"items"`
}

// KitsuIDs returns an array of all Kitsu IDs.
func (index *AnimeIndex) KitsuIDs() []string {
	idList := make([]string, 0, len(index.Items))

	for _, anime := range index.Items {
		if anime.KitsuID == 0 {
			continue
		}

		idList = append(idList, strconv.Itoa(anime.KitsuID))
	}

	return idList
}
