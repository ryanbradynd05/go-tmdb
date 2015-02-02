package tmdb

import (
	. "gopkg.in/check.v1"
)

const darkKnightReviewID string = "5013bc76760ee372cb00253e"

func (s *TmdbSuite) TestGetReviewInfo(c *C) {
	result, err := s.tmdb.GetReviewInfo(darkKnightReviewID)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, darkKnightReviewID)
	c.Assert(result.Author, Equals, "Chris")
	c.Assert(result.Iso639_1, Equals, "en")
	c.Assert(result.MediaID, Equals, 49026)
	c.Assert(result.MediaTitle, Equals, "The Dark Knight Rises")
	c.Assert(result.MediaType, Equals, "Movie")
	c.Assert(result.URL, Equals, "http://j.mp/P18dg1")
}
