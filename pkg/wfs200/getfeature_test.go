package wfs200

import (
	"encoding/xml"
	"net/url"
	"testing"

	"github.com/pdok/ogc-specifications/pkg/common"
	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

func sp(s string) *string {
	return &s
}

func ip(i int) *int {
	return &i
}

func TestGetFeatureType(t *testing.T) {
	dft := GetFeatureRequest{}
	if dft.Type() != `GetFeature` {
		t.Errorf("test: %d, expected: %s,\n got: %s", 0, `GetFeature`, dft.Type())
	}
}

func TestGetFeatureBuildXML(t *testing.T) {
	var tests = []struct {
		gf     GetFeatureRequest
		result string
	}{
		0: {gf: GetFeatureRequest{StandardPresentationParameters: &StandardPresentationParameters{OutputFormat: sp("application/gml+xml; version=3.2"), Count: ip(3), Startindex: ip(0)}, BaseRequest: BaseRequest{Service: "WFS", Version: "2.0.0"}, Query: Query{TypeNames: "test", SrsName: sp("urn:ogc:def:crs:EPSG::28992")}},
			result: `<?xml version="1.0" encoding="UTF-8"?>
<GetFeature service="WFS" version="2.0.0" outputFormat="application/gml+xml; version=3.2" count="3" startindex="0">
 <Query typeNames="test" srsName="urn:ogc:def:crs:EPSG::28992"></Query>
</GetFeature>`},
		1: {gf: GetFeatureRequest{StandardPresentationParameters: &StandardPresentationParameters{OutputFormat: sp("application/gml+xml; version=3.2"), Count: ip(3), Startindex: ip(0)}, BaseRequest: BaseRequest{
			Attr: common.XMLAttribute{
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "kadastralekaartv4"}, Value: "http://kadastralekaartv4.geonovum.nl"}},
			Service: "WFS", Version: "2.0.0"}, Query: Query{TypeNames: "test", SrsName: sp("urn:ogc:def:crs:EPSG::28992")}},
			result: `<?xml version="1.0" encoding="UTF-8"?>
<GetFeature service="WFS" version="2.0.0" xmlns:_xmlns="xmlns" _xmlns:kadastralekaartv4="http://kadastralekaartv4.geonovum.nl" outputFormat="application/gml+xml; version=3.2" count="3" startindex="0">
 <Query typeNames="test" srsName="urn:ogc:def:crs:EPSG::28992"></Query>
</GetFeature>`},
	}

	for k, test := range tests {
		body := test.gf.ToXML()

		if string(body) != test.result {
			t.Errorf("test: %d, Expected body %s but was not \n got: %s", k, test.result, string(body))
		}
	}
}

// TODO
// Merge TestParseBodyGetFeature & TestParseQueryParameters GetFeature comporison into single func, like with WMS GetMap
func TestGetFeatureParseXML(t *testing.T) {
	var tests = []struct {
		body      []byte
		result    GetFeatureRequest
		exception []wsc110.Exception
	}{
		// Get 3 features
		0: {body: []byte(`<GetFeature outputFormat="application/gml+xml; version=3.2" count="3" startindex="0" service="WFS" version="2.0.0">
		<Query typeNames="kadastralekaart:kadastralegrens" srsName="urn:ogc:def:crs:EPSG::28992"/>
	   </GetFeature>`),
			result: GetFeatureRequest{XMLName: xml.Name{Local: "GetFeature"}, StandardPresentationParameters: &StandardPresentationParameters{OutputFormat: sp("application/gml+xml; version=3.2"), Count: ip(3), Startindex: ip(0)},
				BaseRequest: BaseRequest{Service: "WFS", Version: "2.0.0"}}},
		// Get feature by resourceid
		1: {body: []byte(`<?xml version="1.0" encoding="UTF-8"?>
		<GetFeature outputFormat="application/gml+xml; version=3.2" count="3" startindex="0" service="WFS" version="2.0.0" xmlns:kadastralekaartv4="http://kadastralekaartv4.geonovum.nl">
		 <Query typeNames="kadastralekaartv4:kadastralegrens" srsName="urn:ogc:def:crs:EPSG::28992">
		  <fes:Filter>
		   <fes:ResourceId rid="kadastralegrens.29316bf0-b87f-4e8d-bf00-21f894bdf655"/>
		  </fes:Filter>
		 </Query>
		</GetFeature>`),
			result: GetFeatureRequest{XMLName: xml.Name{Local: "GetFeature"}, StandardPresentationParameters: &StandardPresentationParameters{OutputFormat: sp("application/gml+xml; version=3.2"), Count: ip(3), Startindex: ip(0)},
				Query: Query{Filter: &Filter{
					ResourceID: &ResourceIDs{{Rid: "kadastralegrens.29316bf0-b87f-4e8d-bf00-21f894bdf655"}}}},
				BaseRequest: BaseRequest{
					Attr: []xml.Attr{
						{Name: xml.Name{Space: "xmlns", Local: "kadastralekaartv4"}, Value: "http://kadastralekaartv4.geonovum.nl"}},
					Service: "WFS",
					Version: "2.0.0"}}},
		// Get feature by PropertyIsEqualTo
		2: {body: []byte(`<?xml version="1.0" encoding="UTF-8"?>
			<GetFeature outputFormat="application/gml+xml; version=3.2" count="3" startindex="0" service="WFS" version="2.0.0" xmlns:kadastralekaartv4="http://kadastralekaartv4.geonovum.nl">
			 <Query typeNames="kadastralekaartv4:kadastralegrens" srsName="urn:ogc:def:crs:EPSG::28992">
			  <fes:Filter>
			   <fes:PropertyIsEqualTo matchCase="true">
				<fes:ValueReference>id</fes:ValueReference>
				<fes:Literal>29316bf0-b87f-4e8d-bf00-21f894bdf655</fes:Literal>
			   </fes:PropertyIsEqualTo>
			  </fes:Filter>
			 </Query>
			</GetFeature>`),
			result: GetFeatureRequest{XMLName: xml.Name{Local: "GetFeature"}, StandardPresentationParameters: &StandardPresentationParameters{OutputFormat: sp("application/gml+xml; version=3.2"), Count: ip(3), Startindex: ip(0)},
				Query: Query{Filter: &Filter{
					ComparisonOperator: ComparisonOperator{PropertyIsEqualTo: &[]PropertyIsEqualTo{{
						ComparisonOperatorAttribute: ComparisonOperatorAttribute{MatchCase: sp("true"), ValueReference: sp("id"), Literal: "29316bf0-b87f-4e8d-bf00-21f894bdf655"},
					}}},
				}},
				BaseRequest: BaseRequest{
					Attr: []xml.Attr{
						{Name: xml.Name{Space: "xmlns", Local: "kadastralekaartv4"}, Value: "http://kadastralekaartv4.geonovum.nl"}},
					Service: "WFS",
					Version: "2.0.0"}}},
		// Not a XML document
		3: {body: []byte(`GetFeature`),
			result:    GetFeatureRequest{XMLName: xml.Name{Local: "GetFeature"}, BaseRequest: BaseRequest{Service: "WFS", Version: "2.0.0"}},
			exception: wsc110.NoApplicableCode("Could not process XML, is it XML?").ToExceptions(),
		},
		// No document
		4: {result: GetFeatureRequest{XMLName: xml.Name{Local: "GetFeature"}, BaseRequest: BaseRequest{Service: "WFS", Version: "2.0.0"}},
			exception: wsc110.NoApplicableCode("Could not process XML, is it XML?").ToExceptions(),
		},
	}

	for k, test := range tests {
		var gf GetFeatureRequest
		exception := gf.ParseXML(test.body)
		if exception != nil {
			if exception[0].Error() != test.exception[0].Error() {
				t.Errorf("test: %d, expected: %s,\n got: %s", k, test.exception, exception)
			}
		} else {
			if gf.BaseRequest.Service != test.result.BaseRequest.Service {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, test.result.Service, gf.BaseRequest.Service)
			}
			if gf.BaseRequest.Version != test.result.BaseRequest.Version {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, test.result.Version, gf.Version)
			}
			if gf.Query.Filter != nil {
				if gf.Query.Filter.ResourceID != nil {
					var r, e []ResourceID
					r = *gf.Query.Filter.ResourceID
					e = *test.result.Query.Filter.ResourceID
					if r[0] != e[0] {
						t.Errorf("test: %d, expected: %s ,\n got: %s", k, e, r)
					}
				}
				if gf.Query.Filter.PropertyIsEqualTo != nil {
					var r, e []PropertyIsEqualTo
					r = *gf.Query.Filter.PropertyIsEqualTo
					e = *test.result.Query.Filter.PropertyIsEqualTo
					if *r[0].ValueReference != *e[0].ValueReference {
						t.Errorf("test: %d, expected: %+v ,\n got: %+v", k, *e[0].ValueReference, *r[0].ValueReference)
					}
				}
			}
			if len(test.result.BaseRequest.Attr) == len(gf.BaseRequest.Attr) {
				c := false
				for _, expected := range test.result.BaseRequest.Attr {
					for _, result := range gf.BaseRequest.Attr {
						if result.Name.Local == expected.Name.Local && result.Value == expected.Value {
							c = true
						}
					}
					if !c {
						t.Errorf("test: %d, expected: %s ,\n got: %s", k, test.result.BaseRequest.Attr, gf.BaseRequest.Attr)
					}
					c = false
				}
			} else {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, test.result.BaseRequest.Attr, gf.BaseRequest.Attr)
			}
		}
	}
}

func TestProcesNamespaces(t *testing.T) {
	var tests = []struct {
		namespace string
		result    []xml.Attr
	}{ // Two namespaces
		0: {namespace: "xmlns(ns1,http://www.someserver.com/ns1),xmlns(ns2,http://someserver.com/ns2)",
			result: []xml.Attr{{Name: xml.Name{Local: "ns1"}, Value: "http://www.someserver.com/ns1"},
				{Name: xml.Name{Local: "ns2"}, Value: "http://someserver.com/ns2"}}},
		// Random string
		1: {namespace: "randomstring",
			result: []xml.Attr{}},
		// Empty string
		2: {namespace: "",
			result: []xml.Attr{}},
		// Duplicate namespace with the same value
		3: {namespace: "xmlns(ns1,http://www.someserver.com/ns1),xmlns(ns1,http://www.someserver.com/ns1)",
			result: []xml.Attr{{Name: xml.Name{Local: "ns1"}, Value: "http://www.someserver.com/ns1"}}},
		// Duplicate namespace with the different values
		4: {namespace: "xmlns(ns1,http://www.someserver.com/ns1),xmlns(ns1,http://someserver.com/ns2)",
			result: []xml.Attr{{Name: xml.Name{Local: "ns1"}, Value: "http://someserver.com/ns2"}}}, //takes the last matched result
		// A namespace with a trailing string outside the xmlns()
		5: {namespace: "xmlns(ns1,http://www.someserver.com/ns1),not a correct,namespace query,string",
			result: []xml.Attr{{Name: xml.Name{Local: "ns1"}, Value: "http://www.someserver.com/ns1"}}},
		// A namespace with a trailing string inside the xmlns()
		6: {namespace: "xmlns(ns1,http://www.someserver.com/ns1,not a correct,namespace query,string)",
			result: []xml.Attr{{Name: xml.Name{Local: "ns1"}, Value: "http://www.someserver.com/ns1"}}},
	}

	for k, test := range tests {
		results := procesNamespaces(test.namespace)
		c := false
		for _, expected := range test.result {
			for _, result := range results {
				if result.Name.Local == expected.Name.Local && result.Value == expected.Value {
					c = true
				}
			}
			if !c {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, test.result, results)
			}
			c = false
		}
	}
}

func TestGetFeatureParseKVP(t *testing.T) {

	var tests = []struct {
		queryParams url.Values
		result      GetFeatureRequest
		exception   wsc110.Exception
	}{ // Standaard getfeature request with count
		0: {queryParams: map[string][]string{REQUEST: {getfeature}, SERVICE: {Service}, VERSION: {Version}, OUTPUTFORMAT: {"application/xml"}, TYPENAMES: {"dummy"}, COUNT: {"3"}},
			result: GetFeatureRequest{XMLName: xml.Name{Local: getfeature}, StandardPresentationParameters: &StandardPresentationParameters{OutputFormat: sp("application/xml"), Count: ip(3)}, BaseRequest: BaseRequest{Service: Service, Version: Version}, Query: Query{TypeNames: "dummy"}}},
		// Invalid getfeature request: missing REQUEST, SERVICE, VERSION
		// But object should still build
		1: {queryParams: map[string][]string{OUTPUTFORMAT: {"application/xml"}, TYPENAMES: {"dummy"}, COUNT: {"3"}, VERSION: {Version}},
			result: GetFeatureRequest{XMLName: xml.Name{Local: getfeature}, BaseRequest: BaseRequest{Version: Version, Service: Service}, StandardPresentationParameters: &StandardPresentationParameters{OutputFormat: sp("application/xml"), Count: ip(3)}, Query: Query{TypeNames: "dummy"}}},
		// Namespacesn
		2: {queryParams: map[string][]string{OUTPUTFORMAT: {"application/xml"}, TYPENAMES: {"dummy"}, COUNT: {"3"}, NAMESPACES: {"xmlns(ns1,http://www.someserver.com/ns1),xmlns(ns2,http://someserver.com/ns2)"}, VERSION: {Version}},
			result: GetFeatureRequest{StandardPresentationParameters: &StandardPresentationParameters{OutputFormat: sp("application/xml"), Count: ip(3)}, XMLName: xml.Name{Local: getfeature},
				BaseRequest: BaseRequest{
					Version: Version,
					Service: Service,
					Attr: []xml.Attr{
						{Name: xml.Name{Space: "xmlns", Local: "ns1"}, Value: "http://www.someserver.com/ns1"},
						{Name: xml.Name{Space: "xmlns", Local: "ns2"}, Value: "http://someserver.com/ns2"}},
				},
				Query: Query{TypeNames: "dummy"}}},
		// Startindex & resulttype
		3: {queryParams: map[string][]string{OUTPUTFORMAT: {"application/xml"}, STARTINDEX: {"1000"}, RESULTTYPE: {"hits"}, TYPENAMES: {"dummy"}, COUNT: {"3"}, VERSION: {Version}},
			result: GetFeatureRequest{
				StandardPresentationParameters: &StandardPresentationParameters{OutputFormat: sp("application/xml"),
					Count:      ip(3),
					Startindex: ip(1000), ResultType: sp("hits"),
				},
				XMLName:     xml.Name{Local: getfeature},
				BaseRequest: BaseRequest{Service: Service, Version: Version}, Query: Query{TypeNames: "dummy"}},
		},
		4: {queryParams: map[string][]string{},
			exception: wsc110.MissingParameterValue(VERSION),
		},
		// Resourceids
		5: {queryParams: map[string][]string{RESOURCEID: {"one,two,three"}, VERSION: {Version}},
			result: GetFeatureRequest{Query: Query{Filter: &Filter{ResourceID: &ResourceIDs{{Rid: "one"}, {Rid: "two"}, {Rid: "three"}}}},
				XMLName:     xml.Name{Local: getfeature},
				BaseRequest: BaseRequest{Service: Service, Version: Version}},
		},
		// Resourceids through Filter
		6: {queryParams: map[string][]string{FILTER: {`<Filter><ResourceId rid="one"/><ResourceId rid="two"/><ResourceId rid="three"/></Filter>`}, VERSION: {Version}},
			result: GetFeatureRequest{Query: Query{Filter: &Filter{ResourceID: &ResourceIDs{{Rid: "one"}, {Rid: "two"}, {Rid: "three"}}}},
				XMLName:     xml.Name{Local: getfeature},
				BaseRequest: BaseRequest{Service: Service, Version: Version}},
		},
		// Resourceids through Filter and RESOURCEID parameter,.. should this be possible? <- NO
		7: {queryParams: map[string][]string{RESOURCEID: {"one,two,three"}, FILTER: {`<Filter><ResourceId rid="four"/><ResourceId rid="five"/><ResourceId rid="six"/></Filter>`}, VERSION: {Version}},
			exception: wsc110.NoApplicableCode(`Only one of the following selectionclauses can be used RESOURCEID,FILTER`),
		},
		// Resourceids through Filter
		8: {queryParams: map[string][]string{FILTER: {`<Filter><ResourceId rid="one"/><ResourceId rid="two"/><ResourceId rid="three"/></Filter>`}, SRSNAME: {"srsname"}, VERSION: {Version}},
			result: GetFeatureRequest{Query: Query{SrsName: sp("srsname"), Filter: &Filter{ResourceID: &ResourceIDs{{Rid: "one"}, {Rid: "two"}, {Rid: "three"}}}},
				XMLName:     xml.Name{Local: getfeature},
				BaseRequest: BaseRequest{Service: Service, Version: Version}},
		},
		// // Complex Filter
		// 0: {QueryParams: map[string][]string{FILTER: []string{`<Filter><OR><AND><PropertyIsLike wildcard='*' singleChar='.' escape='!'><PropertyName>NAME</PropertyName><Literal>Syd*</Literal></PropertyIsLike><PropertyIsEqualTo><PropertyName>POPULATION</PropertyName><Literal>4250065</Literal></PropertyIsEqualTo></AND><DWithin><PropertyName>Geometry</PropertyName><Point srsName="mekker"><coordinates>135.500000,34.666667</coordinates></Point><Distance units='m'>10000</Distance></DWithin></OR></Filter>`}, SRSNAME: []string{"srsname"}},
		// 	Result: GetFeatureRequest{Query: Query{SrsName: sp("srsname"), Filter: &Filter{}}, BaseRequest: BaseRequest{Version: Version}}},
		9: {queryParams: map[string][]string{BBOX: {`1,1,2,2`}, FILTER: {`<Filter><ResourceId rid="one"/><ResourceId rid="two"/><ResourceId rid="three"/></Filter>`}, SRSNAME: {"srsname"}, VERSION: {Version}},
			exception: wsc110.NoApplicableCode(`Only one of the following selectionclauses can be used FILTER,BBOX`),
		},
	}

	for k, test := range tests {
		var gf GetFeatureRequest
		if exception := gf.ParseQueryParameters(test.queryParams); exception != nil {
			if exception[0] != test.exception {
				t.Errorf("test: %d, expected: %+v ,\n got: %+v", k, test.exception, exception)
			}
		} else {
			compareGetFeatureQuery(gf, test.result, k, t)
		}
	}
}

func compareGetFeatureQuery(result, expected GetFeatureRequest, tid int, t *testing.T) {
	if result.BaseRequest.Service != expected.BaseRequest.Service || result.BaseRequest.Version != expected.BaseRequest.Version {
		t.Errorf("test: %d, expected: %+v ,\n got: %+v", tid, expected.BaseRequest, result.BaseRequest)
	}
	if expected.BaseRequest.Attr != nil {
		for _, r := range expected.BaseRequest.Attr {
			found := false
			for _, a := range result.BaseRequest.Attr {
				if r.Name.Local == a.Name.Local && r.Value == a.Value {
					found = true
				}
			}
			if !found {
				t.Errorf("test: %d, expected: %+v ,\n got: %+v", tid, expected.BaseRequest.Attr, result.BaseRequest.Attr)
			}
		}
	}

	if expected.StandardPresentationParameters != nil {
		if expected.Count != nil {
			if *result.Count != *expected.Count {
				t.Errorf("test: %d, expected: %+v ,\n got: %+v", tid, expected.Count, result.Count)
			}
		}
		if expected.OutputFormat != nil {
			if *result.OutputFormat != *expected.OutputFormat {
				t.Errorf("test: %d, expected: %+v ,\n got: %+v", tid, expected.OutputFormat, result.OutputFormat)
			}
		}
		if expected.Startindex != nil {
			if *result.Startindex != *expected.Startindex {
				t.Errorf("test: %d, expected: %+v ,\n got: %+v", tid, expected.Startindex, result.Startindex)
			}
		}
		if expected.ResultType != nil {
			if *result.ResultType != *expected.ResultType {
				t.Errorf("test: %d, expected: %+v ,\n got: %+v", tid, expected.ResultType, result.ResultType)
			}
		}
	}

	if result.XMLName.Local != expected.XMLName.Local {
		t.Errorf("test: %d, expected: %+v ,\n got: %+v", tid, expected.XMLName.Local, result.XMLName.Local)
	}
	if result.Query.TypeNames != expected.Query.TypeNames {
		t.Errorf("test: %d, expected: %+v ,\n got: %+v", tid, expected.Query.TypeNames, result.Query.TypeNames)
	}

	if expected.Query.Filter != nil {
		if expected.Query.SrsName != nil {
			if *expected.Query.SrsName != *result.Query.SrsName {
				t.Errorf("test: %d, expected: %+v ,\n got: %+v", tid, *expected.Query.SrsName, *result.Query.SrsName)
			}
		}
		if expected.Query.Filter.ResourceID != nil {
			for _, erid := range *expected.Query.Filter.ResourceID {
				found := false
				for _, rid := range *result.Query.Filter.ResourceID {
					if erid.Rid == rid.Rid {
						found = true
					}
				}
				if !found {
					t.Errorf("test: %d, expected: %+v ,\n got: %+v", tid, *expected.Query.Filter.ResourceID, *result.Query.Filter.ResourceID)
				}
			}
		}
	}
}

func TestParseQueryInnerXML(t *testing.T) {

	// TODO: this test only works for this given example!

	point := `<Point srsName="asrsname"><coordinates>135.500000,34.666667</coordinates></Point>`
	var tests = []struct {
		queryParams url.Values
		result      GetFeatureRequest
	}{
		0: {queryParams: map[string][]string{VERSION: {Version}, FILTER: {`<Filter><OR><AND><PropertyIsLike wildcard='*' singleChar='.' escape='!'><PropertyName>NAME</PropertyName><Literal>Syd*</Literal></PropertyIsLike><PropertyIsEqualTo><PropertyName>POPULATION</PropertyName><Literal>4250065</Literal></PropertyIsEqualTo></AND><DWithin><PropertyName>Geometry</PropertyName>` + point + `<Distance units='m'>10000</Distance></DWithin></OR></Filter>`}, SRSNAME: {"srsname"}},
			result: GetFeatureRequest{Query: Query{SrsName: sp("srsname"), Filter: &Filter{OR: &OR{SpatialOperator: SpatialOperator{
				DWithin: &DWithin{PropertyName: "Geometry", GeometryOperand: GeometryOperand{Point: &Point{Geometry: Geometry{SrsName: "asrsname", Content: "<coordinates>135.500000,34.666667</coordinates>"}}}, Distance: Distance{Units: "m", Text: "10000"}}}}}}, BaseRequest: BaseRequest{Version: Version}}},
	}

	for k, test := range tests {
		var gf GetFeatureRequest
		gf.ParseQueryParameters(test.queryParams)

		if test.result.Query.Filter.OR.DWithin.GeometryOperand.Point.Geometry.Content != gf.Query.Filter.OR.DWithin.GeometryOperand.Point.Geometry.Content {
			t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, test.result.Query.Filter.OR.DWithin.GeometryOperand.Point.Geometry.Content, gf.Query.Filter.OR.DWithin.GeometryOperand.Point.Geometry.Content)
		}

		if test.result.Query.Filter.OR.DWithin.GeometryOperand.Point.Geometry.SrsName != gf.Query.Filter.OR.DWithin.GeometryOperand.Point.Geometry.SrsName {
			t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, test.result.Query.Filter.OR.DWithin.GeometryOperand.Point.Geometry.SrsName, gf.Query.Filter.OR.DWithin.GeometryOperand.Point.Geometry.SrsName)
		}
	}
}

func TestMergeResourceIDGroups(t *testing.T) {
	var tests = []struct {
		inputRids  [][]ResourceID
		outputRids []ResourceID
	}{
		0: {inputRids: [][]ResourceID{{{Rid: "one"}}, {{Rid: "four"}, {Rid: "five"}}, {{Rid: "two"}, {Rid: "three"}}},
			outputRids: []ResourceID{{Rid: "one"}, {Rid: "two"}, {Rid: "three"}, {Rid: "four"}, {Rid: "five"}}},
		1: {inputRids: [][]ResourceID{{{Rid: "one"}, {Rid: "one"}}},
			outputRids: []ResourceID{{Rid: "one"}, {Rid: "one"}}},
	}

	for k, test := range tests {
		mergedRids := mergeResourceIDGroups(test.inputRids...)

		if len(mergedRids) != len(test.outputRids) {
			t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, test.outputRids, mergedRids)
		} else {
			for _, rid := range mergedRids {
				found := false
				for _, erid := range test.outputRids {
					if rid == erid {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, test.outputRids, mergedRids)
				}
			}
		}
	}
}

func TestGetFeatureBuildKVP(t *testing.T) {
	var tests = []struct {
		getfeature    GetFeatureRequest
		expectedquery url.Values
	}{
		0: {getfeature: GetFeatureRequest{XMLName: xml.Name{Local: getfeature},
			BaseRequest: BaseRequest{Service: Service, Version: Version},
			Query:       Query{TypeNames: `ns1:F1`},
		},
			expectedquery: map[string][]string{REQUEST: {getfeature}, SERVICE: {Service}, VERSION: {Version}, TYPENAMES: {`ns1:F1`}}},
		1: {getfeature: GetFeatureRequest{XMLName: xml.Name{Local: getfeature},
			BaseRequest:                    BaseRequest{Service: Service, Version: Version},
			StandardPresentationParameters: &StandardPresentationParameters{Startindex: ip(100), Count: ip(21)},
			Query: Query{
				TypeNames: `ns1:F1`,
				Filter:    &Filter{ResourceID: &ResourceIDs{{Rid: "one"}, {Rid: "two"}}}},
		},
			expectedquery: map[string][]string{REQUEST: {getfeature}, SERVICE: {Service}, VERSION: {Version}, TYPENAMES: {`ns1:F1`}, STARTINDEX: {"100"}, COUNT: {"21"},
				RESOURCEID: {`one,two`}}},
		2: {getfeature: GetFeatureRequest{XMLName: xml.Name{Local: getfeature},
			BaseRequest:                    BaseRequest{Service: Service, Version: Version},
			Query:                          Query{TypeNames: `ns1:F1`},
			StandardPresentationParameters: &StandardPresentationParameters{OutputFormat: sp("xml"), ResultType: sp("hits")}},
			expectedquery: map[string][]string{REQUEST: {getfeature}, SERVICE: {Service}, VERSION: {Version}, TYPENAMES: {`ns1:F1`}, OUTPUTFORMAT: {"xml"}, RESULTTYPE: {"hits"}},
		},
	}

	for k, test := range tests {
		result := test.getfeature.ToQueryParameters()
		if len(test.expectedquery) != len(result) {
			t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, test.expectedquery, result)
		} else {
			for _, rid := range result {
				found := false
				for _, erid := range test.expectedquery {
					if rid[0] == erid[0] {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, test.expectedquery, result)
				}
			}
		}
	}
}

func TestUnmarshalTextGeoBOXX(t *testing.T) {
	var tests = []struct {
		query     string
		expected  GEOBBOX
		exception wsc110.Exception
	}{
		0: {query: "18.54,-72.3544,18.62,-72.2564",
			expected: GEOBBOX{Envelope: Envelope{LowerCorner: wsc110.Position{18.54, -72.3544}, UpperCorner: wsc110.Position{18.62, -72.2564}}}},
		1: {query: "49.1874,-123.2778,49.3504,-122.8892,urn:ogc:def:crs:EPSG::4326",
			expected: GEOBBOX{SrsName: sp("urn:ogc:def:crs:EPSG::4326"), Envelope: Envelope{LowerCorner: wsc110.Position{49.1874, -123.2778}, UpperCorner: wsc110.Position{49.3504, -122.8892}}}},
		2: {query: "",
			exception: wsc110.MissingParameterValue(`BBOX`, ``),
		},
		3: {query: "18.54;-72.3544;18.62;-72.2564",
			exception: wsc110.MissingParameterValue(`BBOX`, `18.54;-72.3544;18.62;-72.2564`),
		},
		// Needs a beter solution
		4: {query: "error,-72.3544,18.62,-72.2564",
			exception: InvalidValue(`BBOX`)},
		5: {query: "18.54,error,18.62,-72.2564",
			exception: InvalidValue(`BBOX`)},
		6: {query: "18.54,-72.3544,error,-72.2564",
			exception: InvalidValue(`BBOX`)},
		7: {query: "18.54,-72.3544,18.62,error",
			exception: InvalidValue(`BBOX`)},
	}

	for k, test := range tests {
		var gb GEOBBOX
		exception := gb.parseString(test.query)

		if exception != nil {
			if exception[0] != test.exception {
				t.Errorf("test: %d, expected: %+v,\n got: %+v", k, test.exception, exception)
			}
		}

		if gb.Envelope != test.expected.Envelope {
			t.Errorf("test: %d, expected: %+v,\n got: %+v", k, test.expected.Envelope, gb.Envelope)
		}
		if gb.SrsName != nil {
			if *gb.SrsName != *test.expected.SrsName {
				t.Errorf("test: %d, expected: %+v,\n got: %+v", k, &test.expected.SrsName, &gb.SrsName)
			}
		}
	}
}

func TestMarshalTextGeoBOXX(t *testing.T) {
	var tests = []struct {
		bbox     GEOBBOX
		expected string
	}{
		0: {expected: "18.540000,-72.354400,18.620000,-72.256400",
			bbox: GEOBBOX{Envelope: Envelope{LowerCorner: wsc110.Position{18.54, -72.3544}, UpperCorner: wsc110.Position{18.62, -72.2564}}}},
		1: {expected: "49.187400,-123.277800,49.350400,-122.889200,urn:ogc:def:crs:EPSG::4326",
			bbox: GEOBBOX{SrsName: sp("urn:ogc:def:crs:EPSG::4326"), Envelope: Envelope{LowerCorner: wsc110.Position{49.1874, -123.2778}, UpperCorner: wsc110.Position{49.3504, -122.8892}}}},
		2: {expected: "",
			bbox: GEOBBOX{}},
		3: {expected: "",
			bbox: GEOBBOX{SrsName: sp("urn:ogc:def:crs:EPSG::4326")}},
	}

	for k, test := range tests {
		result := test.bbox.toString()
		if result != test.expected {
			t.Errorf("test: %d, expected: %s,\n got: %s", k, test.expected, result)
		}
	}
}

// ----------
// Benchmarks
// ----------

func BenchmarkGetFeatureBuildKVP(b *testing.B) {
	gf := GetFeatureRequest{Query: Query{SrsName: sp("srsname"), Filter: &Filter{SpatialOperator: SpatialOperator{BBOX: &GEOBBOX{Envelope: Envelope{LowerCorner: wsc110.Position{1, 1}, UpperCorner: wsc110.Position{2, 2}}}}, ResourceID: &ResourceIDs{{Rid: "one"}, {Rid: "two"}, {Rid: "three"}}}}, BaseRequest: BaseRequest{Version: Version}}
	for i := 0; i < b.N; i++ {
		gf.ToQueryParameters()
	}
}

func BenchmarkGetFeatureBuildXML(b *testing.B) {
	gf := GetFeatureRequest{Query: Query{SrsName: sp("srsname"), Filter: &Filter{SpatialOperator: SpatialOperator{BBOX: &GEOBBOX{Envelope: Envelope{LowerCorner: wsc110.Position{1, 1}, UpperCorner: wsc110.Position{2, 2}}}}, ResourceID: &ResourceIDs{{Rid: "one"}, {Rid: "two"}, {Rid: "three"}}}}, BaseRequest: BaseRequest{Version: Version}}
	for i := 0; i < b.N; i++ {
		gf.ToXML()
	}
}

func BenchmarkGetFeatureParseKVP(b *testing.B) {
	kvp := map[string][]string{REQUEST: {getfeature}, SERVICE: {Service}, VERSION: {Version}, OUTPUTFORMAT: {"application/xml"}, TYPENAMES: {"dummy"}, COUNT: {"3"}}

	for i := 0; i < b.N; i++ {
		gm := GetFeatureRequest{}
		gm.ParseQueryParameters(kvp)
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
