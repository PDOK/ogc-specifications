package utils

import (
	"encoding/xml"
)

// XMLAttribute wrapper around the array of xml.Attr
type XMLAttribute []xml.Attr

// StripDuplicateAttr removes the duplicate Attributes from a []Attribute
func StripDuplicateAttr(attr []xml.Attr) []xml.Attr {
	attributemap := make(map[xml.Name]string)
	for _, a := range attr {
		attributemap[xml.Name{Space: a.Name.Space, Local: a.Name.Local}] = a.Value
	}

	var strippedAttr []xml.Attr
	for k, v := range attributemap {
		strippedAttr = append(strippedAttr, xml.Attr{Name: k, Value: v})
	}
	return strippedAttr
}

// UnmarshalXML func for the XMLAttr struct
func (xmlattr *XMLAttribute) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var newattributes XMLAttribute
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		default:
			newattributes = append(newattributes, xml.Attr{Name: attr.Name, Value: attr.Value})
		}
	}
	*xmlattr = newattributes

	for {
		// if it got this far the XML is 'valid' and the xmlattr are set
		// so we ignore the err
		token, _ := d.Token()
		switch el := token.(type) {
		case xml.EndElement:
			if el == start.End() {
				return nil
			}
		}
	}
}
