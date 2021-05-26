package wsc110

import (
	"fmt"
	"strconv"
	"strings"
)

// BoundingBoxUnmarshal struct
// BoundingBox struct to be used for Unmarshalling of boundingbox object
// regarding the issue https://stackoverflow.com/questions/48609596/xml-namespace-prefix-issue-at-go
// There are 2 structs needed one for Unmarshalling and one for Marshalling
type BoundingBoxUnmarshal struct {
	Crs         string   `xml:"crs,attr,omitempty" yaml:"crs,omitempty"`
	Dimensions  string   `xml:"dimensions,attr,omitempty" yaml:"dimensions,omitempty"`
	LowerCorner Position `xml:"LowerCorner" yaml:"lowercorner"`
	UpperCorner Position `xml:"UpperCorner" yaml:"uppercorner"`
}

// BoundingBox struct
type BoundingBox struct {
	Crs         string   `xml:"crs,attr,omitempty" yaml:"crs,omitempty"`
	Dimensions  string   `xml:"dimensions,attr,omitempty" yaml:"dimensions,omitempty"`
	LowerCorner Position `xml:"ows:LowerCorner" yaml:"lowercorner"`
	UpperCorner Position `xml:"ows:UpperCorner" yaml:"uppercorner"`
}

// WGS84BoundingBox layers on the wsc110.BoundingBox
// with a predefined crs "urn:ogc:def:crs:OGC::84"
// and bound dimensions, lowercorner and uppercorner, all 2 values
type WGS84BoundingBox BoundingBox

// Position type
type Position [2]float64

// ToQueryParameters function for getting a string value from a Query BBOX
func (b *BoundingBox) ToQueryParameters() string {
	return fmt.Sprintf("%f,%f,%f,%f", b.LowerCorner[0], b.LowerCorner[1], b.UpperCorner[0], b.UpperCorner[1])
}

//ParseString builds a BoundingBox based on a string
func (b *BoundingBox) ParseString(boundingbox string) Exception {
	result := strings.Split(boundingbox, ",")
	var lx, ly, ux, uy float64
	var err error

	if len(result) < 4 {
		return InvalidParameterValue(boundingbox, `boundingbox`)
	}

	if len(result) == 4 || len(result) == 5 {
		if lx, err = strconv.ParseFloat(result[0], 64); err != nil {
			return InvalidParameterValue(boundingbox, `boundingbox`)
		}
		if ly, err = strconv.ParseFloat(result[1], 64); err != nil {
			return InvalidParameterValue(boundingbox, `boundingbox`)
		}
		if ux, err = strconv.ParseFloat(result[2], 64); err != nil {
			return InvalidParameterValue(boundingbox, `boundingbox`)
		}
		if uy, err = strconv.ParseFloat(result[3], 64); err != nil {
			return InvalidParameterValue(boundingbox, `boundingbox`)
		}
	}

	b.LowerCorner = [2]float64{lx, ly}
	b.UpperCorner = [2]float64{ux, uy}

	if len(result) == 5 {
		b.Crs = result[4]
	}

	return nil
}
