package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestConfiguration(c *C) {
	result, err := s.tmdb.Configuration()
	s.baseTest(&result, err, c)
	c.Assert(result.Images.BaseURL, Equals, "http://image.tmdb.org/t/p/")
	c.Assert(result.Images.SecureBaseURL, Equals, "https://image.tmdb.org/t/p/")
	c.Assert(len(result.Images.BackdropSizes), Equals, 4)
	c.Assert(len(result.ChangeKeys), Equals, 53)
}
