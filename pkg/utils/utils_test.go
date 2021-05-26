package utils

import (
	"net/url"
	"testing"
)

func TestKeysToUpper(t *testing.T) {
	var tests = []struct {
		query         url.Values
		expectedquery url.Values
	}{
		// Default GetCapbilities request
		0: {query: map[string][]string{"SERVICE": {"WFS"}, "REQUEST": {"GetCapabilities"}, "VERSION": {"2.0.0"}}, expectedquery: map[string][]string{"SERVICE": {"WFS"}, "REQUEST": {"GetCapabilities"}, "VERSION": {"2.0.0"}}},
		// UPPER, lower, MiXeDcAsE GetCapbilities request
		1: {query: map[string][]string{"SERVICE": {"WFS"}, "request": {"GetCapabilities"}, "VeRsIoN": {"2.0.0"}}, expectedquery: map[string][]string{"SERVICE": {"WFS"}, "REQUEST": {"GetCapabilities"}, "VERSION": {"2.0.0"}}},
		// empty request
		2: {query: map[string][]string{}, expectedquery: map[string][]string{}},
		// nothing in nothing out same as empty request
		3: {},
		// Multiple parameters
		4: {query: map[string][]string{"SERVICE": {"WFS"}, "SeRvIcE": {"WMS"}, "service": {"wmts"}}, expectedquery: map[string][]string{"SERVICE": {"WFS", "wmts", "WMS"}}},
	}

	for k, test := range tests {
		q := KeysToUpper(test.query)
		if len(q) != len(test.expectedquery) {
			t.Errorf("test: %d, expected: %s \ngot: %s", k, test.expectedquery, q)
		}
	}
}
