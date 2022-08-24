package wsc110

import (
	"encoding/xml"
	"net/url"
	"testing"
)

func sp(s string) *string {
	return &s
}

// An example of a GetCapabilities request message encoded using KVP is
// take from OGC 06-121r3 page 16
const getcapabilities_request = `SERVICE=WCS&REQUEST=GetCapabilities&ACCEPTVERSIONS=1.0.0,0.8.3&SECTIONS=Contents&UPDATESEQUENCE=XYZ123&ACCEPTFORMATS=text/xml`

func TestCapabilitiesParseQueryParameters(t *testing.T) {
	var tests = []struct {
		query      url.Values
		result     GetCapabilitiesRequest
		exceptions []Exception
	}{
		0: {query: map[string][]string{"SERVICE": {"WCS"},
			"REQUEST":        {"GetCapabilities"},
			"ACCEPTVERSIONS": {"1.0.0,0.8.3"},
			"SECTIONS":       {"Contents"},
			"UPDATESEQUENCE": {"XYZ123"},
			"ACCEPTFORMATS":  {"text/xml"}},
			result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"},
				Service:        "WCS",
				UpdateSequence: sp("XYZ123"),
				AcceptFormats:  AcceptFormats{OutputFormat: []string{"text/xml"}},
				AcceptVersions: AcceptVersions{Version: []string{"1.0.0", "0.8.3"}}},
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
			if test.result.AcceptVersions.Version[0] != gc.AcceptVersions.Version[0] {
				t.Errorf("test: %d, expected: %s ,\n got: %s", k, test.result.AcceptVersions, gc.AcceptVersions)
			}
		}
	}
}

// An example of a GetCapabilities request message encoded in XML is
// take from OGC 06-121r3 page 19
const getcapabilities_xml_request = `<?xml version="1.0" encoding="UTF-8"?>
<GetCapabilities xmlns="http://www.opengis.net/ows/1.1" xmlns:ows="http://www.opengis.net/ows/1.1" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.opengis.net/ows/1.1 fragmentGetCapabilitiesRequest.xsd" service="WCS" updateSequence="XYZ123">
<!-- Maximum example for WCS. Primary editor: Arliss Whiteside -->
<AcceptVersions>
<Version>1.0.0</Version>
<Version>0.8.3</Version>
</AcceptVersions>
<Sections>
<Section>Contents</Section>
</Sections>
<AcceptFormats>
<OutputFormat>text/xml</OutputFormat>
</AcceptFormats>
</GetCapabilities>`

func TestGetCapabilitiesParseXML(t *testing.T) {
	var tests = []struct {
		body      []byte
		result    GetCapabilitiesRequest
		exception Exception
	}{

		0: {body: []byte(getcapabilities_xml_request),
			result: GetCapabilitiesRequest{XMLName: xml.Name{Local: "GetCapabilities"},
				Service:        "WCS",
				AcceptVersions: AcceptVersions{Version: []string{"1.0.0", "0.8.3"}},
				Sections:       Sections{Section: []string{"Contents"}},
				AcceptFormats:  AcceptFormats{OutputFormat: []string{"text/xml"}},
				UpdateSequence: sp("XYZ123"),
				Attr: []xml.Attr{{Name: xml.Name{Local: "xmlns"}, Value: "http://www.opengis.net/ows/1.1"},
					{Name: xml.Name{Space: "xmlns", Local: "ows"}, Value: "http://www.opengis.net/ows/1.1"},
					{Name: xml.Name{Space: "xmlns", Local: "xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
					{Name: xml.Name{Space: "http://www.w3.org/2001/XMLSchema-instance", Local: "schemaLocation"}, Value: "http://www.opengis.net/ows/1.1 fragmentGetCapabilitiesRequest.xsd"}}},
		},
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
				t.Errorf("test: %d, expected: %v ,\n got: %v", k, test.result, gc)
			}
			if gc.AcceptVersions.Version[0] != test.result.AcceptVersions.Version[0] {
				t.Errorf("test: %d, expected: %v ,\n got: %v", k, test.result, gc)
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
