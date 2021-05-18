package wms130

import (
	"testing"
)

var capabilities = Capabilities{
	WMSCapabilities: WMSCapabilities{
		Layer: []Layer{
			{Name: sp(`depthOneLayerOne`),
				Layer: []*Layer{
					{Name: sp(`depthTwoLayerThree`), Style: []*Style{{Name: `StyleOne`}, {Name: `StyleTwo`}}},
					{Name: sp(`depthTwoLayerFour`),
						Layer: []*Layer{
							{Name: sp(`depthThreeLayerSix`)},
							{Name: sp(`depthThreeLayerSeven`), Style: []*Style{{Name: `StyleThree`}}},
						},
					},
				},
			},
			{Name: sp(`depthOneLayerTwo`),
				Layer: []*Layer{
					{Name: sp(`depthTwoLayerFive`), Style: []*Style{{Name: `StyleFour`}, {Name: `StyleFive`}}}},
			},
		},
	},
}

func TestGetLayerNames(t *testing.T) {
	expected := []string{`depthOneLayerOne`, `depthOneLayerTwo`, `depthTwoLayerThree`, `depthTwoLayerFour`, `depthTwoLayerFive`, `depthThreeLayerSix`, `depthThreeLayerSeven`}

	for _, n := range capabilities.GetLayerNames() {
		found := false
		for _, e := range expected {
			if n == e {
				found = true
			}
		}
		if !found {
			t.Errorf(" got: %s", n)
		}
	}
}

func TestStyleDefined(t *testing.T) {
	var tests = []struct {
		layer   string
		style   string
		defined bool
	}{
		0: {layer: `depthOneLayerOne`, style: `none`, defined: false},
		1: {layer: `depthTwoLayerThree`, style: `StyleTwo`, defined: true},
		2: {layer: `depthTwoLayerFive`, style: `StyleFour`, defined: true},
	}

	for k, test := range tests {
		d := capabilities.StyleDefined(test.layer, test.style)
		if test.defined != d {
			t.Errorf("test: %d, expected: %t \ngot: %t", k, test.defined, d)
		}
	}
}

func TestGetLayer(t *testing.T) {
	var tests = []struct {
		layername string
		exception Exceptions
	}{
		0: {layername: `depthTwoLayerThree`},
		1: {layername: `depthThreeLayerSeven`},
		2: {layername: `unknownLayer`, exception: Exceptions{LayerNotDefined(`unknownLayer`)}},
	}

	for k, test := range tests {
		layerfound, exception := capabilities.GetLayer(test.layername)
		if exception != nil {
			if test.exception != nil {
				if test.exception[0].Code() != exception[0].Code() {
					t.Errorf("test: %d, expected: %s \ngot: %v", k, test.layername, capabilities.GetLayerNames())
				}
			}
		} else {
			if *layerfound.Name != test.layername {
				t.Errorf("test: %d, expected: %s \ngot: %s", k, capabilities.GetLayerNames(), *layerfound.Name)
			}
		}
	}
}
