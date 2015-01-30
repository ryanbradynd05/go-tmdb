package tmdb

import (
	. "gopkg.in/check.v1"
)

func (s *TmdbSuite) TestGetCertificationsMovieList(c *C) {
	movieResult, err := s.tmdb.GetCertificationsMovieList()
	s.baseTest(&movieResult, err, c)
	usMovieCerts := movieResult.Certifications["US"]
	usMovieCertsOpts := "NR|G|PG|PG-13|R|NC-17"
	for _, movieCert := range usMovieCerts {
		c.Assert(movieCert.Certification, Matches, usMovieCertsOpts)
	}
}

func (s *TmdbSuite) TestGetCertificationsTvList(c *C) {
	tvResult, err := s.tmdb.GetCertificationsTvList()
	s.baseTest(&tvResult, err, c)
	usTvCerts := tvResult.Certifications["US"]
	usTvCertsOpts := "NR|TV-Y|TV-Y7|TV-G|TV-PG|TV-14|TV-MA"
	for _, tvCert := range usTvCerts {
		c.Assert(tvCert.Certification, Matches, usTvCertsOpts)
	}
}
