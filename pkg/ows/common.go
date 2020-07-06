package ows

import (
	"encoding/xml"
	"fmt"
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
	Crs         string   `xml:"crs,attr"`
	Dimensions  string   `xml:"dimensions,attr"`
	LowerCorner Position `xml:"LowerCorner"`
	UpperCorner Position `xml:"UpperCorner"`
}

// Position type
type Position [2]float64

// BuildQueryString function for getting a KVP Query BBOX value
func (b *BoundingBox) BuildQueryString() string {
	return fmt.Sprintf("%f,%f,%f,%f", b.LowerCorner[0], b.LowerCorner[1], b.UpperCorner[0], b.UpperCorner[1])
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
