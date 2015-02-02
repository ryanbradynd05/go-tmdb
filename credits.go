package tmdb

import (
	"fmt"
)

// Credit struct
type Credit struct {
	CreditType string `json:"credit_type"`
	Department string
	Job        string
	Media      struct {
		ID           int
		Name         string
		OriginalName string `json:"original_name"`
		Character    string
		Episodes     []struct {
			AirDate       string `json:"air_date"`
			EpisodeNumber int    `json:"episode_number"`
			Name          string
			Overview      string
			SeasonNumber  int    `json:"season_number"`
			StillPath     string `json:"still_path"`
		}
		Seasons []struct {
			AirDate      string `json:"air_date"`
			PosterPath   string `json:"poster_path"`
			SeasonNumber int    `json:"season_number"`
		}
	}
	MediaType string `json:"media_type"`
	ID        string
	Person    struct {
		Name string
		ID   string
	}
}

// GetCreditInfo gets the detailed information about a particular credit record
// http://docs.themoviedb.apiary.io/#reference/credits/creditcreditid/get
func (tmdb *TMDb) GetCreditInfo(id string, options map[string]string) (*Credit, error) {
	var availableOptions = map[string]struct{}{
		"language": {}}
	var creditInfo Credit
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/credit/%v?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &creditInfo)
	return result.(*Credit), err
}
