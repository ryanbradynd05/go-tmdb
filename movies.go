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
	Genres              []Genre
	Homepage            string
	ID                  int
	ImdbID              string `json:"imdb_id"`
	OriginalLanguage    string `json:"original_language"`
	OriginalTitle       string `json:"original_title"`
	Overview            string
	Popularity          float32
	PosterPath          string              `json:"poster_path"`
	ProductionCompanies []ProductionCompany `json:"production_companies"`
	ProductionCountries []InternationalName `json:"production_countries"`
	ReleaseDate         string              `json:"release_date"`
	Revenue             uint32
	Runtime             uint32
	SpokenLanguages     []SpokenLanguage `json:"spoken_languages"`
	Status              string
	Tagline             string
	Title               string
	Video               bool
	VoteAverage         float32 `json:"vote_average"`
	VoteCount           uint32  `json:"vote_count"`
}

// Genre struct
type Genre struct {
	ID   int
	Name string
}

// ProductionCompany struct
type ProductionCompany struct {
	ID   int
	Name string
}

// InternationalName struct
type InternationalName struct {
	Iso3166_1 string `json:"iso_3166_1"`
	Name      string
}

// SpokenLanguage struct
type SpokenLanguage struct {
	Iso639_1 string `json:"iso_639_1"`
	Name     string
}

// MovieInfo gets the basic movie information for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieid/get
func (tmdb *TMDb) MovieInfo(id int) (*Movie, error) {
	var movie Movie
	url := fmt.Sprintf("%s/movie/%v?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &movie)
	return result.(*Movie), err
}

// InternationalTitle struct
type InternationalTitle struct {
	Iso3166_1 string `json:"iso_3166_1"`
	Title     string
}

// AlternativeTitles struct
type AlternativeTitles struct {
	ID     int
	Titles []InternationalTitle
}

// MovieAlternativeTitles gets the alternative titles for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidalternativetitles/get
func (tmdb *TMDb) MovieAlternativeTitles(id int) (*AlternativeTitles, error) {
	var titles AlternativeTitles
	url := fmt.Sprintf("%s/movie/%v/alternative_titles?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &titles)
	return result.(*AlternativeTitles), err
}

// Cast struct
type Cast struct {
	CastID      int `json:"cast_id"`
	Character   string
	CreditID    string `json:"credit_id"`
	ID          int
	Name        string
	Order       int
	ProfilePath string `json:"profile_path"`
}

// Credits struct
type Credits struct {
	ID   int
	Cast []Cast
}

// MovieCredits gets the credits for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidcredits/get
func (tmdb *TMDb) MovieCredits(id int) (*Credits, error) {
	var credits Credits
	url := fmt.Sprintf("%s/movie/%v/credits?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &credits)
	return result.(*Credits), err
}

// Backdrop struct
type Backdrop struct {
	FilePath    string `json:"file_path"`
	Width       int
	Height      int
	Iso639_1    string  `json:"iso_639_1"`
	AspectRatio float32 `json:"aspect_ratio"`
	VoteAverage float32 `json:"vote_average"`
	VoteCount   uint32  `json:"vote_count"`
}

// Images struct
type Images struct {
	ID        int
	Backdrops []Backdrop
}

// MovieImages gets the images for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidimages/get
func (tmdb *TMDb) MovieImages(id int) (*Images, error) {
	var images Images
	url := fmt.Sprintf("%s/movie/%v/images?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &images)
	return result.(*Images), err
}

// Keyword struct
type Keyword struct {
	ID   int
	Name string
}

// Keywords struct
type Keywords struct {
	ID       int
	Keywords []Keyword
}

// MovieKeywords gets the keywords for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidkeywords/get
func (tmdb *TMDb) MovieKeywords(id int) (*Keywords, error) {
	var keywords Keywords
	url := fmt.Sprintf("%s/movie/%v/keywords?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &keywords)
	return result.(*Keywords), err
}

// Country struct
type Country struct {
	Iso3166_1     string `json:"iso_3166_1"`
	Certification string
	ReleaseDate   string `json:"release_date"`
}

// Releases struct
type Releases struct {
	ID        int
	Countries []Country
}

// MovieReleases gets the releases for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidreleases/get
func (tmdb *TMDb) MovieReleases(id int) (*Releases, error) {
	var releases Releases
	url := fmt.Sprintf("%s/movie/%v/releases?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &releases)
	return result.(*Releases), err
}

// VideoResult struct
type VideoResult struct {
	ID       int
	Iso639_1 string `json:"iso_639_1"`
	Key      string
	Name     string
	Site     string
	Size     int
	Type     string
}

// Videos struct
type Videos struct {
	ID      int
	Results []VideoResult
}

// MovieVideos gets the videos for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidvideos/get
func (tmdb *TMDb) MovieVideos(id int) (*Videos, error) {
	var videos Videos
	url := fmt.Sprintf("%s/movie/%v/videos?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &videos)
	return result.(*Videos), err
}

// Translation struct
type Translation struct {
	Iso639_1    string `json:"iso_639_1"`
	Name        string
	EnglishName string `json:"english_name"`
}

// Translations struct
type Translations struct {
	ID           int
	Translations []Translation
}

// MovieTranslations gets the translations for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidtranslations/get
func (tmdb *TMDb) MovieTranslations(id int) (*Translations, error) {
	var translations Translations
	url := fmt.Sprintf("%s/movie/%v/translations?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &translations)
	return result.(*Translations), err
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

// ReviewResult struct
type ReviewResult struct {
	ID      string
	Author  string
	Content string
	URL     string
}

// Reviews struct
type Reviews struct {
	ID           int
	Page         int
	Results      []ReviewResult
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

// ListResult struct
type ListResult struct {
	Description   string
	FavoriteCount int `json:"favorite_count"`
	ID            string
	ItemCount     int    `json:"item_count"`
	Iso639_1      string `json:"iso_639_1"`
	Name          string
	PosterPath    string `json:"poster_path"`
}

// Lists struct
type Lists struct {
	ID           int
	Page         int
	Results      []ListResult
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

// ChangeItem struct
type ChangeItem struct {
	ID     string
	Action string
	Time   string
}

// Change struct
type Change struct {
	Key   string
	Items []ChangeItem
}

// Changes struct
type Changes struct {
	Changes []Change
}

// MovieChanges gets the changes for a specific movie id
// http://docs.themoviedb.apiary.io/#reference/movies/movieidchanges/get
func (tmdb *TMDb) MovieChanges(id int) (*Changes, error) {
	var changes Changes
	url := fmt.Sprintf("%s/movie/%v/changes?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(url, &changes)
	return result.(*Changes), err
}

// MovieLatest gets the latest movie
// http://docs.themoviedb.apiary.io/#reference/movies/movielatest/get
func (tmdb *TMDb) MovieLatest() (*Movie, error) {
	var movie Movie
	url := fmt.Sprintf("%s/movie/latest?api_key=%s", baseURL, tmdb.apiKey)
	result, err := callTmdb(url, &movie)
	return result.(*Movie), err
}

// DateRange struct
type DateRange struct {
	Minimum string
	Maximum string
}

// DatedResults struct
type DatedResults struct {
	Dates        DateRange
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
