package tmdb

import (
	. "gopkg.in/check.v1"
	"testing"
)

const testKey string = "b12bd26ee15e6972f2f4c03363b74fa2"

func Test(t *testing.T) { TestingT(t) }

type TmdbSuite struct {
	tmdb *TMDb
}

var _ = Suite(&TmdbSuite{})

func (s *TmdbSuite) SetUpSuite(c *C) {
	s.tmdb = Init(testKey)
}

func (s *TmdbSuite) baseTest(input interface{}, err error, c *C) {
	c.Assert(err, IsNil)
	jsonRes, err := ToJSON(input)
	c.Assert(err, IsNil)
	c.Assert(jsonRes, NotNil)
}
