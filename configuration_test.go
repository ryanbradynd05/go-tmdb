package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestGetConfiguration(c *C) {
	result, err := s.tmdb.GetConfiguration()
	s.baseTest(&result, err, c)
	c.Assert(result.Images.BaseURL, Equals, "http://image.tmdb.org/t/p/")
	c.Assert(result.Images.SecureBaseURL, Equals, "https://image.tmdb.org/t/p/")
	c.Assert(result.Images.BackdropSizes, HasLen, 4)
	c.Assert(result.ChangeKeys, HasLen, 53)
}
