package utils

import (
	"encoding/xml"
)

// XMLAttribute wrapper around the array of xml.Attr
type XMLAttribute []xml.Attr

// StripDuplicateAttr removes the duplicate Attributes from a []Attribute
func StripDuplicateAttr(attr []xml.Attr) []xml.Attr {
	attributeMap := make(map[xml.Name]string)
	for _, a := range attr {
		attributeMap[xml.Name{Space: a.Name.Space, Local: a.Name.Local}] = a.Value
	}

	var strippedAttr []xml.Attr
	for k, v := range attributeMap {
		strippedAttr = append(strippedAttr, xml.Attr{Name: k, Value: v})
	}
	return strippedAttr
}

// UnmarshalXML func for the XMLAttr struct
func (xmlattr *XMLAttribute) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var newAttributes XMLAttribute
	for _, attr := range start.Attr {
		newAttributes = append(newAttributes, xml.Attr{Name: attr.Name, Value: attr.Value})
	}
	*xmlattr = newAttributes

	for {
		// if it got this far the XML is 'valid' and the xmlattr are set
		// so we ignore the err
		token, _ := d.Token()

		if el, ok := token.(xml.EndElement); ok {
			if el == start.End() {
				return nil
			}
		}
	}
}
