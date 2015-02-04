package tmdb

import (
	"fmt"
)

// TV struct
type TV struct {
	BackdropPath string `json:"backdrop_path"`
	CreatedBy    []struct {
		ID          int
		Name        string
		ProfilePath string `json:"profile_path"`
	}
	EpisodeRunTime []int  `json:"episode_run_time"`
	FirstAirDate   string `json:"first_air_date"`
	Genres         []struct {
		ID   int
		Name string
	}
	Homepage     string
	ID           int
	InProduction bool `json:"in_production"`
	Languages    []string
	LastAirDate  string `json:"last_air_date"`
	Name         string
	Networks     []struct {
		ID   int
		Name string
	}
	NumberOfEpisodes    int      `json:"number_of_episodes"`
	NumberOfSeasons     int      `json:"number_of_seasons"`
	OriginCountry       []string `json:"origin_country"`
	OriginalLanguage    string   `json:"original_language"`
	OriginalName        string   `json:"original_name"`
	Overview            string
	Popularity          float32
	PosterPath          string `json:"poster_path"`
	ProductionCompanies []struct {
		ID   int
		Name string
	} `json:"production_companies"`
	Seasons []struct {
		AirDate      string `json:"air_date"`
		EpisodeCount int    `json:"episode_count"`
		ID           int
		PosterPath   string `json:"poster_path"`
		SeasonNumber int    `json:"season_number"`
	}
	Status            string
	Type              string
	VoteAverage       float32              `json:"vote_average"`
	VoteCount         uint32               `json:"vote_count"`
	AlternativeTitles *TvAlternativeTitles `json:"alternative_titles,omitempty"`
	Changes           *TvChanges           `json:",omitempty"`
	Credits           *TvCredits           `json:",omitempty"`
}

// TvAccountState struct
type TvAccountState struct {
	ID        int
	Favorite  bool
	Watchlist bool
	Rated     struct {
		Value float32
	}
}

// TvAlternativeTitles struct
type TvAlternativeTitles struct {
	ID      int
	Results []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Title     string
	}
}

// TvChanges struct
type TvChanges struct {
	Changes []struct {
		Key   string
		Items []struct {
			ID     string
			Action string
			Time   string
		}
	}
}

// TvCredits struct
type TvCredits struct {
	ID   int
	Cast []struct {
		Character   string
		CreditID    string `json:"credit_id"`
		ID          int
		Name        string
		Order       int
		ProfilePath string `json:"profile_path"`
	}
	AlternativeTitles *TvAlternativeTitles `json:"alternative_titles,omitempty"`
	Changes           *TvChanges           `json:",omitempty"`
}

// GetTvInfo gets the primary information about a TV series by id
// http://docs.themoviedb.apiary.io/#reference/tv/tvid/get
func (tmdb *TMDb) GetTvInfo(id int, options map[string]string) (*TV, error) {
	var availableOptions = map[string]struct{}{
		"language":           {},
		"append_to_response": {}}
	var tvInfo TV
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &tvInfo)
	return result.(*TV), err
}

// // GetTvAccountStates gets the status of whether or not the TV show has been rated or added to their favourite or watch lists
// // http://docs.themoviedb.apiary.io/#reference/tv/tvidaccountstates/get
// func (tmdb *TMDb) GetTvAccountStates(id int) (*TvAccountState, error) {
//  // TODO
//  var state TvAccountState
//  var err error
//  return &state, err
// }

// GetTvAlternativeTitles for a specific show id
// http://docs.themoviedb.apiary.io/#reference/tv/tvidalternativetitles/get
func (tmdb *TMDb) GetTvAlternativeTitles(id int) (*TvAlternativeTitles, error) {
	var titles TvAlternativeTitles
	uri := fmt.Sprintf("%s/tv/%v/alternative_titles?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := getTmdb(uri, &titles)
	return result.(*TvAlternativeTitles), err
}

// GetTvChanges for a specific show id
// http://docs.themoviedb.apiary.io/#reference/tv/tvidchanges/get
func (tmdb *TMDb) GetTvChanges(id int, options map[string]string) (*TvChanges, error) {
	var availableOptions = map[string]struct{}{
		"start_date": {},
		"end_date":   {}}
	var changes TvChanges
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v/changes?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &changes)
	json, _ := ToJSON(result)
	fmt.Printf("Result: %v\nURL: %v", json, uri)
	return result.(*TvChanges), err
}

// GetTvCredits for a specific TV show id
// http://docs.themoviedb.apiary.io/#reference/tv/tvidcredits/get
func (tmdb *TMDb) GetTvCredits(id int, options map[string]string) (*TvCredits, error) {
	var availableOptions = map[string]struct{}{
		"language":           {},
		"append_to_response": {}}
	var credits TvCredits
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/tv/%v/credits?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &credits)
	return result.(*TvCredits), err
}
