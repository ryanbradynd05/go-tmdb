package tmdb

import (
	"fmt"
)

// DiscoverMovie discovers movies by different types of data like average rating, number of votes, genres and certifications
// https://developers.themoviedb.org/3/discover/movie-discover
func (tmdb *TMDb) DiscoverMovie(options map[string]string) (*MoviePagedResults, error) {
	var availableOptions = map[string]struct{}{
		"certification_country":    {},
		"certification":            {},
		"certification.lte":        {},
		"include_adult":            {},
		"include_video":            {},
		"language":                 {},
		"page":                     {},
		"primary_release_year":     {},
		"primary_release_date.gte": {},
		"primary_release_date.lte": {},
		"release_date.gte":         {},
		"release_date.lte":         {},
		"sort_by":                  {},
		"vote_count.gte":           {},
		"vote_count.lte":           {},
		"vote_average.gte":         {},
		"vote_average.lte":         {},
		"with_cast":                {},
		"with_crew":                {},
		"with_companies":           {},
		"with_genres":              {},
		"with_keywords":            {},
		"with_people":              {},
		"year":                     {}}
	optionsString := getOptionsString(options, availableOptions)
	var results MoviePagedResults
	uri := fmt.Sprintf("%s/discover/movie?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &results)
	return result.(*MoviePagedResults), err
}

// DiscoverTV discovers TV shows by different types of data like average rating, number of votes, genres, the network they aired on and air dates
// https://developers.themoviedb.org/3/discover/tv-discover
func (tmdb *TMDb) DiscoverTV(options map[string]string) (*TvPagedResults, error) {
	var availableOptions = map[string]struct{}{
		"page":                {},
		"language":            {},
		"sort_by":             {},
		"first_air_date_year": {},
		"first_air_date.gte":  {},
		"first_air_date.lte":  {},
		"vote_count.gte":      {},
		"vote_average.gte":    {},
		"with_genres":         {},
		"with_networks":       {}}
	optionsString := getOptionsString(options, availableOptions)
	var results TvPagedResults
	uri := fmt.Sprintf("%s/discover/tv?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &results)
	return result.(*TvPagedResults), err
}
