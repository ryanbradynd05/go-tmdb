package tmdb

import (
	. "gopkg.in/check.v1"
)

const bradPittID int = 287

func (s *TmdbSuite) TestGetPersonInfo(c *C) {
	result, err := s.tmdb.GetPersonInfo(bradPittID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Name, Equals, "Brad Pitt")
	c.Assert(result.ID, Equals, bradPittID)
	c.Assert(result.Adult, Equals, false)
	c.Assert(result.Birthday, Equals, "1963-12-18")
	c.Assert(result.Deathday, Equals, "")
	c.Assert(result.AlsoKnownAs[0], Equals, "William Bradley Pitt")
	c.Assert(result.PlaceOfBirth, Equals, "Shawnee - Oklahoma - USA")

	var options = make(map[string]string)
	options["append_to_response"] = "changes,movie_credits,tv_credits,combined_credits,external_ids,images,tagged_images"
	allResult, err := s.tmdb.GetPersonInfo(bradPittID, options)
	s.baseTest(&allResult, err, c)
	c.Assert(allResult.ID, Equals, bradPittID)
	c.Assert(allResult.Name, Equals, "Brad Pitt")
	c.Assert(allResult.Changes, NotNil)
	c.Assert(allResult.MovieCredits, NotNil)
	c.Assert(allResult.TvCredits, NotNil)
	c.Assert(allResult.CombinedCredits, NotNil)
	c.Assert(allResult.ExternalIds, NotNil)
	c.Assert(allResult.Images, NotNil)
	c.Assert(allResult.TaggedImages, NotNil)
}

func (s *TmdbSuite) TestGetPersonChanges(c *C) {
	result, err := s.tmdb.GetPersonChanges(bradPittID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Changes, NotNil)
	allResultLength := len(result.Changes)

	var janOptions = make(map[string]string)
	janOptions["start_date"] = "2015-01-01"
	janOptions["end_date"] = "2015-01-14"
	janResult, err := s.tmdb.GetPersonChanges(bradPittID, janOptions)
	s.baseTest(&janResult, err, c)
	c.Assert(janResult.Changes, NotNil)
	c.Assert(len(janResult.Changes) >= allResultLength, Equals, true)
}

func (s *TmdbSuite) TestGetPersonMovieCredits(c *C) {
	result, err := s.tmdb.GetPersonMovieCredits(bradPittID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, bradPittID)
	c.Assert(result.Cast, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetPersonTvCredits(c *C) {
	result, err := s.tmdb.GetPersonTvCredits(bradPittID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, bradPittID)
	c.Assert(result.Cast, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetPersonCombinedCredits(c *C) {
	result, err := s.tmdb.GetPersonCombinedCredits(bradPittID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, bradPittID)
	c.Assert(result.Cast, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetPersonExternalIds(c *C) {
	result, err := s.tmdb.GetPersonExternalIds(bradPittID)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, bradPittID)
	c.Assert(result.ImdbID, Equals, "nm0000093")
	c.Assert(result.TvrageID, Equals, 59436)
}

func (s *TmdbSuite) TestGetPersonImages(c *C) {
	result, err := s.tmdb.GetPersonImages(bradPittID)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, bradPittID)
	c.Assert(len(result.Profiles) > 5, Equals, true)
}

func (s *TmdbSuite) TestGetPersonTaggedImages(c *C) {
	result, err := s.tmdb.GetPersonTaggedImages(bradPittID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, bradPittID)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, HasLen, 20)
	c.Assert(result.TotalPages >= 3, Equals, true)
	c.Assert(result.TotalResults >= 45, Equals, true)

	var page2Options = make(map[string]string)
	page2Options["page"] = "2"
	page2Result, err := s.tmdb.GetPersonTaggedImages(bradPittID, page2Options)
	s.baseTest(&page2Result, err, c)
	c.Assert(page2Result.Page, Equals, 2)
	c.Assert(page2Result.Results, HasLen, 20)
	c.Assert(page2Result.TotalPages, Equals, result.TotalPages)
	c.Assert(page2Result.TotalResults, Equals, result.TotalResults)
}

func (s *TmdbSuite) TestGetPersonPopular(c *C) {
	result, err := s.tmdb.GetPersonPopular(nil)
	s.baseTest(&result, err, c)
	c.Assert(result.Page, Equals, 1)
	c.Assert(result.Results, HasLen, 20)
	c.Assert(result.TotalPages >= 3, Equals, true)
	c.Assert(result.TotalResults >= 45, Equals, true)

	var page2Options = make(map[string]string)
	page2Options["page"] = "2"
	page2Result, err := s.tmdb.GetPersonPopular(page2Options)
	s.baseTest(&page2Result, err, c)
	c.Assert(page2Result.Page, Equals, 2)
	c.Assert(page2Result.Results, HasLen, 20)
}

func (s *TmdbSuite) TestGetPersonLatest(c *C) {
	result, err := s.tmdb.GetPersonLatest()
	s.baseTest(&result, err, c)
	c.Assert(result.ID, NotNil)
}
