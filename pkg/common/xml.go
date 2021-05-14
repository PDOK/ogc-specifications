package common

import (
	"encoding/xml"
	"fmt"
)

// XMLAttribute wrapper around the array of xml.Attr
type XMLAttribute []xml.Attr

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

// MarshalXML Position
func (c *CRS) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var s = ``
	if c.Namespace != `` {
		s = fmt.Sprintf("%s:%d", c.Namespace, c.Code)
	}

	return e.EncodeElement(s, start)
}

// UnmarshalXML Position
func (c *CRS) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var crs CRS
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch el := token.(type) {
		case xml.CharData:
			crs.parseString(string([]byte(el)))
		case xml.EndElement:
			if el == start.End() {
				*c = crs
				return nil
			}
		}
	}
}