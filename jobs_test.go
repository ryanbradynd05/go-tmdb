package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestGetJobList(c *C) {
	result, err := s.tmdb.GetJobList()
	s.baseTest(&result, err, c)
	c.Assert(len(result.Jobs), Equals, 12)
}
