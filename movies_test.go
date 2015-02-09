package tmdb

import (
	. "gopkg.in/check.v1"
)

const darkKnightID int = 49026
const fightClubID int = 550
const takenThreeID int = 260346

func (s *TmdbSuite) TestGetMovieInfo(c *C) {
	result, err := s.tmdb.GetMovieInfo(fightClubID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Title, Equals, "Fight Club")
	c.Assert(result.ID, Equals, fightClubID)

	var options = make(map[string]string)
	options["append_to_response"] = "alternative_titles,credits,images,keywords,releases,videos,translations,similar,reviews,lists,changes,ratings"
	allResult, err := s.tmdb.GetMovieInfo(fightClubID, options)
	s.baseTest(&allResult, err, c)
	c.Assert(allResult.Title, Equals, "Fight Club")
	c.Assert(allResult.ID, Equals, fightClubID)
	c.Assert(allResult.AlternativeTitles, NotNil)
	c.Assert(allResult.Credits, NotNil)
	c.Assert(allResult.Images, NotNil)
	c.Assert(allResult.Keywords, NotNil)
	c.Assert(allResult.Releases, NotNil)
	c.Assert(allResult.Videos, NotNil)
	c.Assert(allResult.Translations, NotNil)
	c.Assert(allResult.Similar, NotNil)
	c.Assert(allResult.Reviews, NotNil)
	c.Assert(allResult.Lists, NotNil)
	c.Assert(allResult.Changes, NotNil)
}

func (s *TmdbSuite) TestGetMovieAccountStates(c *C) {
	result, err := s.tmdb.GetMovieAccountStates(fightClubID, session)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, fightClubID)
}

func (s *TmdbSuite) TestGetMovieAlternativeTitles(c *C) {
	result, err := s.tmdb.GetMovieAlternativeTitles(fightClubID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, fightClubID)
	c.Assert(result.Titles, Not(HasLen), 0)
	c.Assert(result.Titles[0].Iso3166_1, Equals, "PL")
}

func (s *TmdbSuite) TestGetMovieChanges(c *C) {
	result, err := s.tmdb.GetMovieChanges(takenThreeID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Changes, NotNil)
	allResultLength := len(result.Changes)

	var janOptions = make(map[string]string)
	janOptions["start_date"] = "2015-01-01"
	janOptions["end_date"] = "2015-01-14"
	janResult, err := s.tmdb.GetMovieChanges(takenThreeID, janOptions)
	s.baseTest(&janResult, err, c)
	c.Assert(janResult.Changes, NotNil)
	c.Assert(len(janResult.Changes) >= allResultLength, Equals, true)
}

func (s *TmdbSuite) TestGetMovieCredits(c *C) {
	result, err := s.tmdb.GetMovieCredits(fightClubID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, fightClubID)
	c.Assert(result.Cast, Not(HasLen), 0)
	c.Assert(result.Cast[0].CastID, Equals, 4)
	c.Assert(result.Cast[0].Character, Equals, "The Narrator")
	c.Assert(result.Cast[0].Name, Equals, "Edward Norton")
}

func (s *TmdbSuite) TestGetMovieImages(c *C) {
	result, err := s.tmdb.GetMovieImages(fightClubID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, fightClubID)
	c.Assert(result.Backdrops, Not(HasLen), 0)
	c.Assert(result.Posters, Not(HasLen), 0)
	allBackdropsLength := len(result.Backdrops)
	allPostersLength := len(result.Posters)

	var engOptions = make(map[string]string)
	engOptions["language"] = "en"
	engResult, err := s.tmdb.GetMovieImages(fightClubID, engOptions)
	s.baseTest(&engResult, err, c)
	c.Assert(engResult.ID, Equals, fightClubID)
	c.Assert(engResult.Backdrops, Not(HasLen), 0)
	c.Assert(engResult.Posters, Not(HasLen), 0)
	c.Assert(engResult.Backdrops[0].Iso639_1, Equals, "en")
	c.Assert(len(engResult.Backdrops) <= allBackdropsLength, Equals, true)
	c.Assert(len(engResult.Posters) <= allPostersLength, Equals, true)
}

func (s *TmdbSuite) TestGetMovieKeywords(c *C) {
	result, err := s.tmdb.GetMovieKeywords(fightClubID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, fightClubID)
	c.Assert(result.Keywords, Not(HasLen), 0)
	c.Assert(result.Keywords[0].ID, Equals, 825)
	c.Assert(result.Keywords[0].Name, Equals, "support group")
}

func (s *TmdbSuite) TestGetMovieLatest(c *C) {
	result, err := s.tmdb.GetMovieLatest()
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Not(Equals), fightClubID)
}

func (s *TmdbSuite) TestGetMovieLists(c *C) {
	result, err := s.tmdb.GetMovieLists(fightClubID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, fightClubID)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results[0].ID, Equals, "522effe419c2955e9922fcf3")
	c.Assert(result.Results[0].Iso639_1, Equals, "en")
	c.Assert(result.Results[0].Name, Equals, "IMDb Top 250")
	allResultLength := len(result.Results)

	var spanishOptions = make(map[string]string)
	spanishOptions["language"] = "es"
	spanishResult, err := s.tmdb.GetMovieLists(fightClubID, spanishOptions)
	s.baseTest(&spanishResult, err, c)
	c.Assert(spanishResult.ID, Equals, fightClubID)
	c.Assert(spanishResult.Page, Equals, 1)
	c.Assert(len(spanishResult.Results) < allResultLength, Equals, true)
}

func (s *TmdbSuite) TestGetMovieNowPlaying(c *C) {
	result, err := s.tmdb.GetMovieNowPlaying(nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)

	var page2Options = make(map[string]string)
	page2Options["page"] = "2"
	page2Result, err := s.tmdb.GetMovieNowPlaying(page2Options)
	s.baseTest(&page2Result, err, c)
	c.Assert(page2Result.Page, Equals, 2)
	c.Assert(page2Result.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetMoviePopular(c *C) {
	result, err := s.tmdb.GetMoviePopular(nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)

	var page2Options = make(map[string]string)
	page2Options["page"] = "2"
	page2Result, err := s.tmdb.GetMoviePopular(page2Options)
	s.baseTest(&page2Result, err, c)
	c.Assert(page2Result.Page, Equals, 2)
	c.Assert(page2Result.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetMovieReleases(c *C) {
	result, err := s.tmdb.GetMovieReleases(fightClubID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, fightClubID)
	c.Assert(result.Countries, Not(HasLen), 0)
	c.Assert(result.Countries[0].Iso3166_1, Equals, "US")
	c.Assert(result.Countries[0].Certification, Equals, "R")
	c.Assert(result.Countries[0].ReleaseDate, Equals, "1999-10-14")
}

func (s *TmdbSuite) TestGetMovieReviews(c *C) {
	result, err := s.tmdb.GetMovieReviews(darkKnightID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, darkKnightID)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)

	var engOptions = make(map[string]string)
	engOptions["language"] = "en"
	engResult, err := s.tmdb.GetMovieReviews(darkKnightID, engOptions)
	s.baseTest(&engResult, err, c)
	c.Assert(engResult.ID, Equals, darkKnightID)
	c.Assert(engResult.Page, Equals, 1)
}

func (s *TmdbSuite) TestGetMovieSimilar(c *C) {
	result, err := s.tmdb.GetMovieSimilar(fightClubID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.Results[0].ID, Equals, 807)
	c.Assert(result.Results[0].Title, Equals, "Se7en")

	var engOptions = make(map[string]string)
	engOptions["language"] = "en"
	engResult, err := s.tmdb.GetMovieSimilar(fightClubID, engOptions)
	s.baseTest(&engResult, err, c)
	c.Assert(engResult.Page, Equals, 1)
	c.Assert(engResult.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetMovieTopRated(c *C) {
	result, err := s.tmdb.GetMovieTopRated(nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)

	var page2Options = make(map[string]string)
	page2Options["page"] = "2"
	page2Result, err := s.tmdb.GetMovieTopRated(page2Options)
	s.baseTest(&page2Result, err, c)
	c.Assert(page2Result.Page, Equals, 2)
	c.Assert(page2Result.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetMovieTranslations(c *C) {
	result, err := s.tmdb.GetMovieTranslations(fightClubID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, fightClubID)
	c.Assert(result.Translations, Not(HasLen), 0)
	c.Assert(result.Translations[0].Iso639_1, Equals, "en")
	c.Assert(result.Translations[0].EnglishName, Equals, "English")
	c.Assert(result.Translations[0].Name, Equals, "English")
}

func (s *TmdbSuite) TestGetMovieUpcoming(c *C) {
	result, err := s.tmdb.GetMovieUpcoming(nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)

	var page2Options = make(map[string]string)
	page2Options["page"] = "2"
	page2Result, err := s.tmdb.GetMovieUpcoming(page2Options)
	s.baseTest(&page2Result, err, c)
	c.Assert(page2Result.Page, Equals, 2)
	c.Assert(page2Result.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetMovieVideos(c *C) {
	result, err := s.tmdb.GetMovieVideos(fightClubID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, fightClubID)
	c.Assert(result.Results, Not(HasLen), 0)
	allResultsLength := len(result.Results)

	var engOptions = make(map[string]string)
	engOptions["language"] = "en"
	engResult, err := s.tmdb.GetMovieVideos(fightClubID, engOptions)
	s.baseTest(&engResult, err, c)
	c.Assert(engResult.ID, Equals, fightClubID)
	c.Assert(engResult.Results, Not(HasLen), 0)
	c.Assert(engResult.Results[0].Iso639_1, Equals, "en")
	c.Assert(len(engResult.Results) <= allResultsLength, Equals, true)
}

func (s *TmdbSuite) TestSetMovieRating(c *C) {
	// result, err := s.tmdb.SetMovieRating(takenThreeID)
	// s.baseTest(&result, err, c)
	// TODO
}
