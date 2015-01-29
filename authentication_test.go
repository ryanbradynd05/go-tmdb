package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestNewToken(c *C) {
	result, err := s.tmdb.NewToken()
	s.baseTest(&result, err, c)
	c.Assert(result.Success, Equals, true)
	c.Assert(result.RequestToken, HasLen, 40)
}
