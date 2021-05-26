package wsc110

import (
	"testing"

	"gopkg.in/yaml.v3"
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

func TestUnmarshalYAMLCrs(t *testing.T) {

	var stringYAML = `defaultcrs: urn:ogc:def:crs:EPSG::4326
othercrs:
- urn:ogc:def:crs:EPSG::4258 
- urn:ogc:def:crs:EPSG::3857`

	type FeatureType struct {
		DefaultCRS CRS   `yaml:"defaultcrs"`
		OtherCRS   []CRS `yaml:"othercrs"`
	}

	var tests = []struct {
		yaml        []byte
		expectedcrs CRS
	}{
		0: {yaml: []byte(stringYAML), expectedcrs: CRS{Code: 4326, Namespace: EPSG}},
		1: {yaml: []byte(`defaultcrs: urn:ogc:def:crs:EPSG::4326`), expectedcrs: CRS{Code: 4326, Namespace: EPSG}},
		2: {yaml: []byte(`defaultcrs: EPSG:4326`), expectedcrs: CRS{Code: 4326, Namespace: EPSG}},
	}
	for k, test := range tests {
		var ftl FeatureType
		err := yaml.Unmarshal(test.yaml, &ftl)
		if err != nil {
			t.Errorf("test: %d, yaml.UnMarshal failed with '%s'\n", k, err)
		} else {
			if ftl.DefaultCRS.Code != test.expectedcrs.Code || ftl.DefaultCRS.Namespace != test.expectedcrs.Namespace {
				t.Errorf("test: %d, expected: %v+,\n got: %v+", k, test.expectedcrs, ftl.DefaultCRS)
			}
		}
	}
}
