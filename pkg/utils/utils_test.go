package utils

import (
	"net/url"
	"testing"

	"github.com/pdok/ogc-specifications/pkg/common"
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

func TestIdentifyRequest(t *testing.T) {
	var tests = []struct {
		doc     []byte
		request string
		errors  common.Exceptions
	}{
		0: {doc: []byte(`<Mekker/>`), request: `Mekker`},
		1: {doc: []byte(`<GetCapabilities/>`), request: `GetCapabilities`},
		2: {doc: []byte(`</>`), errors: common.Exceptions{common.MissingParameterValue()}},
		3: {doc: []byte(`<|\/|>`), errors: common.Exceptions{common.MissingParameterValue()}},
		4: {doc: nil, errors: common.Exceptions{common.MissingParameterValue()}},
	}

	for k, i := range tests {
		request, errs := IdentifyRequest(i.doc)
		if errs != nil {
			if errs[0].Error() != i.errors[0].Error() {
				t.Errorf("test: %d, expected: %s \ngot: %s", k, i.errors[0].Error(), errs[0].Error())
			}
		} else {
			if request != i.request {
				t.Errorf("test: %d, expected: %s \ngot: %s", k, i.request, request)
			}
		}

	}
}

func TestIdentifyRequestKVP(t *testing.T) {
	var tests = []struct {
		url     map[string][]string
		request string
		errors  common.Exceptions
	}{
		0: {url: map[string][]string{REQUEST: {`Mekker`}}, request: `Mekker`},
		1: {url: map[string][]string{REQUEST: {`GetCapabilities`}}, request: `GetCapabilities`},
		2: {url: map[string][]string{`SERVICE`: {`NoREQUESTKey`}}, errors: common.Exceptions{common.MissingParameterValue()}},
		3: {url: map[string][]string{}, errors: common.Exceptions{common.MissingParameterValue()}},
		4: {url: nil, errors: common.Exceptions{common.MissingParameterValue()}},
	}

	for k, i := range tests {
		request, errs := IdentifyRequestKVP(i.url)
		if errs != nil {
			if errs[0].Error() != i.errors[0].Error() {
				t.Errorf("test: %d, expected: %s \ngot: %s", k, i.errors[0].Error(), errs[0].Error())
			}
		} else {
			if request != i.request {
				t.Errorf("test: %d, expected: %s \ngot: %s", k, i.request, request)
			}
		}
	}
}
