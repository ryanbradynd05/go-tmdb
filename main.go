package tmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const baseURL string = "https://api.themoviedb.org/3"

const MAX_REQUEST_PER_SECOND = 4

var (
	// hack: add some millisecond for dont get 429 error
	rate = time.Second / MAX_REQUEST_PER_SECOND + time.Millisecond * 20
	throttle = time.Tick(rate)
)

// TMDb container struct for global properties
type TMDb struct {
	apiKey string
}

type apiStatus struct {
	Code    int    `json:"status_code"`
	Message string `json:"status_message"`
}

// Init setup the apiKey
func Init(apiKey string) *TMDb {
	return &TMDb{apiKey: apiKey}
}

// ToJSON converts from struct to JSON
func ToJSON(payload interface{}) (string, error) {
	jsonRes := []byte("{}") //Default value in case of error
	jsonRes, err := json.MarshalIndent(payload, "", "  ")
	return string(jsonRes), err
}

func getTmdb(url string, payload interface{}) (interface{}, error) {
	<-throttle

	res, err := http.Get(url)
	if err != nil { // HTTP connection error
		return payload, err
	}

	defer res.Body.Close()  //Clean up
  
	body, err := ioutil.ReadAll(res.Body)
	if err != nil { // Failed to read body
		return payload, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 { // Success!
		json.Unmarshal(body, &payload)
		return payload, nil
	}

	// Handle failure modes
	var status apiStatus
	err = json.Unmarshal(body, &status)
	if err != nil {
		return payload, err
	}
	return payload, fmt.Errorf("Code (%d): %s", status.Code, status.Message)
}

func getOptionsString(options map[string]string, availableOptions map[string]struct{}) string {
	var optionsString = ""
	for key, val := range options {
		if _, ok := availableOptions[key]; ok {
			newString := fmt.Sprintf("%s&%s=%s", optionsString, key, val)
			optionsString = newString
		}
	}
	return optionsString
}
