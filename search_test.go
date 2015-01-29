package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestSearchMovie(c *C) {
	result, err := s.tmdb.SearchMovie("Fight Club")
	s.baseTest(&result, err, c)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.TotalResults, Not(Equals), 0)
}
