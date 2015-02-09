package tmdb

import (
	"fmt"
)

// AuthenticationToken struct
type AuthenticationToken struct {
	Success      bool
	RequestToken string `json:"request_token"`
	ExpiresAt    string `json:"expires_at,omitempty"`
}

// AuthenticationSession struct
type AuthenticationSession struct {
	Success   bool
	SessionID string `json:"session_id"`
}

// AuthenticationGuestSession struct
type AuthenticationGuestSession struct {
	Success        bool
	GuestSessionID string `json:"guest_session_id"`
	ExpiresAt      string `json:"expires_at"`
}

// GetAuthToken generates a valid request token for user based authentication
// http://docs.themoviedb.apiary.io/#reference/authentication/authenticationtokennew/get
func (tmdb *TMDb) GetAuthToken() (*AuthenticationToken, error) {
	var token AuthenticationToken
	uri := fmt.Sprintf("%s/authentication/token/new?api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &token)
	return result.(*AuthenticationToken), err
}

// GetAuthValidateToken authenticates a user with a TMDb username and password
// http://docs.themoviedb.apiary.io/#reference/authentication/authenticationtokenvalidatewithlogin/get
func (tmdb *TMDb) GetAuthValidateToken(token, user, password string) (*AuthenticationToken, error) {
	var validToken AuthenticationToken
	uri := fmt.Sprintf("%s/authentication/token/validate_with_login?api_key=%s&request_token=%s&username=%s&password=%s", baseURL, tmdb.apiKey, token, user, password)
	result, err := getTmdb(uri, &validToken)
	return result.(*AuthenticationToken), err
}

// GetAuthSession generates a session id for user based authentication
// http://docs.themoviedb.apiary.io/#reference/authentication/authenticationsessionnew/get
func (tmdb *TMDb) GetAuthSession(token string) (*AuthenticationSession, error) {
	var session AuthenticationSession
	uri := fmt.Sprintf("%s/authentication/session/new?api_key=%s&request_token=%s", baseURL, tmdb.apiKey, token)
	result, err := getTmdb(uri, &session)
	return result.(*AuthenticationSession), err
}

// GetAuthGuestSession generates a valid request token for user based authentication
// http://docs.themoviedb.apiary.io/#reference/authentication/authenticationguestsessionnew/get
func (tmdb *TMDb) GetAuthGuestSession() (*AuthenticationGuestSession, error) {
	var session AuthenticationGuestSession
	uri := fmt.Sprintf("%s/authentication/guest_session/new?api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &session)
	return result.(*AuthenticationGuestSession), err
}
