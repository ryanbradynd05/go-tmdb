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
func (tmdb *TMDb) SearchMovie(name string) (*MovieSearchResults, error) {
	var movies MovieSearchResults
	safeName := url.QueryEscape(name)
	uri := fmt.Sprintf("%s/search/movie?query=%s&api_key=%s", baseURL, safeName, tmdb.apiKey)
	result, err := callTmdb(uri, &movies)
	return result.(*MovieSearchResults), err
}
