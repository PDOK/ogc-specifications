package wms130

import (
	"testing"

	"gopkg.in/yaml.v3"
)

var capabilities = Capabilities{
	WMSCapabilities: WMSCapabilities{
		Layer: []Layer{
			{Name: new(`depthOneLayerOne`),
				Layer: []*Layer{
					{Name: new(`depthTwoLayerThree`), Style: []*Style{{Name: `StyleOne`}, {Name: `StyleTwo`}}},
					{Name: new(`depthTwoLayerFour`),
						Layer: []*Layer{
							{Name: new(`depthThreeLayerSix`)},
							{Name: new(`depthThreeLayerSeven`), Style: []*Style{{Name: `StyleThree`}}},
						},
					},
				},
			},
			{Name: new(`depthOneLayerTwo`),
				Layer: []*Layer{
					{Name: new(`depthTwoLayerFive`), Style: []*Style{{Name: `StyleFour`}, {Name: `StyleFive`}}}},
			},
		},
	},
}

var capabilitiesWithException = Capabilities{
	WMSCapabilities: WMSCapabilities{
		Exception: ExceptionType{
			Format: []string{"XML"},
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
		layerName string
		exception Exceptions
	}{
		0: {layerName: `depthTwoLayerThree`},
		1: {layerName: `depthThreeLayerSeven`},
		2: {layerName: `unknownLayer`, exception: Exceptions{LayerNotDefined(`unknownLayer`)}},
	}

	for k, test := range tests {
		layerFound, exception := capabilities.GetLayer(test.layerName)
		if exception != nil {
			if test.exception != nil {
				if test.exception[0].Code() != exception[0].Code() {
					t.Errorf("test: %d, expected: %s \ngot: %v", k, test.layerName, capabilities.GetLayerNames())
				}
			}
		} else {
			if *layerFound.Name != test.layerName {
				t.Errorf("test: %d, expected: %s \ngot: %s", k, capabilities.GetLayerNames(), *layerFound.Name)
			}
		}
	}
}

func TestException(t *testing.T) {

	out, err := yaml.Marshal(capabilitiesWithException)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var got map[string]any
	if err = yaml.Unmarshal(out, &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	wmsRaw, ok := got["wmsCapabilities"]
	if !ok {
		t.Fatalf("expected 'wmsCapabilities' key")
	}

	wms, ok := wmsRaw.(map[string]any)
	if !ok {
		t.Fatalf("'wmsCapabilities' is not a map, got %T", wmsRaw)
	}

	exceptionRaw, ok := wms["exception"]
	if !ok {
		t.Fatalf("expected 'exception' key to exist")
	}

	exception, ok := exceptionRaw.(map[string]any)
	if !ok {
		t.Fatalf("'exception' is not a map, got %T", exceptionRaw)
	}

	formatsRaw, ok := exception["format"]
	if !ok {
		t.Fatalf("expected 'format' key to exist")
	}

	formats, ok := formatsRaw.([]any)
	if !ok {
		t.Fatalf("'format' is not a slice, got %T", formatsRaw)
	}

	if len(formats) != 1 {
		t.Fatalf("expected exactly 1 format, got %d", len(formats))
	}

	formatStr, ok := formats[0].(string)
	if !ok {
		t.Fatalf("format value is not string, got %T", formats[0])
	}

	if formatStr != "XML" {
		t.Errorf("expected 'XML', got %s", formatStr)
	}
}
