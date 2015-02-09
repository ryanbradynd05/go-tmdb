package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestGetFind(c *C) {
	movieResult, err := s.tmdb.GetFind("tt0137523", "imdb_id", nil)
	s.baseTest(&movieResult, err, c)
	c.Assert(movieResult.MovieResults, NotNil)
	c.Assert(movieResult.MovieResults, HasLen, 1)
	c.Assert(movieResult.MovieResults[0].ID, Equals, fightClubID)
	c.Assert(movieResult.MovieResults[0].Title, Equals, "Fight Club")

	personResult, err := s.tmdb.GetFind("nm0000093", "imdb_id", nil)
	s.baseTest(&personResult, err, c)
	c.Assert(personResult.PersonResults, NotNil)
	c.Assert(personResult.PersonResults, HasLen, 1)
	c.Assert(personResult.PersonResults[0].ID, Equals, bradPittID)
	c.Assert(personResult.PersonResults[0].Name, Equals, "Brad Pitt")

	tvResult, err := s.tmdb.GetFind("tt0944947", "imdb_id", nil)
	s.baseTest(&tvResult, err, c)
	c.Assert(tvResult.TvResults, NotNil)
	c.Assert(tvResult.TvResults, HasLen, 1)
	c.Assert(tvResult.TvResults[0].ID, Equals, gameOfThronesID)
	c.Assert(tvResult.TvResults[0].Name, Equals, "Game of Thrones")

	tvEpisodeResult, err := s.tmdb.GetFind("tt1480055", "imdb_id", nil)
	s.baseTest(&tvEpisodeResult, err, c)
	c.Assert(tvEpisodeResult.TvEpisodeResults, NotNil)
	c.Assert(tvEpisodeResult.TvEpisodeResults, HasLen, 1)
	c.Assert(tvEpisodeResult.TvEpisodeResults[0].ID, Equals, gameOfThronesPilotID)
	c.Assert(tvEpisodeResult.TvEpisodeResults[0].Name, Equals, "Winter Is Coming")
	c.Assert(tvEpisodeResult.TvEpisodeResults[0].ShowID, Equals, gameOfThronesID)

	tvSeasonResult, err := s.tmdb.GetFind("364731", "tvdb_id", nil)
	s.baseTest(&tvSeasonResult, err, c)
	c.Assert(tvSeasonResult.TvSeasonResults, NotNil)
	c.Assert(tvSeasonResult.TvSeasonResults, HasLen, 1)
	c.Assert(tvSeasonResult.TvSeasonResults[0].ID, Equals, gameOfThronesFirstSeasonID)
	c.Assert(tvSeasonResult.TvSeasonResults[0].Name, Equals, "Season 1")
	c.Assert(tvSeasonResult.TvSeasonResults[0].ShowID, Equals, gameOfThronesID)
}
