package wms130

import (
	"encoding/xml"
	"net/url"
	"testing"
)

func TestGetCapabilitiesParseXML(t *testing.T) {
	var tests = []struct {
		body      []byte
		result    GetCapabilitiesRequest
		exception error
	}{
		// GetCapabilities
		0: {body: []byte(`<GetCapabilities service="wms" version="1.3.0" xmlns="http://www.opengis.net/wms"/>`),
			result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, BaseRequest: BaseRequest{Service: "wms", Version: "1.3.0", Attr: []xml.Attr{{Name: xml.Name{Local: "xmlns"}, Value: "http://www.opengis.net/wms"}}}}},
		// Unknown XML document
		1: {body: []byte("<Unknown/>"),
			exception: MissingParameterValue("REQUEST")},
		// no XML document
		2: {body: []byte("no XML document, just a string"),
			exception: MissingParameterValue()},
		// document at all
		3: {exception: MissingParameterValue()},
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

func TestGetCapabilitiesParseKVP(t *testing.T) {
	var tests = []struct {
		query      url.Values
		result     GetCapabilitiesRequest
		exceptions Exceptions
	}{
		// "Normal" query request with UPPER/lower/MiXeD case
		0: {query: map[string][]string{"SERVICE": {"wms"}, "Request": {"GetCapabilities"}, "version": {"1.3.0"}},
			result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, BaseRequest: BaseRequest{Service: Service, Version: Version}}},
		// Missing mandatory SERVICE attribute
		1: {query: map[string][]string{"Request": {"GetCapabilities"}},
			result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, BaseRequest: BaseRequest{Service: Service}}},
		// Missing optional VERSION attribute
		2: {query: map[string][]string{"SERVICE": {"wms"}, "Request": {"GetCapabilities"}},
			result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, BaseRequest: BaseRequest{Service: Service}}},
		// Unknown optional VERSION attribute
		3: {query: map[string][]string{"SERVICE": {"wms"}, "Request": {"GetCapabilities"}, "version": {"3.4.5"}},
			result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, BaseRequest: BaseRequest{Service: Service, Version: "3.4.5"}}},
		4: {query: map[string][]string{"SERVICE": {"wms"}, "Request": {"GetCapabilities"}, "version": {"no version found"}},
			result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"}, BaseRequest: BaseRequest{Service: Service, Version: "no version found"}}},
		// No mandatory SERVICE, REQUEST attribute only optional VERSION
		5: {exceptions: Exceptions{MissingParameterValue(REQUEST), MissingParameterValue(SERVICE)}},
	}

	for k, test := range tests {
		var gc GetCapabilitiesRequest
		exceptions := gc.ParseQueryParameters(test.query)
		if len(exceptions) > 0 {
			for _, exception := range exceptions {
				found := false
				for _, testexception := range test.exceptions {
					if exception == testexception {
						found = true
					}
				}
				if !found {
					t.Errorf("test exception: %d, expected one of: %s ,\n got: %s", k, test.exceptions, exception.Error())
				}
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
		object   GetCapabilitiesRequest
		excepted url.Values
	}{
		0: {object: GetCapabilitiesRequest{BaseRequest: BaseRequest{Service: Service, Version: Version}, XMLName: xml.Name{Local: `GetCapabilities`}},
			excepted: map[string][]string{
				VERSION: {Version},
				SERVICE: {Service},
				REQUEST: {`GetCapabilities`},
			}},
	}

	for k, test := range tests {
		url := test.object.ToQueryParameters()
		if len(test.excepted) != len(url) {
			t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, test.excepted, url)
		} else {
			for _, rid := range url {
				found := false
				for _, erid := range test.excepted {
					if rid[0] == erid[0] {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("test: %d, expected: %+v,\n got: %+v: ", k, test.excepted, url)
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

func BenchmarkGetCapabilitiesToQueryParameters(b *testing.B) {
	gc := GetCapabilitiesRequest{XMLName: xml.Name{Local: getcapabilities}, BaseRequest: BaseRequest{Service: Service, Version: Version}}
	for i := 0; i < b.N; i++ {
		gc.ToQueryParameters()
	}
}

func BenchmarkGetCapabilitiesToXML(b *testing.B) {
	gc := GetCapabilitiesRequest{XMLName: xml.Name{Local: getcapabilities}, BaseRequest: BaseRequest{Service: Service, Version: Version}}
	for i := 0; i < b.N; i++ {
		gc.ToQueryParameters()
	}
}
