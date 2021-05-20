package wms130

import (
	"encoding/xml"
	"net/url"
	"testing"
)

func TestGetCapabilitiesParseXML(t *testing.T) {
	var tests = []struct {
		Body      []byte
		Result    GetCapabilitiesRequest
		Exception error
	}{
		// GetCapabilities
		0: {Body: []byte(`<GetCapabilities service="wms" version="1.3.0" xmlns="http://www.opengis.net/wms"/>`),
			Result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, BaseRequest: BaseRequest{Service: "wms", Version: "1.3.0", Attr: []xml.Attr{{Name: xml.Name{Local: "xmlns"}, Value: "http://www.opengis.net/wms"}}}}},
		// Unknown XML document
		1: {Body: []byte("<Unknown/>"), Exception: MissingParameterValue("REQUEST")},
		// no XML document
		2: {Body: []byte("no XML document, just a string"), Exception: MissingParameterValue()},
		// document at all
		3: {Exception: MissingParameterValue()},
	}

	for k, n := range tests {
		var gc GetCapabilitiesRequest
		exception := gc.ParseXML(n.Body)
		if exception != nil {
			if exception[0].Error() != n.Exception.Error() {
				t.Errorf("test: %d, expected: %s,\n got: %s", k, n.Exception, exception)
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

func TestGetCapabilitiesParseKVP(t *testing.T) {
	var tests = []struct {
		Query      url.Values
		Result     GetCapabilitiesRequest
		Exceptions Exceptions
	}{
		// "Normal" query request with UPPER/lower/MiXeD case
		0: {Query: map[string][]string{"SERVICE": {"wms"}, "Request": {"GetCapabilities"}, "version": {"1.3.0"}},
			Result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, BaseRequest: BaseRequest{Service: Service, Version: Version}}},
		// Missing mandatory SERVICE attribute
		1: {Query: map[string][]string{"Request": {"GetCapabilities"}},
			Result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}}},
		// Missing optional VERSION attribute
		2: {Query: map[string][]string{"SERVICE": {"wms"}, "Request": {"GetCapabilities"}},
			Result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, BaseRequest: BaseRequest{Service: Service}}},
		// Unknown optional VERSION attribute
		3: {Query: map[string][]string{"SERVICE": {"wms"}, "Request": {"GetCapabilities"}, "version": {"3.4.5"}},
			Result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, BaseRequest: BaseRequest{Service: Service, Version: "3.4.5"}}},
		4: {Query: map[string][]string{"SERVICE": {"wms"}, "Request": {"GetCapabilities"}, "version": {"no version found"}},
			Result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, BaseRequest: BaseRequest{Service: Service, Version: "no version found"}}},
		// No mandatory SERVICE, REQUEST attribute only optional VERSION
		5: {
			Exceptions: Exceptions{MissingParameterValue(REQUEST), MissingParameterValue(SERVICE)}},
	}

	for k, test := range tests {
		var gc GetCapabilitiesRequest
		exceptions := gc.ParseQueryParameters(test.Query)
		if len(exceptions) > 0 {
			for _, exception := range exceptions {
				found := false
				for _, testexception := range test.Exceptions {
					if exception == testexception {
						found = true
					}
				}
				if !found {
					t.Errorf("test exception: %d, expected one of: %s ,\n got: %s", k, test.Exceptions, exception.Error())
				}
			}
		} else {
			if test.Result.XMLName.Local != gc.XMLName.Local {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, test.Result.XMLName.Local, gc.XMLName.Local)
			}
			if test.Result.Service != gc.Service {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, test.Result.Service, gc.Service)
			}
			if test.Result.Version != gc.Version {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, test.Result.Version, gc.Version)
			}
		}
	}
}

func TestGetCapabilitiesBuildKVP(t *testing.T) {
	var tests = []struct {
		Object    GetCapabilitiesRequest
		Excepted  url.Values
		Exception Exceptions
	}{
		0: {Object: GetCapabilitiesRequest{BaseRequest: BaseRequest{Service: Service, Version: Version}, XMLName: xml.Name{Local: `GetCapabilities`}},
			Excepted: map[string][]string{
				VERSION: {Version},
				SERVICE: {Service},
				REQUEST: {`GetCapabilities`},
			}},
	}

	for k, n := range tests {
		url := n.Object.ToQueryParameters()
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

func TestGetCapabilitiesBuildXML(t *testing.T) {
	var tests = []struct {
		gc     GetCapabilitiesRequest
		result string
	}{
		0: {gc: GetCapabilitiesRequest{BaseRequest: BaseRequest{Service: Service, Version: Version}, XMLName: xml.Name{Local: `GetCapabilities`}},
			result: `<?xml version="1.0" encoding="UTF-8"?>
<GetCapabilities service="WMS" version="1.3.0"/>`},
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

func BenchmarkGetCapabilitiesBuildKVP(b *testing.B) {
	gc := GetCapabilitiesRequest{XMLName: xml.Name{Local: getcapabilities}, BaseRequest: BaseRequest{Service: Service, Version: Version}}
	for i := 0; i < b.N; i++ {
		gc.ToQueryParameters()
	}
}

func BenchmarkGetCapabilitiesBuildXML(b *testing.B) {
	gc := GetCapabilitiesRequest{XMLName: xml.Name{Local: getcapabilities}, BaseRequest: BaseRequest{Service: Service, Version: Version}}
	for i := 0; i < b.N; i++ {
		gc.ToQueryParameters()
	}
}
