package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestGetNewToken(c *C) {
	result, err := s.tmdb.GetNewToken()
	s.baseTest(&result, err, c)
	c.Assert(result.Success, Equals, true)
	c.Assert(result.RequestToken, HasLen, 40)
}
