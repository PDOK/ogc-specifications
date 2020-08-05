package ows

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

// BoundingBox struct
// Base BoundingBox struct to be used for OGC Boundingbox object
// What todo with Geoserver implementation...
// 'cause the http://schemas.opengis.net/ows/1.0.0/owsCommon.xsd is quite clear... PositionTypes and CRS not srsName and coords....
// <BoundingBox srsName="http://www.opengis.net/gml/srs/epsg.xml#4326">
//   <gml:coord><gml:X>-130</gml:X><gml:Y>24</gml:Y></gml:coord>
//   <gml:coord><gml:X>-55</gml:X><gml:Y>50</gml:Y></gml:coord>
// </BoundingBox>
type BoundingBox struct {
	Crs         string   `xml:"crs,attr,omitempty" yaml:"crs,omitempty"`
	Dimensions  string   `xml:"dimensions,attr,omitempty" yaml:"dimensions,omitempty"`
	LowerCorner Position `xml:"LowerCorner" yaml:"lowercorner"`
	UpperCorner Position `xml:"UpperCorner" yaml:"uppercorner"`
}

// Position type
type Position [2]float64

// BuildKVP function for getting a KVP Query BBOX value
func (b *BoundingBox) BuildKVP() string {
	return fmt.Sprintf("%f,%f,%f,%f", b.LowerCorner[0], b.LowerCorner[1], b.UpperCorner[0], b.UpperCorner[1])
}

//Build builds a BoundingBox
func (b *BoundingBox) Build(boundingbox string) Exception {
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

// Keywords in struct for repeatability
type Keywords struct {
	Keyword []string `xml:"Keyword" yaml:"keyword"`
}

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
