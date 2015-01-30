package tmdb

import (
	. "gopkg.in/check.v1"
)

const avengersCollectionID int = 86311

func (s *TmdbSuite) TestGetCollectionInfo(c *C) {
	collection, err := s.tmdb.GetCollectionInfo(avengersCollectionID)
	s.baseTest(&collection, err, c)
	c.Assert(collection.ID, Equals, avengersCollectionID)
	c.Assert(collection.Parts, NotNil)
	c.Assert(collection.Parts, Not(HasLen), 0)
}

func (s *TmdbSuite) TestGetCollectionImages(c *C) {
	images, err := s.tmdb.GetCollectionImages(avengersCollectionID)
	s.baseTest(&images, err, c)
	c.Assert(images.ID, Equals, avengersCollectionID)
}
