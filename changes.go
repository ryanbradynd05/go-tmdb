package tmdb

import (
	"fmt"
)

// Changes struct
type Changes struct {
	Results []struct {
		ID    int
		Adult bool
	}
}

var changeOptions = map[string]struct{}{
	"page":       {},
	"start_date": {},
	"end_date":   {}}

// GetChangesMovie gets a list of movie ids that have been edited
// https://developers.themoviedb.org/3/changes/get-movie-change-list
func (tmdb *TMDb) GetChangesMovie(options map[string]string) (*Changes, error) {
	var movieChanges Changes
	optionsString := getOptionsString(options, changeOptions)
	uri := fmt.Sprintf("%s/movie/changes?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &movieChanges)
	return result.(*Changes), err
}

// GetChangesPerson gets a list of people ids that have been edited
// https://developers.themoviedb.org/3/changes/get-person-change-list
func (tmdb *TMDb) GetChangesPerson(options map[string]string) (*Changes, error) {
	var personChanges Changes
	optionsString := getOptionsString(options, changeOptions)
	uri := fmt.Sprintf("%s/person/changes?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &personChanges)
	return result.(*Changes), err
}

// GetChangesTv gets a list of tv show ids that have been edited
// https://developers.themoviedb.org/3/changes/get-tv-change-list
func (tmdb *TMDb) GetChangesTv(options map[string]string) (*Changes, error) {
	var tvChanges Changes
	optionsString := getOptionsString(options, changeOptions)
	uri := fmt.Sprintf("%s/tv/changes?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &tvChanges)
	return result.(*Changes), err
}
