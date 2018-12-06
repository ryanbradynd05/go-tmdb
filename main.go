package tmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const baseURL string = "https://api.themoviedb.org/3"

const maxRequestPerSecond = 4

var (
	// hack: add some millisecond for don`t get 429 error
	rate     = time.Second/maxRequestPerSecond + time.Millisecond*20
	throttle = time.Tick(rate)
)

// Config struct
type Config struct {
	APIKey   string
	UseProxy bool
	Proxies  []Proxy
}

// Proxy struct
type Proxy struct {
	Host     string
	Port     string
	Login    string
	Password string
	Auth     bool
	throttle <-chan time.Time
}

// TMDb container struct for global properties
type TMDb struct {
	apiKey string
}

var internalConfig tmdbConfig

type tmdbConfig struct {
	useProxy   bool
	proxies    []Proxy
	roundRobin RoundRobin
}

type apiStatus struct {
	Code    int    `json:"status_code"`
	Message string `json:"status_message"`
}

// Init setup the apiKey
func Init(config Config) *TMDb {
	internalConfig := new(tmdbConfig)
	if config.UseProxy == true && len(config.Proxies) > 1 {
		internalConfig.useProxy = config.UseProxy
		internalConfig.proxies = prepareProxies(config.Proxies)
		internalConfig.roundRobin = InitRoundRobin(len(internalConfig.proxies))
	}

	return &TMDb{apiKey: config.APIKey}
}

// ToJSON converts from struct to JSON
func ToJSON(payload interface{}) (string, error) {
	jsonRes := []byte("{}") // Default value in case of error
	jsonRes, err := json.MarshalIndent(payload, "", "  ")
	return string(jsonRes), err
}

func getTmdb(url string, payload interface{}) (interface{}, error) {
	var httpRequest http.Client
	var blocker <-chan time.Time

	if internalConfig.useProxy {
		roundRobin := internalConfig.roundRobin.GetTicker()
		proxy := internalConfig.proxies[roundRobin]

		if proxy.Host == "localhost" {
			httpRequest = getHTTPClient()
		} else {
			httpRequest = getHTTPClientWithProxy(proxy)
		}

		blocker = proxy.throttle
	} else {
		httpRequest = getHTTPClient()
		blocker = throttle
	}

	<-blocker

	res, err := httpRequest.Get(url)
	if err != nil { // HTTP connection error
		return payload, err
	}

	defer res.Body.Close() // Clean up

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

func prepareProxies(proxies []Proxy) []Proxy {
	preparedProxies := make([]Proxy, len(proxies))
	for i, proxy := range proxies {
		proxy.throttle = time.Tick(rate)
		preparedProxies[i] = proxy
	}

	return preparedProxies
}

func getHTTPClient() http.Client {
	return http.Client{
		Transport: &http.Transport{},
	}
}

func getHTTPClientWithProxy(proxy Proxy) http.Client {
	return http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(makeProxyURL(proxy)),
		},
	}
}

func makeProxyURL(proxy Proxy) *url.URL {
	proxyURL := ""
	if proxy.Auth {
		proxyURL = fmt.Sprintf("https://%s:%s@%s:%s",
			proxy.Login,
			proxy.Password,
			proxy.Host,
			proxy.Port)
	} else {
		proxyURL = fmt.Sprintf("https://%s:%s",
			proxy.Host,
			proxy.Port)
	}

	proxyURLInterface, err := url.Parse(proxyURL)

	if err != nil {
		panic(err)
	}

	return proxyURLInterface
}
