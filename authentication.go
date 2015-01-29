package tmdb

import (
	"fmt"
)

// Token struct
type Token struct {
	Success      bool
	RequestToken string `json:"request_token"`
	ExpiresAt    string `json:"expires_at"`
}

// NewToken generates a valid request token for user based authentication
// http://docs.themoviedb.apiary.io/#reference/authentication/authenticationtokennew/get
func (tmdb *TMDb) NewToken() (*Token, error) {
	var token Token
	url := fmt.Sprintf("%s/authentication/token/new?api_key=%s", baseURL, tmdb.apiKey)
	result, err := callTmdb(url, &token)
	return result.(*Token), err
}
