package main

import (
	"encoding/xml"
	"fmt"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

func main() {

	var bbox wsc110.BoundingBox
	bbox.Crs = `EPSG:4326`
	bbox.LowerCorner = [2]float64{-180.0, -90.0}
	bbox.UpperCorner = [2]float64{180.0, 90.0}

	xml, _ := xml.MarshalIndent(bbox, "", " ")

	fmt.Println(string(xml))

}
