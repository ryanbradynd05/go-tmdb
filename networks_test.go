package tmdb

import (
	. "gopkg.in/check.v1"
)

const hboID int = 49

func (s *TmdbSuite) TestGetNetworkInfo(c *C) {
	result, err := s.tmdb.GetNetworkInfo(hboID)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, hboID)
	c.Assert(result.Name, Equals, "HBO")
}
