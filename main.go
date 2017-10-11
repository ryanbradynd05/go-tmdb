package tmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const baseURL string = "https://api.themoviedb.org/3"

var (
	requestMutex   = &sync.Mutex{}
	rateLimitReset time.Time
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
	// Go single-threaded so we can deal with the rate limit
	requestMutex.Lock()
	defer requestMutex.Unlock()

	now := time.Now()
	if rateLimitReset.After(now) { // We have a reset time in the future, so we're out of requests
		// Wait for rate limiter to be reset
		<-time.After(rateLimitReset.Sub(now))
	}

	res, err := http.Get(url)
	if err != nil { // HTTP connection error
		return payload, err
	}

	if res.Header.Get(`x-ratelimit-remaining`) == `0` { // Out of requests for this period
		reset := res.Header.Get(`x-ratelimit-reset`)
		iReset, err := strconv.ParseInt(reset, 10, 64)
		if err == nil {
			// Set the reset time here, the next request will trip it
			rateLimitReset = time.Unix(iReset+1, 0)
		}
	}

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
