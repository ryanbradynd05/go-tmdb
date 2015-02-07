package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestGetTvEpisodeInfo(c *C) {
	result, err := s.tmdb.GetTvEpisodeInfo(gameOfThronesID, 1, 1, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Name, Equals, "Winter Is Coming")
	c.Assert(result.AirDate, Equals, "2011-04-17")
	c.Assert(result.ID, Equals, 63056)
	c.Assert(result.SeasonNumber, Equals, 1)
	c.Assert(result.EpisodeNumber, Equals, 1)
}
