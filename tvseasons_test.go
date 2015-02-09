package tmdb

import (
	. "gopkg.in/check.v1"
)

const gameOfThronesFirstSeasonID int = 3624

func (s *TmdbSuite) TestGetTvSeasonInfo(c *C) {
	result, err := s.tmdb.GetTvSeasonInfo(gameOfThronesID, 1, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Name, Equals, "Season 1")
	c.Assert(result.AirDate, Equals, "2011-04-17")
	c.Assert(result.Episodes, HasLen, 10)
	c.Assert(result.ID, Equals, 3624)
	c.Assert(result.SeasonNumber, Equals, 1)
}

func (s *TmdbSuite) TestGetTvSeasonChanges(c *C) {
	result, err := s.tmdb.GetTvSeasonChanges(gameOfThronesFirstSeasonID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Changes, NotNil)
	allResultLength := len(result.Changes)

	var janOptions = make(map[string]string)
	janOptions["start_date"] = "2015-01-01"
	janOptions["end_date"] = "2015-01-14"
	janResult, err := s.tmdb.GetTvSeasonChanges(gameOfThronesFirstSeasonID, janOptions)
	s.baseTest(&janResult, err, c)
	c.Assert(janResult.Changes, NotNil)
	c.Assert(len(janResult.Changes) >= allResultLength, Equals, true)
}

func (s *TmdbSuite) TestGetTvSeasonCredits(c *C) {
	result, err := s.tmdb.GetTvSeasonCredits(gameOfThronesID, 1)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, gameOfThronesFirstSeasonID)
	c.Assert(result.Cast, Not(HasLen), 0)
	c.Assert(result.Cast[0].Character, Equals, "Eddard Stark")
	c.Assert(result.Cast[0].Name, Equals, "Sean Bean")
}

func (s *TmdbSuite) TestGetTvSeasonExternalIds(c *C) {
	result, err := s.tmdb.GetTvSeasonExternalIds(gameOfThronesID, 1, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, gameOfThronesFirstSeasonID)
	c.Assert(result.TvdbID, Equals, 364731)
}

func (s *TmdbSuite) TestGetTvSeasonImages(c *C) {
	result, err := s.tmdb.GetTvSeasonImages(gameOfThronesID, 1, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, gameOfThronesFirstSeasonID)
	c.Assert(result.Posters, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetTvSeasonVideos(c *C) {
	result, err := s.tmdb.GetTvSeasonVideos(gameOfThronesID, 1, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, gameOfThronesFirstSeasonID)
	c.Assert(result.Results, NotNil)
}
