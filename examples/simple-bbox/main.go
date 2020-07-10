package main

import (
	"encoding/xml"
	"fmt"

	ows "github.com/pdok/ogc-specifications/pkg/ows"
)

func main() {

	var bbox ows.BoundingBox
	bbox.Crs = `EPSG:4326`
	bbox.LowerCorner = [2]float64{-180.0, -90.0}
	bbox.UpperCorner = [2]float64{180.0, 90.0}

	xml, _ := xml.MarshalIndent(bbox, "", " ")

	fmt.Println(string(xml))

}
