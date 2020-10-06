package tmdb

import (
	"fmt"
)

// ListInfo struct
type ListInfo struct {
	CreatedBy     string `json:"created_by"`
	FavoriteCount int    `json:"favorite_count"`
	ID            string
	Description   string
	Items         []MovieShort
	ItemCount     int    `json:"item_count"`
	Iso639_1      string `json:"iso_639_1"`
	Name          string
	PosterPath    string `json:"poster_path"`
}

// ListItemStatus struct
type ListItemStatus struct {
	ID          string
	ItemPresent bool `json:"item_present"`
}

// GetListInfo gets a list by id
// https://developers.themoviedb.org/3/lists/get-list-details
func (tmdb *TMDb) GetListInfo(id string) (*ListInfo, error) {
	var listInfo ListInfo
	uri := fmt.Sprintf("%s/list/%v?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := getTmdb(uri, &listInfo)
	return result.(*ListInfo), err
}

// DeleteList deletes a list by id
// https://developers.themoviedb.org/3/lists/delete-list
// func (tmdb *TMDb) DeleteList(id string) {
// TODO
// }

// GetListItemStatus checks to see if a movie ID is already added to a list
// hhttps://developers.themoviedb.org/3/lists/check-item-status
func (tmdb *TMDb) GetListItemStatus(id string, movieID int) (*ListItemStatus, error) {
	var itemStatus ListItemStatus
	uri := fmt.Sprintf("%s/list/%v/item_status?api_key=%s&movie_id=%v", baseURL, id, tmdb.apiKey, movieID)
	result, err := getTmdb(uri, &itemStatus)
	return result.(*ListItemStatus), err
}

// PostListAddItem lets users add new movies to a list that they created
// https://developers.themoviedb.org/3/lists/add-movie
// func (tmdb *TMDb) PostListAddItem(id string) {
// TODO
// }

// PostListRemoveItem lets users remove movies from a list that they created
// https://developers.themoviedb.org/3/lists/remove-movie
// func (tmdb *TMDb) PostListRemoveItem(id string) {
// TODO
// }

// PostListClear clears all of the items in a list
// https://developers.themoviedb.org/3/lists/clear-list
// func (tmdb *TMDb) PostListRemoveItem(id string) {
// TODO
// }
