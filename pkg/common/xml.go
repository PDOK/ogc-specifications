package common

import (
	"encoding/xml"
	"fmt"
)

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
