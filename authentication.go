package tmdb

import (
	"fmt"
)

// AuthenticationToken struct
type AuthenticationToken struct {
	Success      bool
	RequestToken string `json:"request_token"`
	ExpiresAt    string `json:"expires_at"`
}

// GetNewToken generates a valid request token for user based authentication
// http://docs.themoviedb.apiary.io/#reference/authentication/authenticationtokennew/get
func (tmdb *TMDb) GetNewToken() (*AuthenticationToken, error) {
	var token AuthenticationToken
	url := fmt.Sprintf("%s/authentication/token/new?api_key=%s", baseURL, tmdb.apiKey)
	result, err := callTmdb(url, &token)
	return result.(*AuthenticationToken), err
}
