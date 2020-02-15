# go-wikimedia
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fciehanski%2Fgo-wikimedia.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fciehanski%2Fgo-wikimedia?ref=badge_shield)

go-wikimedia is an interface to the Wikimedia (Wikipedia, Wiktionary, etc.) API 
implemented in the Go programming language.

This project was originally created by [Patrick Mylund Nielsen](https://github.com/patrickmn). I 
forked his repo for my project [pastime](https://github.com/ciehanski/pastime). If you notice a bug, feel free to submit an issue on this repo. 

## Installation

```bash
go get github.com/ciehanski/go-wikimedia
```

## Documentation

[https://godoc.org/github.com/ciehanski/go-wikimedia](https://godoc.org/github.com/ciehanski/go-wikimedia)  

or from the CLI:   

```bash
go doc github.com/ciehanski/go-wikimedia
```

## Usage

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	
	"github.com/ciehanski/go-wikimedia"
)

func main() {
    wiki, err := wikimedia.New(wikimedia.Options{
    	Client:    http.DefaultClient,
    	URL:       "https://en.wikipedia.org/w/api.php",
    	UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 " +
    		"(KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
    })
    if err != nil {
    	log.Fatal(err.Error())
    }
    resp, err := wiki.Query(url.Values{
        "action":      {"query"},
        "prop":        {"extracts"},
        "titles":      {"Osmosis|Procrastination"},
        "exsentences": {"5"},
        "explaintext": {"1"},
        "original":    {"source"},
    })
    if err != nil {
    	log.Fatalf("Error executing query: %s", err.Error())
    }
    for _, v := range resp.Query.Pages {
    	fmt.Println(v.Title, "-", v.Extract)
    }
}
```

## Contributing

Please feel free to contribute and submit any PRs to this project.

## License

MIT

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fciehanski%2Fgo-wikimedia.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fciehanski%2Fgo-wikimedia?ref=badge_large)