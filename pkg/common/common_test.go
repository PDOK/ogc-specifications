package common

import (
	"encoding/xml"
	"testing"
)

func TestStripDuplicateAttr(t *testing.T) {
	var tests = []struct {
		attributes []xml.Attr
		expected   []xml.Attr
	}{
		0: {attributes: []xml.Attr{{Name: xml.Name{Local: "gml"}, Value: "http://www.opengis.net/gml/3.2"}}, expected: []xml.Attr{{Name: xml.Name{Local: "gml"}, Value: "http://www.opengis.net/gml/3.2"}}},
		1: {attributes: []xml.Attr{{Name: xml.Name{Local: "gml"}, Value: "http://www.opengis.net/gml/3.2"}, {Name: xml.Name{Local: "gml"}, Value: "http://www.opengis.net/gml/3.2"}, {Name: xml.Name{Local: "gml"}, Value: "http://www.opengis.net/gml/3.2"}},
			expected: []xml.Attr{{Name: xml.Name{Local: "gml"}, Value: "http://www.opengis.net/gml/3.2"}}},
	}

	for k, test := range tests {
		stripped := StripDuplicateAttr(test.attributes)
		if len(test.expected) != len(stripped) {
			t.Errorf("test: %d, expected: %s,\n got: %s", k, test.expected, stripped)
		} else {
			c := false
			for _, exceptedAttr := range test.expected {
				for _, result := range stripped {
					if exceptedAttr == result {
						c = true
					}
				}
				if !c {
					t.Errorf("test: %d, expected: %s,\n got: %s", k, test.expected, stripped)
				}
				c = false
			}
		}
	}
}

func TestCRSParseString(t *testing.T) {
	var tests = []struct {
		input       string
		expectedCRS CRS
	}{
		0: {}, // Empty input == empty struct
		1: {input: `urn:ogc:def:crs:EPSG::4326`, expectedCRS: CRS{Code: 4326, Namespace: `EPSG`}},
		2: {input: `EPSG:4326`, expectedCRS: CRS{Code: 4326, Namespace: `EPSG`}},
	}

	for k, test := range tests {
		var crs CRS
		crs.parseString(test.input)

		if crs.Code != test.expectedCRS.Code || crs.Namespace != test.expectedCRS.Namespace {
			t.Errorf("test: %d, expected: %v,\n got: %v", k, test.expectedCRS, crs)
		}
	}
}
