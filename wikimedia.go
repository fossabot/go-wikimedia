// Package wikimedia is an interface to the
// Wikimedia (Wikipedia, Wiktionary, etc.) API built in Go.
package wikimedia

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/json-iterator/go"
)

// ApiPage model struct as defined by any
// Wikimedia API JSON response.
type ApiPage struct {
	PageId    int       `json:"pageid"`
	Ns        int       `json:"ns"`
	Title     string    `json:"title"`
	Extract   string    `json:"extract"`
	Thumbnail Thumbnail `json:"thumbnail"`
	Original  Original  `json:"original"`
}

// ApiQuery model struct as defined by any
// Wikimedia API JSON response.
type ApiQuery struct {
	Pages      map[string]ApiPage `json:"pages"`
	Search     []ApiSearch        `json:"search"`
	SearchInfo ApiSearchInfo      `json:"searchinfo"`
}

// ApiQueryContinue model struct as defined by any
// Wikimedia API JSON response.
type ApiQueryContinue struct {
	Search ApiQueryContinueSearch `json:"search"`
}

// ApiQueryContinueSearch model struct as defined by any
// Wikimedia API JSON response.
type ApiQueryContinueSearch struct {
	SrOffset int `json:"sroffset"`
}

// ApiResponse model struct as defined by any
// Wikimedia API JSON response.
type ApiResponse struct {
	Query         ApiQuery         `json:"query"`
	QueryContinue ApiQueryContinue `json:"query-continue"`
}

// ApiSearch model struct as defined by any
// Wikimedia API JSON response.
type ApiSearch struct {
	Ns        int       `json:"ns"`
	Title     string    `json:"title"`
	Snippet   string    `json:"snippet"`
	Size      int       `json:"size"`
	WordCount int       `json:"wordcount"`
	Timestamp time.Time `json:"timestamp"`
}

// ApiSearchInfo model struct as defined by any
// Wikimedia API JSON response.
type ApiSearchInfo struct {
	TotalHits int `json:"source"`
}

// Original model struct as defined by any
// Wikimedia API JSON response.
type Original struct {
	Source string `json:"source"`
}

// Thumbnail model struct as defined by any
// Wikimedia API JSON response.
type Thumbnail struct {
	Source string `json:"source"`
}

// Wikimedia is an API client struct.
type Wikimedia struct {
	Options Options
}

// Options is a collection of configurable options for the Wikimedia client.
type Options struct {
	// Full URL of the Wikimedia API, e.g. url.Parse("https://en.wikipedia.org/w/api.php")
	URL string
	// HTTP client to use (defaults to http.DefaultClient)
	Client *http.Client
	// User-Agent header to provide
	UserAgent string
}

// json replaces the standard JSON library instance with a faster JSON implementation
// ref: https://github.com/json-iterator/go.
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// New initializes a Wikimedia object that queries the specified API URL,
// e.g. http://en.wikipedia.org/w/api.php or http://da.wiktionary.org/w/api.php.
// Returns a pointer to a Wikimedia struct and an error.
func New(options ...Options) (*Wikimedia, error) {
	var opts Options
	if len(options) == 0 {
		opts = Options{}
	} else {
		opts = options[0]
	}

	if opts.URL == "" {
		return nil, errors.New("URL cannot be nil")
	}
	_, err := url.Parse(opts.URL)
	if err != nil {
		return nil, err
	}

	// Return the new Wikimedia object
	return &Wikimedia{
		Options: opts,
	}, nil
}

// Query quires the Wikimedia API using the user-specified query.
// See https://en.wikipedia.org/w/api.php for a reference.
// Returns a pointer to an ApiResponse and an error.
func (wiki *Wikimedia) Query(query url.Values) (*ApiResponse, error) {
	// Construct the query string and make a new request to the wiki's URL
	query["format"] = []string{"json"}
	queryURL := fmt.Sprintf("%s?%s", wiki.Options.URL, query.Encode())
	resp, err := wiki.get(queryURL)
	if err != nil {
		return nil, err
	}

	// Decode the Wikimedia JSON API response
	var apiResp ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	// Return the Wikimedia API response
	return &apiResp, nil
}

// get sends a request to the specific Wikimedia API utilizing the
// user specified input. Can be any Wikimedia API, not just Wikipedia.
// Returns a pointer to an http.Response and an error.
func (wiki *Wikimedia) get(url string) (*http.Response, error) {
	// Create the request var from the URL specified
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set user agent, if specified
	if wiki.Options.UserAgent != "" {
		req.Header.Set("User-Agent", wiki.Options.UserAgent)
	}

	// Set custom http client to complete the request, if specified
	if wiki.Options.Client != nil {
		return wiki.Options.Client.Do(req)
	}

	// Use the default http client to complete the request otherwise
	return http.DefaultClient.Do(req)
}
