package tmdb

import (
	"fmt"
)

// AccountInfo struct
type AccountInfo struct {
	ID           int
	IncludeAdult bool   `json:"include_adult"`
	Iso3166_1    string `json:"iso_3166_1"`
	Iso639_1     string `json:"iso_639_1"`
	Name         string
	Username     string
}

// GetAccountInfo gets the basic information for an account
// https://developers.themoviedb.org/3/account/get-account-details
func (tmdb *TMDb) GetAccountInfo(sessionID string) (*AccountInfo, error) {
	var account AccountInfo
	uri := fmt.Sprintf("%s/account?api_key=%s&session_id=%s", baseURL, tmdb.apiKey, sessionID)
	result, err := getTmdb(uri, &account)
	return result.(*AccountInfo), err
}

// GetAccountLists gets the lists that you have created and marked as a favorite
// https://developers.themoviedb.org/3/account/get-created-lists
func (tmdb *TMDb) GetAccountLists(id int, sessionID string, options map[string]string) (*MovieLists, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"language": {}}
	var lists MovieLists
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/account/%v/lists?api_key=%s&session_id=%s%s", baseURL, id, tmdb.apiKey, sessionID, optionsString)
	result, err := getTmdb(uri, &lists)
	return result.(*MovieLists), err
}

// GetAccountFavoriteMovies gets the list of favorite movies for an account
// https://developers.themoviedb.org/3/account/get-favorite-movies
func (tmdb *TMDb) GetAccountFavoriteMovies(id int, sessionID string, options map[string]string) (*MoviePagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"sort_by":  {},
		"language": {}}
	var favorites MoviePagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/account/%v/favorite/movies?api_key=%s&session_id=%s%s", baseURL, id, tmdb.apiKey, sessionID, optionsString)
	result, err := getTmdb(uri, &favorites)
	return result.(*MoviePagedResults), err
}

// GetAccountFavoriteTv gets the list of favorite movies for an account
// https://developers.themoviedb.org/3/account/get-favorite-tv-shows
func (tmdb *TMDb) GetAccountFavoriteTv(id int, sessionID string, options map[string]string) (*TvPagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"sort_by":  {},
		"language": {}}
	var favorites TvPagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/account/%v/favorite/tv?api_key=%s&session_id=%s%s", baseURL, id, tmdb.apiKey, sessionID, optionsString)
	result, err := getTmdb(uri, &favorites)
	return result.(*TvPagedResults), err
}

// GetAccountRatedMovies gets the list of rated movies (and associated rating) for an account
// https://developers.themoviedb.org/3/account/get-rated-movies
func (tmdb *TMDb) GetAccountRatedMovies(id int, sessionID string, options map[string]string) (*MoviePagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"sort_by":  {},
		"language": {}}
	var favorites MoviePagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/account/%v/rated/movies?api_key=%s&session_id=%s%s", baseURL, id, tmdb.apiKey, sessionID, optionsString)
	result, err := getTmdb(uri, &favorites)
	return result.(*MoviePagedResults), err
}

// GetAccountRatedTv gets the list of rated TV shows (and associated rating) for an account
// https://developers.themoviedb.org/3/account/get-rated-tv-shows
func (tmdb *TMDb) GetAccountRatedTv(id int, sessionID string, options map[string]string) (*TvPagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"sort_by":  {},
		"language": {}}
	var favorites TvPagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/account/%v/rated/tv?api_key=%s&session_id=%s%s", baseURL, id, tmdb.apiKey, sessionID, optionsString)
	result, err := getTmdb(uri, &favorites)
	return result.(*TvPagedResults), err
}

// GetAccountWatchlistMovies gets the list of movies on an accounts watchlist
// https://developers.themoviedb.org/3/account/get-movie-watchlist
func (tmdb *TMDb) GetAccountWatchlistMovies(id int, sessionID string, options map[string]string) (*MoviePagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"sort_by":  {},
		"language": {}}
	var favorites MoviePagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/account/%v/watchlist/movies?api_key=%s&session_id=%s%s", baseURL, id, tmdb.apiKey, sessionID, optionsString)
	result, err := getTmdb(uri, &favorites)
	return result.(*MoviePagedResults), err
}

// GetAccountWatchlistTv gets the list of TV series on an accounts watchlist
// https://developers.themoviedb.org/3/account/get-tv-show-watchlist
func (tmdb *TMDb) GetAccountWatchlistTv(id int, sessionID string, options map[string]string) (*TvPagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"sort_by":  {},
		"language": {}}
	var favorites TvPagedResults
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/account/%v/watchlist/tv?api_key=%s&session_id=%s%s", baseURL, id, tmdb.apiKey, sessionID, optionsString)
	result, err := getTmdb(uri, &favorites)
	return result.(*TvPagedResults), err
}
