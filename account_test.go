package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestGetAccountInfo(c *C) {
	result, err := s.tmdb.GetAccountInfo(session)
	s.baseTest(&result, err, c)
	c.Assert(result.Username, Equals, user)
}

func (s *TmdbSuite) TestGetAccountLists(c *C) {
	result, err := s.tmdb.GetAccountLists(accountID, session, nil)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestGetAccountFavoriteMovies(c *C) {
	result, err := s.tmdb.GetAccountFavoriteMovies(accountID, session, nil)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestGetAccountFavoriteTv(c *C) {
	result, err := s.tmdb.GetAccountFavoriteTv(accountID, session, nil)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestGetAccountRatedMovies(c *C) {
	result, err := s.tmdb.GetAccountRatedMovies(accountID, session, nil)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestGetAccountRatedTv(c *C) {
	result, err := s.tmdb.GetAccountRatedTv(accountID, session, nil)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestGetAccountWatchlistMovies(c *C) {
	result, err := s.tmdb.GetAccountWatchlistMovies(accountID, session, nil)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestGetAccountWatchlistTv(c *C) {
	result, err := s.tmdb.GetAccountWatchlistTv(accountID, session, nil)
	s.baseTest(&result, err, c)
}
