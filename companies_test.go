package tmdb

import (
	. "gopkg.in/check.v1"
)

const columbiaID int = 5

func (s *TmdbSuite) TestGetCompanyInfo(c *C) {
	company, err := s.tmdb.GetCompanyInfo(columbiaID)
	s.baseTest(&company, err, c)
	c.Assert(company.ID, Equals, columbiaID)
	c.Assert(company.Name, Equals, "Columbia Pictures")
}

func (s *TmdbSuite) TestGetCompanyMovies(c *C) {
	movies, err := s.tmdb.GetCompanyMovies(columbiaID)
	s.baseTest(&movies, err, c)
	c.Assert(movies.ID, Equals, columbiaID)
	c.Assert(movies.Page, Equals, 1)
	c.Assert(movies.Results, NotNil)
	c.Assert(movies.Results, Not(HasLen), 0)
}
