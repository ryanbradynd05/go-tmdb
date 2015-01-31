package tmdb

import (
	. "gopkg.in/check.v1"
)

const (
	kennyBaniaID      string = "5256c98119c2956ff604faab"
	tyrionLannisterID string = "5256c8b219c2956ff6047cd8"
)

func (s *TmdbSuite) TestGetCreditInfo(c *C) {
	tyrion, err := s.tmdb.GetCreditInfo(tyrionLannisterID)
	s.baseTest(&tyrion, err, c)
	c.Assert(tyrion.ID, Equals, tyrionLannisterID)
	c.Assert(tyrion.Person.Name, Equals, "Peter Dinklage")
	c.Assert(tyrion.Media.Character, Equals, "Tyrion Lannister")
	c.Assert(tyrion.Media.Seasons, NotNil)
	c.Assert(tyrion.Media.Seasons, Not(HasLen), 0)

	bania, err := s.tmdb.GetCreditInfo(kennyBaniaID)
	s.baseTest(&bania, err, c)
	c.Assert(bania.ID, Equals, kennyBaniaID)
	c.Assert(bania.Person.Name, Equals, "Steve Hytner")
	c.Assert(bania.Media.OriginalName, Equals, "Seinfeld")
	c.Assert(bania.Media.Episodes, NotNil)
	c.Assert(bania.Media.Episodes, Not(HasLen), 0)
}
