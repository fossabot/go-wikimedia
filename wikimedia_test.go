package wikimedia

import (
	"fmt"
	"net/url"
	"testing"
)

func TestQuery(t *testing.T) {
	wiki, _ := New(Options{
		URL: "https://en.wikipedia.org/w/api.php",
	})
	resp, err := wiki.Query(url.Values{
		"action":      {"query"},
		"prop":        {"extracts"},
		"titles":      {"Google"},
		"exsentences": {"5"},
		"explaintext": {"1"},
	})
	if err != nil {
		t.Error(err)
	}

	for _, v := range resp.Query.Pages {
		if v.Title != "Google" {
			t.Error("Incorrect response value")
		}
	}
}

func TestGet(t *testing.T) {
	wiki, _ := New(Options{
		URL: "https://en.wikipedia.org/w/api.php",
	})
	query := url.Values{
		"action":      {"query"},
		"prop":        {"extracts"},
		"titles":      {"Google"},
		"exsentences": {"5"},
		"explaintext": {"1"},
	}
	queryURL := fmt.Sprintf("%s?%s", wiki.Options.URL, query.Encode())
	_, err := wiki.get(queryURL)
	if err != nil {
		t.Error(err)
	}

	resp, err := wiki.get("oF#.com.()389HPFDI*&$@")
	if resp != nil && err == nil {
		t.Error(err)
	}
}

func TestNew(t *testing.T) {
	wiki, err := New(Options{
		URL: "https://en.wikipedia.org/w/api.php",
	})
	if err != nil {
		t.Error(err)
	}
	if wiki == nil {
		t.Error("New() error, blank Wikimedia object returned")
	}

	// Test blank URL
	wiki, err = New(Options{
		URL: "",
	})
	if wiki != nil && err == nil {
		t.Error("Expected an error and blank wiki object")
	}
}

func BenchmarkQuery(b *testing.B) {
	wiki, _ := New(Options{
		URL: "https://en.wikipedia.org/w/api.php",
	})
	for i := 0; i < b.N; i++ {
		wiki.Query(url.Values{
			"action":      {"query"},
			"prop":        {"extracts"},
			"titles":      {"Google"},
			"exsentences": {"5"},
			"explaintext": {"1"},
		})
	}
}

func BenchmarkGet(b *testing.B) {
	wiki, _ := New(Options{
		URL: "https://en.wikipedia.org/w/api.php",
	})
	query := url.Values{
		"action":      {"query"},
		"prop":        {"extracts"},
		"titles":      {"Google"},
		"exsentences": {"5"},
		"explaintext": {"1"},
	}
	queryURL := fmt.Sprintf("%s?%s", wiki.Options.URL, query.Encode())
	for i := 0; i < b.N; i++ {
		wiki.get(queryURL)
	}
}
