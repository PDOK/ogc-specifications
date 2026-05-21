package wms130

import (
	"github.com/pdok/ogc-specifications/pkg/common"
	"testing"
)

func TestWFSException(t *testing.T) {
	var tests = []struct {
		exception     exception
		exceptionText string
		exceptionCode string
		locatorCode   string
	}{
		0: {exception: exception{ExceptionDetails: common.ExceptionDetails{ExceptionCode: "", ExceptionText: "", LocatorCode: ""}},
			exceptionText: "",
			exceptionCode: "",
			locatorCode:   "",
		},
		1: {exception: InvalidFormat(`unknownimage`),
			exceptionText: "The format: unknownimage, is a invalid image format",
			exceptionCode: "InvalidFormat",
		},
		2: {exception: InvalidCRS(),
			exceptionCode: "InvalidCRS",
		},
		3: {exception: LayerNotDefined(),
			exceptionCode: "LayerNotDefined",
		},
		4: {exception: StyleNotDefined(),
			exceptionCode: "StyleNotDefined",
			exceptionText: `There is a one-to-one correspondence between the values in the LAYERS parameter and the values in the STYLES parameter. 
	Expecting an empty string for the STYLES like STYLES= or comma-separated list STYLES=,,, or using keyword default STYLES=default,default,...`,
		},
		5: {exception: LayerNotQueryable(),
			exceptionCode: "LayerNotQueryable",
		},
		6: {exception: InvalidPoint("0", "0"),
			exceptionCode: "InvalidPoint",
			exceptionText: "The parameters I and J are invalid, given: 0 for I and 0 for J",
		},
		7: {exception: CurrentUpdateSequence(),
			exceptionCode: "CurrentUpdateSequence",
		},
		8: {exception: InvalidUpdateSequence(),
			exceptionCode: "InvalidUpdateSequence",
		},
		9: {exception: MissingDimensionValue(),
			exceptionCode: "MissingDimensionValue",
		},
		10: {exception: InvalidDimensionValue(),
			exceptionCode: "InvalidDimensionValue",
		},
		11: {exception: InvalidCRS(`ESPG:UNKNOWN`),
			exceptionCode: "InvalidCRS",
			exceptionText: `CRS is not known by this service: ESPG:UNKNOWN`,
		},
		12: {exception: LayerNotDefined(`unknown:layer`),
			exceptionCode: "LayerNotDefined",
			exceptionText: `The layer: unknown:layer is not known by the server`,
		},
	}

	for k, test := range tests {
		if test.exception.Error() != test.exceptionText {
			t.Errorf("test: %d, expected: %s\n got: %s", k, test.exceptionText, test.exception.Error())
		}
		if test.exception.Code() != test.exceptionCode {
			t.Errorf("test: %d, expected: %s\n got: %s", k, test.exceptionCode, test.exception.Code())
		}
		if test.exception.Locator() != test.locatorCode {
			t.Errorf("test: %d, expected: %s\n got: %s", k, test.locatorCode, test.exception.Locator())
		}
	}
}

func TestReport(t *testing.T) {
	var tests = []struct {
		exceptions Exceptions
		result     []byte
		err        error
	}{
		0: {exceptions: Exceptions{exception{ExceptionDetails: common.ExceptionDetails{ExceptionCode: "", ExceptionText: "", LocatorCode: ""}}},
			result: []byte(`<?xml version="1.0" encoding="UTF-8"?>
<ServiceExceptionReport version="1.3.0" xmlns="http://www.opengis.net/ogc" xsi="http://www.w3.org/2001/XMLSchema-instance" schemaLocation="http://www.opengis.net/ogc http://schemas.opengis.net/wms/1.3.0/exceptions_1_3_0.xsd">
 <ServiceException code=""></ServiceException>
</ServiceExceptionReport>`)},
		1: {exceptions: Exceptions{
			LayerNotQueryable(`unknown:layer`),
			InvalidPoint("0", "0"),
		},
			result: []byte(`<?xml version="1.0" encoding="UTF-8"?>
<ServiceExceptionReport version="1.3.0" xmlns="http://www.opengis.net/ogc" xsi="http://www.w3.org/2001/XMLSchema-instance" schemaLocation="http://www.opengis.net/ogc http://schemas.opengis.net/wms/1.3.0/exceptions_1_3_0.xsd">
 <ServiceException code="LayerNotQueryable" locator="unknown:layer">Layer: unknown:layer, can not be queried</ServiceException>
 <ServiceException code="InvalidPoint">The parameters I and J are invalid, given: 0 for I and 0 for J</ServiceException>
</ServiceExceptionReport>`)},
	}

	for k, test := range tests {
		r := test.exceptions.ToReport().ToBytes()

		if string(r) != string(test.result) {
			t.Errorf("test: %d, expected: %s\n got: %s", k, test.result, r)
		}
	}
}
