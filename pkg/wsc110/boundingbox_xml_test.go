package wsc110

import (
	"encoding/xml"
	"errors"
	"testing"
)

func TestUnMarshalXMLBoundingBox(t *testing.T) {
	var tests = []struct {
		xmlraw      string
		boundingbox BoundingBoxUnmarshal
		exception   error
	}{
		// BoundingBox from GetMap schemas.opengis.net/sld/1.1.0/example_getmap.xml example request
		0: {xmlraw: `<BoundingBox crs="http://www.opengis.net/gml/srs/epsg.xml#4326">
		<ows:LowerCorner>-180.0 -90.0</ows:LowerCorner>
		<ows:UpperCorner>180.0 90.0</ows:UpperCorner>
		</BoundingBox>`,
			boundingbox: BoundingBoxUnmarshal{Crs: "http://www.opengis.net/gml/srs/epsg.xml#4326", LowerCorner: [2]float64{-180.0, -90.0}, UpperCorner: [2]float64{180.0, 90.0}}},
		1: {xmlraw: `<BoundingBox crs="http://www.opengis.net/gml/srs/epsg.xml#4326" dimensions="2">
			<ows:LowerCorner>-180.0 -90.0</ows:LowerCorner>
			<ows:UpperCorner>180.0 90.0</ows:UpperCorner>
			</BoundingBox>`,
			boundingbox: BoundingBoxUnmarshal{Crs: "http://www.opengis.net/gml/srs/epsg.xml#4326", Dimensions: "2", LowerCorner: [2]float64{-180.0, -90.0}, UpperCorner: [2]float64{180.0, 90.0}}},
		2: {xmlraw: `<BoundingBox crs="http://www.opengis.net/gml/srs/epsg.xml#4326" dimensions="2">
			<ows:LowerCorner/>
			<ows:UpperCorner/>
			</BoundingBox>`,
			boundingbox: BoundingBoxUnmarshal{Crs: "http://www.opengis.net/gml/srs/epsg.xml#4326", Dimensions: "2"}},
		3: {xmlraw: `<BoundingBox/>`,
			boundingbox: BoundingBoxUnmarshal{}},
		4: {xmlraw: `<BoundingBox crs="http://www.opengis.net/gml/srs/epsg.xml#4326" dimensions="2">
			<ows:LowerCorner>Not a coord</ows:LowerCorner>
			<ows:UpperCorner/>
			</BoundingBox>`,
			boundingbox: BoundingBoxUnmarshal{Crs: "http://www.opengis.net/gml/srs/epsg.xml#4326", Dimensions: "2"}},
		5: {xmlraw: `<BoundingBox crs="http://www.opengis.net/gml/srs/epsg.xml#4326" dimensions="2">
			<ows:LowerCorner>Not a coord</ows:LowerCorner>
			<ows:UpperCorner/>
			corrupt xml"`,
			exception: errors.New("XML syntax error on line 4: unexpected EOF")},
	}
	for k, a := range tests {
		var bbox BoundingBoxUnmarshal
		if err := xml.Unmarshal([]byte(a.xmlraw), &bbox); err != nil {
			if err.Error() != a.exception.Error() {
				t.Errorf("test: %d, expected no error,\n got: %s", k, err.Error())
			}

		} else {
			if a.boundingbox != bbox {
				t.Errorf("test: %d, expected: %v+,\n got: %v+", k, a.boundingbox, bbox)
			}
		}
	}
}

func TestMarshalXMLPosition(t *testing.T) {
	var tests = []struct {
		position Position
		xml      string
	}{
		0: {position: Position{}, xml: "<Position>0.000000 0.000000</Position>"},
		1: {position: Position{-180.0, 90.0}, xml: "<Position>-180.000000 90.000000</Position>"},
	}
	for k, a := range tests {
		d, err := xml.Marshal(&a.position)
		if err != nil {
			t.Errorf("xml.Marshal failed with '%s'\n", err)
		}
		str := string(d)
		if str != a.xml {
			t.Errorf("test: %d, expected: %v+,\n got: %v+", k, a.xml, str)
		}
	}
}

func TestUnMarshalXMLPosition(t *testing.T) {
	var tests = []struct {
		position  Position
		xml       string
		exception error
	}{
		0: {position: Position{}, xml: "<Position>0.000000 0.000000</Position>", exception: errors.New("")},
		1: {position: Position{-180.0, 90.0}, xml: "<Position>-180.000000 90.000000</Position>", exception: errors.New("")},
		2: {position: Position{}, xml: "<Position/>", exception: errors.New("")},
		3: {position: Position{}, xml: "EOF", exception: errors.New("EOF")},
	}
	for k, a := range tests {
		var position Position
		if err := xml.Unmarshal([]byte(a.xml), &position); err != nil {
			if err.Error() != a.exception.Error() {
				t.Errorf("test: %d, expected no error,\n got: %s", k, err.Error())
			}

		} else {
			if a.position != position {
				t.Errorf("test: %d, expected: %v+,\n got: %v+", k, a.position, position)
			}
		}
	}
}
