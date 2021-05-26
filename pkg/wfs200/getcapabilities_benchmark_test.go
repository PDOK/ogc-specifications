package wfs200

import (
	"encoding/xml"
	"testing"
)

// ----------
// Benchmarks
// ----------

func BenchmarkGetCapabilitiesToQueryParameters(b *testing.B) {
	gc := GetCapabilitiesRequest{XMLName: xml.Name{Local: getcapabilities}, Service: Service, Version: Version}
	for i := 0; i < b.N; i++ {
		gc.ToQueryParameters()
	}
}

func BenchmarkGetCapabilitiesToXML(b *testing.B) {
	gc := GetCapabilitiesRequest{XMLName: xml.Name{Local: getcapabilities}, Service: Service, Version: Version}
	for i := 0; i < b.N; i++ {
		gc.ToXML()
	}
}
