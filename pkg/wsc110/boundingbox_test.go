package wsc110

import (
	"testing"

	"github.com/pdok/ogc-specifications/pkg/common"
)

func TestBoundingBoxBuildQueryString(t *testing.T) {
	var tests = []struct {
		boundingbox       BoundingBox
		boundingboxstring string
	}{
		// While 'not' correct this will we checked in the validation step
		0: {boundingbox: BoundingBox{}, boundingboxstring: `0.000000,0.000000,0.000000,0.000000`},
		1: {boundingbox: BoundingBox{LowerCorner: [2]float64{-180.0, -90.0}, UpperCorner: [2]float64{180.0, 90.0}}, boundingboxstring: `-180.000000,-90.000000,180.000000,90.000000`},
	}
	for k, test := range tests {
		str := test.boundingbox.ToQueryParameters()
		if str != test.boundingboxstring {
			t.Errorf("test: %d, expected: %v+,\n got: %v+", k, test.boundingboxstring, str)
		}
	}
}

func TestBuildBoundingBox(t *testing.T) {
	var tests = []struct {
		boundingbox string
		bbox        BoundingBox
		exception   common.Exception
	}{
		0: {boundingbox: "0,0,100,100", bbox: BoundingBox{LowerCorner: [2]float64{0, 0}, UpperCorner: [2]float64{100, 100}}},
		1: {boundingbox: "0,0,-100,-100", bbox: BoundingBox{LowerCorner: [2]float64{0, 0}, UpperCorner: [2]float64{-100, -100}}}, // while this isn't correct, this will be 'addressed' in the validation step
		2: {boundingbox: "0,0,100", exception: InvalidParameterValue(`0,0,100`, `boundingbox`)},
		3: {boundingbox: ",,,", exception: InvalidParameterValue(`,,,`, `boundingbox`)},
		4: {boundingbox: ",,,100", exception: InvalidParameterValue(`,,,100`, `boundingbox`)},
		5: {boundingbox: "number,,,100", exception: InvalidParameterValue(`number,,,100`, `boundingbox`)},
	}

	for k, test := range tests {
		var bbox BoundingBox
		if exception := bbox.ParseString(test.boundingbox); exception != nil {
			if exception != test.exception {
				t.Errorf("test: %d, expected: %+v \ngot: %+v", k, test.exception, exception)
			}
		} else {
			if bbox != test.bbox {
				t.Errorf("test: %d, expected: %+v \ngot: %+v", k, test.bbox, bbox)
			}
		}
	}
}
