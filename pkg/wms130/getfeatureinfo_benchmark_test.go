package wms130

import (
	"encoding/xml"
	"testing"
)

// ----------
// Benchmarks
// ----------

func BenchmarkGetFeatureInfoToQueryParameters(b *testing.B) {
	gfi := GetFeatureInfoRequest{
		XMLName: xml.Name{Local: `GetFeatureInfo`},
		BaseRequest: BaseRequest{
			Service: Service,
			Version: Version},
		StyledLayerDescriptor: StyledLayerDescriptor{
			NamedLayer: []NamedLayer{
				{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
			}},
		CRS: "EPSG:4326",
		BoundingBox: BoundingBox{
			LowerCorner: [2]float64{-180.0, -90.0},
			UpperCorner: [2]float64{180.0, 90.0},
		},
		Size:        Size{Width: 1024, Height: 512},
		QueryLayers: []string{`CenterLine`},
		InfoFormat:  `application/json`,
		I:           1,
		J:           1,
	}
	for i := 0; i < b.N; i++ {
		gfi.ToQueryParameters()
	}
}

func BenchmarkGetFeatureInfoToXML(b *testing.B) {
	gfi := GetFeatureInfoRequest{
		XMLName: xml.Name{Local: `GetFeatureInfo`},
		BaseRequest: BaseRequest{
			Service: Service,
			Version: Version},
		StyledLayerDescriptor: StyledLayerDescriptor{
			NamedLayer: []NamedLayer{
				{Name: "Rivers", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Roads", NamedStyle: &NamedStyle{Name: "CenterLine"}},
				{Name: "Houses", NamedStyle: &NamedStyle{Name: "Outline"}},
			}},
		CRS: "EPSG:4326",
		BoundingBox: BoundingBox{
			LowerCorner: [2]float64{-180.0, -90.0},
			UpperCorner: [2]float64{180.0, 90.0},
		},
		Size:        Size{Width: 1024, Height: 512},
		QueryLayers: []string{`CenterLine`},
		InfoFormat:  `application/json`,
		I:           1,
		J:           1,
	}
	for i := 0; i < b.N; i++ {
		gfi.ToXML()
	}
}
