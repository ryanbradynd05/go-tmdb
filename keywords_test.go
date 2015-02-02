package tmdb

import (
	. "gopkg.in/check.v1"
)

const fightKeywordId int = 1721

func (s *TmdbSuite) TestGetKeywordInfo(c *C) {
	result, err := s.tmdb.GetKeywordInfo(fightKeywordId)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, fightKeywordId)
	c.Assert(result.Name, Equals, "fight")
}

func (s *TmdbSuite) TestGetKeywordMovies(c *C) {
	result, err := s.tmdb.GetKeywordMovies(fightKeywordId, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, fightKeywordId)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.TotalPages > 7, Equals, true)
	c.Assert(result.TotalResults > 150, Equals, true)

	var page2Options = make(map[string]string)
	page2Options["page"] = "2"
	page2Result, err := s.tmdb.GetKeywordMovies(fightKeywordId, page2Options)
	s.baseTest(&page2Result, err, c)
	c.Assert(page2Result.ID, Equals, fightKeywordId)
	c.Assert(page2Result.Page, Equals, 2)
	c.Assert(page2Result.TotalPages, Equals, result.TotalPages)
	c.Assert(page2Result.TotalResults, Equals, result.TotalResults)
}
