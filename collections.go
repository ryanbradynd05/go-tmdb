package tmdb

import (
	"fmt"
)

// Collection struct
type Collection struct {
	BackdropPath string `json:"backdrop_path"`
	ID           int
	Name         string
	PosterPath   string `json:"poster_path"`
	Parts        []struct {
		BackdropPath string `json:"backdrop_path"`
		ID           int
		PosterPath   string `json:"poster_path"`
		ReleaseDate  string `json:"release_date"`
		Name         string
	}
}

// GetCollectionInfo gets the basic collection information for a specific collection id
// http://docs.themoviedb.apiary.io/#reference/collections/collectionid/get
func (tmdb *TMDb) GetCollectionInfo(id int) (*Collection, error) {
	var collection Collection
	uri := fmt.Sprintf("%s/collection/%v?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(uri, &collection)
	return result.(*Collection), err
}

// CollectionImage struct
type CollectionImage struct {
	AspectRatio float32 `json:"aspect_ratio"`
	FilePath    string  `json:"file_path"`
	Height      int
	Iso639_1    string `json:"iso_639_1"`
	Width       int
}

// CollectionImages struct
type CollectionImages struct {
	Backdrops  []CollectionImage
	ID         int
	Name       string
	PosterPath string `json:"poster_path"`
	Posters    []CollectionImage
}

// GetCollectionImages gets a list of people ids that have been edited
// http://docs.themoviedb.apiary.io/#reference/collections/collectionidimages/get
func (tmdb *TMDb) GetCollectionImages(id int) (*CollectionImages, error) {
	var images CollectionImages
	uri := fmt.Sprintf("%s/collection/%v/images?api_key=%s", baseURL, id, tmdb.apiKey)
	result, err := callTmdb(uri, &images)
	return result.(*CollectionImages), err
}
