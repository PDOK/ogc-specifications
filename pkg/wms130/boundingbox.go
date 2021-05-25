package wms130

import (
	"fmt"
	"strconv"
	"strings"
)

type BoundingBox struct {
	Crs         string   `xml:"crs,attr,omitempty" yaml:"crs,omitempty"`
	Dimensions  string   `xml:"dimensions,attr,omitempty" yaml:"dimensions,omitempty"`
	LowerCorner Position `xml:"LowerCorner" yaml:"lowercorner"`
	UpperCorner Position `xml:"UpperCorner" yaml:"uppercorner"`
}

// Position type
type Position [2]float64

// ToQueryParameters function for getting a KVP Query BBOX value
func (b *BoundingBox) BuildQueryParameters() string {
	return fmt.Sprintf("%f,%f,%f,%f", b.LowerCorner[0], b.LowerCorner[1], b.UpperCorner[0], b.UpperCorner[1])
}

//ParseString builds a BoundingBox based on a string
func (b *BoundingBox) parseString(boundingbox string) Exceptions {
	result := strings.Split(boundingbox, ",")
	var lx, ly, ux, uy float64
	var err error

	if len(result) < 4 {
		return InvalidParameterValue(boundingbox, `boundingbox`).ToExceptions()
	}

	if len(result) == 4 || len(result) == 5 {
		if lx, err = strconv.ParseFloat(result[0], 64); err != nil {
			return InvalidParameterValue(boundingbox, `boundingbox`).ToExceptions()
		}
		if ly, err = strconv.ParseFloat(result[1], 64); err != nil {
			return InvalidParameterValue(boundingbox, `boundingbox`).ToExceptions()
		}
		if ux, err = strconv.ParseFloat(result[2], 64); err != nil {
			return InvalidParameterValue(boundingbox, `boundingbox`).ToExceptions()
		}
		if uy, err = strconv.ParseFloat(result[3], 64); err != nil {
			return InvalidParameterValue(boundingbox, `boundingbox`).ToExceptions()
		}
	}

	b.LowerCorner = [2]float64{lx, ly}
	b.UpperCorner = [2]float64{ux, uy}

	if len(result) == 5 {
		b.Crs = result[4]
	}

	return nil
}
