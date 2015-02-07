package tmdb

import (
	"fmt"
)

// TvSeason struct
type TvSeason struct {
	ID           int
	AirDate      string `json:"air_date"`
	Name         string
	Overview     string
	PosterPath   string `json:"poster_path"`
	SeasonNumber int    `json:"season_number"`
	Episodes     []TvEpisode
}

// GetTvSeasonInfo the primary information about a TV season by its season number
// http://docs.themoviedb.apiary.io/#reference/tv-seasons/tvidseasonseasonnumber/get
func (tmdb *TMDb) GetTvSeasonInfo(showID, seasonID int, options map[string]string) (*TvSeason, error) {
	var availableOptions = map[string]struct{}{
		"language":           {},
		"append_to_response": {}}
	var season TvSeason
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v/season/%v?api_key=%s%s", baseURL, showID, seasonID, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &season)
	return result.(*TvSeason), err
}
