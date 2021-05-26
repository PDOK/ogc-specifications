package wfs200

import (
	"encoding/xml"
	"testing"
)

// ----------
// Benchmarks
// ----------

func BenchmarkDescribeFeatureTypeToQueryParameters(b *testing.B) {
	df := DescribeFeatureTypeRequest{
		XMLName:     xml.Name{Local: `DescribeFeatureType`},
		BaseRequest: BaseRequest{Version: Version, Service: Service},
		BaseDescribeFeatureTypeRequest: BaseDescribeFeatureTypeRequest{
			OutputFormat: sp(`application/json`),
			TypeName:     sp(`example:example`)}}
	for i := 0; i < b.N; i++ {
		df.ToQueryParameters()
	}
}

func BenchmarkDescribeFeatureTypeToXML(b *testing.B) {
	df := DescribeFeatureTypeRequest{
		XMLName:     xml.Name{Local: `DescribeFeatureType`},
		BaseRequest: BaseRequest{Version: Version, Service: Service},
		BaseDescribeFeatureTypeRequest: BaseDescribeFeatureTypeRequest{
			OutputFormat: sp(`application/json`),
			TypeName:     sp(`example:example`)}}
	for i := 0; i < b.N; i++ {
		df.ToXML()
	}
}
