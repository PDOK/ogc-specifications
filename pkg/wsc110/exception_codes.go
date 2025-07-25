package wsc110

import (
	"fmt"

	"github.com/pdok/ogc-specifications/pkg/common"
)

// OperationNotSupported exception
func OperationNotSupported(message string) Exception {
	return exception{
		ExceptionDetails: common.ExceptionDetails{
			ExceptionText: "This service does not know the operation: " + message,
			ExceptionCode: `OperationNotSupported`,
			LocatorCode:   message,
		},
	}
}

// MissingParameterValue exception
func MissingParameterValue(s ...string) Exception {
	if len(s) >= 2 {
		return exception{ExceptionDetails: common.ExceptionDetails{ExceptionText: fmt.Sprintf("%s key got incorrect value: %s", s[0], s[1]), ExceptionCode: "MissingParameterValue", LocatorCode: s[0]}}
	}
	if len(s) == 1 {
		return exception{ExceptionDetails: common.ExceptionDetails{ExceptionText: "Missing key: " + s[0], ExceptionCode: "MissingParameterValue", LocatorCode: s[0]}}
	}

	return exception{ExceptionDetails: common.ExceptionDetails{ExceptionText: `Could not determine REQUEST`, ExceptionCode: "MissingParameterValue", LocatorCode: "REQUEST"}}
}

// InvalidParameterValue exception
func InvalidParameterValue(value, locator string) Exception {
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionText: fmt.Sprintf("%s contains a invalid value: %s", locator, value),
		LocatorCode:   value,
		ExceptionCode: `InvalidParameterValue`,
	}}
}

// VersionNegotiationFailed exception
func VersionNegotiationFailed(version string) Exception {
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionText: version + " is an invalid version number",
		ExceptionCode: `VersionNegotiationFailed`,
		LocatorCode:   "VERSION",
	}}
}

// InvalidUpdateSequence exception
func InvalidUpdateSequence() Exception {
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionCode: `InvalidUpdateSequence`,
	}}
}

// OptionNotSupported exception
func OptionNotSupported(s ...string) Exception {
	if len(s) == 1 {
		return exception{ExceptionDetails: common.ExceptionDetails{
			ExceptionText: s[0],
			ExceptionCode: `OptionNotSupported`,
		}}
	}
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionCode: `OptionNotSupported`,
	}}
}

// NoApplicableCode exception
func NoApplicableCode(message string) Exception {
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionText: message,
		ExceptionCode: `NoApplicableCode`,
	}}
}
