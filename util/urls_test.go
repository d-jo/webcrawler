package util_test

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/d-jo/webcrawler/util"
)

func TestURLs1(t *testing.T) {
	input := `<a href="http://www.example.com/1"></a>`
	expected := []string{"http://www.example.com/1"}

	actual, err := util.SearchForURLs(input, "example.com")

	if err != nil {
		t.Errorf("error searching for urls: %s", err)
	}

	if len(actual) != len(expected) {
		t.Errorf("expected %d urls, got %d", len(expected), len(actual))
	}

	for i := range expected {
		if actual[i][0] != expected[i] {
			t.Errorf("expected %s, got %s", expected[i], actual[i])
		}
	}
}

func TestURLs2(t *testing.T) {
	input := `filler text ... http://www.example.com/1?param=2 ... filler text`
	expected := []string{"http://www.example.com/1?param=2"}

	actual, err := util.SearchForURLs(input, "example.com")

	if err != nil {
		t.Errorf("error searching for urls: %s", err)
	}

	if len(actual) != len(expected) {
		t.Errorf("expected %d urls, got %d", len(expected), len(actual))
	}

	for i := range expected {
		if actual[i][0] != expected[i] {
			t.Errorf("expected %s, got %s", expected[i], actual[i])
		}
	}
}

func TestURLs3(t *testing.T) {
	input := "http://localhost:5555/"
	expected := "http://localhost:5555"

	parsedUrl, err := url.Parse(input)

	if err != nil {
		t.Errorf("error parsing url: %s", err)
	}

	host := parsedUrl.Scheme + "://" + parsedUrl.Hostname()
	port := parsedUrl.Port()

	if len(port) > 0 {
		host = fmt.Sprintf("%s:%s", host, port)
	}

	if host != expected {
		t.Errorf("expected %s, got %s", expected, host)
	}
}

func TestURLs4(t *testing.T) {
	input := "http://example.com/"
	expected := "http://example.com"

	parsedUrl, err := url.Parse(input)

	if err != nil {
		t.Errorf("error parsing url: %s", err)
	}

	host := parsedUrl.Scheme + "://" + parsedUrl.Hostname()
	port := parsedUrl.Port()

	if len(port) > 0 {
		host = fmt.Sprintf("%s:%s", host, port)
	}

	if host != expected {
		t.Errorf("expected %s, got %s", expected, host)
	}
}
