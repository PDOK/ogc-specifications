package wfs200

import (
	"encoding/xml"
	"net/url"
	"testing"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

func TestDescribeFeatureTypeType(t *testing.T) {
	dft := DescribeFeatureTypeRequest{}
	if dft.Type() != `DescribeFeatureType` {
		t.Errorf("test: %d, expected: %s,\n got: %s", 0, `DescribeFeatureType`, dft.Type())
	}
}

func TestDescribeFeatureTypeParseXML(t *testing.T) {
	var tests = []struct {
		body      []byte
		result    DescribeFeatureTypeRequest
		exception wsc110.Exceptions
	}{
		// Lots of attribute declarations
		0: {body: []byte(`<DescribeFeatureType service="wfs" version="2.0.0" xmlns:gml="http://www.opengis.net/gml/3.2" xmlns:wfs="http://www.opengis.net/wfs/2.0" xmlns:ows="http://www.opengis.net/ows/1.1" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:fes="http://www.opengis.net/fes/2.0" xmlns:inspire_common="http://inspire.ec.europa.eu/schemas/common/1.0" xmlns:inspire_dls="http://inspire.ec.europa.eu/schemas/inspire_dls/1.0" xmlns:kadastralekaartv4="http://kadastralekaartv4.geonovum.nl" xsi:schemaLocation="http://www.opengis.net/wfs/2.0 http://schemas.opengis.net/wfs/2.0/wfs.xsd http://inspire.ec.europa.eu/schemas/inspire_dls/1.0 http://inspire.ec.europa.eu/schemas/inspire_dls/1.0/inspire_dls.xsd http://inspire.ec.europa.eu/schemas/common/1.0 http://inspire.ec.europa.eu/schemas/common/1.0/common.xsd"/>`),
			result: DescribeFeatureTypeRequest{XMLName: xml.Name{Local: describefeaturetype}, BaseRequest: BaseRequest{Service: "wfs", Version: "2.0.0",
				Attr: []xml.Attr{{Name: xml.Name{Space: "xmlns", Local: "gml"}, Value: "http://www.opengis.net/gml/3.2"},
					{Name: xml.Name{Space: "xmlns", Local: "wfs"}, Value: "http://www.opengis.net/wfs/2.0"},
					{Name: xml.Name{Space: "xmlns", Local: "ows"}, Value: "http://www.opengis.net/ows/1.1"},
					{Name: xml.Name{Space: "xmlns", Local: "xlink"}, Value: "http://www.w3.org/1999/xlink"},
					{Name: xml.Name{Space: "xmlns", Local: "xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
					{Name: xml.Name{Space: "xmlns", Local: "fes"}, Value: "http://www.opengis.net/fes/2.0"},
					{Name: xml.Name{Space: "xmlns", Local: "inspire_common"}, Value: "http://inspire.ec.europa.eu/schemas/common/1.0"},
					{Name: xml.Name{Space: "xmlns", Local: "inspire_dls"}, Value: "http://inspire.ec.europa.eu/schemas/inspire_dls/1.0"},
					{Name: xml.Name{Space: "xmlns", Local: "kadastralekaartv4"}, Value: "http://kadastralekaartv4.geonovum.nl"},
					{Name: xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "schemaLocation"}, Value: "http://www.opengis.net/wfs/2.0 http://schemas.opengis.net/wfs/2.0/wfs.xsd http://inspire.ec.europa.eu/schemas/inspire_dls/1.0 http://inspire.ec.europa.eu/schemas/inspire_dls/1.0/inspire_dls.xsd http://inspire.ec.europa.eu/schemas/common/1.0 http://inspire.ec.europa.eu/schemas/common/1.0/common.xsd"}}}}},
		// Unknown XML document
		1: {body: []byte("<Unknown/>"),
			exception: wsc110.Exception{ExceptionText: "This service does not know the operation: expected element type <DescribeFeatureType> but have <Unknown>"}.ToExceptions()},
		// no XML document
		2: {body: []byte("no XML document, just a string"),
			exception: wsc110.Exception{ExceptionText: "Could not process XML, is it XML?"}.ToExceptions()},
		// document at all
		3: {exception: wsc110.Exception{ExceptionText: "Could not process XML, is it XML?"}.ToExceptions()},
		// Duplicate attributes in XML message with the same value
		4: {body: []byte(`<DescribeFeatureType service="wfs" version="2.0.0" xmlns:wfs="http://www.opengis.net/wfs/2.0"  xmlns:wfs="http://www.opengis.net/wfs/2.0" xmlns:wfs="http://www.opengis.net/wfs/2.0"/>`),
			result: DescribeFeatureTypeRequest{XMLName: xml.Name{Local: describefeaturetype}, BaseRequest: BaseRequest{Service: "wfs", Version: "2.0.0",
				Attr: []xml.Attr{{Name: xml.Name{Space: "xmlns", Local: "wfs"}, Value: "http://www.opengis.net/wfs/2.0"}}}}},
		// Duplicate attributes in XML message with different values
		5: {body: []byte(`<DescribeFeatureType service="wfs" version="2.0.0" xmlns:wfs="http://www.opengis.net/ows/1.1"  xmlns:wfs="http://www.w3.org/2001/XMLSchema-instance" xmlns:wfs="http://www.opengis.net/wfs/2.0"/>`),
			result: DescribeFeatureTypeRequest{XMLName: xml.Name{Local: describefeaturetype}, BaseRequest: BaseRequest{Service: "wfs", Version: "2.0.0",
				Attr: []xml.Attr{{Name: xml.Name{Space: "xmlns", Local: "wfs"}, Value: "http://www.opengis.net/wfs/2.0"}}}}},
		6: {body: []byte(`<DescribeFeatureType service="wfs" version="2.0.0" typeName="acme:anvils"/>`),
			result: DescribeFeatureTypeRequest{XMLName: xml.Name{Local: describefeaturetype},
				BaseDescribeFeatureTypeRequest: BaseDescribeFeatureTypeRequest{TypeName: sp("acme:anvils")},
				BaseRequest:                    BaseRequest{Service: "wfs", Version: "2.0.0"}}},
	}

	for k, n := range tests {
		var dft DescribeFeatureTypeRequest
		exception := dft.ParseXML(n.body)
		if exception != nil {
			if n.exception != nil {
				if exception[0].Error() != n.exception[0].Error() {
					t.Errorf("test: %d, expected: %s,\n got: %s", k, n.exception, exception)
				}
			} else {
				t.Errorf("test: %d, expected NO exception,\n got: %s", k, exception)
			}

		} else {
			if dft.BaseRequest.Service != n.result.BaseRequest.Service {
				t.Errorf("test: %d, expected: %+v ,\n got: %+v", k, n.result, dft)
			}
			if dft.BaseRequest.Version != n.result.BaseRequest.Version {
				t.Errorf("test: %d, expected: %+v ,\n got: %+v", k, n.result, dft)
			}
			if dft.BaseDescribeFeatureTypeRequest.TypeName != nil {
				if *dft.BaseDescribeFeatureTypeRequest.TypeName != *n.result.BaseDescribeFeatureTypeRequest.TypeName {
					t.Errorf("test: %d, expected: %+v ,\n got: %+v", k, *n.result.BaseDescribeFeatureTypeRequest.TypeName, *dft.BaseDescribeFeatureTypeRequest.TypeName)
				}
			}
			if len(n.result.BaseRequest.Attr) == len(dft.BaseRequest.Attr) {
				c := false
				for _, expected := range n.result.BaseRequest.Attr {
					for _, result := range dft.BaseRequest.Attr {
						if result.Name.Local == expected.Name.Local && result.Value == expected.Value {
							c = true
						}
					}
					if !c {
						t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.result.BaseRequest.Attr, dft.BaseRequest.Attr)
					}
					c = false
				}
			} else {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.result.BaseRequest.Attr, dft.BaseRequest.Attr)
			}
		}
	}
}

func TestDescribeFeatureTypeParseKVP(t *testing.T) {
	var tests = []struct {
		query     url.Values
		result    DescribeFeatureTypeRequest
		exception wsc110.Exceptions
	}{
		// "Normal" query request with UPPER/lower/MiXeD case
		0: {query: map[string][]string{"SERVICE": {Service}, "Request": {describefeaturetype}, "version": {"2.0.0"}},
			result: DescribeFeatureTypeRequest{XMLName: xml.Name{Local: describefeaturetype}, BaseRequest: BaseRequest{Service: "WFS", Version: "2.0.0"}}},
		// Missing mandatory SERVICE attribute
		1: {query: map[string][]string{"Request": {describefeaturetype}},
			exception: wsc110.Exceptions{wsc110.MissingParameterValue(VERSION)}},
		// Missing optional VERSION attribute
		2: {query: map[string][]string{"SERVICE": {"WFS"}, "Request": {describefeaturetype}, "Version": {"2.0.0"}},
			result: DescribeFeatureTypeRequest{XMLName: xml.Name{Local: describefeaturetype}, BaseRequest: BaseRequest{Service: "WFS", Version: Version}}},
		// Unknown optional VERSION attribute
		3: {query: map[string][]string{"SERVICE": {"WFS"}, "Request": {describefeaturetype}, "version": {"no version supplied"}},
			result: DescribeFeatureTypeRequest{XMLName: xml.Name{Local: describefeaturetype}, BaseRequest: BaseRequest{Service: "WFS", Version: "no version supplied"}}},
		// Not configured optional VERSION attribute
		4: {query: map[string][]string{"SERVICE": {"WFS"}, "Request": {describefeaturetype}, "version": {"1.1.0"}},
			result: DescribeFeatureTypeRequest{XMLName: xml.Name{Local: describefeaturetype}, BaseRequest: BaseRequest{Service: "WFS", Version: "1.1.0"}}},
		5: {query: map[string][]string{VERSION: {Version}, SERVICE: {Service}, REQUEST: {describefeaturetype}, TYPENAME: {"acme:anvils"}},
			result: DescribeFeatureTypeRequest{XMLName: xml.Name{Local: describefeaturetype},
				BaseDescribeFeatureTypeRequest: BaseDescribeFeatureTypeRequest{TypeName: sp("acme:anvils")},
				BaseRequest:                    BaseRequest{Service: Service, Version: Version}}},
		6: {query: map[string][]string{},
			exception: wsc110.Exceptions{wsc110.MissingParameterValue(VERSION)},
		},
	}

	for k, n := range tests {
		var dft DescribeFeatureTypeRequest
		exception := dft.ParseKVP(n.query)
		if exception != nil {
			if exception[0].Error() != n.exception[0].Error() {
				t.Errorf("test: %d, expected: %s,\n got: %s", k, n.exception, exception)
			}
		} else {
			if n.result.XMLName.Local != dft.XMLName.Local {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.result.XMLName.Local, dft.XMLName.Local)
			}
			if n.result.BaseRequest.Service != dft.BaseRequest.Service {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.result.BaseRequest.Service, dft.BaseRequest.Service)
			}
			if n.result.BaseRequest.Version != dft.BaseRequest.Version {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.result.BaseRequest.Version, dft.BaseRequest.Version)
			}
			if dft.BaseDescribeFeatureTypeRequest.TypeName != nil {
				if *dft.BaseDescribeFeatureTypeRequest.TypeName != *n.result.BaseDescribeFeatureTypeRequest.TypeName {
					t.Errorf("test: %d, expected: %+v ,\n got: %+v", k, *n.result.BaseDescribeFeatureTypeRequest.TypeName, *dft.BaseDescribeFeatureTypeRequest.TypeName)
				}
			}
		}
	}
}
func TestDescribeFeatureTypeBuildKVP(t *testing.T) {
	var tests = []struct {
		dft   DescribeFeatureTypeRequest
		query url.Values
	}{
		0: {dft: DescribeFeatureTypeRequest{
			XMLName:     xml.Name{Local: `DescribeFeatureType`},
			BaseRequest: BaseRequest{Version: Version, Service: Service},
			BaseDescribeFeatureTypeRequest: BaseDescribeFeatureTypeRequest{
				OutputFormat: sp(`application/json`),
				TypeName:     sp(`example:example`)}},
			query: map[string][]string{"OUTPUTFORMAT": {"application/json"}, "VERSION": {`2.0.0`}, "REQUEST": {"DescribeFeatureType"}, "SERVICE": {"WFS"}, "TYPENAME": {"example:example"}},
		},
	}

	for k, v := range tests {
		values := v.dft.BuildKVP()
		c := false
		for _, value := range values {
			for _, q := range v.query {
				if value[0] == q[0] {
					c = true
				}
			}
			if !c {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, v.query, values)
			}
			c = false
		}
	}
}

func TestDescribeFeatureTypeBuildXML(t *testing.T) {
	var tests = []struct {
		dft  DescribeFeatureTypeRequest
		body string
	}{
		0: {dft: DescribeFeatureTypeRequest{
			XMLName:     xml.Name{Local: `DescribeFeatureType`},
			BaseRequest: BaseRequest{Version: Version, Service: Service},
			BaseDescribeFeatureTypeRequest: BaseDescribeFeatureTypeRequest{
				OutputFormat: sp(`application/json`),
				TypeName:     sp(`example:example`)}},
			body: `<?xml version="1.0" encoding="UTF-8"?>
<DescribeFeatureType service="WFS" version="2.0.0" outputFormat="application/json" typeNames="example:example"/>`,
		},
	}
	for k, v := range tests {
		b := string(v.dft.BuildXML())
		if b != v.body {
			t.Errorf("test: %d, expected: %s ,\n got: %s", k, v.body, b)
		}
	}
}

// ----------
// Benchmarks
// ----------

func BenchmarkDescribeFeatureTypeBuildKVP(b *testing.B) {
	df := DescribeFeatureTypeRequest{
		XMLName:     xml.Name{Local: `DescribeFeatureType`},
		BaseRequest: BaseRequest{Version: Version, Service: Service},
		BaseDescribeFeatureTypeRequest: BaseDescribeFeatureTypeRequest{
			OutputFormat: sp(`application/json`),
			TypeName:     sp(`example:example`)}}
	for i := 0; i < b.N; i++ {
		df.BuildKVP()
	}
}

func BenchmarkDescribeFeatureTypeBuildXML(b *testing.B) {
	df := DescribeFeatureTypeRequest{
		XMLName:     xml.Name{Local: `DescribeFeatureType`},
		BaseRequest: BaseRequest{Version: Version, Service: Service},
		BaseDescribeFeatureTypeRequest: BaseDescribeFeatureTypeRequest{
			OutputFormat: sp(`application/json`),
			TypeName:     sp(`example:example`)}}
	for i := 0; i < b.N; i++ {
		df.BuildXML()
	}
}
