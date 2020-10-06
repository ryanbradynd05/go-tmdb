package tmdb

import (
	"fmt"
)

// Certification struct
type Certification struct {
	Certifications map[string][]struct {
		Certification string
		Meaning       string
		Order         int
	}
}

// GetCertificationsMovieList for movies
// https://developers.themoviedb.org/3/certifications/get-movie-certifications
func (tmdb *TMDb) GetCertificationsMovieList() (*Certification, error) {
	var movieCert Certification
	uri := fmt.Sprintf("%s/certification/movie/list?api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &movieCert)
	return result.(*Certification), err
}

// GetCertificationsTvList for tv shows
// https://developers.themoviedb.org/3/certifications/get-tv-certifications
func (tmdb *TMDb) GetCertificationsTvList() (*Certification, error) {
	var tvCert Certification
	uri := fmt.Sprintf("%s/certification/tv/list?api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &tvCert)
	return result.(*Certification), err
}
