package tmdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// ErrUnknownMediaType var
var ErrUnknownMediaType = errors.New("Unknown media type. Unable to unmarhal to MultiSearchResultsInfo")

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

// MultiSearchBase interface
type MultiSearchBase interface {
	interfaceMarkerMethod() // We need a method to be
}

// MultiSearchMovieInfo struct
type MultiSearchMovieInfo struct {
	PosterPath       string `json:"poster_path"`
	Adult            bool
	Overview         string   `json:"overview"`
	ReleaseDate      string   `json:"release_date"`
	OriginalTitle    string   `json:"original_title"`
	GenreIDs         []uint32 `json:"ganre_ids"`
	OriginalLanguage string   `json:"original_language"`
	Title            string   `json:"title"`
	BackdropPath     string   `json:"backdrop_path"`
	Popularity       float32
	VoteCount        uint32 `json:"vote_count"`
	Video            bool
	VoteAverage      float32 `json:"vote_average"`
	ID               int
	MediaType        string `json:"media_type"`
}

func (MultiSearchMovieInfo) interfaceMarkerMethod() { return }

// MultiSearchTvInfo struct
type MultiSearchTvInfo struct {
	BackdropPath     string   `json:"backdrop_path"`
	OriginalName     string   `json:"original_name"`
	OriginalTitle    string   `json:"original_title"`
	OriginalLanguage string   `json:"original_language"`
	Overview         string   `json:"overview"`
	FirstAirDate     string   `json:"first_air_date"`
	OriginCountry    []string `json:"origin_country"`
	GenreIDs         []uint32 `json:"ganre_ids"`
	PosterPath       string   `json:"poster_path"`
	Popularity       float32
	Name             string
	VoteAverage      float32 `json:"vote_average"`
	VoteCount        uint32  `json:"vote_count"`
	ID               int
	MediaType        string `json:"media_type"`
}

func (MultiSearchTvInfo) interfaceMarkerMethod() { return }

// MultiSearchPersonInfo struct
type MultiSearchPersonInfo struct {
	ProfilePath string `json:"profile_path"`
	Adult       bool
	KnownFor    MultiSearchResultsInfo `json:"known_for"`
	ID          int
	MediaType   string `json:"media_type"`
}

func (MultiSearchPersonInfo) interfaceMarkerMethod() { return }

// GetMoviesKnownFor func
func (res MultiSearchPersonInfo) GetMoviesKnownFor() (movieResults []MultiSearchMovieInfo) {
	movieResults = make([]MultiSearchMovieInfo, 0)

	for i := 0; i < len(res.KnownFor); i++ {
		var base interface{} = res.KnownFor[i]
		if casted, ok := base.(*MultiSearchMovieInfo); ok {
			movieResults = append(movieResults, *casted)
		}
	}

	return
}

// GetTvKnownFor func
func (res MultiSearchPersonInfo) GetTvKnownFor() (tvResults []MultiSearchTvInfo) {
	tvResults = make([]MultiSearchTvInfo, 0)

	for i := 0; i < len(res.KnownFor); i++ {
		var base interface{} = res.KnownFor[i]
		if casted, ok := base.(*MultiSearchTvInfo); ok {
			tvResults = append(tvResults, *casted)
		}
	}

	return
}

// MultiSearchResultsInfo type
type MultiSearchResultsInfo []MultiSearchBase

// UnmarshalJSON func splits up the JSON array into the raw JSON
// for each object then unamrshal into a map to check the "type"
// field and unmarshal again into the correct type.
func (v *MultiSearchResultsInfo) UnmarshalJSON(data []byte) error {
	var raw []json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	for _, r := range raw {
		var obj map[string]interface{}
		err := json.Unmarshal(r, &obj)
		if err != nil {
			return err
		}

		mediaType := ""
		if t, ok := obj["media_type"].(string); ok {
			mediaType = t
		}

		var actual MultiSearchBase
		switch mediaType {
		case "movie":
			actual = &MultiSearchMovieInfo{}
		case "tv":
			actual = &MultiSearchTvInfo{}
		case "person":
			actual = &MultiSearchPersonInfo{}
		default:
			return ErrUnknownMediaType
		}

		err = json.Unmarshal(r, actual)
		if err != nil {
			return err
		}
		*v = append(*v, actual)
	}
	return nil
}

// MultiSearchResults struct
type MultiSearchResults struct {
	Page         int
	Results      MultiSearchResultsInfo
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// GetMoviesResults func
func (res MultiSearchResults) GetMoviesResults() (movieResults []MultiSearchMovieInfo) {

	for i := 0; i < len(res.Results); i++ {
		var base interface{} = res.Results[i]
		if casted, ok := base.(*MultiSearchMovieInfo); ok {
			movieResults = append(movieResults, *casted)
		}
	}

	return
}

// GetTvResults func
func (res MultiSearchResults) GetTvResults() (tvResults []MultiSearchTvInfo) {

	for i := 0; i < len(res.Results); i++ {
		var base interface{} = res.Results[i]

		if casted, ok := base.(*MultiSearchTvInfo); ok {
			tvResults = append(tvResults, *casted)
		}
	}

	return
}

// GetPersonResults func
func (res MultiSearchResults) GetPersonResults() (personResults []MultiSearchPersonInfo) {

	for i := 0; i < len(res.Results); i++ {
		var base interface{} = res.Results[i]
		if casted, ok := base.(*MultiSearchPersonInfo); ok {
			personResults = append(personResults, *casted)
		}
	}

	return
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
// https://developers.themoviedb.org/3/search/search-collections
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
// https://developers.themoviedb.org/3/search/search-companies
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
// https://developers.themoviedb.org/3/search/search-keywords
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
// No current documentation exists for this endpoint
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
// https://developers.themoviedb.org/3/search/search-movies
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
// https://developers.themoviedb.org/3/search/multi-search
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
// https://developers.themoviedb.org/3/search/search-people
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
// https://developers.themoviedb.org/3/search/search-tv-shows
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
