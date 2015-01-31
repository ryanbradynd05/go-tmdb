package tmdb

import (
	"fmt"
)

// DiscoverMovie discovers movies by different types of data like average rating, number of votes, genres and certifications
// http://docs.themoviedb.apiary.io/#reference/discover/discovermovie/get
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
	result, err := callTmdb(uri, &results)
	return result.(*MoviePagedResults), err
}
