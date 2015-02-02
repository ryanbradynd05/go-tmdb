package tmdb

import (
	. "gopkg.in/check.v1"
)

const oscarWinnerListID string = "509ec17b19c2950a0600050d"

func (s *TmdbSuite) TestGetListInfo(c *C) {
	result, err := s.tmdb.GetListInfo(oscarWinnerListID)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, oscarWinnerListID)
	c.Assert(result.CreatedBy, Equals, "Travis Bell")
	c.Assert(result.Iso639_1, Equals, "en")
	c.Assert(result.ItemCount, Equals, 86)
	c.Assert(result.Name, Equals, "Best Picture Winners - The Academy Awards")
}

func (s *TmdbSuite) TestGetListItemStatus(c *C) {
	fightClubResult, err := s.tmdb.GetListItemStatus(oscarWinnerListID, 550)
	s.baseTest(&fightClubResult, err, c)
	c.Assert(fightClubResult.ID, Equals, oscarWinnerListID)
	c.Assert(fightClubResult.ItemPresent, Equals, false)

	argoResult, err := s.tmdb.GetListItemStatus(oscarWinnerListID, 68734)
	s.baseTest(&argoResult, err, c)
	c.Assert(argoResult.ID, Equals, oscarWinnerListID)
	c.Assert(argoResult.ItemPresent, Equals, true)
}
