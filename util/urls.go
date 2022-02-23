package util

import (
	"regexp"
	"strings"
)

// generic regex for all links:

// https?:\/\/(www)?\.?([^(\s|\")]*)

// need one for just a specific domain

// https?:\/\/(www.?)?.+(google)\.?([^(\s|\")]*)

// improved:

// (https?:\/\/)(www.?)?[^\"\s]*(stackoverflow.com)\.?([^(\s|\")]*)

// even more improved:

// (https?:\/\/)(www.?)?[^\"\s]*(stackoverflow\.com)\.?([^\s\"]*)

// final version:

// (https?:\/\/)(www.?)?[^\"\s\']*(stackoverflow\.com)\.?([^\s\"\']*)

/*
Given a string, return a list of all URLs in the string that belong to the domain given with host. host should
include the domain and tld only

returns the result from the regex as a slice of string slices
the first element in each is the full match
*/
func SearchForURLs(text string, host string) ([][]string, error) {
	// group 0: protocol
	// group 1: www ? if it exists
	//   	 : then we match some url characters only. not whitespace or quotes
	// group 2: host, including top level domain
	// group 3: path and arguments
	// end the regex capturing URL characters until a non-URL character is found like a quote in
	// link tags or whitespace
	rx, err := regexp.Compile(`(https?:\/\/)(www.?)?[^\"\s\']*(` + strings.ReplaceAll(host, ".", "\\.") + `)\.?([^\s\"\']*)`)

	if err != nil {
		return nil, err
	}

	matches := rx.FindAllStringSubmatch(text, -1)

	return matches, nil
}
