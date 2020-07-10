package wfs200

import (
	"encoding/xml"
	"net/url"
	"testing"

	"github.com/pdok/ogc-specifications/pkg/ows"
)

func TestGetCapabilitiesType(t *testing.T) {
	dft := GetCapabilities{}
	if dft.Type() != `GetCapabilities` {
		t.Errorf("test: %d, expected: %s,\n got: %s", 0, `GetCapabilities`, dft.Type())
	}
}

func TestParseBodyGetCapabilities(t *testing.T) {
	var tests = []struct {
		Body   []byte
		Result GetCapabilities
		Error  error
	}{
		// Lots of attribute declarations
		0: {Body: []byte(`<GetCapabilities service="wfs" version="2.0.0" xmlns:gml="http://www.opengis.net/gml/3.2" xmlns:wfs="http://www.opengis.net/wfs/2.0" xmlns:ows="http://www.opengis.net/ows/1.1" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:fes="http://www.opengis.net/fes/2.0" xmlns:inspire_common="http://inspire.ec.europa.eu/schemas/common/1.0" xmlns:inspire_dls="http://inspire.ec.europa.eu/schemas/inspire_dls/1.0" xmlns:kadastralekaartv4="http://kadastralekaartv4.geonovum.nl" xsi:schemaLocation="http://www.opengis.net/wfs/2.0 http://schemas.opengis.net/wfs/2.0/wfs.xsd http://inspire.ec.europa.eu/schemas/inspire_dls/1.0 http://inspire.ec.europa.eu/schemas/inspire_dls/1.0/inspire_dls.xsd http://inspire.ec.europa.eu/schemas/common/1.0 http://inspire.ec.europa.eu/schemas/common/1.0/common.xsd"/>`),
			Result: GetCapabilities{XMLName: xml.Name{Local: "GetCapabilities"}, Service: "wfs", Version: "2.0.0",
				Attr: []xml.Attr{{Name: xml.Name{Space: "xmlns", Local: "gml"}, Value: "http://www.opengis.net/gml/3.2"},
					{Name: xml.Name{Space: "xmlns", Local: "wfs"}, Value: "http://www.opengis.net/wfs/2.0"},
					{Name: xml.Name{Space: "xmlns", Local: "ows"}, Value: "http://www.opengis.net/ows/1.1"},
					{Name: xml.Name{Space: "xmlns", Local: "xlink"}, Value: "http://www.w3.org/1999/xlink"},
					{Name: xml.Name{Space: "xmlns", Local: "xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
					{Name: xml.Name{Space: "xmlns", Local: "fes"}, Value: "http://www.opengis.net/fes/2.0"},
					{Name: xml.Name{Space: "xmlns", Local: "inspire_common"}, Value: "http://inspire.ec.europa.eu/schemas/common/1.0"},
					{Name: xml.Name{Space: "xmlns", Local: "inspire_dls"}, Value: "http://inspire.ec.europa.eu/schemas/inspire_dls/1.0"},
					{Name: xml.Name{Space: "xmlns", Local: "kadastralekaartv4"}, Value: "http://kadastralekaartv4.geonovum.nl"},
					{Name: xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "schemaLocation"}, Value: "http://www.opengis.net/wfs/2.0 http://schemas.opengis.net/wfs/2.0/wfs.xsd http://inspire.ec.europa.eu/schemas/inspire_dls/1.0 http://inspire.ec.europa.eu/schemas/inspire_dls/1.0/inspire_dls.xsd http://inspire.ec.europa.eu/schemas/common/1.0 http://inspire.ec.europa.eu/schemas/common/1.0/common.xsd"}}}},
		// Unknown XML document
		1: {Body: []byte("<Unknown/>"), Error: &WFSException{ExceptionText: "This service does not know the operation: expected element type <GetCapabilities> but have <Unknown>"}},
		// no XML document
		2: {Body: []byte("no XML document, just a string"), Error: &WFSException{ExceptionText: "Could not process XML, is it XML?"}},
		// document at all
		3: {Error: &WFSException{ExceptionText: "Could not process XML, is it XML?"}},
		// Duplicate attributes in XML message with the same value
		4: {Body: []byte(`<GetCapabilities service="wfs" version="2.0.0" xmlns:wfs="http://www.opengis.net/wfs/2.0"  xmlns:wfs="http://www.opengis.net/wfs/2.0" xmlns:wfs="http://www.opengis.net/wfs/2.0"/>`),
			Result: GetCapabilities{XMLName: xml.Name{Local: "GetCapabilities"}, Service: "wfs", Version: "2.0.0",
				Attr: []xml.Attr{{Name: xml.Name{Space: "xmlns", Local: "wfs"}, Value: "http://www.opengis.net/wfs/2.0"}}}},
		// Duplicate attributes in XML message with different values
		5: {Body: []byte(`<GetCapabilities service="wfs" version="2.0.0" xmlns:wfs="http://www.opengis.net/ows/1.1"  xmlns:wfs="http://www.w3.org/2001/XMLSchema-instance" xmlns:wfs="http://www.opengis.net/wfs/2.0"/>`),
			Result: GetCapabilities{XMLName: xml.Name{Local: "GetCapabilities"}, Service: "wfs", Version: "2.0.0",
				Attr: []xml.Attr{{Name: xml.Name{Space: "xmlns", Local: "wfs"}, Value: "http://www.opengis.net/wfs/2.0"}}}},
	}

	for k, n := range tests {
		var gc GetCapabilities
		err := gc.ParseBody(n.Body)
		if err != nil {
			if err.Error() != n.Error.Error() {
				t.Errorf("test: %d, expected: %s,\n got: %s", k, n.Error, err)
			}
		} else {
			if gc.Service != n.Result.Service {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.Result, gc)
			}
			if gc.Version != n.Result.Version {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.Result, gc)
			}
			if len(n.Result.Attr) == len(gc.Attr) {
				c := false
				for _, expected := range n.Result.Attr {
					for _, result := range gc.Attr {
						if result.Name.Local == expected.Name.Local && result.Value == expected.Value {
							c = true
						}
					}
					if !c {
						t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.Result.Attr, gc.Attr)
					}
					c = false
				}
			} else {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.Result.Attr, gc.Attr)
			}
		}
	}
}

func TestParseQueryParametersGetCapabilities(t *testing.T) {
	var tests = []struct {
		Query  url.Values
		Result GetCapabilities
		Error  error
	}{
		// "Normal" query request with UPPER/lower/MiXeD case
		0: {Query: map[string][]string{"SERVICE": {"wfs"}, "Request": {"GetCapabilities"}, "version": {"2.0.0"}},
			Result: GetCapabilities{XMLName: xml.Name{Local: "GetCapabilities"}, Service: "WFS", Version: "2.0.0"}},
		// Missing mandatory SERVICE attribute
		1: {Query: map[string][]string{"Request": {"GetCapabilities"}},
			Result: GetCapabilities{XMLName: xml.Name{Local: "GetCapabilities"}}},
		// Missing optional VERSION attribute
		2: {Query: map[string][]string{"SERVICE": {"wfs"}, "Request": {"GetCapabilities"}},
			Result: GetCapabilities{XMLName: xml.Name{Local: "GetCapabilities"}, Service: "WFS"}},
		// Unknown optional VERSION attribute
		3: {Query: map[string][]string{"SERVICE": {"wfs"}, "Request": {"GetCapabilities"}, "version": {"3.4.5"}},
			Result: GetCapabilities{XMLName: xml.Name{Local: "GetCapabilities"}, Service: "WFS", Version: "3.4.5"}},
		4: {Query: map[string][]string{"SERVICE": {"wfs"}, "Request": {"GetCapabilities"}, "version": {"no version found"}},
			Result: GetCapabilities{XMLName: xml.Name{Local: "GetCapabilities"}, Service: "WFS", Version: "NO VERSION FOUND"}},
		// No mandatory SERVICE, REQUEST attribute only optional VERSION
		5: {
			Error: &WFSException{ExceptionText: "Failed to parse the operation, found: "}},
	}

	for k, n := range tests {
		var gc GetCapabilities
		err := gc.ParseQuery(n.Query)
		if err != nil {
			if err.Error() != n.Error.Error() {
				t.Errorf("test: %d, expected: %s,\n got: %s", k, n.Error, err)
			}
		} else {
			if n.Result.XMLName.Local != gc.XMLName.Local {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.Result.XMLName.Local, gc.XMLName.Local)
			}
			if n.Result.Service != gc.Service {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.Result.Service, gc.Service)
			}
			if n.Result.Version != gc.Version {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, n.Result.Version, gc.Version)
			}
		}
	}
}

func TestGetCapabilitiesBuildQuery(t *testing.T) {
	var tests = []struct {
		Object   GetCapabilities
		Excepted url.Values
		Error    ows.Exception
	}{
		0: {Object: GetCapabilities{Service: Service, Version: Version, XMLName: xml.Name{Local: `GetCapabilities`}},
			Excepted: map[string][]string{
				VERSION: {Version},
				SERVICE: {Service},
				REQUEST: {`GetCapabilities`},
			}},
	}

	for k, n := range tests {
		url := n.Object.BuildQuery()
		if len(n.Excepted) != len(url) {
			t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, n.Excepted, url)
		} else {
			for _, rid := range url {
				found := false
				for _, erid := range n.Excepted {
					if rid[0] == erid[0] {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, n.Excepted, url)
				}
			}
		}
	}
}

func TestGetCapabilitiesBuildBody(t *testing.T) {
	var tests = []struct {
		gc     GetCapabilities
		result string
	}{
		0: {gc: GetCapabilities{Service: Service, Version: Version, XMLName: xml.Name{Local: `GetCapabilities`}},
			result: `<?xml version="1.0" encoding="UTF-8"?>
<GetCapabilities service="WFS" version="2.0.0"/>`},
	}

	for k, v := range tests {
		body := v.gc.BuildBody()

		if string(body) != v.result {
			t.Errorf("test: %d, Expected body %s but was not \n got: %s", k, v.result, string(body))
		}
	}
}
