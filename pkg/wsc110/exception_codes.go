package wsc110

import (
	"fmt"
)

// OperationNotSupported exception
func OperationNotSupported(message string) Exception {
	return exception{
		ExceptionText: fmt.Sprintf("This service does not know the operation: %s", message),
		ExceptionCode: `OperationNotSupported`,
		LocatorCode:   message,
	}
}

// MissingParameterValue exception
func MissingParameterValue(s ...string) Exception {
	if len(s) >= 2 {
		return exception{ExceptionText: fmt.Sprintf("%s key got incorrect value: %s", s[0], s[1]), ExceptionCode: "MissingParameterValue", LocatorCode: s[0]}
	}
	if len(s) == 1 {
		return exception{ExceptionText: fmt.Sprintf("Missing key: %s", s[0]), ExceptionCode: "MissingParameterValue", LocatorCode: s[0]}
	}

	return exception{ExceptionText: `Could not determine REQUEST`, ExceptionCode: "MissingParameterValue", LocatorCode: "REQUEST"}
}

// InvalidParameterValue exception
func InvalidParameterValue(value, locator string) Exception {
	return exception{
		ExceptionText: fmt.Sprintf("%s contains a invalid value: %s", locator, value),
		LocatorCode:   value,
		ExceptionCode: `InvalidParameterValue`,
	}
}

// VersionNegotiationFailed exception
func VersionNegotiationFailed(version string) Exception {
	return exception{
		ExceptionText: fmt.Sprintf("%s is an invalid version number", version),
		ExceptionCode: `VersionNegotiationFailed`,
		LocatorCode:   "VERSION",
	}
}

// InvalidUpdateSequence exception
func InvalidUpdateSequence() Exception {
	return exception{
		ExceptionCode: `InvalidUpdateSequence`,
	}
}

// OptionNotSupported exception
func OptionNotSupported(s ...string) Exception {
	if len(s) == 1 {
		return exception{
			ExceptionText: s[0],
			ExceptionCode: `OptionNotSupported`,
		}
	}
	return exception{
		ExceptionCode: `OptionNotSupported`,
	}
}

// NoApplicableCode exception
func NoApplicableCode(message string) Exception {
	return exception{
		ExceptionText: message,
		ExceptionCode: `NoApplicableCode`,
	}
}
