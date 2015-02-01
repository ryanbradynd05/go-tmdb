package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestGetMovieGenres(c *C) {
	movieGenres, err := s.tmdb.GetMovieGenres(nil)
	s.baseTest(&movieGenres, err, c)
	c.Assert(movieGenres.Genres, NotNil)
	c.Assert(movieGenres.Genres, Not(HasLen), 0)

	var movieOptions = make(map[string]string)
	movieOptions["language"] = "es"
	spanishMovieGenres, err := s.tmdb.GetMovieGenres(movieOptions)
	s.baseTest(&spanishMovieGenres, err, c)
	c.Assert(spanishMovieGenres.Genres, NotNil)
	c.Assert(spanishMovieGenres.Genres, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetTvGenres(c *C) {
	tvGenres, err := s.tmdb.GetTvGenres(nil)
	s.baseTest(&tvGenres, err, c)
	c.Assert(tvGenres.Genres, NotNil)
	c.Assert(tvGenres.Genres, Not(HasLen), 0)

	var tvOptions = make(map[string]string)
	tvOptions["language"] = "es"
	spanishTvGenres, err := s.tmdb.GetTvGenres(tvOptions)
	s.baseTest(&spanishTvGenres, err, c)
	c.Assert(spanishTvGenres.Genres, NotNil)
	c.Assert(spanishTvGenres.Genres, Not(HasLen), 0)
}
