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

// TvSeasonImages struct
type TvSeasonImages struct {
	ID      int
	Posters []TvImage
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

// GetTvSeasonChanges gets a TV season's changes by season ID
// http://docs.themoviedb.apiary.io/#reference/tv-seasons/tvseasonidchanges/get
func (tmdb *TMDb) GetTvSeasonChanges(id int, options map[string]string) (*TvChanges, error) {
	var availableOptions = map[string]struct{}{
		"start_date": {},
		"end_date":   {}}
	var changes TvChanges
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/season/%v/changes?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &changes)
	return result.(*TvChanges), err
}

// GetTvSeasonCredits gets the cast & crew credits for a TV season by season number
// http://docs.themoviedb.apiary.io/#reference/tv-seasons/tvidseasonseasonnumbercredits/get
func (tmdb *TMDb) GetTvSeasonCredits(showID, seasonNum int) (*TvCredits, error) {
	var credits TvCredits
	uri := fmt.Sprintf("%s/tv/%v/season/%v/credits?api_key=%s", baseURL, showID, seasonNum, tmdb.apiKey)
	result, err := getTmdb(uri, &credits)
	return result.(*TvCredits), err
}

// GetTvSeasonExternalIds gets the external ids for a TV season by season number
// http://docs.themoviedb.apiary.io/#reference/tv-seasons/tvidseasonseasonnumberexternalids/get
func (tmdb *TMDb) GetTvSeasonExternalIds(showID, seasonNum int, options map[string]string) (*TvExternalIds, error) {
	var availableOptions = map[string]struct{}{
		"language": {}}
	var ids TvExternalIds
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v/season/%v/external_ids?api_key=%s%s", baseURL, showID, seasonNum, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &ids)
	return result.(*TvExternalIds), err
}

// GetTvSeasonImages gets the images (posters) that we have stored for a TV season by season number
// http://docs.themoviedb.apiary.io/#reference/tv-seasons/tvidseasonseasonnumberimages/get
func (tmdb *TMDb) GetTvSeasonImages(showID, seasonNum int, options map[string]string) (*TvSeasonImages, error) {
	var availableOptions = map[string]struct{}{
		"language":               {},
		"include_image_language": {}}
	var images TvSeasonImages
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v/season/%v/images?api_key=%s%s", baseURL, showID, seasonNum, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &images)
	return result.(*TvSeasonImages), err
}

// GetTvSeasonVideos gets the videos that have been added to a TV season
// http://docs.themoviedb.apiary.io/#reference/tv-seasons/tvidseasonseasonnumbervideos/get
func (tmdb *TMDb) GetTvSeasonVideos(showID, seasonNum int, options map[string]string) (*TvVideos, error) {
	var availableOptions = map[string]struct{}{
		"language": {}}
	var videos TvVideos
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v/season/%v/videos?api_key=%s%s", baseURL, showID, seasonNum, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &videos)
	return result.(*TvVideos), err
}
