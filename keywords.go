package tmdb

import (
	"fmt"
)

// Keyword struct
type Keyword struct {
	ID   int
	Name string
}

// GetKeywordInfo gets the basic information for a specific keyword id
// https://developers.themoviedb.org/3/keywords/get-keyword-details
func (tmdb *TMDb) GetKeywordInfo(id int) (*Keyword, error) {
	var keywordInfo Keyword
	uri := fmt.Sprintf("%s/keyword/%v?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := getTmdb(uri, &keywordInfo)
	return result.(*Keyword), err
}

// GetKeywordMovies gets the list of movies for a particular keyword by id
// https://developers.themoviedb.org/3/keywords/get-movies-by-keyword
func (tmdb *TMDb) GetKeywordMovies(id int, options map[string]string) (*MoviePagedResults, error) {
	var availableOptions = map[string]struct{}{
		"language": {},
		"page":     {}}
	var movies MoviePagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/keyword/%v/movies?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &movies)
	return result.(*MoviePagedResults), err
}
