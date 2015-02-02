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
// http://docs.themoviedb.apiary.io/#reference/certifications/certificationmovielist/get
func (tmdb *TMDb) GetCertificationsMovieList() (*Certification, error) {
	var movieCert Certification
	uri := fmt.Sprintf("%s/certification/movie/list?api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &movieCert)
	return result.(*Certification), err
}

// GetCertificationsTvList for tv shows
// http://docs.themoviedb.apiary.io/#reference/certifications/certificationtvlist/get
func (tmdb *TMDb) GetCertificationsTvList() (*Certification, error) {
	var tvCert Certification
	uri := fmt.Sprintf("%s/certification/tv/list?api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &tvCert)
	return result.(*Certification), err
}
