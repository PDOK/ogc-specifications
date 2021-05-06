package common

import (
	"testing"

	"gopkg.in/yaml.v3"
)

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
		expectedCrs CRS
	}{
		0: {yaml: []byte(stringYAML), expectedCrs: CRS{Code: 4326, Namespace: EPSG}},
		1: {yaml: []byte(`defaultcrs: urn:ogc:def:crs:EPSG::4326`), expectedCrs: CRS{Code: 4326, Namespace: EPSG}},
		2: {yaml: []byte(`defaultcrs: EPSG:4326`), expectedCrs: CRS{Code: 4326, Namespace: EPSG}},
	}
	for k, test := range tests {
		var ftl FeatureType
		err := yaml.Unmarshal(test.yaml, &ftl)
		if err != nil {
			t.Errorf("test: %d, yaml.UnMarshal failed with '%s'\n", k, err)
		} else {
			if ftl.DefaultCRS.Code != test.expectedCrs.Code || ftl.DefaultCRS.Namespace != test.expectedCrs.Namespace {
				t.Errorf("test: %d, expected: %v+,\n got: %v+", k, test.expectedCrs, ftl.DefaultCRS)
			}
		}
	}
}
