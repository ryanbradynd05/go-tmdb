package tmdb

import (
	"fmt"
)

// Person struct
type Person struct {
	ID           int
	Name         string
	Overview     string
	Adult        bool
	Biography    string
	Birthday     string
	Deathday     string
	Homepage     string
	AlsoKnownAs  []string       `json:"also_known_as"`
	PlaceOfBirth string         `json:"place_of_birth"`
	ProfilePath  string         `json:"profile_path"`
	Credits      *PersonCredits `json:",omitempty"`
}

// PersonCredits struct
type PersonCredits struct {
	ID   int
	Cast []struct {
		Adult         bool
		Character     string
		CreditID      string `json:"credit_id"`
		ID            int
		OriginalTitle string `json:"original_title"`
		PosterPath    string `json:"poster_path"`
		ReleaseDate   string `json:"release_date"`
		Title         string
	}
}

// GetPersonInfo gets the general person information for a specific id
// http://docs.themoviedb.apiary.io/#reference/people/personid/get
func (tmdb *TMDb) GetPersonInfo(id int, options map[string]string) (*Person, error) {
	var availableOptions = map[string]struct{}{
		"append_to_response": {}}
	var personInfo Person
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/person/%v?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &personInfo)
	return result.(*Person), err
}

// GetPersonCredits for a specific person id
// http://docs.themoviedb.apiary.io/#reference/people/personidmoviecredits/get
func (tmdb *TMDb) GetPersonCredits(id int, options map[string]string) (*PersonCredits, error) {
	var availableOptions = map[string]struct{}{
		"language":           {},
		"append_to_response": {}}
	var credits PersonCredits
	optionsString := getOptionsString(options, availableOptions)
	uri := fmt.Sprintf("%s/person/%v/credits?api_key=%s%s", baseURL, id, tmdb.apiKey, optionsString)
	result, err := getTmdb(uri, &credits)
	json, _ := ToJSON(result)
	fmt.Printf("%v\n%v\n", uri, json)
	return result.(*PersonCredits), err
}
