package tmdb

import (
	"fmt"
)

// Collection struct
type Collection struct {
	BackdropPath string `json:"backdrop_path"`
	ID           int
	Name         string
	PosterPath   string            `json:"poster_path"`
	Images       *CollectionImages `json:",omitempty"`
	Parts        []struct {
		BackdropPath string `json:"backdrop_path"`
		ID           int
		PosterPath   string `json:"poster_path"`
		ReleaseDate  string `json:"release_date"`
		Title        string
	}
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

// GetCollectionInfo gets the basic collection information for a specific collection id
// http://docs.themoviedb.apiary.io/#reference/collections/collectionid/get
func (tmdb *TMDb) GetCollectionInfo(id int, options map[string]string) (*Collection, error) {
	var availableOptions = map[string]struct{}{
		"language":           {},
		"append_to_response": {}}
	var collection Collection
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/collection/%v?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &collection)
	return result.(*Collection), err
}

// GetCollectionImages gets a list of people ids that have been edited
// http://docs.themoviedb.apiary.io/#reference/collections/collectionidimages/get
func (tmdb *TMDb) GetCollectionImages(id int, options map[string]string) (*CollectionImages, error) {
	var availableOptions = map[string]struct{}{
		"language":               {},
		"append_to_response":     {},
		"include_image_language": {}}
	var images CollectionImages
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/collection/%v/images?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &images)
	return result.(*CollectionImages), err
}
