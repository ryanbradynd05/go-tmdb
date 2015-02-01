package tmdb

import (
	. "gopkg.in/check.v1"
)

const columbiaID int = 5

func (s *TmdbSuite) TestGetCompanyInfo(c *C) {
	company, err := s.tmdb.GetCompanyInfo(columbiaID, nil)
	s.baseTest(&company, err, c)
	c.Assert(company.ID, Equals, columbiaID)
	c.Assert(company.Name, Equals, "Columbia Pictures")
	c.Assert(company.Movies, Not(NotNil))

	var options = make(map[string]string)
	options["append_to_response"] = "movies"
	companyMovies, err := s.tmdb.GetCompanyInfo(columbiaID, options)
	s.baseTest(&companyMovies, err, c)
	c.Assert(companyMovies.ID, Equals, columbiaID)
	c.Assert(companyMovies.Name, Equals, "Columbia Pictures")
	c.Assert(companyMovies.Movies, NotNil)
}

func (s *TmdbSuite) TestGetCompanyMovies(c *C) {
	movies, err := s.tmdb.GetCompanyMovies(columbiaID, nil)
	s.baseTest(&movies, err, c)
	c.Assert(movies.ID, Equals, columbiaID)
	c.Assert(movies.Page, Equals, 1)
	c.Assert(movies.TotalPages > 20, Equals, true)
	c.Assert(movies.TotalResults > 400, Equals, true)
	c.Assert(movies.Results, NotNil)
	c.Assert(movies.Results, Not(HasLen), 0)

	var options = make(map[string]string)
	options["page"] = "2"
	moviesPage2, err := s.tmdb.GetCompanyMovies(columbiaID, options)
	s.baseTest(&moviesPage2, err, c)
	c.Assert(moviesPage2.ID, Equals, columbiaID)
	c.Assert(moviesPage2.Page, Equals, 2)
	c.Assert(moviesPage2.TotalPages > 20, Equals, true)
	c.Assert(moviesPage2.TotalResults > 400, Equals, true)
	c.Assert(moviesPage2.Results, NotNil)
	c.Assert(moviesPage2.Results, Not(HasLen), 0)
}
