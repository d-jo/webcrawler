package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func main() {
	var handlerFunc http.HandlerFunc

	var pages map[string]string

	handlerFunc = func(w http.ResponseWriter, r *http.Request) {
		p, ok := pages[r.URL.Path]
		if ok {
			w.Write([]byte(p))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}

	srv := httptest.NewServer(handlerFunc)
	defer srv.Close()

	pages = map[string]string{
		"/":      fmt.Sprintf("<a href=\"%s/page1\">page1</a>", srv.URL),
		"/page1": fmt.Sprintf("<a href=\"%s/page2\">page2</a><a href=\"%s/page3\">page3</a>", srv.URL, srv.URL),
		"/page2": fmt.Sprintf("<a href=\"%s/page1\">page1</a>", srv.URL),
		"/page3": "empty page wiith no links",
	}

	fmt.Println(srv.URL)

	// wait for input
	var input string
	fmt.Scanln(&input)
}
