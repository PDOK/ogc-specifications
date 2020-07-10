package utils

import (
	"net/url"
	"strings"
)

// KeysToUpper convert all the keys to lowercase
func KeysToUpper(query url.Values) url.Values {
	newquery := url.Values{}
	for key, values := range query {
		newquery[strings.ToUpper(key)] = append(newquery[strings.ToUpper(key)], values...)
	}
	return newquery
}
