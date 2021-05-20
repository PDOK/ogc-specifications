package wfs200

import (
	"testing"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

func TestWFSException(t *testing.T) {
	var tests = []struct {
		exception     wsc110.Exception
		exceptionText string
		exceptionCode string
		locatorCode   string
	}{
		0: {exception: wsc110.Exception{ExceptionCode: "", ExceptionText: "", LocatorCode: ""},
			exceptionText: "",
			exceptionCode: "",
			locatorCode:   "",
		},
		1: {exception: CannotLockAllFeatures(),
			exceptionCode: "CannotLockAllFeatures",
		},
		2: {exception: DuplicateStoredQueryIDValue(),
			exceptionCode: "DuplicateStoredQueryIDValue",
		},
		3: {exception: DuplicateStoredQueryParameterName(),
			exceptionCode: "DuplicateStoredQueryParameterName",
		},
		4: {exception: FeaturesNotLocked(),
			exceptionCode: "FeaturesNotLocked",
		},
		5: {exception: InvalidLockID(),
			exceptionCode: "InvalidLockID",
		},
		6: {exception: InvalidValue(),
			exceptionCode: "InvalidValue",
		},
		7: {exception: LockHasExpired(),
			exceptionCode: "LockHasExpired",
		},
		8: {exception: OperationParsingFailed("PARAMETER", "VALUE"),
			exceptionCode: "OperationParsingFailed",
			exceptionText: "Failed to parse the operation, found: PARAMETER",
			locatorCode:   "VALUE",
		},
		9: {exception: OperationProcessingFailed(),
			exceptionCode: "OperationProcessingFailed",
		},
		10: {exception: ResponseCacheExpired(),
			exceptionCode: "ResponseCacheExpired",
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
	}{
		0: {exceptions: Exceptions{wsc110.Exception{ExceptionCode: "", ExceptionText: "", LocatorCode: ""}},
			result: []byte(`<?xml version="1.0" encoding="UTF-8"?>
<ows:ExceptionReport xmlns:ows="http://www.opengis.net/ows/1.1" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.opengis.net/ows/1.1 http://schemas.opengis.net/ows/1.1.0/owsExceptionReport.xsd" version="2.0.0" xml:lang="en">
 <ows:Exception exceptionCode=""></ows:Exception>
</ows:ExceptionReport>`)},
		1: {exceptions: Exceptions{
			CannotLockAllFeatures(),
			DuplicateStoredQueryIDValue(),
		},
			result: []byte(`<?xml version="1.0" encoding="UTF-8"?>
<ows:ExceptionReport xmlns:ows="http://www.opengis.net/ows/1.1" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.opengis.net/ows/1.1 http://schemas.opengis.net/ows/1.1.0/owsExceptionReport.xsd" version="2.0.0" xml:lang="en">
 <ows:Exception exceptionCode="CannotLockAllFeatures"></ows:Exception>
 <ows:Exception exceptionCode="DuplicateStoredQueryIDValue"></ows:Exception>
</ows:ExceptionReport>`)},
	}

	for k, test := range tests {
		r := test.exceptions.ToReport().ToBytes()

		if string(r) != string(test.result) {
			t.Errorf("test: %d, expected: %s\n got: %s", k, test.result, r)
		}
	}
}
