package ows

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
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

// UnmarshalXML for BoundingBox
func (b *BoundingBox) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var boundingbox BoundingBox
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "crs":
			boundingbox.Crs = attr.Value
		case "dimensions":
			boundingbox.Dimensions = attr.Value
		}
	}

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch el := token.(type) {
		case xml.StartElement:
			//TODO compact code
			if el.Name.Local == "LowerCorner" {
				var position string
				elmt := xml.StartElement(el)
				d.DecodeElement(&position, &elmt)

				coords := getPositionFromString(position)
				if len(coords) >= 2 {
					// take first 2 positions (xy)
					boundingbox.LowerCorner = [2]float64{coords[0], coords[1]}
				}
			}
			if el.Name.Local == "UpperCorner" {
				var position string
				elmt := xml.StartElement(el)
				d.DecodeElement(&position, &elmt)

				coords := getPositionFromString(position)
				if len(coords) >= 2 {
					// take first 2 positions (xy)
					boundingbox.UpperCorner = [2]float64{coords[0], coords[1]}
				}
			}

		case xml.EndElement:
			if el == start.End() {
				*b = boundingbox
				return nil
			}
		}
	}
}

// MarshalXML Postion
func (p *Position) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf("%f %f", p[0], p[1])
	return e.EncodeElement(s, start)
}

func getPositionFromString(position string) []float64 {
	regex := regexp.MustCompile(` `)
	result := regex.Split(position, -1)
	var ps []float64 //slice because lenght can be 2 or more

	// check if 'strings' are parsable to float64
	// if one is not return nothing
	for _, fs := range result {
		f, err := strconv.ParseFloat(fs, 64)
		if err != nil {
			return nil
		}
		ps = append(ps, f)
	}
	return ps
}
