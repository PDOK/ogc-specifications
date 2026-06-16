package wms130

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
)

// MarshalXML Position
func (p Position) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf("%f %f", p[0], p[1])
	return e.EncodeElement(s, start)
}

// UnmarshalXML Position
func (p *Position) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var position Position
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch el := token.(type) {
		case xml.CharData:
			coords := getPositionFromString(string([]byte(el)))
			if len(coords) >= 2 {
				// take first 2 positions (xy)
				position = [2]float64{coords[0], coords[1]}
			}
		case xml.EndElement:
			if el == start.End() {
				*p = position
				return nil
			}
		}
	}
}

func getPositionFromString(position string) []float64 {
	regex := regexp.MustCompile(` `)
	result := regex.Split(position, -1)
	var ps []float64 // slice because length can be 2 or more

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
