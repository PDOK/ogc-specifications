package utils

import (
	"net/url"
	"strings"
)

// KeysToUpper convert all the keys to UpperCase
func KeysToUpper(q url.Values) url.Values {
	r := url.Values{}
	for k, v := range q {
		r[strings.ToUpper(k)] = append(r[strings.ToUpper(k)], v...)
	}
	return r
}

// KeysToLower convert all the keys to LowerCase
func KeysToLower(q url.Values) url.Values {
	r := url.Values{}
	for k, v := range q {
		r[strings.ToLower(k)] = append(r[strings.ToLower(k)], v...)
	}
	return r
}
