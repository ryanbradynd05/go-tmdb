package tmdb

import (
	"fmt"
)

// Movie struct
type Movie struct {
	Adult               bool
	BackdropPath        string `json:"backdrop_path"`
	BelongsToCollection bool   `json:"belongs_to_collection"`
	Budget              uint32
	Genres              []struct {
		ID   int
		Name string
	}
	Homepage            string
	ID                  int
	ImdbID              string `json:"imdb_id"`
	OriginalLanguage    string `json:"original_language"`
	OriginalTitle       string `json:"original_title"`
	Overview            string
	Popularity          float32
	PosterPath          string `json:"poster_path"`
	ProductionCompanies []struct {
		ID   int
		Name string
	} `json:"production_companies"`
	ProductionCountries []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Name      string
	} `json:"production_countries"`
	ReleaseDate     string `json:"release_date"`
	Revenue         uint32
	Runtime         uint32
	SpokenLanguages []struct {
		Iso639_1 string `json:"iso_639_1"`
		Name     string
	} `json:"spoken_languages"`
	Status      string
	Tagline     string
	Title       string
	Video       bool
	VoteAverage float32 `json:"vote_average"`
	VoteCount   uint32  `json:"vote_count"`
}

// ShortMovie struct
type ShortMovie struct {
	Adult         bool
	BackdropPath  string `json:"backdrop_path"`
	ID            int
	OriginalTitle string `json:"original_title"`
	Popularity    float32
	PosterPath    string `json:"poster_path"`
	ReleaseDate   string `json:"release_date"`
	Title         string
	Video         bool
	VoteAverage   float32 `json:"vote_average"`
	VoteCount     uint32  `json:"vote_count"`
}

// MovieInfo gets the basic movie information for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieid/get
func (tmdb *TMDb) MovieInfo(id int) (*Movie, error) {
	var movie Movie
	url := fmt.Sprintf("%s/movie/%v?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &movie)
	return result.(*Movie), err
}

// AccountState struct
type AccountState struct {
	ID        int
	Favorite  bool
	Watchlist bool
	Rated     struct {
		Value float32
	}
}

// MovieAccountStates gets the status of whether or not the movie has been rated or added to their favourite or movie watch list
// http://docs.themoviedb.apiary.io/#reference/movies/movieidaccountstates/get
func (tmdb *TMDb) MovieAccountStates(id int) (*AccountState, error) {
	// TODO
	var state AccountState
	var err error
	return &state, err
}

// AlternativeTitles struct
type AlternativeTitles struct {
	ID     int
	Titles []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Title     string
	}
}

// MovieAlternativeTitles gets the alternative titles for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidalternativetitles/get
func (tmdb *TMDb) MovieAlternativeTitles(id int) (*AlternativeTitles, error) {
	var titles AlternativeTitles
	url := fmt.Sprintf("%s/movie/%v/alternative_titles?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &titles)
	return result.(*AlternativeTitles), err
}

// Credits struct
type Credits struct {
	ID   int
	Cast []struct {
		CastID      int `json:"cast_id"`
		Character   string
		CreditID    string `json:"credit_id"`
		ID          int
		Name        string
		Order       int
		ProfilePath string `json:"profile_path"`
	}
}

// MovieCredits gets the credits for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidcredits/get
func (tmdb *TMDb) MovieCredits(id int) (*Credits, error) {
	var credits Credits
	url := fmt.Sprintf("%s/movie/%v/credits?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &credits)
	return result.(*Credits), err
}

// Images struct
type Images struct {
	ID        int
	Backdrops []struct {
		FilePath    string `json:"file_path"`
		Width       int
		Height      int
		Iso639_1    string  `json:"iso_639_1"`
		AspectRatio float32 `json:"aspect_ratio"`
		VoteAverage float32 `json:"vote_average"`
		VoteCount   uint32  `json:"vote_count"`
	}
}

// MovieImages gets the images for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidimages/get
func (tmdb *TMDb) MovieImages(id int) (*Images, error) {
	var images Images
	url := fmt.Sprintf("%s/movie/%v/images?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &images)
	return result.(*Images), err
}

// Keywords struct
type Keywords struct {
	ID       int
	Keywords []struct {
		ID   int
		Name string
	}
}

// MovieKeywords gets the keywords for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidkeywords/get
func (tmdb *TMDb) MovieKeywords(id int) (*Keywords, error) {
	var keywords Keywords
	url := fmt.Sprintf("%s/movie/%v/keywords?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &keywords)
	return result.(*Keywords), err
}

// Releases struct
type Releases struct {
	ID        int
	Countries []struct {
		Iso3166_1     string `json:"iso_3166_1"`
		Certification string
		ReleaseDate   string `json:"release_date"`
	}
}

// MovieReleases gets the releases for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidreleases/get
func (tmdb *TMDb) MovieReleases(id int) (*Releases, error) {
	var releases Releases
	url := fmt.Sprintf("%s/movie/%v/releases?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &releases)
	return result.(*Releases), err
}

// Videos struct
type Videos struct {
	ID      int
	Results []struct {
		ID       int
		Iso639_1 string `json:"iso_639_1"`
		Key      string
		Name     string
		Site     string
		Size     int
		Type     string
	}
}

// MovieVideos gets the videos for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidvideos/get
func (tmdb *TMDb) MovieVideos(id int) (*Videos, error) {
	var videos Videos
	url := fmt.Sprintf("%s/movie/%v/videos?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &videos)
	return result.(*Videos), err
}

// TranslationsList struct
type TranslationsList struct {
	ID           int
	Translations []struct {
		Iso639_1    string `json:"iso_639_1"`
		Name        string
		EnglishName string `json:"english_name"`
	}
}

// MovieTranslations gets the translations for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidtranslations/get
func (tmdb *TMDb) MovieTranslations(id int) (*TranslationsList, error) {
	var translations TranslationsList
	url := fmt.Sprintf("%s/movie/%v/translations?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &translations)
	return result.(*TranslationsList), err
}

// PagedResults struct
type PagedResults struct {
	Page         int
	Results      []ShortMovie
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// MovieSimilar gets the similar movies for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidsimilar/get
func (tmdb *TMDb) MovieSimilar(id, page int) (*PagedResults, error) {
	var similar PagedResults
	url := fmt.Sprintf("%s/movie/%v/similar?page=%v&api_key=%s", baseURL, id, page, tmdb.apiKey)
	result, err := callTmdb(url, &similar)
	return result.(*PagedResults), err
}

// Reviews struct
type Reviews struct {
	ID      int
	Page    int
	Results []struct {
		ID      string
		Author  string
		Content string
		URL     string
	}
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// MovieReviews gets the reviews for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidreviews/get
func (tmdb *TMDb) MovieReviews(id, page int) (*Reviews, error) {
	var reviews Reviews
	url := fmt.Sprintf("%s/movie/%v/reviews?page=%v&api_key=%s", baseURL, id, page, tmdb.apiKey)
	result, err := callTmdb(url, &reviews)
	return result.(*Reviews), err
}

// Lists struct
type Lists struct {
	ID      int
	Page    int
	Results []struct {
		Description   string
		FavoriteCount int `json:"favorite_count"`
		ID            string
		ItemCount     int    `json:"item_count"`
		Iso639_1      string `json:"iso_639_1"`
		Name          string
		PosterPath    string `json:"poster_path"`
	}
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// MovieLists gets the lists that the movie belongs to
// http://docs.themoviedb.apiary.io/#reference/movies/movieidlists/get
func (tmdb *TMDb) MovieLists(id, page int) (*Lists, error) {
	var lists Lists
	url := fmt.Sprintf("%s/movie/%v/lists?page=%v&api_key=%s", baseURL, id, page, tmdb.apiKey)
	result, err := callTmdb(url, &lists)
	return result.(*Lists), err
}

// ChangeList struct
type ChangeList struct {
	Changes []struct {
		Key   string
		Items []struct {
			ID     string
			Action string
			Time   string
		}
	}
}

// MovieChanges gets the changes for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidchanges/get
func (tmdb *TMDb) MovieChanges(id int) (*ChangeList, error) {
	var changes ChangeList
	url := fmt.Sprintf("%s/movie/%v/changes?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &changes)
	return result.(*ChangeList), err
}

// RatingResponse struct
type RatingResponse struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

// MovieRating lets users rate a movie
// http://docs.themoviedb.apiary.io/#reference/movies/movieidrating/post
func (tmdb *TMDb) MovieRating(id int) (*RatingResponse, error) {
	// TODO
	var rating RatingResponse
	var err error
	return &rating, err
}

// MovieLatest gets the latest movie
// http://docs.themoviedb.apiary.io/#reference/movies/movielatest/get
func (tmdb *TMDb) MovieLatest() (*Movie, error) {
	var movie Movie
	url := fmt.Sprintf("%s/movie/latest?api_key=%s", baseURL, tmdb.apiKey)
	result, err := callTmdb(url, &movie)
	return result.(*Movie), err
}

// DatedResults struct
type DatedResults struct {
	Dates struct {
		Minimum string
		Maximum string
	}
	Page         int
	Results      []ShortMovie
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// MovieUpcoming gets the list of upcoming movies by release date
// http://docs.themoviedb.apiary.io/#reference/movies/movieupcoming/get
func (tmdb *TMDb) MovieUpcoming() (*DatedResults, error) {
	var upcoming DatedResults
	url := fmt.Sprintf("%s/movie/upcoming?api_key=%s", baseURL, tmdb.apiKey)
	result, err := callTmdb(url, &upcoming)
	return result.(*DatedResults), err
}

// NowPlaying gets the list of movies playing that have been, or are being released this week
// http://docs.themoviedb.apiary.io/#reference/movies/movienowplaying/get
func (tmdb *TMDb) NowPlaying() (*DatedResults, error) {
	var nowPlaying DatedResults
	url := fmt.Sprintf("%s/movie/now_playing?api_key=%s", baseURL, tmdb.apiKey)
	result, err := callTmdb(url, &nowPlaying)
	return result.(*DatedResults), err
}

// Popular gets the list of popular movies on The Movie Database
// http://docs.themoviedb.apiary.io/#reference/movies/moviepopular/get
func (tmdb *TMDb) Popular() (*PagedResults, error) {
	var popular PagedResults
	url := fmt.Sprintf("%s/movie/popular?api_key=%s", baseURL, tmdb.apiKey)
	result, err := callTmdb(url, &popular)
	return result.(*PagedResults), err
}

// TopRated gets the list of top rated movies
// http://docs.themoviedb.apiary.io/#reference/movies/movietoprated/get
func (tmdb *TMDb) TopRated() (*PagedResults, error) {
	var topRated PagedResults
	url := fmt.Sprintf("%s/movie/top_rated?api_key=%s", baseURL, tmdb.apiKey)
	result, err := callTmdb(url, &topRated)
	return result.(*PagedResults), err
}
