package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestGetMovieInfo(c *C) {
	result, err := s.tmdb.GetMovieInfo(550)
	s.baseTest(&result, err, c)
	c.Assert(result.Title, Equals, "Fight Club")
	c.Assert(result.ID, Equals, 550)
}

func (s *TmdbSuite) TestGetMovieAccountStates(c *C) {
	// result, err := s.tmdb.GetMovieAccountStates(260346)
	// s.baseTest(&result, err, c)
	// TODO
}

func (s *TmdbSuite) TestGetMovieAlternativeTitles(c *C) {
	result, err := s.tmdb.GetMovieAlternativeTitles(550)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, 550)
	c.Assert(result.Titles, Not(HasLen), 0)
	c.Assert(result.Titles[0].Iso3166_1, Equals, "PL")
}

func (s *TmdbSuite) TestGetMovieCredits(c *C) {
	result, err := s.tmdb.GetMovieCredits(550)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, 550)
	c.Assert(result.Cast, Not(HasLen), 0)
	c.Assert(result.Cast[0].CastID, Equals, 4)
	c.Assert(result.Cast[0].Character, Equals, "The Narrator")
	c.Assert(result.Cast[0].Name, Equals, "Edward Norton")
}

func (s *TmdbSuite) TestGetMovieImages(c *C) {
	result, err := s.tmdb.GetMovieImages(550)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, 550)
	c.Assert(result.Backdrops, Not(HasLen), 0)
	c.Assert(result.Backdrops[0].Height, Equals, 720)
	c.Assert(result.Backdrops[0].Iso639_1, Equals, "xx")
}

func (s *TmdbSuite) TestGetMovieKeywords(c *C) {
	result, err := s.tmdb.GetMovieKeywords(550)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, 550)
	c.Assert(result.Keywords, Not(HasLen), 0)
	c.Assert(result.Keywords[0].ID, Equals, 825)
	c.Assert(result.Keywords[0].Name, Equals, "support group")
}

func (s *TmdbSuite) TestGetMovieReleases(c *C) {
	result, err := s.tmdb.GetMovieReleases(550)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, 550)
	c.Assert(result.Countries, Not(HasLen), 0)
	c.Assert(result.Countries[0].Iso3166_1, Equals, "US")
	c.Assert(result.Countries[0].Certification, Equals, "R")
	c.Assert(result.Countries[0].ReleaseDate, Equals, "1999-10-14")
}

func (s *TmdbSuite) TestGetMovieVideos(c *C) {
	result, err := s.tmdb.GetMovieVideos(550)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, 550)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.Results[0].Iso639_1, Equals, "en")
	c.Assert(result.Results[0].Site, Equals, "YouTube")
	c.Assert(result.Results[0].Type, Equals, "Trailer")
}

func (s *TmdbSuite) TestGetMovieTranslations(c *C) {
	result, err := s.tmdb.GetMovieTranslations(550)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, 550)
	c.Assert(result.Translations, Not(HasLen), 0)
	c.Assert(result.Translations[0].Iso639_1, Equals, "en")
	c.Assert(result.Translations[0].EnglishName, Equals, "English")
	c.Assert(result.Translations[0].Name, Equals, "English")
}

func (s *TmdbSuite) TestGetSimilarMovies(c *C) {
	result, err := s.tmdb.GetSimilarMovies(550, 1)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.Results[0].ID, Equals, 807)
	c.Assert(result.Results[0].Title, Equals, "Se7en")
}

func (s *TmdbSuite) TestGetMovieReviews(c *C) {
	result, err := s.tmdb.GetMovieReviews(49026, 1)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetMovieLists(c *C) {
	result, err := s.tmdb.GetMovieLists(550, 1)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, 550)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results[0].ID, Equals, "522effe419c2955e9922fcf3")
	c.Assert(result.Results[0].Iso639_1, Equals, "en")
	c.Assert(result.Results[0].Name, Equals, "IMDb Top 250")
}

func (s *TmdbSuite) TestGetMovieChanges(c *C) {
	result, err := s.tmdb.GetMovieChanges(260346)
	s.baseTest(&result, err, c)
	c.Assert(result.Changes, Not(HasLen), 0)
}

func (s *TmdbSuite) TestSetMovieRating(c *C) {
	// result, err := s.tmdb.SetMovieRating(260346)
	// s.baseTest(&result, err, c)
	// TODO
}

func (s *TmdbSuite) TestGetLatestMovie(c *C) {
	result, err := s.tmdb.GetLatestMovie()
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Not(Equals), 550)
}

func (s *TmdbSuite) TestGetUpcomingMovies(c *C) {
	result, err := s.tmdb.GetUpcomingMovies()
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetNowPlayingMovies(c *C) {
	result, err := s.tmdb.GetNowPlayingMovies()
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetPopularMovies(c *C) {
	result, err := s.tmdb.GetPopularMovies()
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetTopRatedMovies(c *C) {
	result, err := s.tmdb.GetTopRatedMovies()
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
}
