package wfs200

import (
	"encoding/xml"
	"net/url"
	"testing"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

func TestGetCapabilitiesType(t *testing.T) {
	dft := GetCapabilitiesRequest{}
	if dft.Type() != `GetCapabilities` {
		t.Errorf("test: %d, expected: %s,\n got: %s", 0, `GetCapabilities`, dft.Type())
	}
}

func TestGetCapabilitiesParseXML(t *testing.T) {
	var tests = []struct {
		body      []byte
		result    GetCapabilitiesRequest
		exception wsc110.Exception
	}{
		// Lots of attribute declarations
		0: {body: []byte(`<GetCapabilities service="wfs" version="2.0.0" xmlns:gml="http://www.opengis.net/gml/3.2" xmlns:wfs="http://www.opengis.net/wfs/2.0" xmlns:ows="http://www.opengis.net/ows/1.1" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:fes="http://www.opengis.net/fes/2.0" xmlns:inspire_common="http://inspire.ec.europa.eu/schemas/common/1.0" xmlns:inspire_dls="http://inspire.ec.europa.eu/schemas/inspire_dls/1.0" xmlns:kadastralekaartv4="http://kadastralekaartv4.geonovum.nl" xsi:schemaLocation="http://www.opengis.net/wfs/2.0 http://schemas.opengis.net/wfs/2.0/wfs.xsd http://inspire.ec.europa.eu/schemas/inspire_dls/1.0 http://inspire.ec.europa.eu/schemas/inspire_dls/1.0/inspire_dls.xsd http://inspire.ec.europa.eu/schemas/common/1.0 http://inspire.ec.europa.eu/schemas/common/1.0/commotest.xsd"/>`),
			result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, Service: "wfs", Version: "2.0.0",
				Attr: []xml.Attr{{Name: xml.Name{Space: "xmlns", Local: "gml"}, Value: "http://www.opengis.net/gml/3.2"},
					{Name: xml.Name{Space: "xmlns", Local: "wfs"}, Value: "http://www.opengis.net/wfs/2.0"},
					{Name: xml.Name{Space: "xmlns", Local: "ows"}, Value: "http://www.opengis.net/ows/1.1"},
					{Name: xml.Name{Space: "xmlns", Local: "xlink"}, Value: "http://www.w3.org/1999/xlink"},
					{Name: xml.Name{Space: "xmlns", Local: "xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
					{Name: xml.Name{Space: "xmlns", Local: "fes"}, Value: "http://www.opengis.net/fes/2.0"},
					{Name: xml.Name{Space: "xmlns", Local: "inspire_common"}, Value: "http://inspire.ec.europa.eu/schemas/common/1.0"},
					{Name: xml.Name{Space: "xmlns", Local: "inspire_dls"}, Value: "http://inspire.ec.europa.eu/schemas/inspire_dls/1.0"},
					{Name: xml.Name{Space: "xmlns", Local: "kadastralekaartv4"}, Value: "http://kadastralekaartv4.geonovum.nl"},
					{Name: xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "schemaLocation"}, Value: "http://www.opengis.net/wfs/2.0 http://schemas.opengis.net/wfs/2.0/wfs.xsd http://inspire.ec.europa.eu/schemas/inspire_dls/1.0 http://inspire.ec.europa.eu/schemas/inspire_dls/1.0/inspire_dls.xsd http://inspire.ec.europa.eu/schemas/common/1.0 http://inspire.ec.europa.eu/schemas/common/1.0/commotest.xsd"}}}},
		// Unknown XML document
		1: {body: []byte("<Unknown/>"),
			exception: wsc110.NoApplicableCode("This service does not know the operation: expected element type <GetCapabilities> but have <Unknown>")},
		// no XML document
		2: {body: []byte("no XML document, just a string"),
			exception: wsc110.NoApplicableCode("Could not process XML, is it XML?")},
		// document at all
		3: {exception: wsc110.NoApplicableCode("Could not process XML, is it XML?")},
		// Duplicate attributes in XML message with the same value
		4: {body: []byte(`<GetCapabilities service="wfs" version="2.0.0" xmlns:wfs="http://www.opengis.net/wfs/2.0"  xmlns:wfs="http://www.opengis.net/wfs/2.0" xmlns:wfs="http://www.opengis.net/wfs/2.0"/>`),
			result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, Service: "wfs", Version: "2.0.0",
				Attr: []xml.Attr{{Name: xml.Name{Space: "xmlns", Local: "wfs"}, Value: "http://www.opengis.net/wfs/2.0"}}}},
		// Duplicate attributes in XML message with different values
		5: {body: []byte(`<GetCapabilities service="wfs" version="2.0.0" xmlns:wfs="http://www.opengis.net/ows/1.1"  xmlns:wfs="http://www.w3.org/2001/XMLSchema-instance" xmlns:wfs="http://www.opengis.net/wfs/2.0"/>`),
			result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, Service: "wfs", Version: "2.0.0",
				Attr: []xml.Attr{{Name: xml.Name{Space: "xmlns", Local: "wfs"}, Value: "http://www.opengis.net/wfs/2.0"}}}},
	}

	for k, test := range tests {
		var gc GetCapabilitiesRequest
		exception := gc.ParseXML(test.body)
		if exception != nil {
			if exception[0].Error() != test.exception.Error() {
				t.Errorf("test: %d, expected: %s,\n got: %s", k, test.exception, exception)
			}
		} else {
			if gc.Service != test.result.Service {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, test.result, gc)
			}
			if gc.Version != test.result.Version {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, test.result, gc)
			}
			if len(test.result.Attr) == len(gc.Attr) {
				c := false
				for _, expected := range test.result.Attr {
					for _, result := range gc.Attr {
						if result.Name.Local == expected.Name.Local && result.Value == expected.Value {
							c = true
						}
					}
					if !c {
						t.Errorf("test: %d, expected: %s ,\n got: %s", k, test.result.Attr, gc.Attr)
					}
					c = false
				}
			} else {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, test.result.Attr, gc.Attr)
			}
		}
	}
}

func TestGetCapabilitiesParseQueryParameters(t *testing.T) {
	var tests = []struct {
		query      url.Values
		result     GetCapabilitiesRequest
		exceptions []wsc110.Exception
	}{
		// "Normal" query request with UPPER/lower/MiXeD case
		0: {query: map[string][]string{"SERVICE": {"wfs"}, "Request": {"GetCapabilities"}, "version": {"2.0.0"}},
			result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, Service: "WFS", Version: "2.0.0"}},
		// Missing mandatory SERVICE attribute
		1: {query: map[string][]string{"Request": {"GetCapabilities"}},
			result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}}},
		// Missing optional VERSION attribute
		2: {query: map[string][]string{"SERVICE": {"wfs"}, "Request": {"GetCapabilities"}},
			result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, Service: "WFS"}},
		// Unknown optional VERSION attribute
		3: {query: map[string][]string{"SERVICE": {"wfs"}, "Request": {"GetCapabilities"}, "version": {"3.4.5"}},
			result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, Service: "WFS", Version: "3.4.5"}},
		4: {query: map[string][]string{"SERVICE": {"wfs"}, "Request": {"GetCapabilities"}, "version": {"no version found"}},
			result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, Service: "WFS", Version: "no version found"}},
		// No mandatory SERVICE, REQUEST attribute only optional VERSION
		5: {
			exceptions: []wsc110.Exception{wsc110.MissingParameterValue(SERVICE), wsc110.MissingParameterValue(REQUEST)},
		},
	}

	for k, test := range tests {
		var gc GetCapabilitiesRequest
		exception := gc.ParseQueryParameters(test.query)
		if exception != nil {
			if exception[0].Error() != test.exceptions[0].Error() {
				t.Errorf("test: %d, expected: %s,\n got: %s", k, test.exceptions, exception)
			}
		} else {
			if test.result.XMLName.Local != gc.XMLName.Local {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, test.result.XMLName.Local, gc.XMLName.Local)
			}
			if test.result.Service != gc.Service {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, test.result.Service, gc.Service)
			}
			if test.result.Version != gc.Version {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, test.result.Version, gc.Version)
			}
		}
	}
}

func TestGetCapabilitiesToQueryParameters(t *testing.T) {
	var tests = []struct {
		Object    GetCapabilitiesRequest
		Excepted  url.Values
		Exception wsc110.Exception
	}{
		0: {Object: GetCapabilitiesRequest{Service: Service, Version: Version, XMLName: xml.Name{Local: `GetCapabilities`}},
			Excepted: map[string][]string{
				VERSION: {Version},
				SERVICE: {Service},
				REQUEST: {`GetCapabilities`},
			}},
	}

	for k, test := range tests {
		url := test.Object.ToQueryParameters()
		if len(test.Excepted) != len(url) {
			t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, test.Excepted, url)
		} else {
			for _, rid := range url {
				found := false
				for _, erid := range test.Excepted {
					if rid[0] == erid[0] {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, test.Excepted, url)
				}
			}
		}
	}
}

func TestGetCapabilitiesToXML(t *testing.T) {
	var tests = []struct {
		gc     GetCapabilitiesRequest
		result string
	}{
		0: {gc: GetCapabilitiesRequest{Service: Service, Version: Version, XMLName: xml.Name{Local: `GetCapabilities`}},
			result: `<?xml version="1.0" encoding="UTF-8"?>
<GetCapabilities service="WFS" version="2.0.0"/>`},
	}

	for k, v := range tests {
		body := v.gc.ToXML()

		if string(body) != v.result {
			t.Errorf("test: %d, Expected body %s but was not \n got: %s", k, v.result, string(body))
		}
	}
}

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
