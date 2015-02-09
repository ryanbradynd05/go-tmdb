package tmdb

import (
	. "gopkg.in/check.v1"
)

const gameOfThronesPilotID int = 63056

func (s *TmdbSuite) TestGetTvEpisodeInfo(c *C) {
	result, err := s.tmdb.GetTvEpisodeInfo(gameOfThronesID, 1, 1, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Name, Equals, "Winter Is Coming")
	c.Assert(result.AirDate, Equals, "2011-04-17")
	c.Assert(result.ID, Equals, gameOfThronesPilotID)
	c.Assert(result.SeasonNumber, Equals, 1)
	c.Assert(result.EpisodeNumber, Equals, 1)
}

func (s *TmdbSuite) TestGetTvEpisodeChanges(c *C) {
	result, err := s.tmdb.GetTvEpisodeChanges(gameOfThronesPilotID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Changes, NotNil)
	allResultLength := len(result.Changes)

	var janOptions = make(map[string]string)
	janOptions["start_date"] = "2015-01-01"
	janOptions["end_date"] = "2015-01-14"
	janResult, err := s.tmdb.GetTvEpisodeChanges(gameOfThronesPilotID, janOptions)
	s.baseTest(&janResult, err, c)
	c.Assert(janResult.Changes, NotNil)
	c.Assert(len(janResult.Changes) >= allResultLength, Equals, true)
}

func (s *TmdbSuite) TestGetTvEpisodeCredits(c *C) {
	result, err := s.tmdb.GetTvEpisodeCredits(gameOfThronesID, 1, 1)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, gameOfThronesPilotID)
	c.Assert(result.Cast, Not(HasLen), 0)
	c.Assert(result.Cast[0].Character, Equals, "Eddard Stark")
	c.Assert(result.Cast[0].Name, Equals, "Sean Bean")
}

func (s *TmdbSuite) TestGetTvEpisodeExternalIds(c *C) {
	result, err := s.tmdb.GetTvEpisodeExternalIds(gameOfThronesID, 1, 1, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, gameOfThronesPilotID)
	c.Assert(result.ImdbID, Equals, "tt1480055")
	c.Assert(result.TvdbID, Equals, 3254641)
	c.Assert(result.TvrageID, Equals, 1065008299)
}

func (s *TmdbSuite) TestGetTvEpisodeImages(c *C) {
	result, err := s.tmdb.GetTvEpisodeImages(gameOfThronesID, 1, 1)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, gameOfThronesPilotID)
	c.Assert(result.Stills, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetTvEpisodeVideos(c *C) {
	result, err := s.tmdb.GetTvEpisodeVideos(gameOfThronesID, 1, 1, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, gameOfThronesPilotID)
	c.Assert(result.Results, NotNil)
}
