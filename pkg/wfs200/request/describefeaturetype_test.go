package request

import (
	"encoding/xml"
	"net/url"
	"testing"

	"github.com/pdok/ogc-specifications/pkg/ows"
	"github.com/pdok/ogc-specifications/pkg/wfs200/exception"
)

func TestDescribeFeatureType(t *testing.T) {
	dft := DescribeFeatureType{}
	if dft.Type() != `DescribeFeatureType` {
		t.Errorf("test: %d, expected: %s,\n got: %s", 0, `DescribeFeatureType`, dft.Type())
	}
}

func TestParseBodyDescribeFeatureType(t *testing.T) {
	var tests = []struct {
		Body   []byte
		Result DescribeFeatureType
		Error  ows.Exception
	}{
		// Lots of attribute declarations
		0: {Body: []byte(`<DescribeFeatureType service="wfs" version="2.0.0" xmlns:gml="http://www.opengis.net/gml/3.2" xmlns:wfs="http://www.opengis.net/wfs/2.0" xmlns:ows="http://www.opengis.net/ows/1.1" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:fes="http://www.opengis.net/fes/2.0" xmlns:inspire_common="http://inspire.ec.europa.eu/schemas/common/1.0" xmlns:inspire_dls="http://inspire.ec.europa.eu/schemas/inspire_dls/1.0" xmlns:kadastralekaartv4="http://kadastralekaartv4.geonovum.nl" xsi:schemaLocation="http://www.opengis.net/wfs/2.0 http://schemas.opengis.net/wfs/2.0/wfs.xsd http://inspire.ec.europa.eu/schemas/inspire_dls/1.0 http://inspire.ec.europa.eu/schemas/inspire_dls/1.0/inspire_dls.xsd http://inspire.ec.europa.eu/schemas/common/1.0 http://inspire.ec.europa.eu/schemas/common/1.0/common.xsd"/>`),
			Result: DescribeFeatureType{XMLName: xml.Name{Local: describefeaturetype}, BaseRequest: BaseRequest{Service: "wfs", Version: "2.0.0",
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
		1: {Body: []byte("<Unknown/>"), Error: &exception.WFSException{ExceptionText: "This service does not know the operation: expected element type <DescribeFeatureType> but have <Unknown>"}},
		// no XML document
		2: {Body: []byte("no XML document, just a string"), Error: &exception.WFSException{ExceptionText: "Could not process XML, is it XML?"}},
		// document at all
		3: {Error: &exception.WFSException{ExceptionText: "Could not process XML, is it XML?"}},
		// Duplicate attributes in XML message with the same value
		4: {Body: []byte(`<DescribeFeatureType service="wfs" version="2.0.0" xmlns:wfs="http://www.opengis.net/wfs/2.0"  xmlns:wfs="http://www.opengis.net/wfs/2.0" xmlns:wfs="http://www.opengis.net/wfs/2.0"/>`),
			Result: DescribeFeatureType{XMLName: xml.Name{Local: describefeaturetype}, BaseRequest: BaseRequest{Service: "wfs", Version: "2.0.0",
				Attr: []xml.Attr{{Name: xml.Name{Space: "xmlns", Local: "wfs"}, Value: "http://www.opengis.net/wfs/2.0"}}}}},
		// Duplicate attributes in XML message with different values
		5: {Body: []byte(`<DescribeFeatureType service="wfs" version="2.0.0" xmlns:wfs="http://www.opengis.net/ows/1.1"  xmlns:wfs="http://www.w3.org/2001/XMLSchema-instance" xmlns:wfs="http://www.opengis.net/wfs/2.0"/>`),
			Result: DescribeFeatureType{XMLName: xml.Name{Local: describefeaturetype}, BaseRequest: BaseRequest{Service: "wfs", Version: "2.0.0",
				Attr: []xml.Attr{{Name: xml.Name{Space: "xmlns", Local: "wfs"}, Value: "http://www.opengis.net/wfs/2.0"}}}}},
		6: {Body: []byte(`<DescribeFeatureType service="wfs" version="2.0.0" typeName="acme:anvils"/>`),
			Result: DescribeFeatureType{XMLName: xml.Name{Local: describefeaturetype},
				BaseDescribeFeatureTypeRequest: BaseDescribeFeatureTypeRequest{TypeName: sp("acme:anvils")},
				BaseRequest:                    BaseRequest{Service: "wfs", Version: "2.0.0"}}},
	}

	for k, n := range tests {
		var dft DescribeFeatureType
		err := dft.ParseXML(n.Body)
		if err != nil {
			if n.Error != nil {
				if err.Error() != n.Error.Error() {
					t.Errorf("test: %d, expected: %s,\n got: %s", k, n.Error, err)
				}
			} else {
				t.Errorf("test: %d, expected NO error,\n got: %s", k, err)
			}

		} else {
			if dft.BaseRequest.Service != n.Result.BaseRequest.Service {
				t.Errorf("test: %d, expected: %+v ,\n got: %+v", k, n.Result, dft)
			}
			if dft.BaseRequest.Version != n.Result.BaseRequest.Version {
				t.Errorf("test: %d, expected: %+v ,\n got: %+v", k, n.Result, dft)
			}
			if dft.BaseDescribeFeatureTypeRequest.TypeName != nil {
				if *dft.BaseDescribeFeatureTypeRequest.TypeName != *n.Result.BaseDescribeFeatureTypeRequest.TypeName {
					t.Errorf("test: %d, expected: %+v ,\n got: %+v", k, *n.Result.BaseDescribeFeatureTypeRequest.TypeName, *dft.BaseDescribeFeatureTypeRequest.TypeName)
				}
			}
			if len(n.Result.BaseRequest.Attr) == len(dft.BaseRequest.Attr) {
				c := false
				for _, expected := range n.Result.BaseRequest.Attr {
					for _, result := range dft.BaseRequest.Attr {
						if result.Name.Local == expected.Name.Local && result.Value == expected.Value {
							c = true
						}
					}
					if !c {
						t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.Result.BaseRequest.Attr, dft.BaseRequest.Attr)
					}
					c = false
				}
			} else {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.Result.BaseRequest.Attr, dft.BaseRequest.Attr)
			}
		}
	}
}

func TestParseQueryParametersDescribeFeatureType(t *testing.T) {
	var tests = []struct {
		Query     url.Values
		Result    DescribeFeatureType
		Exception ows.Exception
	}{
		// "Normal" query request with UPPER/lower/MiXeD case
		0: {Query: map[string][]string{"SERVICE": {Service}, "Request": {describefeaturetype}, "version": {"2.0.0"}},
			Result: DescribeFeatureType{XMLName: xml.Name{Local: describefeaturetype}, BaseRequest: BaseRequest{Service: "WFS", Version: "2.0.0"}}},
		// Missing mandatory SERVICE attribute
		1: {Query: map[string][]string{"Request": {describefeaturetype}},
			Exception: ows.MissingParameterValue(VERSION)},
		// Missing optional VERSION attribute
		2: {Query: map[string][]string{"SERVICE": {"WFS"}, "Request": {describefeaturetype}, "Version": {"2.0.0"}},
			Result: DescribeFeatureType{XMLName: xml.Name{Local: describefeaturetype}, BaseRequest: BaseRequest{Service: "WFS", Version: Version}}},
		// Unknown optional VERSION attribute
		3: {Query: map[string][]string{"SERVICE": {"WFS"}, "Request": {describefeaturetype}, "version": {"no version supplied"}},
			Result: DescribeFeatureType{XMLName: xml.Name{Local: describefeaturetype}, BaseRequest: BaseRequest{Service: "WFS", Version: "no version supplied"}}},
		// Not configured optional VERSION attribute
		4: {Query: map[string][]string{"SERVICE": {"WFS"}, "Request": {describefeaturetype}, "version": {"1.1.0"}},
			Result: DescribeFeatureType{XMLName: xml.Name{Local: describefeaturetype}, BaseRequest: BaseRequest{Service: "WFS", Version: "1.1.0"}}},
		5: {Query: map[string][]string{VERSION: {Version}, SERVICE: {Service}, REQUEST: {describefeaturetype}, TYPENAME: {"acme:anvils"}},
			Result: DescribeFeatureType{XMLName: xml.Name{Local: describefeaturetype},
				BaseDescribeFeatureTypeRequest: BaseDescribeFeatureTypeRequest{TypeName: sp("acme:anvils")},
				BaseRequest:                    BaseRequest{Service: Service, Version: Version}}},
		6: {Query: map[string][]string{},
			Exception: ows.MissingParameterValue(VERSION),
		},
	}

	for k, n := range tests {
		var dft DescribeFeatureType
		err := dft.ParseQuery(n.Query)
		if err != nil {
			if err.Error() != n.Exception.Error() {
				t.Errorf("test: %d, expected: %s,\n got: %s", k, n.Exception, err)
			}
		} else {
			if n.Result.XMLName.Local != dft.XMLName.Local {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.Result.XMLName.Local, dft.XMLName.Local)
			}
			if n.Result.BaseRequest.Service != dft.BaseRequest.Service {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.Result.BaseRequest.Service, dft.BaseRequest.Service)
			}
			if n.Result.BaseRequest.Version != dft.BaseRequest.Version {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.Result.BaseRequest.Version, dft.BaseRequest.Version)
			}
			if dft.BaseDescribeFeatureTypeRequest.TypeName != nil {
				if *dft.BaseDescribeFeatureTypeRequest.TypeName != *n.Result.BaseDescribeFeatureTypeRequest.TypeName {
					t.Errorf("test: %d, expected: %+v ,\n got: %+v", k, *n.Result.BaseDescribeFeatureTypeRequest.TypeName, *dft.BaseDescribeFeatureTypeRequest.TypeName)
				}
			}
		}
	}
}
func TestBuildQuery(t *testing.T) {
	var tests = []struct {
		dft   DescribeFeatureType
		query url.Values
	}{
		0: {dft: DescribeFeatureType{
			XMLName:     xml.Name{Local: `DescribeFeatureType`},
			BaseRequest: BaseRequest{Version: Version, Service: Service},
			BaseDescribeFeatureTypeRequest: BaseDescribeFeatureTypeRequest{
				OutputFormat: sp(`application/json`),
				TypeName:     sp(`example:example`)}},
			query: map[string][]string{"OUTPUTFORMAT": {"application/json"}, "VERSION": {`2.0.0`}, "REQUEST": {"DescribeFeatureType"}, "SERVICE": {"WFS"}, "TYPENAME": {"example:example"}},
		},
	}

	for k, v := range tests {
		values := v.dft.BuildQuery()
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

func TestBuildBodyDescribeFeatureType(t *testing.T) {
	var tests = []struct {
		dft  DescribeFeatureType
		body string
	}{
		0: {dft: DescribeFeatureType{
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
