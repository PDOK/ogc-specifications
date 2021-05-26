package wfs200

import (
	"testing"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

// ----------
// Benchmarks
// ----------

func BenchmarkGetFeatureToQueryParameters(b *testing.B) {
	gf := GetFeatureRequest{Query: Query{SrsName: sp("srsname"), Filter: &Filter{SpatialOperator: SpatialOperator{BBOX: &GEOBBOX{Envelope: Envelope{LowerCorner: wsc110.Position{1, 1}, UpperCorner: wsc110.Position{2, 2}}}}, ResourceID: &ResourceIDs{{Rid: "one"}, {Rid: "two"}, {Rid: "three"}}}}, BaseRequest: BaseRequest{Version: Version}}
	for i := 0; i < b.N; i++ {
		gf.ToQueryParameters()
	}
}

func BenchmarkGetFeatureToXML(b *testing.B) {
	gf := GetFeatureRequest{
		Query: Query{
			SrsName: sp("srsname"),
			Filter: &Filter{
				SpatialOperator: SpatialOperator{
					BBOX: &GEOBBOX{Envelope: Envelope{
						LowerCorner: wsc110.Position{1, 1},
						UpperCorner: wsc110.Position{2, 2}}},
				},
				ResourceID: &ResourceIDs{
					{Rid: "one"},
					{Rid: "two"},
					{Rid: "three"},
				},
			},
		},
		BaseRequest: BaseRequest{Service: Service, Version: Version}}
	for i := 0; i < b.N; i++ {
		gf.ToXML()
	}
}

func BenchmarkGetFeatureParseQueryParameters(b *testing.B) {
	pv := map[string][]string{REQUEST: {getfeature}, SERVICE: {Service}, VERSION: {Version}, OUTPUTFORMAT: {"application/xml"}, TYPENAMES: {"dummy"}, COUNT: {"3"}}

	for i := 0; i < b.N; i++ {
		f := GetFeatureRequest{}
		f.ParseQueryParameters(pv)
	}
}

func BenchmarkGetFeatureParseXML(b *testing.B) {
	doc := []byte(`<?xml version="1.0" encoding="UTF-8"?>
	<GetFeature outputFormat="application/gml+xml; version=3.2" count="3" startindex="0" service="WFS" version="2.0.0" xmlns:kadastralekaartv4="http://kadastralekaartv4.geonovum.nl">
	 <Query typeNames="kadastralekaartv4:kadastralegrens" srsName="urn:ogc:def:crs:EPSG::28992">
	  <fes:Filter>
	   <fes:PropertyIsEqualTo matchCase="true">
		<fes:ValueReference>id</fes:ValueReference>
		<fes:Literal>29316bf0-b87f-4e8d-bf00-21f894bdf655</fes:Literal>
	   </fes:PropertyIsEqualTo>
	  </fes:Filter>
	 </Query>
	</GetFeature>`)

	for i := 0; i < b.N; i++ {
		gm := GetFeatureRequest{}
		gm.ParseXML(doc)
	}
}
