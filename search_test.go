package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestSearchMovie(c *C) {
	result, err := s.tmdb.SearchMovie("Fight Club", nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.TotalResults, Not(Equals), 0)
	c.Assert(result.Results[0].Title, Equals, "Fight Club")

	var options = make(map[string]string)
	options["language"] = "es"
	spanishResult, err := s.tmdb.SearchMovie("Fight Club", options)
	s.baseTest(&spanishResult, err, c)
	c.Assert(spanishResult.Results, Not(HasLen), 0)
	c.Assert(spanishResult.TotalResults, Not(Equals), 0)
	c.Assert(spanishResult.Results[0].OriginalTitle, Equals, "Fight Club")
	c.Assert(spanishResult.Results[0].Title, Equals, "El club de la lucha")
}
