package ows

import (
	"encoding/xml"
	"testing"
)

func TestBoundingBoxBuildQueryString(t *testing.T) {
	var tests = []struct {
		boundingbox       BoundingBox
		boundingboxstring string
	}{
		// While 'not' correct this will we checked in the validation step
		0: {boundingbox: BoundingBox{}, boundingboxstring: `0.000000,0.000000,0.000000,0.000000`},
		1: {boundingbox: BoundingBox{LowerCorner: [2]float64{-180.0, -90.0}, UpperCorner: [2]float64{180.0, 90.0}}, boundingboxstring: `-180.000000,-90.000000,180.000000,90.000000`},
	}
	for k, a := range tests {
		str := a.boundingbox.BuildQueryString()
		if str != a.boundingboxstring {
			t.Errorf("test: %d, expected: %v+,\n got: %v+", k, a.boundingboxstring, str)
		}
	}
}

func TestStripDuplicateAttr(t *testing.T) {
	var tests = []struct {
		attributes []xml.Attr
		expected   []xml.Attr
	}{
		0: {attributes: []xml.Attr{{Name: xml.Name{Local: "gml"}, Value: "http://www.opengis.net/gml/3.2"}}, expected: []xml.Attr{{Name: xml.Name{Local: "gml"}, Value: "http://www.opengis.net/gml/3.2"}}},
		1: {attributes: []xml.Attr{{Name: xml.Name{Local: "gml"}, Value: "http://www.opengis.net/gml/3.2"}, {Name: xml.Name{Local: "gml"}, Value: "http://www.opengis.net/gml/3.2"}, {Name: xml.Name{Local: "gml"}, Value: "http://www.opengis.net/gml/3.2"}},
			expected: []xml.Attr{{Name: xml.Name{Local: "gml"}, Value: "http://www.opengis.net/gml/3.2"}}},
	}

	for k, a := range tests {
		stripped := StripDuplicateAttr(a.attributes)
		if len(a.expected) != len(stripped) {
			t.Errorf("test: %d, expected: %s,\n got: %s", k, a.expected, stripped)
		} else {
			c := false
			for _, exceptedattr := range a.expected {
				for _, result := range stripped {
					if exceptedattr == result {
						c = true
					}
				}
				if !c {
					t.Errorf("test: %d, expected: %s,\n got: %s", k, a.expected, stripped)
				}
				c = false
			}
		}
	}
}
