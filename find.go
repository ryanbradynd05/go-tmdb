package tmdb

import (
	"fmt"
)

// FindResults struct
type FindResults struct {
	MovieResults     []MovieShort  `json:"movie_results,omitempty"`
	PersonResults    []PersonShort `json:"person_results,omitempty"`
	TvResults        []TvShort     `json:"tv_results,omitempty"`
	TvEpisodeResults []struct {
		AirDate       string `json:"air_date"`
		EpisodeNumber int    `json:"episode_number"`
		Name          string
		ID            int
		SeasonNumber  int     `json:"season_number"`
		StillPath     string  `json:"still_path"`
		ShowID        int     `json:"show_id"`
		VoteAverage   float32 `json:"vote_average"`
		VoteCount     int     `json:"vote_count"`
	} `json:"tv_episode_results,omitempty"`
	TvSeasonResults []struct {
		AirDate      string `json:"air_date"`
		Name         string
		ID           int
		SeasonNumber int `json:"season_number"`
		ShowID       int `json:"show_id"`
	} `json:"tv_season_results,omitempty"`
}

// GetFind makes it easy to search for objects in our database by an external id
// https://developers.themoviedb.org/3/find/find-by-id
func (tmdb *TMDb) GetFind(id, source string, options map[string]string) (*FindResults, error) {
	var availableOptions = map[string]struct{}{
		"language": {}}
	var results FindResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/find/%s?api_key=%s&external_source=%s%s", baseURL, id, tmdb.apiKey, source, optionsString)
	result, err := getTmdb(uri, &results)
	return result.(*FindResults), err
}
