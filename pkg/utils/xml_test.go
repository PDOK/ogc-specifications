package utils

import (
	"encoding/xml"
	"errors"
	"testing"
)

func TestUnMarshalXMLAttribute(t *testing.T) {
	var tests = []struct {
		xmlraw   string
		expected XMLAttribute
		err      error
	}{
		0: {xmlraw: `<startelement attr="one"/>`, expected: XMLAttribute{xml.Attr{Name: xml.Name{Local: "attr"}, Value: "one"}}},
		1: {xmlraw: `<startelement attr="two" attr="three"/>`, expected: XMLAttribute{xml.Attr{Name: xml.Name{Local: "attr"}, Value: "two"}, xml.Attr{Name: xml.Name{Local: "attr"}, Value: "three"}}},
		2: {xmlraw: `<startelement b:attr="two" b:item="three"/>`, expected: XMLAttribute{xml.Attr{Name: xml.Name{Space: "b", Local: "attr"}, Value: "two"}, xml.Attr{Name: xml.Name{Space: "b", Local: "item"}, Value: "three"}}},
		3: {xmlraw: `<startelement attr="one"`, err: errors.New("XML syntax error on line 1: unexpected EOF")},
	}

	for k, test := range tests {
		var xmlattr XMLAttribute
		if err := xml.Unmarshal([]byte(test.xmlraw), &xmlattr); err != nil {
			if err.Error() != test.err.Error() {
				t.Errorf("test: %d, expected no error,\n got: %s", k, err.Error())
			}
		}

		if len(test.expected) != len(xmlattr) {
			t.Errorf("test: %d, expected: %s,\n got: %s", k, test.expected, xmlattr)
		} else {
			c := false
			for _, exceptedAttr := range test.expected {
				for _, result := range xmlattr {
					if exceptedAttr == result {
						c = true
					}
				}
				if !c {
					t.Errorf("test: %d, expected: %s,\n got: %s", k, test.expected, xmlattr)
				}
				c = false
			}
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
