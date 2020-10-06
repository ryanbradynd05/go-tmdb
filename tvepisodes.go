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

// TvEpisodeImages struct
type TvEpisodeImages struct {
	ID     int
	Stills []TvImage
}

// GetTvEpisodeInfo gets the primary information about a TV episode by combination of a season and episode number
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-details
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

// GetTvEpisodeChanges gets a TV episode's changes by episode ID
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-changes
func (tmdb *TMDb) GetTvEpisodeChanges(id int, options map[string]string) (*TvChanges, error) {
	var availableOptions = map[string]struct{}{
		"start_date": {},
		"end_date":   {}}
	var changes TvChanges
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/episode/%v/changes?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &changes)
	return result.(*TvChanges), err
}

// GetTvEpisodeCredits gets the TV episode credits by combination of season and episode number
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-credits
func (tmdb *TMDb) GetTvEpisodeCredits(showID, seasonNum, episodeNum int) (*TvCredits, error) {
	var credits TvCredits
	uri := fmt.Sprintf("%s/tv/%v/season/%v/episode/%v/credits?api_key=%s", baseURL, showID, seasonNum, episodeNum, tmdb.apiKey)
	result, err := getTmdb(uri, &credits)
	return result.(*TvCredits), err
}

// GetTvEpisodeExternalIds gets the external ids for a TV episode by comabination of a season and episode number
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-external-ids
func (tmdb *TMDb) GetTvEpisodeExternalIds(showID, seasonNum, episodeNum int, options map[string]string) (*TvExternalIds, error) {
	var availableOptions = map[string]struct{}{
		"language": {}}
	var ids TvExternalIds
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v/season/%v/episode/%v/external_ids?api_key=%s%s", baseURL, showID, seasonNum, episodeNum, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &ids)
	return result.(*TvExternalIds), err
}

// GetTvEpisodeImages gets the images (episode stills) for a TV episode by combination of a season and episode number
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-images
func (tmdb *TMDb) GetTvEpisodeImages(showID, seasonNum, episodeNum int) (*TvEpisodeImages, error) {
	var images TvEpisodeImages
	uri := fmt.Sprintf("%s/tv/%v/season/%v/episode/%v/images?api_key=%s", baseURL, showID, seasonNum, episodeNum, tmdb.apiKey)
	result, err := getTmdb(uri, &images)
	return result.(*TvEpisodeImages), err
}

// GetTvEpisodeVideos gets the videos that have been added to a TV episode
// https://developers.themoviedb.org/3/tv-episodes/get-tv-episode-videos
func (tmdb *TMDb) GetTvEpisodeVideos(showID, seasonNum, episodeNum int, options map[string]string) (*TvVideos, error) {
	var availableOptions = map[string]struct{}{
		"language": {}}
	var videos TvVideos
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v/season/%v/episode/%v/videos?api_key=%s%s", baseURL, showID, seasonNum, episodeNum, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &videos)
	return result.(*TvVideos), err
}
