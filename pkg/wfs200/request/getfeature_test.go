package request

import (
	"encoding/xml"
	"net/url"
	"testing"

	"github.com/pdok/ogc-specifications/pkg/common"
	"github.com/pdok/ogc-specifications/pkg/wfs200/exception"
)

func sp(s string) *string {
	return &s
}

func ip(i int) *int {
	return &i
}

func TestGetFeatureType(t *testing.T) {
	dft := GetFeature{}
	if dft.Type() != `GetFeature` {
		t.Errorf("test: %d, expected: %s,\n got: %s", 0, `GetFeature`, dft.Type())
	}
}

func TestGetFeatureBuildXML(t *testing.T) {
	var tests = []struct {
		gf     GetFeature
		result string
	}{
		0: {gf: GetFeature{BaseGetFeatureRequest: BaseGetFeatureRequest{OutputFormat: sp("application/gml+xml; version=3.2"), Count: ip(3), Startindex: ip(0)}, BaseRequest: BaseRequest{Service: "WFS", Version: "2.0.0"}, Query: Query{TypeNames: "test", SrsName: sp("urn:ogc:def:crs:EPSG::28992")}},
			result: `<?xml version="1.0" encoding="UTF-8"?>
<GetFeature service="WFS" version="2.0.0" outputFormat="application/gml+xml; version=3.2" count="3" startindex="0">
 <Query typeNames="test" srsName="urn:ogc:def:crs:EPSG::28992"></Query>
</GetFeature>`},
		1: {gf: GetFeature{BaseGetFeatureRequest: BaseGetFeatureRequest{OutputFormat: sp("application/gml+xml; version=3.2"), Count: ip(3), Startindex: ip(0)}, BaseRequest: BaseRequest{
			Attr: common.XMLAttribute{
				xml.Attr{Name: xml.Name{Space: "xmlns", Local: "kadastralekaartv4"}, Value: "http://kadastralekaartv4.geonovum.nl"}},
			Service: "WFS", Version: "2.0.0"}, Query: Query{TypeNames: "test", SrsName: sp("urn:ogc:def:crs:EPSG::28992")}},
			result: `<?xml version="1.0" encoding="UTF-8"?>
<GetFeature service="WFS" version="2.0.0" xmlns:_xmlns="xmlns" _xmlns:kadastralekaartv4="http://kadastralekaartv4.geonovum.nl" outputFormat="application/gml+xml; version=3.2" count="3" startindex="0">
 <Query typeNames="test" srsName="urn:ogc:def:crs:EPSG::28992"></Query>
</GetFeature>`},
	}

	for k, v := range tests {
		body := v.gf.BuildXML()

		if string(body) != v.result {
			t.Errorf("test: %d, Expected body %s but was not \n got: %s", k, v.result, string(body))
		}
	}
}

// TODO
// Merge TestParseBodyGetFeature & TestParseQueryParameters GetFeature comporison into single func, like with WMS GetMap
func TestGetFeatureParseXML(t *testing.T) {
	var tests = []struct {
		Body      []byte
		Result    GetFeature
		Exception common.Exception
	}{
		// Get 3 features
		0: {Body: []byte(`<GetFeature outputFormat="application/gml+xml; version=3.2" count="3" startindex="0" service="WFS" version="2.0.0">
		<Query typeNames="kadastralekaart:kadastralegrens" srsName="urn:ogc:def:crs:EPSG::28992"/>
	   </GetFeature>`),
			Result: GetFeature{XMLName: xml.Name{Local: "GetFeature"}, BaseGetFeatureRequest: BaseGetFeatureRequest{OutputFormat: sp("application/gml+xml; version=3.2"), Count: ip(3), Startindex: ip(0)},
				BaseRequest: BaseRequest{Service: "WFS", Version: "2.0.0"}}},
		// Get feature by resourceid
		1: {Body: []byte(`<?xml version="1.0" encoding="UTF-8"?>
		<GetFeature outputFormat="application/gml+xml; version=3.2" count="3" startindex="0" service="WFS" version="2.0.0" xmlns:kadastralekaartv4="http://kadastralekaartv4.geonovum.nl">
		 <Query typeNames="kadastralekaartv4:kadastralegrens" srsName="urn:ogc:def:crs:EPSG::28992">
		  <fes:Filter>
		   <fes:ResourceId rid="kadastralegrens.29316bf0-b87f-4e8d-bf00-21f894bdf655"/>
		  </fes:Filter>
		 </Query>
		</GetFeature>`),
			Result: GetFeature{XMLName: xml.Name{Local: "GetFeature"}, BaseGetFeatureRequest: BaseGetFeatureRequest{OutputFormat: sp("application/gml+xml; version=3.2"), Count: ip(3), Startindex: ip(0)},
				Query: Query{Filter: &Filter{
					ResourceID: &[]ResourceID{{Rid: "kadastralegrens.29316bf0-b87f-4e8d-bf00-21f894bdf655"}}}},
				BaseRequest: BaseRequest{
					Attr: []xml.Attr{
						{Name: xml.Name{Space: "xmlns", Local: "kadastralekaartv4"}, Value: "http://kadastralekaartv4.geonovum.nl"}},
					Service: "WFS",
					Version: "2.0.0"}}},
		// Get feature by PropertyIsEqualTo
		2: {Body: []byte(`<?xml version="1.0" encoding="UTF-8"?>
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
			Result: GetFeature{XMLName: xml.Name{Local: "GetFeature"}, BaseGetFeatureRequest: BaseGetFeatureRequest{OutputFormat: sp("application/gml+xml; version=3.2"), Count: ip(3), Startindex: ip(0)},
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
		3: {Body: []byte(`GetFeature`),
			Result:    GetFeature{XMLName: xml.Name{Local: "GetFeature"}, BaseRequest: BaseRequest{Service: "WFS", Version: "2.0.0"}},
			Exception: common.NoApplicableCode("Could not process XML, is it XML?"),
		},
		// No document
		4: {Result: GetFeature{XMLName: xml.Name{Local: "GetFeature"}, BaseRequest: BaseRequest{Service: "WFS", Version: "2.0.0"}},
			Exception: common.NoApplicableCode("Could not process XML, is it XML?"),
		},
	}

	for k, n := range tests {
		var gf GetFeature
		err := gf.ParseXML(n.Body)
		if err != nil {
			if err.Error() != n.Exception.Error() {
				t.Errorf("test: %d, expected: %s,\n got: %s", k, n.Exception, err)
			}
		} else {
			if gf.BaseRequest.Service != n.Result.BaseRequest.Service {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.Result.Service, gf.BaseRequest.Service)
			}
			if gf.BaseRequest.Version != n.Result.BaseRequest.Version {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.Result.Version, gf.Version)
			}
			if gf.Query.Filter != nil {
				if gf.Query.Filter.ResourceID != nil {
					var r, e []ResourceID
					r = *gf.Query.Filter.ResourceID
					e = *n.Result.Query.Filter.ResourceID
					if r[0] != e[0] {
						t.Errorf("test: %d, expected: %s ,\n got: %s", k, e, r)
					}
				}
				if gf.Query.Filter.PropertyIsEqualTo != nil {
					var r, e []PropertyIsEqualTo
					r = *gf.Query.Filter.PropertyIsEqualTo
					e = *n.Result.Query.Filter.PropertyIsEqualTo
					if *r[0].ValueReference != *e[0].ValueReference {
						t.Errorf("test: %d, expected: %+v ,\n got: %+v", k, *e[0].ValueReference, *r[0].ValueReference)
					}
				}
			}
			if len(n.Result.BaseRequest.Attr) == len(gf.BaseRequest.Attr) {
				c := false
				for _, expected := range n.Result.BaseRequest.Attr {
					for _, result := range gf.BaseRequest.Attr {
						if result.Name.Local == expected.Name.Local && result.Value == expected.Value {
							c = true
						}
					}
					if !c {
						t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.Result.BaseRequest.Attr, gf.BaseRequest.Attr)
					}
					c = false
				}
			} else {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.Result.BaseRequest.Attr, gf.BaseRequest.Attr)
			}
		}
	}
}

func TestProcesNamespaces(t *testing.T) {
	var tests = []struct {
		Namespace string
		Result    []xml.Attr
	}{ // Two namespaces
		0: {Namespace: "xmlns(ns1,http://www.someserver.com/ns1),xmlns(ns2,http://someserver.com/ns2)",
			Result: []xml.Attr{{Name: xml.Name{Local: "ns1"}, Value: "http://www.someserver.com/ns1"},
				{Name: xml.Name{Local: "ns2"}, Value: "http://someserver.com/ns2"}}},
		// Random string
		1: {Namespace: "randomstring",
			Result: []xml.Attr{}},
		// Empty string
		2: {Namespace: "",
			Result: []xml.Attr{}},
		// Duplicate namespace with the same value
		3: {Namespace: "xmlns(ns1,http://www.someserver.com/ns1),xmlns(ns1,http://www.someserver.com/ns1)",
			Result: []xml.Attr{{Name: xml.Name{Local: "ns1"}, Value: "http://www.someserver.com/ns1"}}},
		// Duplicate namespace with the different values
		4: {Namespace: "xmlns(ns1,http://www.someserver.com/ns1),xmlns(ns1,http://someserver.com/ns2)",
			Result: []xml.Attr{{Name: xml.Name{Local: "ns1"}, Value: "http://someserver.com/ns2"}}}, //takes the last matched result
		// A namespace with a trailing string outside the xmlns()
		5: {Namespace: "xmlns(ns1,http://www.someserver.com/ns1),not a correct,namespace query,string",
			Result: []xml.Attr{{Name: xml.Name{Local: "ns1"}, Value: "http://www.someserver.com/ns1"}}},
		// A namespace with a trailing string inside the xmlns()
		6: {Namespace: "xmlns(ns1,http://www.someserver.com/ns1,not a correct,namespace query,string)",
			Result: []xml.Attr{{Name: xml.Name{Local: "ns1"}, Value: "http://www.someserver.com/ns1"}}},
	}

	for k, n := range tests {
		results := procesNamespaces(n.Namespace)
		c := false
		for _, expected := range n.Result {
			for _, result := range results {
				if result.Name.Local == expected.Name.Local && result.Value == expected.Value {
					c = true
				}
			}
			if !c {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.Result, results)
			}
			c = false
		}
	}
}

func TestGetFeatureParseKVP(t *testing.T) {

	var tests = []struct {
		QueryParams url.Values
		Result      GetFeature
		Exception   common.Exception
	}{ // Standaard getfeature request with count
		0: {QueryParams: map[string][]string{REQUEST: {getfeature}, SERVICE: {Service}, VERSION: {Version}, OUTPUTFORMAT: {"application/xml"}, TYPENAMES: {"dummy"}, COUNT: {"3"}},
			Result: GetFeature{XMLName: xml.Name{Local: getfeature}, BaseGetFeatureRequest: BaseGetFeatureRequest{OutputFormat: sp("application/xml"), Count: ip(3)}, BaseRequest: BaseRequest{Service: Service, Version: Version}, Query: Query{TypeNames: "dummy"}}},
		// Invalid getfeature request: missing REQUEST, SERVICE, VERSION
		// But object should still build
		1: {QueryParams: map[string][]string{OUTPUTFORMAT: {"application/xml"}, TYPENAMES: {"dummy"}, COUNT: {"3"}, VERSION: {Version}},
			Result: GetFeature{BaseRequest: BaseRequest{Version: Version}, BaseGetFeatureRequest: BaseGetFeatureRequest{OutputFormat: sp("application/xml"), Count: ip(3)}, Query: Query{TypeNames: "dummy"}}},
		// Namespacesn
		2: {QueryParams: map[string][]string{OUTPUTFORMAT: {"application/xml"}, TYPENAMES: {"dummy"}, COUNT: {"3"}, NAMESPACES: {"xmlns(ns1,http://www.someserver.com/ns1),xmlns(ns2,http://someserver.com/ns2)"}, VERSION: {Version}},
			Result: GetFeature{BaseGetFeatureRequest: BaseGetFeatureRequest{OutputFormat: sp("application/xml"), Count: ip(3)}, BaseRequest: BaseRequest{
				Version: Version,
				Attr: []xml.Attr{
					{Name: xml.Name{Space: "xmlns", Local: "ns1"}, Value: "http://www.someserver.com/ns1"},
					{Name: xml.Name{Space: "xmlns", Local: "ns2"}, Value: "http://someserver.com/ns2"}},
			},
				Query: Query{TypeNames: "dummy"}}},
		// Startindex & resulttype
		3: {QueryParams: map[string][]string{OUTPUTFORMAT: {"application/xml"}, STARTINDEX: {"1000"}, RESULTTYPE: {"hits"}, TYPENAMES: {"dummy"}, COUNT: {"3"}, VERSION: {Version}},
			Result: GetFeature{BaseGetFeatureRequest: BaseGetFeatureRequest{OutputFormat: sp("application/xml"), Count: ip(3), Startindex: ip(1000), ResultType: sp("hits")}, BaseRequest: BaseRequest{Version: Version}, Query: Query{TypeNames: "dummy"}}},
		4: {QueryParams: map[string][]string{},
			Exception: common.MissingParameterValue(VERSION),
		},
		// Resourceids
		5: {QueryParams: map[string][]string{RESOURCEID: {"one,two,three"}, VERSION: {Version}},
			Result: GetFeature{Query: Query{Filter: &Filter{ResourceID: &[]ResourceID{{Rid: "one"}, {Rid: "two"}, {Rid: "three"}}}}, BaseRequest: BaseRequest{Version: Version}}},
		// Resourceids through Filter
		6: {QueryParams: map[string][]string{FILTER: {`<Filter><ResourceId rid="one"/><ResourceId rid="two"/><ResourceId rid="three"/></Filter>`}, VERSION: {Version}},
			Result: GetFeature{Query: Query{Filter: &Filter{ResourceID: &[]ResourceID{{Rid: "one"}, {Rid: "two"}, {Rid: "three"}}}}, BaseRequest: BaseRequest{Version: Version}}},
		// Resourceids through Filter and RESOURCEID parameter,.. should this be possible?
		7: {QueryParams: map[string][]string{RESOURCEID: {"one,two,three"}, FILTER: {`<Filter><ResourceId rid="four"/><ResourceId rid="five"/><ResourceId rid="six"/></Filter>`}, VERSION: {Version}},
			Result: GetFeature{Query: Query{Filter: &Filter{ResourceID: &[]ResourceID{{Rid: "one"}, {Rid: "two"}, {Rid: "three"}, {Rid: "four"}, {Rid: "five"}, {Rid: "six"}}}}, BaseRequest: BaseRequest{Version: Version}}},
		// Resourceids through Filter
		8: {QueryParams: map[string][]string{FILTER: {`<Filter><ResourceId rid="one"/><ResourceId rid="two"/><ResourceId rid="three"/></Filter>`}, SRSNAME: {"srsname"}, VERSION: {Version}},
			Result: GetFeature{Query: Query{SrsName: sp("srsname"), Filter: &Filter{ResourceID: &[]ResourceID{{Rid: "one"}, {Rid: "two"}, {Rid: "three"}}}}, BaseRequest: BaseRequest{Version: Version}}},
		// // Complex Filter
		// 0: {QueryParams: map[string][]string{FILTER: []string{`<Filter><OR><AND><PropertyIsLike wildcard='*' singleChar='.' escape='!'><PropertyName>NAME</PropertyName><Literal>Syd*</Literal></PropertyIsLike><PropertyIsEqualTo><PropertyName>POPULATION</PropertyName><Literal>4250065</Literal></PropertyIsEqualTo></AND><DWithin><PropertyName>Geometry</PropertyName><Point srsName="mekker"><coordinates>135.500000,34.666667</coordinates></Point><Distance units='m'>10000</Distance></DWithin></OR></Filter>`}, SRSNAME: []string{"srsname"}},
		// 	Result: GetFeature{Query: Query{SrsName: sp("srsname"), Filter: &Filter{}}, BaseRequest: BaseRequest{Version: Version}}},
		9: {QueryParams: map[string][]string{BBOX: {`1,1,2,2`}, FILTER: {`<Filter><ResourceId rid="one"/><ResourceId rid="two"/><ResourceId rid="three"/></Filter>`}, SRSNAME: {"srsname"}, VERSION: {Version}},
			Result: GetFeature{Query: Query{SrsName: sp("srsname"), Filter: &Filter{SpatialOperator: SpatialOperator{BBOX: &GEOBBOX{Envelope: Envelope{LowerCorner: common.Position{1, 1}, UpperCorner: common.Position{2, 2}}}}, ResourceID: &[]ResourceID{{Rid: "one"}, {Rid: "two"}, {Rid: "three"}}}}, BaseRequest: BaseRequest{Version: Version}}},
	}

	for tid, q := range tests {
		var gf GetFeature
		if err := gf.ParseKVP(q.QueryParams); err != nil {
			if err[0] != q.Exception {
				t.Errorf("test: %d, expected: %+v ,\n got: %+v", tid, q.Exception, err)
			}
		} else {
			compareGetFeatureQuery(gf, q.Result, tid, t)
		}
	}
}

func compareGetFeatureQuery(result, expected GetFeature, tid int, t *testing.T) {
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
	if result.XMLName.Local != expected.XMLName.Local {
		t.Errorf("test: %d, expected: %+v ,\n got: %+v", tid, expected.XMLName.Local, result.XMLName.Local)
	}
	if result.Query.TypeNames != expected.Query.TypeNames {
		t.Errorf("test: %d, expected: %+v ,\n got: %+v", tid, expected.Query.TypeNames, result.Query.TypeNames)
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
		QueryParams url.Values
		Result      GetFeature
	}{
		0: {QueryParams: map[string][]string{VERSION: {Version}, FILTER: {`<Filter><OR><AND><PropertyIsLike wildcard='*' singleChar='.' escape='!'><PropertyName>NAME</PropertyName><Literal>Syd*</Literal></PropertyIsLike><PropertyIsEqualTo><PropertyName>POPULATION</PropertyName><Literal>4250065</Literal></PropertyIsEqualTo></AND><DWithin><PropertyName>Geometry</PropertyName>` + point + `<Distance units='m'>10000</Distance></DWithin></OR></Filter>`}, SRSNAME: {"srsname"}},
			Result: GetFeature{Query: Query{SrsName: sp("srsname"), Filter: &Filter{OR: &OR{SpatialOperator: SpatialOperator{
				DWithin: &DWithin{PropertyName: "Geometry", GeometryOperand: GeometryOperand{Point: &Point{Geometry: Geometry{SrsName: "asrsname", Content: "<coordinates>135.500000,34.666667</coordinates>"}}}, Distance: Distance{Units: "m", Text: "10000"}}}}}}, BaseRequest: BaseRequest{Version: Version}}},
	}

	for k, q := range tests {
		var gf GetFeature
		gf.ParseKVP(q.QueryParams)

		if q.Result.Query.Filter.OR.DWithin.GeometryOperand.Point.Geometry.Content != gf.Query.Filter.OR.DWithin.GeometryOperand.Point.Geometry.Content {
			t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, q.Result.Query.Filter.OR.DWithin.GeometryOperand.Point.Geometry.Content, gf.Query.Filter.OR.DWithin.GeometryOperand.Point.Geometry.Content)
		}

		if q.Result.Query.Filter.OR.DWithin.GeometryOperand.Point.Geometry.SrsName != gf.Query.Filter.OR.DWithin.GeometryOperand.Point.Geometry.SrsName {
			t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, q.Result.Query.Filter.OR.DWithin.GeometryOperand.Point.Geometry.SrsName, gf.Query.Filter.OR.DWithin.GeometryOperand.Point.Geometry.SrsName)
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

	for k, q := range tests {
		mergedRids := mergeResourceIDGroups(q.inputRids...)

		if len(mergedRids) != len(q.outputRids) {
			t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, q.outputRids, mergedRids)
		} else {
			for _, rid := range mergedRids {
				found := false
				for _, erid := range q.outputRids {
					if rid == erid {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, q.outputRids, mergedRids)
				}
			}
		}
	}
}

func TestGetFeatureBuildKVP(t *testing.T) {
	var tests = []struct {
		getfeature    GetFeature
		expectedquery url.Values
	}{
		0: {getfeature: GetFeature{XMLName: xml.Name{Local: getfeature}, BaseRequest: BaseRequest{Service: Service, Version: Version}},
			expectedquery: map[string][]string{REQUEST: {getfeature}, SERVICE: {Service}, VERSION: {Version}}},
		1: {getfeature: GetFeature{XMLName: xml.Name{Local: getfeature}, BaseRequest: BaseRequest{Service: Service, Version: Version}, BaseGetFeatureRequest: BaseGetFeatureRequest{Startindex: ip(100), Count: ip(21)},
			Query: Query{Filter: &Filter{ResourceID: &[]ResourceID{{Rid: "one"}, {Rid: "two"}}}}},
			expectedquery: map[string][]string{REQUEST: {getfeature}, SERVICE: {Service}, VERSION: {Version}, STARTINDEX: {"100"}, COUNT: {"21"},
				FILTER: {url.QueryEscape(`<Filter><ResourceId rid="one"></ResourceId><ResourceId rid="two"></ResourceId></Filter>`)}}},
		2: {getfeature: GetFeature{XMLName: xml.Name{Local: getfeature}, BaseRequest: BaseRequest{Service: Service, Version: Version}, BaseGetFeatureRequest: BaseGetFeatureRequest{OutputFormat: sp("xml"), ResultType: sp("hits")}},
			expectedquery: map[string][]string{REQUEST: {getfeature}, SERVICE: {Service}, VERSION: {Version}, OUTPUTFORMAT: {"xml"}, RESULTTYPE: {"hits"}}},
	}

	for k, q := range tests {
		result := q.getfeature.BuildKVP()
		if len(q.expectedquery) != len(result) {
			t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, q.expectedquery, result)
		} else {
			for _, rid := range result {
				found := false
				for _, erid := range q.expectedquery {
					if rid[0] == erid[0] {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, q.expectedquery, result)
				}
			}
		}
	}
}

func TestUnmarshalTextGeoBOXX(t *testing.T) {
	var tests = []struct {
		Query     string
		Expected  GEOBBOX
		Exception common.Exception
	}{
		0: {Query: "18.54,-72.3544,18.62,-72.2564", Expected: GEOBBOX{Envelope: Envelope{LowerCorner: common.Position{18.54, -72.3544}, UpperCorner: common.Position{18.62, -72.2564}}}},
		1: {Query: "49.1874,-123.2778,49.3504,-122.8892,urn:ogc:def:crs:EPSG::4326", Expected: GEOBBOX{SrsName: sp("urn:ogc:def:crs:EPSG::4326"), Envelope: Envelope{LowerCorner: common.Position{49.1874, -123.2778}, UpperCorner: common.Position{49.3504, -122.8892}}}},
		2: {Query: "", Expected: GEOBBOX{}},
		3: {Query: "18.54;-72.3544;18.62;-72.2564", Expected: GEOBBOX{}},
		// Needs a beter solution
		4: {Query: "error,-72.3544,18.62,-72.2564", Exception: exception.InvalidValue(`BBOX`)},
		5: {Query: "18.54,error,18.62,-72.2564", Exception: exception.InvalidValue(`BBOX`)},
		6: {Query: "18.54,-72.3544,error,-72.2564", Exception: exception.InvalidValue(`BBOX`)},
		7: {Query: "18.54,-72.3544,18.62,error", Exception: exception.InvalidValue(`BBOX`)},
	}

	for k, a := range tests {
		var gb GEOBBOX
		err := gb.UnmarshalText(a.Query)

		if err != nil {
			if err != a.Exception {
				t.Errorf("test: %d, expected: %+v,\n got: %+v", k, a.Exception, err)
			}
		}

		if gb.Envelope != a.Expected.Envelope {
			t.Errorf("test: %d, expected: %+v,\n got: %+v", k, a.Expected.Envelope, gb.Envelope)
		}
		if gb.SrsName != nil {
			if *gb.SrsName != *a.Expected.SrsName {
				t.Errorf("test: %d, expected: %+v,\n got: %+v", k, &a.Expected.SrsName, &gb.SrsName)
			}
		}
	}
}

func TestMarshalTextGeoBOXX(t *testing.T) {
	var tests = []struct {
		GeoBBox  GEOBBOX
		Expected string
	}{
		0: {Expected: "18.540000,-72.354400,18.620000,-72.256400", GeoBBox: GEOBBOX{Envelope: Envelope{LowerCorner: common.Position{18.54, -72.3544}, UpperCorner: common.Position{18.62, -72.2564}}}},
		1: {Expected: "49.187400,-123.277800,49.350400,-122.889200,urn:ogc:def:crs:EPSG::4326", GeoBBox: GEOBBOX{SrsName: sp("urn:ogc:def:crs:EPSG::4326"), Envelope: Envelope{LowerCorner: common.Position{49.1874, -123.2778}, UpperCorner: common.Position{49.3504, -122.8892}}}},
		2: {Expected: "", GeoBBox: GEOBBOX{}},
		3: {Expected: "", GeoBBox: GEOBBOX{SrsName: sp("urn:ogc:def:crs:EPSG::4326")}},
	}

	for k, a := range tests {
		result := a.GeoBBox.MarshalText()
		if result != a.Expected {
			t.Errorf("test: %d, expected: %s,\n got: %s", k, a.Expected, result)
		}
	}
}

// ----------
// Benchmarks
// ----------

func BenchmarkGetFeatureBuildKVP(b *testing.B) {
	gf := GetFeature{Query: Query{SrsName: sp("srsname"), Filter: &Filter{SpatialOperator: SpatialOperator{BBOX: &GEOBBOX{Envelope: Envelope{LowerCorner: common.Position{1, 1}, UpperCorner: common.Position{2, 2}}}}, ResourceID: &[]ResourceID{{Rid: "one"}, {Rid: "two"}, {Rid: "three"}}}}, BaseRequest: BaseRequest{Version: Version}}
	for i := 0; i < b.N; i++ {
		gf.BuildKVP()
	}
}

func BenchmarkGetFeatureBuildXML(b *testing.B) {
	gf := GetFeature{Query: Query{SrsName: sp("srsname"), Filter: &Filter{SpatialOperator: SpatialOperator{BBOX: &GEOBBOX{Envelope: Envelope{LowerCorner: common.Position{1, 1}, UpperCorner: common.Position{2, 2}}}}, ResourceID: &[]ResourceID{{Rid: "one"}, {Rid: "two"}, {Rid: "three"}}}}, BaseRequest: BaseRequest{Version: Version}}
	for i := 0; i < b.N; i++ {
		gf.BuildXML()
	}
}

func BenchmarkGetFeatureParseKVP(b *testing.B) {
	kvp := map[string][]string{REQUEST: {getfeature}, SERVICE: {Service}, VERSION: {Version}, OUTPUTFORMAT: {"application/xml"}, TYPENAMES: {"dummy"}, COUNT: {"3"}}

	for i := 0; i < b.N; i++ {
		gm := GetFeature{}
		gm.ParseKVP(kvp)
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
		gm := GetFeature{}
		gm.ParseXML(doc)
	}
}
