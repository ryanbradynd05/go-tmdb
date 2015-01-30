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

// MovieShort struct
type MovieShort struct {
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

// GetMovieInfo for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieid/get
func (tmdb *TMDb) GetMovieInfo(id int) (*Movie, error) {
	var movie Movie
	uri := fmt.Sprintf("%s/movie/%v?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(uri, &movie)
	return result.(*Movie), err
}

// MovieAccountState struct
type MovieAccountState struct {
	ID        int
	Favorite  bool
	Watchlist bool
	Rated     struct {
		Value float32
	}
}

// GetMovieAccountStates gets the status of whether or not the movie has been rated or added to their favourite or movie watch list
// http://docs.themoviedb.apiary.io/#reference/movies/movieidaccountstates/get
func (tmdb *TMDb) GetMovieAccountStates(id int) (*MovieAccountState, error) {
	// TODO
	var state MovieAccountState
	var err error
	return &state, err
}

// MovieAlternativeTitles struct
type MovieAlternativeTitles struct {
	ID     int
	Titles []struct {
		Iso3166_1 string `json:"iso_3166_1"`
		Title     string
	}
}

// GetMovieAlternativeTitles for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidalternativetitles/get
func (tmdb *TMDb) GetMovieAlternativeTitles(id int) (*MovieAlternativeTitles, error) {
	var titles MovieAlternativeTitles
	uri := fmt.Sprintf("%s/movie/%v/alternative_titles?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(uri, &titles)
	return result.(*MovieAlternativeTitles), err
}

// MovieCredits struct
type MovieCredits struct {
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

// GetMovieCredits for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidcredits/get
func (tmdb *TMDb) GetMovieCredits(id int) (*MovieCredits, error) {
	var credits MovieCredits
	uri := fmt.Sprintf("%s/movie/%v/credits?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(uri, &credits)
	return result.(*MovieCredits), err
}

// MovieImages struct
type MovieImages struct {
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

// GetMovieImages for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidimages/get
func (tmdb *TMDb) GetMovieImages(id int) (*MovieImages, error) {
	var images MovieImages
	uri := fmt.Sprintf("%s/movie/%v/images?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(uri, &images)
	return result.(*MovieImages), err
}

// MovieKeywords struct
type MovieKeywords struct {
	ID       int
	Keywords []struct {
		ID   int
		Name string
	}
}

// GetMovieKeywords for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidkeywords/get
func (tmdb *TMDb) GetMovieKeywords(id int) (*MovieKeywords, error) {
	var keywords MovieKeywords
	uri := fmt.Sprintf("%s/movie/%v/keywords?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(uri, &keywords)
	return result.(*MovieKeywords), err
}

// MovieReleases struct
type MovieReleases struct {
	ID        int
	Countries []struct {
		Iso3166_1     string `json:"iso_3166_1"`
		Certification string
		ReleaseDate   string `json:"release_date"`
	}
}

// GetMovieReleases for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidreleases/get
func (tmdb *TMDb) GetMovieReleases(id int) (*MovieReleases, error) {
	var releases MovieReleases
	uri := fmt.Sprintf("%s/movie/%v/releases?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(uri, &releases)
	return result.(*MovieReleases), err
}

// MovieVideos struct
type MovieVideos struct {
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

// GetMovieVideos for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidvideos/get
func (tmdb *TMDb) GetMovieVideos(id int) (*MovieVideos, error) {
	var videos MovieVideos
	uri := fmt.Sprintf("%s/movie/%v/videos?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(uri, &videos)
	return result.(*MovieVideos), err
}

// MovieTranslations struct
type MovieTranslations struct {
	ID           int
	Translations []struct {
		Iso639_1    string `json:"iso_639_1"`
		Name        string
		EnglishName string `json:"english_name"`
	}
}

// GetMovieTranslations for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidtranslations/get
func (tmdb *TMDb) GetMovieTranslations(id int) (*MovieTranslations, error) {
	var translations MovieTranslations
	uri := fmt.Sprintf("%s/movie/%v/translations?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(uri, &translations)
	return result.(*MovieTranslations), err
}

// MoviePagedResults struct
type MoviePagedResults struct {
	Page         int
	Results      []MovieShort
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// GetSimilarMovies for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidsimilar/get
func (tmdb *TMDb) GetSimilarMovies(id, page int) (*MoviePagedResults, error) {
	var similar MoviePagedResults
	uri := fmt.Sprintf("%s/movie/%v/similar?page=%v&api_key=%s", baseURL, id, page, tmdb.apiKey)
	result, err := callTmdb(uri, &similar)
	return result.(*MoviePagedResults), err
}

// MovieReviews struct
type MovieReviews struct {
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

// GetMovieReviews for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidreviews/get
func (tmdb *TMDb) GetMovieReviews(id, page int) (*MovieReviews, error) {
	var reviews MovieReviews
	uri := fmt.Sprintf("%s/movie/%v/reviews?page=%v&api_key=%s", baseURL, id, page, tmdb.apiKey)
	result, err := callTmdb(uri, &reviews)
	return result.(*MovieReviews), err
}

// MovieLists struct
type MovieLists struct {
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

// GetMovieLists that the movie belongs to
// http://docs.themoviedb.apiary.io/#reference/movies/movieidlists/get
func (tmdb *TMDb) GetMovieLists(id, page int) (*MovieLists, error) {
	var lists MovieLists
	uri := fmt.Sprintf("%s/movie/%v/lists?page=%v&api_key=%s", baseURL, id, page, tmdb.apiKey)
	result, err := callTmdb(uri, &lists)
	return result.(*MovieLists), err
}

// MovieChanges struct
type MovieChanges struct {
	Changes []struct {
		Key   string
		Items []struct {
			ID     string
			Action string
			Time   string
		}
	}
}

// GetMovieChanges for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidchanges/get
func (tmdb *TMDb) GetMovieChanges(id int) (*MovieChanges, error) {
	var changes MovieChanges
	uri := fmt.Sprintf("%s/movie/%v/changes?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(uri, &changes)
	return result.(*MovieChanges), err
}

// MovieRating struct
type MovieRating struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

// SetMovieRating lets users rate a movie
// http://docs.themoviedb.apiary.io/#reference/movies/movieidrating/post
func (tmdb *TMDb) SetMovieRating(id int) (*MovieRating, error) {
	// TODO
	var rating MovieRating
	var err error
	return &rating, err
}

// GetLatestMovie gets the latest movie
// http://docs.themoviedb.apiary.io/#reference/movies/movielatest/get
func (tmdb *TMDb) GetLatestMovie() (*Movie, error) {
	var movie Movie
	uri := fmt.Sprintf("%s/movie/latest?api_key=%s", baseURL, tmdb.apiKey)
	result, err := callTmdb(uri, &movie)
	return result.(*Movie), err
}

// MovieDatedResults struct
type MovieDatedResults struct {
	Dates struct {
		Minimum string
		Maximum string
	}
	Page         int
	Results      []MovieShort
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// GetUpcomingMovies by release date
// http://docs.themoviedb.apiary.io/#reference/movies/movieupcoming/get
func (tmdb *TMDb) GetUpcomingMovies() (*MovieDatedResults, error) {
	var upcoming MovieDatedResults
	uri := fmt.Sprintf("%s/movie/upcoming?api_key=%s", baseURL, tmdb.apiKey)
	result, err := callTmdb(uri, &upcoming)
	return result.(*MovieDatedResults), err
}

// GetNowPlayingMovies that have been, or are being released this week
// http://docs.themoviedb.apiary.io/#reference/movies/movienowplaying/get
func (tmdb *TMDb) GetNowPlayingMovies() (*MovieDatedResults, error) {
	var nowPlaying MovieDatedResults
	uri := fmt.Sprintf("%s/movie/now_playing?api_key=%s", baseURL, tmdb.apiKey)
	result, err := callTmdb(uri, &nowPlaying)
	return result.(*MovieDatedResults), err
}

// GetPopularMovies gets the list of popular movies on The Movie Database
// http://docs.themoviedb.apiary.io/#reference/movies/moviepopular/get
func (tmdb *TMDb) GetPopularMovies() (*MoviePagedResults, error) {
	var popular MoviePagedResults
	uri := fmt.Sprintf("%s/movie/popular?api_key=%s", baseURL, tmdb.apiKey)
	result, err := callTmdb(uri, &popular)
	return result.(*MoviePagedResults), err
}

// GetTopRatedMovies gets the list of top rated movies
// http://docs.themoviedb.apiary.io/#reference/movies/movietoprated/get
func (tmdb *TMDb) GetTopRatedMovies() (*MoviePagedResults, error) {
	var topRated MoviePagedResults
	uri := fmt.Sprintf("%s/movie/top_rated?api_key=%s", baseURL, tmdb.apiKey)
	result, err := callTmdb(uri, &topRated)
	return result.(*MoviePagedResults), err
}
