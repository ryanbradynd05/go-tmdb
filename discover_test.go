package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestDiscoverMovie(c *C) {
	defaultResults, err := s.tmdb.DiscoverMovie(nil)
	s.baseTest(&defaultResults, err, c)
	c.Assert(defaultResults.Page, Equals, 1)
	c.Assert(defaultResults.TotalPages, Not(Equals), 0)
	c.Assert(defaultResults.TotalResults, Not(Equals), 0)
	c.Assert(defaultResults.Results, NotNil)
	c.Assert(defaultResults.Results, Not(HasLen), 0)

	var tomHanksOptions = make(map[string]string)
	tomHanksOptions["with_cast"] = "31"
	tomHanksMovies, err := s.tmdb.DiscoverMovie(tomHanksOptions)
	s.baseTest(&tomHanksMovies, err, c)
	c.Assert(tomHanksMovies.Page, Equals, 1)
	c.Assert(tomHanksMovies.TotalPages > 3, Equals, true)
	c.Assert(tomHanksMovies.TotalResults > 75, Equals, true)
	c.Assert(tomHanksMovies.Results, NotNil)
	c.Assert(tomHanksMovies.Results, Not(HasLen), 0)

	var tomHanksTimAllenOptions = make(map[string]string)
	tomHanksTimAllenOptions["with_cast"] = "31,12898"
	tomHanksTimAllenMovies, err := s.tmdb.DiscoverMovie(tomHanksTimAllenOptions)
	s.baseTest(&tomHanksTimAllenMovies, err, c)
	c.Assert(tomHanksTimAllenMovies.Page, Equals, 1)
	c.Assert(tomHanksTimAllenMovies.TotalPages, Equals, 1)
	c.Assert(tomHanksTimAllenMovies.TotalResults, Equals, 8)
	c.Assert(tomHanksTimAllenMovies.Results, NotNil)
	c.Assert(tomHanksTimAllenMovies.Results, HasLen, 8)

	var goodMcConaugheyOptions = make(map[string]string)
	goodMcConaugheyOptions["with_cast"] = "10297"
	goodMcConaugheyOptions["vote_average.gte"] = "7.5"
	goodMcConaugheyMovies, err := s.tmdb.DiscoverMovie(goodMcConaugheyOptions)
	s.baseTest(&goodMcConaugheyMovies, err, c)
	c.Assert(goodMcConaugheyMovies.Page, Equals, 1)
	c.Assert(goodMcConaugheyMovies.TotalPages, Equals, 1)
	c.Assert(goodMcConaugheyMovies.TotalResults, Equals, 3)
	c.Assert(goodMcConaugheyMovies.Results, NotNil)
	c.Assert(goodMcConaugheyMovies.Results, HasLen, 3)

	var horrible2010Options = make(map[string]string)
	horrible2010Options["vote_average.lte"] = "1.0"
	horrible2010Options["vote_average.gte"] = "0.1"
	horrible2010Options["year"] = "2010"
	horrible2010Movies, err := s.tmdb.DiscoverMovie(horrible2010Options)
	s.baseTest(&horrible2010Movies, err, c)
	c.Assert(horrible2010Movies.Page, Equals, 1)
	c.Assert(horrible2010Movies.TotalPages, Equals, 2)
	c.Assert(horrible2010Movies.TotalResults, Equals, 24)
	c.Assert(horrible2010Movies.Results, NotNil)
	c.Assert(horrible2010Movies.Results, HasLen, 20)

	var horrible2010Page2Options = make(map[string]string)
	horrible2010Page2Options["vote_average.lte"] = "1.0"
	horrible2010Page2Options["vote_average.gte"] = "0.1"
	horrible2010Page2Options["year"] = "2010"
	horrible2010Page2Options["page"] = "2"
	horrible2010Page2Movies, err := s.tmdb.DiscoverMovie(horrible2010Page2Options)
	s.baseTest(&horrible2010Page2Movies, err, c)
	c.Assert(horrible2010Page2Movies.Page, Equals, 2)
	c.Assert(horrible2010Page2Movies.TotalPages, Equals, 2)
	c.Assert(horrible2010Page2Movies.TotalResults, Equals, 24)
	c.Assert(horrible2010Page2Movies.Results, NotNil)
	c.Assert(horrible2010Page2Movies.Results, HasLen, 4)
}

func (s *TmdbSuite) TestDiscoverTV(c *C) {
	defaultResults, err := s.tmdb.DiscoverTV(nil)
	s.baseTest(&defaultResults, err, c)
	c.Assert(defaultResults.Page, Equals, 1)
	c.Assert(defaultResults.TotalPages, Not(Equals), 0)
	c.Assert(defaultResults.TotalResults, Not(Equals), 0)
	c.Assert(defaultResults.Results, NotNil)
	c.Assert(defaultResults.Results, Not(HasLen), 0)

	var fxOptions = make(map[string]string)
	fxOptions["with_networks"] = "88"
	fxTv, err := s.tmdb.DiscoverTV(fxOptions)
	s.baseTest(&fxTv, err, c)
	c.Assert(fxTv.Page, Equals, 1)
	c.Assert(fxTv.TotalPages >= 3, Equals, true)
	c.Assert(fxTv.TotalResults >= 45, Equals, true)
	c.Assert(fxTv.Results, NotNil)
	c.Assert(fxTv.Results, Not(HasLen), 0)

	var fxAmcOptions = make(map[string]string)
	fxAmcOptions["with_networks"] = "88|174"
	fxAmcTv, err := s.tmdb.DiscoverTV(fxAmcOptions)
	s.baseTest(&fxAmcTv, err, c)
	c.Assert(fxAmcTv.Page, Equals, 1)
	c.Assert(fxAmcTv.TotalPages >= 4, Equals, true)
	c.Assert(fxAmcTv.TotalResults >= 75, Equals, true)
	c.Assert(fxAmcTv.Results, NotNil)
	c.Assert(fxAmcTv.Results, Not(HasLen), 0)

	var ratedFxOptions = make(map[string]string)
	ratedFxOptions["with_networks"] = "88"
	ratedFxOptions["vote_count.gte"] = "5"
	ratedFxTv, err := s.tmdb.DiscoverTV(ratedFxOptions)
	s.baseTest(&ratedFxTv, err, c)
	c.Assert(ratedFxTv.Page, Equals, 1)
	c.Assert(ratedFxTv.TotalPages >= 1, Equals, true)
	c.Assert(ratedFxTv.TotalResults >= 10, Equals, true)
	c.Assert(ratedFxTv.Results, NotNil)
	c.Assert(ratedFxTv.Results, Not(HasLen), 0)

	var rated2010Options = make(map[string]string)
	rated2010Options["vote_count.gte"] = "5"
	rated2010Options["first_air_date_year"] = "2010"
	rated2010Tv, err := s.tmdb.DiscoverTV(rated2010Options)
	s.baseTest(&rated2010Tv, err, c)
	c.Assert(rated2010Tv.Page, Equals, 1)
	c.Assert(rated2010Tv.TotalPages >= 1, Equals, true)
	c.Assert(rated2010Tv.TotalResults >= 18, Equals, true)
	c.Assert(rated2010Tv.Results, NotNil)
	c.Assert(rated2010Tv.Results, Not(HasLen), 0)
}
