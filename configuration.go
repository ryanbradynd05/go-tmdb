package tmdb

import (
	"fmt"
)

// Config struct
type Config struct {
	Images     ImageConfig
	ChangeKeys []string `json:"change_keys"`
}

// ImageConfig struct
type ImageConfig struct {
	BaseURL       string   `json:"base_url"`
	SecureBaseURL string   `json:"secure_base_url"`
	BackdropSizes []string `json:"backdrop_sizes"`
	LogoSizes     []string `json:"logo_sizes"`
	PosterSizes   []string `json:"poster_sizes"`
	ProfileSizes  []string `json:"profile_sizes"`
	StillSizes    []string `json:"still_sizes"`
}

// Configuration gets the system wide configuration information
// http://docs.themoviedb.apiary.io/#reference/configuration/configuration/get
func (tmdb *TMDb) Configuration() (*Config, error) {
	var config Config
	url := fmt.Sprintf("%s/configuration?api_key=%s", baseURL, tmdb.apiKey)
	result, err := callTmdb(url, &config)
	return result.(*Config), err
}
