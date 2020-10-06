package tmdb

import (
	"fmt"
)

// Timezones map
type Timezones []map[string][]string

// GetTimezonesList gets the list of supported timezones
// https://developers.themoviedb.org/3/configuration/get-timezones
func (tmdb *TMDb) GetTimezonesList() (*Timezones, error) {
	var timezoneList Timezones
	uri := fmt.Sprintf("%s/timezones/list?api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &timezoneList)
	return result.(*Timezones), err
}
