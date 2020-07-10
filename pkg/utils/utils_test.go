package utils

import (
	"net/url"
	"testing"
)

func TestKeysToUpper(t *testing.T) {
	var testkeysToUpperQuerys = []struct {
		query         url.Values
		expectedQuery url.Values
	}{
		// Default GetCapbilities request
		0: {query: map[string][]string{"SERVICE": {"WFS"}, "REQUEST": {"GetCapabilities"}, "VERSION": {"2.0.0"}}, expectedQuery: map[string][]string{"SERVICE": {"WFS"}, "REQUEST": {"GetCapabilities"}, "VERSION": {"2.0.0"}}},
		// UPPER, lower, MiXeDcAsE GetCapbilities request
		1: {query: map[string][]string{"SERVICE": {"WFS"}, "request": {"GetCapabilities"}, "VeRsIoN": {"2.0.0"}}, expectedQuery: map[string][]string{"SERVICE": {"WFS"}, "REQUEST": {"GetCapabilities"}, "VERSION": {"2.0.0"}}},
		// empty request
		2: {query: map[string][]string{}, expectedQuery: map[string][]string{}},
		// nothing in nothing out same as empty request
		3: {},
		// Multiple parameters
		4: {query: map[string][]string{"SERVICE": {"WFS"}, "SeRvIcE": {"WMS"}, "service": {"wmts"}}, expectedQuery: map[string][]string{"SERVICE": {"WFS", "wmts", "WMS"}}},
	}

	for k, tq := range testkeysToUpperQuerys {
		q := KeysToUpper(tq.query)
		if len(q) != len(tq.expectedQuery) {
			t.Errorf("test: %d, expected: %s \ngot: %s", k, tq.expectedQuery, q)
		}
	}
}
