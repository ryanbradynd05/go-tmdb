package tmdb

import (
	"fmt"
)

// TvEpisode struct
type TvEpisode struct {
	AirDate string `json:"air_date"`
	Crew    []struct {
		ID          int
		CreditID    string `json:"credit_id"`
		Name        string
		Department  string
		Job         string
		ProfilePath string `json:"profile_path"`
	}
	EpisodeNumber int `json:"episode_number"`
	GuestStars    []struct {
		ID          int
		CreditID    string `json:"credit_id"`
		Name        string
		Character   string
		Order       int
		ProfilePath string `json:"profile_path"`
	} `json:"guest_stars"`
	Name           string
	Overview       string
	ID             int
	ProductionCode string  `json:"production_code"`
	SeasonNumber   int     `json:"season_number"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float32 `json:"vote_average"`
	VoteCount      uint32  `json:"vote_count"`
}

// GetTvEpisodeInfo gets the primary information about a TV episode by combination of a season and episode number
// http://docs.themoviedb.apiary.io/#reference/tv-episodes/tvidseasonseasonnumberepisodeepisodenumber/get
func (tmdb *TMDb) GetTvEpisodeInfo(showID, seasonNum, episodeNum int, options map[string]string) (*TvEpisode, error) {
	var availableOptions = map[string]struct{}{
		"language":           {},
		"append_to_response": {}}
	var episode TvEpisode
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v/season/%v/episode/%v?api_key=%s%s", baseURL, showID, seasonNum, episodeNum, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &episode)
	return result.(*TvEpisode), err
}
