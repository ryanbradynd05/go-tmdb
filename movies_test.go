package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestMovieInfo(c *C) {
	result, err := s.tmdb.MovieInfo(550)
	s.baseTest(&result, err, c)
	c.Assert(result.Title, Equals, "Fight Club")
	c.Assert(result.ID, Equals, 550)
}

func (s *TmdbSuite) TestMovieAccountStates(c *C) {
	// result, err := s.tmdb.MovieAccountStates(260346)
	// s.baseTest(&result, err, c)
	// TODO
}

func (s *TmdbSuite) TestMovieAlternativeTitles(c *C) {
	result, err := s.tmdb.MovieAlternativeTitles(550)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, 550)
	c.Assert(result.Titles, Not(HasLen), 0)
	c.Assert(result.Titles[0].Iso3166_1, Equals, "PL")
}

func (s *TmdbSuite) TestMovieCredits(c *C) {
	result, err := s.tmdb.MovieCredits(550)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, 550)
	c.Assert(result.Cast, Not(HasLen), 0)
	c.Assert(result.Cast[0].CastID, Equals, 4)
	c.Assert(result.Cast[0].Character, Equals, "The Narrator")
	c.Assert(result.Cast[0].Name, Equals, "Edward Norton")
}

func (s *TmdbSuite) TestMovieImages(c *C) {
	result, err := s.tmdb.MovieImages(550)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, 550)
	c.Assert(result.Backdrops, Not(HasLen), 0)
	c.Assert(result.Backdrops[0].Height, Equals, 720)
	c.Assert(result.Backdrops[0].Iso639_1, Equals, "xx")
}

func (s *TmdbSuite) TestMovieKeywords(c *C) {
	result, err := s.tmdb.MovieKeywords(550)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, 550)
	c.Assert(result.Keywords, Not(HasLen), 0)
	c.Assert(result.Keywords[0].ID, Equals, 825)
	c.Assert(result.Keywords[0].Name, Equals, "support group")
}

func (s *TmdbSuite) TestMovieReleases(c *C) {
	result, err := s.tmdb.MovieReleases(550)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, 550)
	c.Assert(result.Countries, Not(HasLen), 0)
	c.Assert(result.Countries[0].Iso3166_1, Equals, "US")
	c.Assert(result.Countries[0].Certification, Equals, "R")
	c.Assert(result.Countries[0].ReleaseDate, Equals, "1999-10-14")
}

func (s *TmdbSuite) TestMovieVideos(c *C) {
	result, err := s.tmdb.MovieVideos(550)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, 550)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.Results[0].Iso639_1, Equals, "en")
	c.Assert(result.Results[0].Site, Equals, "YouTube")
	c.Assert(result.Results[0].Type, Equals, "Trailer")
}

func (s *TmdbSuite) TestMovieTranslations(c *C) {
	result, err := s.tmdb.MovieTranslations(550)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, 550)
	c.Assert(result.Translations, Not(HasLen), 0)
	c.Assert(result.Translations[0].Iso639_1, Equals, "en")
	c.Assert(result.Translations[0].EnglishName, Equals, "English")
	c.Assert(result.Translations[0].Name, Equals, "English")
}

func (s *TmdbSuite) TestMovieSimilar(c *C) {
	result, err := s.tmdb.MovieSimilar(550, 1)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.Results[0].ID, Equals, 807)
	c.Assert(result.Results[0].Title, Equals, "Se7en")
}

func (s *TmdbSuite) TestMovieReviews(c *C) {
	result, err := s.tmdb.MovieReviews(49026, 1)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestMovieLists(c *C) {
	result, err := s.tmdb.MovieLists(550, 1)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, 550)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results[0].ID, Equals, "522effe419c2955e9922fcf3")
	c.Assert(result.Results[0].Iso639_1, Equals, "en")
	c.Assert(result.Results[0].Name, Equals, "IMDb Top 250")
}

func (s *TmdbSuite) TestMovieChanges(c *C) {
	result, err := s.tmdb.MovieChanges(260346)
	s.baseTest(&result, err, c)
	c.Assert(result.Changes, Not(HasLen), 0)
}

func (s *TmdbSuite) TestMovieRating(c *C) {
	// result, err := s.tmdb.MovieRating(260346)
	// s.baseTest(&result, err, c)
	// TODO
}

func (s *TmdbSuite) TestMovieLatest(c *C) {
	result, err := s.tmdb.MovieLatest()
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Not(Equals), 550)
}

func (s *TmdbSuite) TestMovieUpcoming(c *C) {
	result, err := s.tmdb.MovieUpcoming()
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestMovieNowPlaying(c *C) {
	result, err := s.tmdb.NowPlaying()
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestMoviePopular(c *C) {
	result, err := s.tmdb.Popular()
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
}

func (s *TmdbSuite) TestMovieTopRated(c *C) {
	result, err := s.tmdb.TopRated()
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
}
