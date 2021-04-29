package common

import (
	"testing"
)

func TestOWSException(t *testing.T) {
	var tests = []struct {
		exception     Exception
		exceptionText string
		exceptionCode string
		locatorCode   string
	}{
		0: {exception: exception{ExceptionCode: "", ExceptionText: "", LocatorCode: ""},
			exceptionText: "",
			exceptionCode: "",
			locatorCode:   "",
		},
		1: {exception: OperationNotSupported("GetCoconut"),
			exceptionText: "This service does not know the operation: GetCoconut",
			exceptionCode: "OperationNotSupported",
			locatorCode:   "GetCoconut",
		},
		2: {exception: MissingParameterValue(),
			exceptionText: "Could not determine REQUEST",
			exceptionCode: "MissingParameterValue",
			locatorCode:   "REQUEST",
		},
		3: {exception: MissingParameterValue("VERSION"),
			exceptionText: "Missing key: VERSION",
			exceptionCode: "MissingParameterValue",
			locatorCode:   "VERSION",
		},
		// TODO: ... is this valid
		4: {exception: MissingParameterValue("SERVICE", "1.3.0"),
			exceptionText: "SERVICE key got incorrect value: 1.3.0",
			exceptionCode: "MissingParameterValue",
			locatorCode:   "SERVICE",
		},
		5: {exception: InvalidParameterValue("WKS", "SERVICE"),
			exceptionText: "SERVICE contains a invalid value: WKS",
			exceptionCode: "InvalidParameterValue",
			locatorCode:   "WKS",
		},
		6: {exception: VersionNegotiationFailed("0.0.0"),
			exceptionText: "0.0.0 is an invalid version number",
			exceptionCode: "VersionNegotiationFailed",
			locatorCode:   "VERSION",
		},
		// TODO: ...
		7: {exception: InvalidUpdateSequence(),
			exceptionCode: "InvalidUpdateSequence",
		},
		// TODO: ...
		8: {exception: OptionNotSupported(),
			exceptionCode: "OptionNotSupported",
		},
		9: {exception: NoApplicableCode("No other exceptionCode specified by this service"),
			exceptionText: "No other exceptionCode specified by this service",
			exceptionCode: "NoApplicableCode",
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

func TestOWSReport(t *testing.T) {
	var tests = []struct {
		exceptions Exceptions
		result     []byte
		err        error
	}{
		0: {exceptions: Exceptions{exception{ExceptionCode: "", ExceptionText: "", LocatorCode: ""}},
			result: []byte(`<?xml version="1.0" encoding="UTF-8"?>
<common:ExceptionReport xmlns:common="http://www.opengis.net/common/1.1" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.opengis.net/common/1.1 http://schemas.opengis.net/common/1.1.0/owsExceptionReport.xsd" version="1.0.0" xml:lang="en">
 <common:Exception exceptionCode=""></common:Exception>
</common:ExceptionReport>`)},
		1: {exceptions: Exceptions{
			OperationNotSupported(`WKS`),
			VersionNegotiationFailed(`0.0.1`),
		},
			result: []byte(`<?xml version="1.0" encoding="UTF-8"?>
<common:ExceptionReport xmlns:common="http://www.opengis.net/common/1.1" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.opengis.net/common/1.1 http://schemas.opengis.net/common/1.1.0/owsExceptionReport.xsd" version="1.0.0" xml:lang="en">
 <common:Exception exceptionCode="OperationNotSupported" locator="WKS">This service does not know the operation: WKS</common:Exception>
 <common:Exception exceptionCode="VersionNegotiationFailed" locator="VERSION">0.0.1 is an invalid version number</common:Exception>
</common:ExceptionReport>`)},
	}

	for k, a := range tests {
		r := a.exceptions.ToReport().ToBytes()

		if string(r) != string(a.result) {
			t.Errorf("test: %d, expected: %s\n got: %s", k, a.result, r)
		}
	}
}
