package wms130

import (
	"encoding/xml"
	"errors"
	"testing"
)

func TestUnMarshalXMLAttribute(t *testing.T) {
	var tests = []struct {
		xmlraw    string
		expected  XMLAttribute
		exception error
	}{
		0: {xmlraw: `<startelement attr="one"/>`, expected: XMLAttribute{xml.Attr{Name: xml.Name{Local: "attr"}, Value: "one"}}},
		1: {xmlraw: `<startelement attr="two" attr="three"/>`, expected: XMLAttribute{xml.Attr{Name: xml.Name{Local: "attr"}, Value: "two"}, xml.Attr{Name: xml.Name{Local: "attr"}, Value: "three"}}},
		2: {xmlraw: `<startelement b:attr="two" b:item="three"/>`, expected: XMLAttribute{xml.Attr{Name: xml.Name{Space: "b", Local: "attr"}, Value: "two"}, xml.Attr{Name: xml.Name{Space: "b", Local: "item"}, Value: "three"}}},
		3: {xmlraw: `<startelement attr="one"`, exception: errors.New("XML syntax error on line 1: unexpected EOF")},
	}

	for k, a := range tests {
		var xmlattr XMLAttribute
		if err := xml.Unmarshal([]byte(a.xmlraw), &xmlattr); err != nil {
			if err.Error() != a.exception.Error() {
				t.Errorf("test: %d, expected no error,\n got: %s", k, err.Error())
			}
		}

		if len(a.expected) != len(xmlattr) {
			t.Errorf("test: %d, expected: %s,\n got: %s", k, a.expected, xmlattr)
		} else {
			c := false
			for _, exceptedAttr := range a.expected {
				for _, result := range xmlattr {
					if exceptedAttr == result {
						c = true
					}
				}
				if !c {
					t.Errorf("test: %d, expected: %s,\n got: %s", k, a.expected, xmlattr)
				}
				c = false
			}
		}
	}
}
