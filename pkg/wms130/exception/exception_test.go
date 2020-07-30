package exception

import (
	"testing"

	"github.com/pdok/ogc-specifications/pkg/ows"
)

func TestWFSException(t *testing.T) {
	var tests = []struct {
		exception     ows.Exception
		exceptionText string
		exceptionCode string
		locatorCode   string
	}{
		0: {exception: WMSException{ExceptionCode: "", ExceptionText: "", LocatorCode: ""},
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
			exceptionText: "The parameters I and J are invalid, given: 0, 0",
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

	for k, a := range tests {
		if a.exception.Error() != a.exceptionText {
			t.Errorf("test: %d, expected: %s\n got: %s", k, a.exceptionText, a.exception.Error())
		}
		if a.exception.Code() != a.exceptionCode {
			t.Errorf("test: %d, expected: %s\n got: %s", k, a.exceptionCode, a.exception.Code())
		}
		if a.exception.Locator() != a.locatorCode {
			t.Errorf("test: %d, expected: %s\n got: %s", k, a.locatorCode, a.exception.Locator())
		}
	}
}

func TestReport(t *testing.T) {
	var tests = []struct {
		exceptions []ows.Exception
		result     []byte
		err        error
	}{
		0: {exceptions: []ows.Exception{WMSException{ExceptionCode: "", ExceptionText: "", LocatorCode: ""}},
			result: []byte(`<?xml version="1.0" encoding="UTF-8"?>
<ServiceExceptionReport version="1.3.0" xmlns="http://www.opengis.net/ogc" xsi="http://www.w3.org/2001/XMLSchema-instance" schemaLocation="http://www.opengis.net/ogc http://schemas.opengis.net/wms/1.3.0/exceptions_1_3_0.xsd">
 <ServiceException code=""></ServiceException>
</ServiceExceptionReport>`)},
		1: {exceptions: []ows.Exception{
			LayerNotQueryable(`unknown:layer`),
			InvalidPoint("0", "0"),
		},
			result: []byte(`<?xml version="1.0" encoding="UTF-8"?>
<ServiceExceptionReport version="1.3.0" xmlns="http://www.opengis.net/ogc" xsi="http://www.w3.org/2001/XMLSchema-instance" schemaLocation="http://www.opengis.net/ogc http://schemas.opengis.net/wms/1.3.0/exceptions_1_3_0.xsd">
 <ServiceException code="LayerNotQueryable" locator="unknown:layer">Layer: unknown:layer, can not be queried</ServiceException>
 <ServiceException code="InvalidPoint">The parameters I and J are invalid, given: 0, 0</ServiceException>
</ServiceExceptionReport>`)},
	}

	for k, a := range tests {
		report := WMSServiceExceptionReport{}
		r := report.Report(a.exceptions)

		if string(r) != string(a.result) {
			t.Errorf("test: %d, expected: %s\n got: %s", k, a.result, r)
		}
	}
}
