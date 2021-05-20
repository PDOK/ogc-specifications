package utils

import (
	"errors"
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

func TestIdentifyRequest(t *testing.T) {
	var tests = []struct {
		doc     []byte
		request string
		errors  error
	}{
		0: {doc: []byte(`<?xml version="1.0" encoding="UTF-8"?>
		<Mekker/>`), request: `Mekker`},
		1: {doc: []byte(`<GetCapabilities/>`), request: `GetCapabilities`},
		2: {doc: []byte(`<ogc:GetMap xmlns:ogc="http://www.opengis.net/ows"
		xmlns:gml="http://www.opengis.net/gml"
		version="1.3.0" service="WMS">
<StyledLayerDescriptor version="1.1.0">
  <NamedLayer>
	<Name>pand</Name>
  </NamedLayer>
</StyledLayerDescriptor>
<BoundingBox srsName="http://www.opengis.net/gml/srs/epsg.xml#3857">
  <coord><X>662489.7241121939151</X><Y>6834200.591356366873</Y></coord>
  <coord><X>663837.270904958481</X><Y>6835015.857165988535</Y></coord>
</BoundingBox>
<Output>
  <Format>image/png</Format>
  <Size>
	<Width>800</Width>
	<Height>450</Height>
  </Size>
</Output>
</ogc:GetMap>`), request: `GetMap`},
		3: {doc: []byte(`</>`), errors: errors.New(`unknown REQUEST parameter`)},
		4: {doc: []byte(`<|\/|>`), errors: errors.New(`unknown REQUEST parameter`)},
		5: {doc: nil, errors: errors.New(`unknown REQUEST parameter`)},
	}

	for k, test := range tests {
		request, errs := IdentifyRequest(test.doc)
		if errs != nil {
			if errs.Error() != test.errors.Error() {
				t.Errorf("test: %d, expected: %s \ngot: %s", k, test.errors.Error(), errs.Error())
			}
		} else {
			if request != test.request {
				t.Errorf("test: %d, expected: %s \ngot: %s", k, test.request, request)
			}
		}
	}
}

func TestIdentifyRequestKVP(t *testing.T) {
	var tests = []struct {
		url     map[string][]string
		request string
		errors  error
	}{
		0: {url: map[string][]string{REQUEST: {`Mekker`}}, request: `Mekker`},
		1: {url: map[string][]string{REQUEST: {`GetCapabilities`}}, request: `GetCapabilities`},
		2: {url: map[string][]string{`SERVICE`: {`NoREQUESTKey`}}, errors: errors.New(`unknown REQUEST parameter`)},
		3: {url: map[string][]string{}, errors: errors.New(`unknown REQUEST parameter`)},
		4: {url: nil, errors: errors.New(`unknown REQUEST parameter`)},
	}

	for k, test := range tests {
		request, errs := IdentifyRequestKVP(test.url)
		if errs != nil {
			if errs.Error() != test.errors.Error() {
				t.Errorf("test: %d, expected: %s \ngot: %s", k, test.errors.Error(), errs.Error())
			}
		} else {
			if request != test.request {
				t.Errorf("test: %d, expected: %s \ngot: %s", k, test.request, request)
			}
		}
	}
}

func BenchmarkIdentifyRequest(b *testing.B) {
	r := []byte(`<ogc:GetMap xmlns:ogc="http://www.opengis.net/ows"
	xmlns:gml="http://www.opengis.net/gml"
	version="1.3.0" service="WMS">
<StyledLayerDescriptor version="1.1.0">
<NamedLayer>
<Name>pand</Name>
</NamedLayer>
</StyledLayerDescriptor>
<BoundingBox srsName="http://www.opengis.net/gml/srs/epsg.xml#3857">
<coord><X>662489.7241121939151</X><Y>6834200.591356366873</Y></coord>
<coord><X>663837.270904958481</X><Y>6835015.857165988535</Y></coord>
</BoundingBox>
<Output>
<Format>image/png</Format>
<Size>
<Width>800</Width>
<Height>450</Height>
</Size>
</Output>
</ogc:GetMap>`)
	for i := 0; i < b.N; i++ {
		IdentifyRequest(r)
		IdentifyRequest(nil)
	}
}

func BenchmarkIdentifyRequestKVP(b *testing.B) {
	kvp := map[string][]string{REQUEST: {`getfeature`}, `SERVICE`: {`WFS`}, `VERSION`: {`2.0.0`}, `OUTPUTFORMAT`: {"application/xml"}, `TYPENAMES`: {"dummy"}, `COUNT`: {"3"}}

	for i := 0; i < b.N; i++ {
		IdentifyRequestKVP(kvp)
		IdentifyRequestKVP(map[string][]string{})
	}
}
