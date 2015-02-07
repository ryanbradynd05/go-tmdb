package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestGetTvSeasonInfo(c *C) {
	result, err := s.tmdb.GetTvSeasonInfo(gameOfThronesID, 1, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Name, Equals, "Season 1")
	c.Assert(result.AirDate, Equals, "2011-04-17")
	c.Assert(result.Episodes, HasLen, 10)
	c.Assert(result.ID, Equals, 3624)
	c.Assert(result.SeasonNumber, Equals, 1)
}
