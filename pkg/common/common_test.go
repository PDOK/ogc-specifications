package common

import (
	"testing"
)

func TestCRSParseString(t *testing.T) {
	var tests = []struct {
		input       string
		expectedCRS CRS
	}{
		0: {}, // Empty input == empty struct
		1: {input: `urn:ogc:def:crs:EPSG::4326`, expectedCRS: CRS{Code: 4326, Namespace: `EPSG`}},
		2: {input: `EPSG:4326`, expectedCRS: CRS{Code: 4326, Namespace: `EPSG`}},
	}

	for k, test := range tests {
		var crs CRS
		crs.parseString(test.input)

		if crs.Code != test.expectedCRS.Code || crs.Namespace != test.expectedCRS.Namespace {
			t.Errorf("test: %d, expected: %v,\n got: %v", k, test.expectedCRS, crs)
		}
	}
}
