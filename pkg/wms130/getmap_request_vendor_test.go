package wms130

import (
	"net/url"
	"testing"
)

func TestVendorRoundTrip(t *testing.T) {
	q := url.Values{
		"SERVICE":        {""},
		"VERSION":        {""},
		"REQUEST":        {""},
		"LAYERS":         {""},
		"STYLES":         {""},
		"CRS":            {""},
		"BBOX":           {"0,0,1,1"},
		"WIDTH":          {"1"},
		"HEIGHT":         {"1"},
		"FORMAT":         {""},
		"map_resolution": {"123"},
	}
	var m GetMapRequest
	if ex := m.ParseQueryParameters(q); ex != nil {
		t.Fatal(ex)
	}
	if m.VendorSpecificParameters["map_resolution"] != "123" || len(m.VendorSpecificParameters) != 1 {
		t.Fatal(m.VendorSpecificParameters)
	}
	if m.ToQueryParameters().Get("map_resolution") != "123" {
		t.Fatal("lost")
	}
}
