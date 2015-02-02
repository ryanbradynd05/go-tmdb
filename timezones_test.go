package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestGetTimezonesList(c *C) {
	result, err := s.tmdb.GetTimezonesList()
	s.baseTest(&result, err, c)
}
