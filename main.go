package tmdb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const baseURL string = "http://api.themoviedb.org/3"

// TMDb container struct for global properties
type TMDb struct {
	apiKey string
}

// Init setup the apiKey
func Init(apiKey string) *TMDb {
	return &TMDb{apiKey: apiKey}
}

// ToJSON converts from struct to JSON
func ToJSON(payload interface{}) (string, error) {
	jsonRes, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return "{}", err
	}
	return string(jsonRes), nil
}

func callTmdb(url string, payload interface{}) (interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		return payload, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return payload, err
	}
	json.Unmarshal(body, &payload)
	return payload, err
}
