package tmdb

import (
	. "gopkg.in/check.v1"
)

const avengersCollectionID int = 86311

func (s *TmdbSuite) TestGetCollectionInfo(c *C) {
	collection, err := s.tmdb.GetCollectionInfo(avengersCollectionID, nil)
	s.baseTest(&collection, err, c)
	c.Assert(collection.ID, Equals, avengersCollectionID)
	c.Assert(collection.Parts, NotNil)
	c.Assert(collection.Parts, Not(HasLen), 0)
	c.Assert(collection.Images, Not(NotNil))

	var options = make(map[string]string)
	options["append_to_response"] = "images"
	collectionImages, err := s.tmdb.GetCollectionInfo(avengersCollectionID, options)
	s.baseTest(&collectionImages, err, c)
	c.Assert(collectionImages.ID, Equals, avengersCollectionID)
	c.Assert(collectionImages.Parts, NotNil)
	c.Assert(collectionImages.Parts, Not(HasLen), 0)
	c.Assert(collectionImages.Images, NotNil)
}

func (s *TmdbSuite) TestGetCollectionImages(c *C) {
	images, err := s.tmdb.GetCollectionImages(avengersCollectionID, nil)
	s.baseTest(&images, err, c)
	c.Assert(images.ID, Equals, avengersCollectionID)
	backdropLength := len(images.Backdrops)

	var options = make(map[string]string)
	options["language"] = "en"
	enImages, err := s.tmdb.GetCollectionImages(avengersCollectionID, options)
	s.baseTest(&enImages, err, c)
	c.Assert(enImages.ID, Equals, avengersCollectionID)
	c.Assert(len(enImages.Backdrops) < backdropLength, Equals, true)
}
