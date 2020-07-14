package request

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
		// GetCapabilities
		0: {Body: []byte(`<GetCapabilities service="wms" version="1.3.0" xmlns="http://www.opengis.net/wms"/>`),
			Result: GetCapabilities{XMLName: xml.Name{Local: "GetCapabilities"}, Service: "wms", Version: "1.3.0",
				Attr: []xml.Attr{{Name: xml.Name{Local: "xmlns"}, Value: "http://www.opengis.net/wms"}}}},
		// Unknown XML document
		1: {Body: []byte("<Unknown/>"), Error: ows.MissingParameterValue("REQUEST")},
		// no XML document
		2: {Body: []byte("no XML document, just a string"), Error: ows.MissingParameterValue()},
		// document at all
		3: {Error: ows.MissingParameterValue()},
	}

	for k, n := range tests {
		var gc GetCapabilities
		err := gc.ParseXML(n.Body)
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
		0: {Query: map[string][]string{"SERVICE": {"wms"}, "Request": {"GetCapabilities"}, "version": {"1.3.0"}},
			Result: GetCapabilities{XMLName: xml.Name{Local: "GetCapabilities"}, Service: Service, Version: Version}},
		// Missing mandatory SERVICE attribute
		1: {Query: map[string][]string{"Request": {"GetCapabilities"}},
			Result: GetCapabilities{XMLName: xml.Name{Local: "GetCapabilities"}}},
		// Missing optional VERSION attribute
		2: {Query: map[string][]string{"SERVICE": {"wms"}, "Request": {"GetCapabilities"}},
			Result: GetCapabilities{XMLName: xml.Name{Local: "GetCapabilities"}, Service: Service}},
		// Unknown optional VERSION attribute
		3: {Query: map[string][]string{"SERVICE": {"wms"}, "Request": {"GetCapabilities"}, "version": {"3.4.5"}},
			Result: GetCapabilities{XMLName: xml.Name{Local: "GetCapabilities"}, Service: Service, Version: "3.4.5"}},
		4: {Query: map[string][]string{"SERVICE": {"wms"}, "Request": {"GetCapabilities"}, "version": {"no version found"}},
			Result: GetCapabilities{XMLName: xml.Name{Local: "GetCapabilities"}, Service: Service, Version: "NO VERSION FOUND"}},
		// No mandatory SERVICE, REQUEST attribute only optional VERSION
		5: {
			Error: ows.MissingParameterValue()},
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
<GetCapabilities service="WMS" version="1.3.0"/>`},
	}

	for k, v := range tests {
		body := v.gc.BuildXML()

		if string(body) != v.result {
			t.Errorf("test: %d, Expected body %s but was not \n got: %s", k, v.result, string(body))
		}
	}
}
