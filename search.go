package tmdb

import (
	"fmt"
	"net/url"
)

// MovieSearchResults struct
type MovieSearchResults struct {
	Page         int
	Results      []MovieShort
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// SearchMovie searches for movies by title
// http://docs.themoviedb.apiary.io/#reference/search/searchmovie/get
func (tmdb *TMDb) SearchMovie(name string, options map[string]string) (*MovieSearchResults, error) {
	var availableOptions = map[string]struct{}{
		"page":                 {},
		"language":             {},
		"include_adult":        {},
		"year":                 {},
		"primary_release_year": {},
		"search_type":          {}}
	var movies MovieSearchResults
	safeName := url.QueryEscape(name)
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/search/movie?query=%s&api_key=%s%s", baseURL, safeName, tmdb.apiKey, optionsString)
	result, err := callTmdb(uri, &movies)
	return result.(*MovieSearchResults), err
}
