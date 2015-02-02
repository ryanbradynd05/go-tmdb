package tmdb

import (
	"fmt"
)

// Timezones map
type Timezones []map[string][]string

// GetTimezonesList gets the list of supported timezones
// http://docs.themoviedb.apiary.io/#reference/timezones/timezoneslist/get
func (tmdb *TMDb) GetTimezonesList() (*Timezones, error) {
	var timezoneList Timezones
	uri := fmt.Sprintf("%s/timezones/list?api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &timezoneList)
	return result.(*Timezones), err
}
