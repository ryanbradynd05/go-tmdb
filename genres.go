package tmdb

import (
	"fmt"
)

// Genre struct
type Genre struct {
	Genres []struct {
		ID   int
		Name string
	}
}

// GetMovieGenres gets the list of movie genres
// https://developers.themoviedb.org/3/genres/get-movie-list
func (tmdb *TMDb) GetMovieGenres(options map[string]string) (*Genre, error) {
	var availableOptions = map[string]struct{}{
		"language": {}}
	var movieGenres Genre
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/genre/movie/list?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &movieGenres)
	return result.(*Genre), err
}

// GetTvGenres gets the list of TV genres
// https://developers.themoviedb.org/3/genres/get-tv-list
func (tmdb *TMDb) GetTvGenres(options map[string]string) (*Genre, error) {
	var availableOptions = map[string]struct{}{
		"language": {}}
	var tvGenres Genre
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/genre/tv/list?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &tvGenres)
	return result.(*Genre), err
}
