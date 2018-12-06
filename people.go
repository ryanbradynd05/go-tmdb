package tmdb

import (
	"fmt"
)

// Person struct
type Person struct {
	ID              int                    `json:"id"`
	Name            string                 `json:"name"`
	Overview        string                 `json:"overview"`
	Adult           bool                   `json:"adult"`
	Biography       string                 `json:"biography"`
	Birthday        string                 `json:"birthday"`
	Deathday        string                 `json:"deathday"`
	Gender          int                    `json:"gender"`
	ImdbID          string                 `json:"imdb_id"`
	Homepage        string                 `json:"homepage"`
	AlsoKnownAs     []string               `json:"also_known_as"`
	PlaceOfBirth    string                 `json:"place_of_birth"`
	ProfilePath     string                 `json:"profile_path"`
	Changes         *PersonChanges         `json:",omitempty"`
	MovieCredits    *PersonMovieCredits    `json:"movie_credits,omitempty"`
	TvCredits       *PersonTvCredits       `json:"tv_credits,omitempty"`
	CombinedCredits *PersonCombinedCredits `json:"combined_credits,omitempty"`
	ExternalIds     *TvExternalIds         `json:"external_ids,omitempty"`
	Images          *PersonImages          `json:",omitempty"`
	TaggedImages    *PersonTaggedImages    `json:"tagged_images,omitempty"`
	Translations    *PersonTranslations    `json:"translations,omitempty"`
}

// PersonTranslations struct
type PersonTranslations struct {
	ID           int
	Translations []struct {
		Iso639_1    string `json:"iso_639_1"`
		Name        string `json:"name"`
		EnglishName string `json:"english_name"`
		Data        struct {
			Title    string `json:"title,omitempty"`
			Overview string `json:"overview,omitempty"`
			Homepage string `json:"homepage,omitempty"`
		} `json:"data"`
	}
}

// PersonShort struct
type PersonShort struct {
	Adult       bool
	ID          int
	Name        string
	Popularity  float32
	ProfilePath string       `json:"profile_path"`
	KnownFor    []MovieShort `json:"known_for"`
}

// PersonChanges struct
type PersonChanges struct {
	Changes []struct {
		Key   string
		Items []struct {
			ID     string
			Action string
			Time   string
		}
	}
}

// PersonCombinedCredits struct
type PersonCombinedCredits struct {
	ID   int
	Cast []struct {
		Adult         bool
		Character     string
		CreditID      string `json:"credit_id"`
		ID            int
		OriginalTitle string `json:"original_title"`
		PosterPath    string `json:"poster_path"`
		ReleaseDate   string `json:"release_date"`
		Title         string
		MediaType     string `json:"media_type"`
	}
	Crew []struct {
		Adult         bool
		CreditID      string `json:"credit_id"`
		Department    string
		ID            int
		Job           string
		OriginalTitle string `json:"original_title"`
		PosterPath    string `json:"poster_path"`
		ReleaseDate   string `json:"release_date"`
		Title         string
		MediaType     string `json:"media_type"`
	}
}

// PersonImages struct
type PersonImages struct {
	ID       int
	Profiles []struct {
		AspectRatio float32 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		ID          string  `json:"id"`
		Width       int     `json:"width"`
		Height      int     `json:"height"`
		Iso639_1    string  `json:"iso_639_1"`
		VoteAverage float32 `json:"vote_average"`
		VoteCount   int     `json:"vote_count"`
	}
}

// PersonLatest struct
type PersonLatest struct {
	ID           int
	Name         string
	Adult        bool
	Biography    string
	Birthday     string
	Deathday     string
	Homepage     string
	AlsoKnownAs  []string `json:"also_known_as"`
	PlaceOfBirth string   `json:"place_of_birth"`
	ProfilePath  string   `json:"profile_path"`
}

// PersonMovieCredits struct
type PersonMovieCredits struct {
	ID   int
	Cast []struct {
		Adult         bool
		Character     string
		CreditID      string `json:"credit_id"`
		ID            int
		OriginalTitle string `json:"original_title"`
		PosterPath    string `json:"poster_path"`
		ReleaseDate   string `json:"release_date"`
		Title         string
	}
	Crew []struct {
		Adult         bool
		CreditID      string `json:"credit_id"`
		Department    string
		ID            int
		Job           string
		OriginalTitle string `json:"original_title"`
		PosterPath    string `json:"poster_path"`
		ReleaseDate   string `json:"release_date"`
		Title         string
	}
}

// PersonPopular struct
type PersonPopular struct {
	Page         int
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
	Results      []PersonShort
}

// PersonTaggedImages struct
type PersonTaggedImages struct {
	ID           int
	Page         int
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
	Results      []struct {
		AspectRatio float32 `json:"aspect_ratio"`
		FilePath    string  `json:"file_path"`
		ID          string
		Width       int
		Height      int
		Iso639_1    string  `json:"iso_639_1"`
		VoteAverage float32 `json:"vote_average"`
		VoteCount   int     `json:"vote_count"`
		ImageType   string  `json:"image_type"`
		MediaType   string  `json:"media_type"`
		Media       struct {
			Adult         bool
			BackdropPath  string `json:"backdrop_path"`
			ID            int
			OriginalTitle string `json:"original_title"`
			PosterPath    string `json:"poster_path"`
			ReleaseDate   string `json:"release_date"`
			Title         string
			Popularity    float32
		}
	}
}

// PersonTvCredits struct
type PersonTvCredits struct {
	ID   int
	Cast []struct {
		Character    string
		CreditID     string `json:"credit_id"`
		EpisodeCount int    `json:"episode_count"`
		FirstAirDate string `json:"first_air_date"`
		ID           int
		Name         string
		OriginalName string `json:"original_name"`
		PosterPath   string `json:"poster_path"`
	}
	Crew []struct {
		CreditID     string `json:"credit_id"`
		Department   string
		FirstAirDate string `json:"first_air_date"`
		ID           int
		Job          string
		Name         string
		OriginalName string `json:"original_name"`
		PosterPath   string `json:"poster_path"`
	}
}

// GetPersonInfo gets the general person information for a specific id
// http://docs.themoviedb.apiary.io/#reference/people/personid/get
func (tmdb *TMDb) GetPersonInfo(id int, options map[string]string) (*Person, error) {
	var availableOptions = map[string]struct{}{
		"append_to_response": {}}
	var personInfo Person
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/person/%v?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &personInfo)
	return result.(*Person), err
}

// GetPersonChanges for a specific person id
// http://docs.themoviedb.apiary.io/#reference/people/personidchanges/get
func (tmdb *TMDb) GetPersonChanges(id int, options map[string]string) (*PersonChanges, error) {
	var availableOptions = map[string]struct{}{
		"start_date": {},
		"end_date":   {}}
	var changes PersonChanges
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/person/%v/changes?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &changes)
	return result.(*PersonChanges), err
}

// GetPersonCombinedCredits gets the combined (movie and TV) credits for a specific person id
// http://docs.themoviedb.apiary.io/#reference/people/personidcombinedcredits/get
func (tmdb *TMDb) GetPersonCombinedCredits(id int, options map[string]string) (*PersonCombinedCredits, error) {
	var availableOptions = map[string]struct{}{
		"language":           {},
		"append_to_response": {}}
	var credits PersonCombinedCredits
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/person/%v/combined_credits?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &credits)
	return result.(*PersonCombinedCredits), err
}

// GetPersonExternalIds gets the external ids for a specific person id
// http://docs.themoviedb.apiary.io/#reference/people/personidexternalids/get
func (tmdb *TMDb) GetPersonExternalIds(id int) (*TvExternalIds, error) {
	var ids TvExternalIds
	uri := fmt.Sprintf("%s/person/%v/external_ids?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := getTmdb(uri, &ids)
	return result.(*TvExternalIds), err
}

// GetPersonImages gets the images for a specific person id
// http://docs.themoviedb.apiary.io/#reference/people/personidimages/get
func (tmdb *TMDb) GetPersonImages(id int) (*PersonImages, error) {
	var images PersonImages
	uri := fmt.Sprintf("%s/person/%v/images?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := getTmdb(uri, &images)
	return result.(*PersonImages), err
}

// GetPersonLatest gets the latest person id
// http://docs.themoviedb.apiary.io/#reference/people/personlatest/get
func (tmdb *TMDb) GetPersonLatest() (*PersonLatest, error) {
	var latest PersonLatest
	uri := fmt.Sprintf("%s/person/latest?api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &latest)
	return result.(*PersonLatest), err
}

// GetPersonMovieCredits gets the movie credits for a specific person id
// http://docs.themoviedb.apiary.io/#reference/people/personidmoviecredits/get
func (tmdb *TMDb) GetPersonMovieCredits(id int, options map[string]string) (*PersonMovieCredits, error) {
	var availableOptions = map[string]struct{}{
		"language":           {},
		"append_to_response": {}}
	var credits PersonMovieCredits
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/person/%v/movie_credits?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &credits)
	return result.(*PersonMovieCredits), err
}

// GetPersonPopular gets the list of popular people on The Movie Database
// http://docs.themoviedb.apiary.io/#reference/people/personpopular/get
func (tmdb *TMDb) GetPersonPopular(options map[string]string) (*PersonPopular, error) {
	var availableOptions = map[string]struct{}{
		"page": {}}
	var popular PersonPopular
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/person/popular?api_key=%s%s", baseURL, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &popular)
	return result.(*PersonPopular), err
}

// GetPersonTaggedImages gets the images that have been tagged with a specific person id
// http://docs.themoviedb.apiary.io/#reference/people/personidtaggedimages/get
func (tmdb *TMDb) GetPersonTaggedImages(id int, options map[string]string) (*PersonTaggedImages, error) {
	var availableOptions = map[string]struct{}{
		"language": {},
		"page":     {}}
	var images PersonTaggedImages
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/person/%v/tagged_images?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &images)
	return result.(*PersonTaggedImages), err
}

// GetPersonTvCredits gets the TV credits for a specific person id
// http://docs.themoviedb.apiary.io/#reference/people/personidtvcredits/get
func (tmdb *TMDb) GetPersonTvCredits(id int, options map[string]string) (*PersonTvCredits, error) {
	var availableOptions = map[string]struct{}{
		"language":           {},
		"append_to_response": {}}
	var credits PersonTvCredits
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/person/%v/tv_credits?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &credits)
	return result.(*PersonTvCredits), err
}
