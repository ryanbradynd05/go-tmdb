package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestMovieInfo(c *C) {
	result, err := s.tmdb.MovieInfo(550)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestMovieAlternativeTitles(c *C) {
	result, err := s.tmdb.MovieAlternativeTitles(550)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestMovieCredits(c *C) {
	result, err := s.tmdb.MovieCredits(550)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestMovieImages(c *C) {
	result, err := s.tmdb.MovieImages(550)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestMovieKeywords(c *C) {
	result, err := s.tmdb.MovieKeywords(550)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestMovieReleases(c *C) {
	result, err := s.tmdb.MovieReleases(550)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestMovieVideos(c *C) {
	result, err := s.tmdb.MovieVideos(550)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestMovieTranslations(c *C) {
	result, err := s.tmdb.MovieTranslations(550)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestMovieSimilar(c *C) {
	result, err := s.tmdb.MovieSimilar(550, 1)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestMovieReviews(c *C) {
	result, err := s.tmdb.MovieReviews(49026, 1)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestMovieLists(c *C) {
	result, err := s.tmdb.MovieLists(550, 1)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestMovieChanges(c *C) {
	result, err := s.tmdb.MovieChanges(260346)
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestMovieLatest(c *C) {
	result, err := s.tmdb.MovieLatest()
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestMovieUpcoming(c *C) {
	result, err := s.tmdb.MovieUpcoming()
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestMovieNowPlaying(c *C) {
	result, err := s.tmdb.NowPlaying()
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestMoviePopular(c *C) {
	result, err := s.tmdb.Popular()
	s.baseTest(&result, err, c)
}

func (s *TmdbSuite) TestMovieTopRated(c *C) {
	result, err := s.tmdb.TopRated()
	s.baseTest(&result, err, c)
}
