package capabilities

import (
	"testing"
)

func sp(s string) *string {
	return &s
}

func TestGetLayerNames(t *testing.T) {
	capabilities := Capability{
		Layer: []Layer{
			{
				Name: sp(`depthOneLayerOne`),
				Layer: []*Layer{
					{
						Name: sp(`depthTwoLayerThree`),
					},
					{
						Name: sp(`depthTwoLayerFour`),
						Layer: []*Layer{
							{
								Name: sp(`depthThreeLayerSix`),
							},
							{
								Name: sp(`depthThreeLayerSeven`),
							},
						},
					},
				},
			},
			{
				Name: sp(`depthOneLayerTwo`),
				Layer: []*Layer{
					{
						Name: sp(`depthTwoLayerFive`),
					},
				},
			},
		},
	}

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
