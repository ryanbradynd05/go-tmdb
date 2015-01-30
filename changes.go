package tmdb

import (
	"fmt"
)

// Changes struct
type Changes struct {
	Results []struct {
		Id    int
		Adult bool
	}
}

// GetChangesMovie gets a list of movie ids that have been edited
// http://docs.themoviedb.apiary.io/#reference/changes/moviechanges/get
func (tmdb *TMDb) GetChangesMovie() (*Changes, error) {
	var movieChanges Changes
	uri := fmt.Sprintf("%s/movie/changes?api_key=%s", baseURL, tmdb.apiKey)
	result, err := callTmdb(uri, &movieChanges)
	return result.(*Changes), err
}

// GetChangesPerson gets a list of people ids that have been edited
// http://docs.themoviedb.apiary.io/#reference/changes/personchanges/get
func (tmdb *TMDb) GetChangesPerson() (*Changes, error) {
	var personChanges Changes
	uri := fmt.Sprintf("%s/person/changes?api_key=%s", baseURL, tmdb.apiKey)
	result, err := callTmdb(uri, &personChanges)
	return result.(*Changes), err
}

// GetChangesTv gets a list of tv show ids that have been edited
// http://docs.themoviedb.apiary.io/#reference/changes/tvchanges/get
func (tmdb *TMDb) GetChangesTv() (*Changes, error) {
	var tvChanges Changes
	uri := fmt.Sprintf("%s/tv/changes?api_key=%s", baseURL, tmdb.apiKey)
	result, err := callTmdb(uri, &tvChanges)
	return result.(*Changes), err
}
