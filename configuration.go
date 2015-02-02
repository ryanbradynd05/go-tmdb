package tmdb

import (
	"fmt"
)

// Configuration struct
type Configuration struct {
	Images struct {
		BaseURL       string   `json:"base_url"`
		SecureBaseURL string   `json:"secure_base_url"`
		BackdropSizes []string `json:"backdrop_sizes"`
		LogoSizes     []string `json:"logo_sizes"`
		PosterSizes   []string `json:"poster_sizes"`
		ProfileSizes  []string `json:"profile_sizes"`
		StillSizes    []string `json:"still_sizes"`
	}
	ChangeKeys []string `json:"change_keys"`
}

// GetConfiguration gets the system wide configuration information
// http://docs.themoviedb.apiary.io/#reference/configuration/configuration/get
func (tmdb *TMDb) GetConfiguration() (*Configuration, error) {
	var config Configuration
	uri := fmt.Sprintf("%s/configuration?api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &config)
	return result.(*Configuration), err
}
