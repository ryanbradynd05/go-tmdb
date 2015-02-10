package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestGetAccountInfo(c *C) {
	result, err := s.tmdb.GetAccountInfo(s.session)
	s.baseTest(&result, err, c)
	c.Assert(result.Username, Equals, s.user)
}

func (s *TmdbSuite) TestGetAccountLists(c *C) {
	result, err := s.tmdb.GetAccountLists(s.accountID, s.session, nil)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestGetAccountFavoriteMovies(c *C) {
	result, err := s.tmdb.GetAccountFavoriteMovies(s.accountID, s.session, nil)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestGetAccountFavoriteTv(c *C) {
	result, err := s.tmdb.GetAccountFavoriteTv(s.accountID, s.session, nil)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestGetAccountRatedMovies(c *C) {
	result, err := s.tmdb.GetAccountRatedMovies(s.accountID, s.session, nil)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestGetAccountRatedTv(c *C) {
	result, err := s.tmdb.GetAccountRatedTv(s.accountID, s.session, nil)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestGetAccountWatchlistMovies(c *C) {
	result, err := s.tmdb.GetAccountWatchlistMovies(s.accountID, s.session, nil)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestGetAccountWatchlistTv(c *C) {
	result, err := s.tmdb.GetAccountWatchlistTv(s.accountID, s.session, nil)
	s.baseTest(&result, err, c)
}
