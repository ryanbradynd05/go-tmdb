package tmdb

import (
	"fmt"
)

// Network struct
type Network struct {
	ID   int
	Name string
}

// GetNetworkInfo gets the basic information about a TV network
// https://developers.themoviedb.org/3/networks/get-network-details
func (tmdb *TMDb) GetNetworkInfo(id int) (*Network, error) {
	var networkInfo Network
	uri := fmt.Sprintf("%s/network/%v?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := getTmdb(uri, &networkInfo)
	return result.(*Network), err
}
