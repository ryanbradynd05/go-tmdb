package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestGetAuthToken(c *C) {
	result, err := s.tmdb.GetAuthToken()
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestGetAuthValidateToken(c *C) {
	result, err := s.tmdb.GetAuthValidateToken(s.userToken, s.user, s.pw)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestGetAuthSession(c *C) {
	result, err := s.tmdb.GetAuthSession(s.userToken)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestGetAuthGuestSession(c *C) {
	result, err := s.tmdb.GetAuthGuestSession()
	s.baseTest(&result, err, c)
}
