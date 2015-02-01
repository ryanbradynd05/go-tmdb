package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestGetChangesMovie(c *C) {
	movieResult, err := s.tmdb.GetChangesMovie(nil)
	s.baseTest(&movieResult, err, c)
	c.Assert(movieResult.Results, NotNil)
	c.Assert(movieResult.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetChangesPerson(c *C) {
	personResult, err := s.tmdb.GetChangesPerson(nil)
	s.baseTest(&personResult, err, c)
	c.Assert(personResult.Results, NotNil)
	c.Assert(personResult.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetChangesTv(c *C) {
	tvResult, err := s.tmdb.GetChangesTv(nil)
	s.baseTest(&tvResult, err, c)
	c.Assert(tvResult.Results, NotNil)
	c.Assert(tvResult.Results, Not(HasLen), 0)
}
