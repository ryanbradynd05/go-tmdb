package tmdb

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/kylelemons/go-gypsy/yaml"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type TmdbSuite struct {
	tmdb         *TMDb
	userToken    string
	user         string
	pw           string
	session      string
	guestSession string
	accountID    int
}

var _ = Suite(&TmdbSuite{})

func (s *TmdbSuite) KeyConfig(apiKey string) Config {
	testKeyConfig := Config{
		APIKey:   apiKey,
		Proxies:  nil,
		UseProxy: false,
	}

	return testKeyConfig
}

func (s *TmdbSuite) SetUpSuite(c *C) {
	pwd, _ := os.Getwd()
	basedir := strings.SplitAfter(pwd, "ryanbradynd05/go-tmdb")
	filename := fmt.Sprintf("%s/local.yml", basedir[0])
	config, _ := yaml.ReadFile(filename)
	s.userToken, _ = config.Get("userToken")
	s.user, _ = config.Get("user")
	s.pw, _ = config.Get("pw")
	s.session, _ = config.Get("session")
	s.guestSession, _ = config.Get("guestSession")
	accountID, _ := config.Get("accountID")
	s.accountID, _ = strconv.Atoi(accountID)
	key, _ := config.Get("testKey")
	testKey := s.KeyConfig(key)
	s.tmdb = Init(testKey)
}

func (s *TmdbSuite) baseTest(input interface{}, err error, c *C) {
	c.Assert(err, IsNil)
	jsonRes, err := ToJSON(input)
	c.Assert(err, IsNil)
	c.Assert(jsonRes, NotNil)
}
