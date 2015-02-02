package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestSearchCollection(c *C) {
	result, err := s.tmdb.SearchCollection("Avengers", nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.TotalResults, Not(Equals), 0)
	c.Assert(result.Results[0].ID, Equals, 86311)
	c.Assert(result.Results[0].Name, Equals, "The Avengers Collection")
}

func (s *TmdbSuite) TestSearchCompany(c *C) {
	result, err := s.tmdb.SearchCompany("Lucas", nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.TotalResults, Not(Equals), 0)
	c.Assert(result.Results[0].Name, Equals, "Lucas Entertainment")
}

func (s *TmdbSuite) TestSearchKeyword(c *C) {
	result, err := s.tmdb.SearchKeyword("action", nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.TotalResults, Not(Equals), 0)
	c.Assert(result.Results[0].ID, Equals, 207600)
	c.Assert(result.Results[0].Name, Equals, "action")
}

func (s *TmdbSuite) TestSearchList(c *C) {
	result, err := s.tmdb.SearchList("best picture", nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.TotalResults, Not(Equals), 0)
	c.Assert(result.Results[0].ID, Equals, "509ec17b19c2950a0600050d")
	c.Assert(result.Results[0].Name, Equals, "Best Picture Winners - The Academy Awards")
}

func (s *TmdbSuite) TestSearchMovie(c *C) {
	result, err := s.tmdb.SearchMovie("Fight Club", nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.TotalResults, Not(Equals), 0)
	c.Assert(result.Results[0].Title, Equals, "Fight Club")

	var options = make(map[string]string)
	options["language"] = "es"
	spanishResult, err := s.tmdb.SearchMovie("Fight Club", options)
	s.baseTest(&spanishResult, err, c)
	c.Assert(spanishResult.Page, Equals, 1)
	c.Assert(spanishResult.Results, Not(HasLen), 0)
	c.Assert(spanishResult.TotalResults, Not(Equals), 0)
	c.Assert(spanishResult.Results[0].OriginalTitle, Equals, "Fight Club")
	c.Assert(spanishResult.Results[0].Title, Equals, "El club de la lucha")
}

func (s *TmdbSuite) TestSearchMulti(c *C) {
	result, err := s.tmdb.SearchMulti("American", nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.TotalPages > 50, Equals, true)
	c.Assert(result.TotalResults > 1100, Equals, true)
	c.Assert(result.Results[1].ID, Equals, 1433)
	c.Assert(result.Results[1].Name, Equals, "American Dad!")

	var options = make(map[string]string)
	options["page"] = "5"
	page5Result, err := s.tmdb.SearchMulti("American", options)
	s.baseTest(&page5Result, err, c)
	c.Assert(page5Result.Page, Equals, 5)
	c.Assert(page5Result.Results, Not(HasLen), 0)
	c.Assert(page5Result.TotalPages, Equals, result.TotalPages)
	c.Assert(page5Result.TotalResults, Equals, result.TotalResults)
	c.Assert(page5Result.Results[0].ID, Equals, 6582)
	c.Assert(page5Result.Results[0].Name, Equals, "American Gladiators")
}

func (s *TmdbSuite) TestSearchPerson(c *C) {
	result, err := s.tmdb.SearchPerson("Brad Pitt", nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.TotalPages, Equals, 1)
	c.Assert(result.TotalResults, Equals, 1)
	c.Assert(result.Results[0].ID, Equals, 287)
	c.Assert(result.Results[0].Name, Equals, "Brad Pitt")

	var options = make(map[string]string)
	options["search_type"] = "ngram"
	ngramResult, err := s.tmdb.SearchPerson("Brad P", options)
	s.baseTest(&ngramResult, err, c)
	c.Assert(ngramResult.Page, Equals, 1)
	c.Assert(ngramResult.Results, Not(HasLen), 0)
	c.Assert(ngramResult.TotalPages > result.TotalPages, Equals, true)
	c.Assert(ngramResult.TotalResults > result.TotalResults, Equals, true)
	c.Assert(ngramResult.Results[0].ID, Equals, 287)
	c.Assert(ngramResult.Results[0].Name, Equals, "Brad Pitt")
}

func (s *TmdbSuite) TestSearchTv(c *C) {
	result, err := s.tmdb.SearchTv("Breaking Bad", nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, Not(HasLen), 0)
	c.Assert(result.TotalPages, Equals, 1)
	c.Assert(result.TotalResults, Equals, 1)
	c.Assert(result.Results[0].ID, Equals, 1396)
	c.Assert(result.Results[0].Name, Equals, "Breaking Bad")

	var options = make(map[string]string)
	options["search_type"] = "ngram"
	ngramResult, err := s.tmdb.SearchTv("Breaki", options)
	s.baseTest(&ngramResult, err, c)
	c.Assert(ngramResult.Page, Equals, 1)
	c.Assert(ngramResult.Results, Not(HasLen), 0)
	c.Assert(ngramResult.TotalPages > result.TotalPages, Equals, true)
	c.Assert(ngramResult.TotalResults > result.TotalResults, Equals, true)
	c.Assert(ngramResult.Results[0].ID, Equals, 1396)
	c.Assert(ngramResult.Results[0].Name, Equals, "Breaking Bad")
}
