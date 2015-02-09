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
	// options["append_to_response"] = "account_states,alternative_titles,changes,credits,images,keywords,similar,translations,videos"
	options["append_to_response"] = "credits"
	allResult, err := s.tmdb.GetPersonInfo(bradPittID, options)
	s.baseTest(&allResult, err, c)
	c.Assert(allResult.ID, Equals, bradPittID)
	c.Assert(allResult.Name, Equals, "Brad Pitt")
	// c.Assert(allResult.Changes, NotNil)
	c.Assert(allResult.Credits, NotNil)
	// c.Assert(allResult.Images, NotNil)
	// c.Assert(allResult.Keywords, NotNil)
	// c.Assert(allResult.Similar, NotNil)
	// c.Assert(allResult.Translations, NotNil)
	// c.Assert(allResult.Videos, NotNil)
}

func (s *TmdbSuite) TestGetPersonCredits(c *C) {
	result, err := s.tmdb.GetPersonCredits(bradPittID, nil)
	s.baseTest(&result, err, c)
	c.Assert(result.ID, Equals, bradPittID)
	c.Assert(result.Cast, Not(HasLen), 0)
}
