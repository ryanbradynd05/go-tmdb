package tmdb

import (
	. "gopkg.in/check.v1"
)

const gameOfThronesID int = 1399
const seinfeldID int = 1400

func (s *TmdbSuite) TestGetTvInfo(c *C) {
	result, err := s.tmdb.GetTvInfo(gameOfThronesID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Name, Equals, "Game of Thrones")
	c.Assert(result.ID, Equals, gameOfThronesID)
	c.Assert(result.FirstAirDate, Equals, "2011-04-17")
	c.Assert(result.Homepage, Equals, "http://www.hbo.com/game-of-thrones/index.html")
	c.Assert(result.Networks[0].ID, Equals, hboID)
	c.Assert(result.Networks[0].Name, Equals, "HBO")
	c.Assert(result.Status, Equals, "Returning Series")
	c.Assert(result.Type, Equals, "Scripted")
	c.Assert(len(result.Seasons), Equals, result.NumberOfSeasons+1)

	var options = make(map[string]string)
	options["append_to_response"] = "account_states,alternative_titles,changes,credits,images,keywords,similar,translations,videos"
	allResult, err := s.tmdb.GetTvInfo(gameOfThronesID, options)
	s.baseTest(&allResult, err, c)
	c.Assert(allResult.ID, Equals, gameOfThronesID)
	c.Assert(allResult.Name, Equals, "Game of Thrones")
	c.Assert(allResult.AlternativeTitles, NotNil)
	c.Assert(allResult.Changes, NotNil)
	c.Assert(allResult.Credits, NotNil)
	c.Assert(allResult.Images, NotNil)
	c.Assert(allResult.Keywords, NotNil)
	c.Assert(allResult.Similar, NotNil)
	c.Assert(allResult.Translations, NotNil)
	c.Assert(allResult.Videos, NotNil)
}

func (s *TmdbSuite) TestGetTvAccountStates(c *C) {
	// result, err := s.tmdb.GetTvAccountStates(gameOfThronesID)
	// s.baseTest(&result, err, c)
	// TODO
}

func (s *TmdbSuite) TestGetTvAiringToday(c *C) {
	result, err := s.tmdb.GetTvAiringToday(nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)

	var pacificOptions = make(map[string]string)
	pacificOptions["timezone"] = "America/Los_Angeles"
	pacificResult, err := s.tmdb.GetTvAiringToday(pacificOptions)
	s.baseTest(&pacificResult, err, c)
	c.Assert(pacificResult.Page, Equals, 1)
	c.Assert(pacificResult.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetTvAlternativeTitles(c *C) {
	result, err := s.tmdb.GetTvAlternativeTitles(gameOfThronesID)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, gameOfThronesID)
	c.Assert(result.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetTvChanges(c *C) {
	result, err := s.tmdb.GetTvChanges(gameOfThronesID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Changes, NotNil)
	allResultLength := len(result.Changes)

	var janOptions = make(map[string]string)
	janOptions["start_date"] = "2015-01-01"
	janOptions["end_date"] = "2015-01-14"
	janResult, err := s.tmdb.GetTvChanges(gameOfThronesID, janOptions)
	s.baseTest(&janResult, err, c)
	c.Assert(janResult.Changes, NotNil)
	c.Assert(len(janResult.Changes) >= allResultLength, Equals, true)
}

func (s *TmdbSuite) TestGetTvCredits(c *C) {
	result, err := s.tmdb.GetTvCredits(gameOfThronesID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, gameOfThronesID)
	c.Assert(result.Cast, Not(HasLen), 0)
	c.Assert(result.Cast[0].Character, Equals, "Tyrion Lannister")
	c.Assert(result.Cast[0].Name, Equals, "Peter Dinklage")
}

func (s *TmdbSuite) TestGetTvImages(c *C) {
	result, err := s.tmdb.GetTvImages(gameOfThronesID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, gameOfThronesID)
	c.Assert(result.Backdrops, Not(HasLen), 0)
	c.Assert(result.Posters, Not(HasLen), 0)
	allBackdropsLength := len(result.Backdrops)
	allPostersLength := len(result.Posters)

	var engOptions = make(map[string]string)
	engOptions["language"] = "en"
	engResult, err := s.tmdb.GetTvImages(gameOfThronesID, engOptions)
	s.baseTest(&engResult, err, c)
	c.Assert(engResult.ID, Equals, gameOfThronesID)
	c.Assert(engResult.Backdrops, Not(HasLen), 0)
	c.Assert(engResult.Posters, Not(HasLen), 0)
	c.Assert(engResult.Backdrops[0].Iso639_1, Equals, "en")
	c.Assert(len(engResult.Backdrops) <= allBackdropsLength, Equals, true)
	c.Assert(len(engResult.Posters) <= allPostersLength, Equals, true)
}

func (s *TmdbSuite) TestGetTvKeywords(c *C) {
	result, err := s.tmdb.GetTvKeywords(gameOfThronesID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, gameOfThronesID)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.Results[0].ID, Equals, 6091)
	c.Assert(result.Results[0].Name, Equals, "war")
}

func (s *TmdbSuite) TestGetTvLatest(c *C) {
	result, err := s.tmdb.GetTvLatest()
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Not(Equals), seinfeldID)
}

func (s *TmdbSuite) TestGetTvOnTheAir(c *C) {
	result, err := s.tmdb.GetTvOnTheAir(nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)

	var page2Options = make(map[string]string)
	page2Options["page"] = "2"
	page2Result, err := s.tmdb.GetTvOnTheAir(page2Options)
	s.baseTest(&page2Result, err, c)
	c.Assert(page2Result.Page, Equals, 2)
	c.Assert(page2Result.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetTvPopular(c *C) {
	result, err := s.tmdb.GetTvPopular(nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)

	var page2Options = make(map[string]string)
	page2Options["page"] = "2"
	page2Result, err := s.tmdb.GetTvPopular(page2Options)
	s.baseTest(&page2Result, err, c)
	c.Assert(page2Result.Page, Equals, 2)
	c.Assert(page2Result.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetTvSimilar(c *C) {
	result, err := s.tmdb.GetTvSimilar(gameOfThronesID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.Results[0].ID, Equals, 44217)
	c.Assert(result.Results[0].Name, Equals, "Vikings")

	var engOptions = make(map[string]string)
	engOptions["language"] = "en"
	engResult, err := s.tmdb.GetTvSimilar(gameOfThronesID, engOptions)
	s.baseTest(&engResult, err, c)
	c.Assert(engResult.Page, Equals, 1)
	c.Assert(engResult.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetTvTopRated(c *C) {
	result, err := s.tmdb.GetTvTopRated(nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)

	var page2Options = make(map[string]string)
	page2Options["page"] = "2"
	page2Result, err := s.tmdb.GetTvTopRated(page2Options)
	s.baseTest(&page2Result, err, c)
	c.Assert(page2Result.Page, Equals, 2)
	c.Assert(page2Result.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetTvTranslations(c *C) {
	result, err := s.tmdb.GetTvTranslations(gameOfThronesID)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, gameOfThronesID)
	c.Assert(result.Translations, Not(HasLen), 0)
	c.Assert(result.Translations[0].Iso639_1, Equals, "en")
	c.Assert(result.Translations[0].EnglishName, Equals, "English")
	c.Assert(result.Translations[0].Name, Equals, "English")
}

func (s *TmdbSuite) TestGetTvVideos(c *C) {
	result, err := s.tmdb.GetTvVideos(gameOfThronesID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, gameOfThronesID)
	c.Assert(result.Results, Not(HasLen), 0)
	allResultsLength := len(result.Results)

	var engOptions = make(map[string]string)
	engOptions["language"] = "en"
	engResult, err := s.tmdb.GetTvVideos(gameOfThronesID, engOptions)
	s.baseTest(&engResult, err, c)
	c.Assert(engResult.ID, Equals, gameOfThronesID)
	c.Assert(engResult.Results, Not(HasLen), 0)
	c.Assert(engResult.Results[0].Iso639_1, Equals, "en")
	c.Assert(len(engResult.Results) <= allResultsLength, Equals, true)
}
