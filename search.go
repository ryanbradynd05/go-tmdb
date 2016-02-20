package tmdb

import (
	"fmt"
	"net/url"
)

// CollectionSearchResults struct
type CollectionSearchResults struct {
	Page    int
	Results []struct {
		ID           int
		BackdropPath string `json:"backdrop_path"`
		Name         string
		PosterPath   string `json:"poster_path"`
	}
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// CompanySearchResults struct
type CompanySearchResults struct {
	Page    int
	Results []struct {
		ID       int
		LogoPath string `json:"logo_path"`
		Name     string
	}
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// KeywordSearchResults struct
type KeywordSearchResults struct {
	Page    int
	Results []struct {
		ID   int
		Name string
	}
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// ListSearchResults struct
type ListSearchResults struct {
	Page    int
	Results []struct {
		Description   string
		FavoriteCount int `json:"favorite_count"`
		ID            string
		ItemCount     int    `json:"item_count"`
		Iso639_1      string `json:"iso_639_1"`
		ListType      string `json:"list_type"`
		Name          string
		PosterPath    string `json:"poster_path"`
	}
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// MovieSearchResults struct
type MovieSearchResults struct {
	Page         int
	Results      []MovieShort
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// MultiSearchResults struct
type MultiSearchResults struct {
	Page    int
	Results []struct {
		BackdropPath  string `json:"backdrop_path"`
		ID            int
		OriginalName  string   `json:"original_name"`
		OriginalTitle string   `json:"original_title"`
		Overview      string   `json:"overview"`
		FirstAirDate  string   `json:"first_air_date"`
		OriginCountry []string `json:"origin_country"`
		PosterPath    string   `json:"poster_path"`
		Popularity    float32
		Name          string
		VoteAverage   float32 `json:"vote_average"`
		VoteCount     uint32  `json:"vote_count"`
		MediaType     string  `json:"media_type"`
	}
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// PersonSearchResults struct
type PersonSearchResults struct {
	Page    int
	Results []struct {
		Adult      bool
		ID         int
		Name       string
		Popularity float32
		PosterPath string `json:"poster_path"`
		KnownFor   []struct {
			Adult         bool
			BackdropPath  string `json:"backdrop_path"`
			ID            int
			OriginalTitle string `json:"original_title"`
			ReleaseDate   string `json:"release_date"`
			PosterPath    string `json:"poster_path"`
			Popularity    float32
			Title         string
			VoteAverage   float32 `json:"vote_average"`
			VoteCount     uint32  `json:"vote_count"`
			MediaType     string  `json:"media_type"`
		} `json:"known_for"`
	}
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// TvSearchResults struct
type TvSearchResults struct {
	Page    int
	Results []struct {
		BackdropPath  string `json:"backdrop_path"`
		ID            int
		OriginalName  string   `json:"original_name"`
		FirstAirDate  string   `json:"first_air_date"`
		OriginCountry []string `json:"origin_country"`
		PosterPath    string   `json:"poster_path"`
		Popularity    float32
		Name          string
		VoteAverage   float32 `json:"vote_average"`
		VoteCount     uint32  `json:"vote_count"`
	}
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// SearchCollection searches for collections by name
// http://docs.themoviedb.apiary.io/#reference/search/searchcollection/get
func (tmdb *TMDb) SearchCollection(name string, options map[string]string) (*CollectionSearchResults, error) {
	var availableOptions = map[string]struct{}{
		"page":     {},
		"language": {}}
	var collections CollectionSearchResults
	safeName := url.QueryEscape(name)
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/search/collection?query=%s&api_key=%s%s", baseURL, safeName, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &collections)
	return result.(*CollectionSearchResults), err
}

// SearchCompany searches for companies by name
// http://docs.themoviedb.apiary.io/#reference/search/searchcompany/get
func (tmdb *TMDb) SearchCompany(name string, options map[string]string) (*CompanySearchResults, error) {
	var availableOptions = map[string]struct{}{
		"page": {}}
	var companies CompanySearchResults
	safeName := url.QueryEscape(name)
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/search/company?query=%s&api_key=%s%s", baseURL, safeName, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &companies)
	return result.(*CompanySearchResults), err
}

// SearchKeyword searches for keywords by name
// http://docs.themoviedb.apiary.io/#reference/search/searchkeyword/get
func (tmdb *TMDb) SearchKeyword(name string, options map[string]string) (*KeywordSearchResults, error) {
	var availableOptions = map[string]struct{}{
		"page": {}}
	var keywords KeywordSearchResults
	safeName := url.QueryEscape(name)
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/search/keyword?query=%s&api_key=%s%s", baseURL, safeName, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &keywords)
	return result.(*KeywordSearchResults), err
}

// SearchList searches for lists by name and description
// http://docs.themoviedb.apiary.io/#reference/search/searchlist/get
func (tmdb *TMDb) SearchList(name string, options map[string]string) (*ListSearchResults, error) {
	var availableOptions = map[string]struct{}{
		"page":          {},
		"include_adult": {}}
	var lists ListSearchResults
	safeName := url.QueryEscape(name)
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/search/list?query=%s&api_key=%s%s", baseURL, safeName, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &lists)
	return result.(*ListSearchResults), err
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
	result, err := getTmdb(uri, &movies)
	return result.(*MovieSearchResults), err
}

// SearchMulti searches the movie, tv show and person collections with a single query
// http://docs.themoviedb.apiary.io/#reference/search/searchmulti/get
func (tmdb *TMDb) SearchMulti(name string, options map[string]string) (*MultiSearchResults, error) {
	var availableOptions = map[string]struct{}{
		"page":          {},
		"language":      {},
		"include_adult": {}}
	var multis MultiSearchResults
	safeName := url.QueryEscape(name)
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/search/multi?query=%s&api_key=%s%s", baseURL, safeName, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &multis)
	return result.(*MultiSearchResults), err
}

// SearchPerson searches for people by name
// http://docs.themoviedb.apiary.io/#reference/search/searchperson/get
func (tmdb *TMDb) SearchPerson(name string, options map[string]string) (*PersonSearchResults, error) {
	var availableOptions = map[string]struct{}{
		"page":          {},
		"search_type":   {},
		"include_adult": {}}
	var people PersonSearchResults
	safeName := url.QueryEscape(name)
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/search/person?query=%s&api_key=%s%s", baseURL, safeName, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &people)
	return result.(*PersonSearchResults), err
}

// SearchTv searches for TV shows by title
// http://docs.themoviedb.apiary.io/#reference/search/searchtv/get
func (tmdb *TMDb) SearchTv(name string, options map[string]string) (*TvSearchResults, error) {
	var availableOptions = map[string]struct{}{
		"page":                {},
		"language":            {},
		"search_type":         {},
		"first_air_date_year": {}}
	var shows TvSearchResults
	safeName := url.QueryEscape(name)
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/search/tv?query=%s&api_key=%s%s", baseURL, safeName, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &shows)
	return result.(*TvSearchResults), err
}
