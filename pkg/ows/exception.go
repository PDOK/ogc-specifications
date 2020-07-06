package ows

import (
	"fmt"
)

// ExceptionReport interface
type ExceptionReport interface {
	Report(Exception) []byte
}

// Exception interfact wraps the two variables:
// Error
// Code
type Exception interface {
	Error() string
	Code() string
}

// OWSException grouping the error message variables together
type OWSException struct {
	ErrorMessage  string `xml:",chardata"`
	ExceptionCode string `xml:"exceptionCode,attr"`
	LocatorCode   string `xml:"locator,attr"`
}

// Error returns available ErrorMessage
func (e OWSException) Error() string {
	return e.ErrorMessage
}

// Code returns available ExceptionCode
func (e OWSException) Code() string {
	return e.ExceptionCode
}

// OperationNotSupported exception
func OperationNotSupported(message string) OWSException {
	return OWSException{
		ErrorMessage:  message,
		ExceptionCode: `OperationNotSupported`,
	}
}

// MissingParameterValue exception
func MissingParameterValue(s ...string) OWSException {
	if len(s) >= 2 {
		return OWSException{ErrorMessage: fmt.Sprintf("%s key got incorrect value: %s", s[0], s[1]), ExceptionCode: "MissingParameterValue"}
	}
	if len(s) == 1 {
		return OWSException{ErrorMessage: fmt.Sprintf("Missing key: %s", s[0]), ExceptionCode: "MissingParameterValue"}
	}

	return OWSException{ErrorMessage: `Could not determine REQUEST`, ExceptionCode: "MissingParameterValue"}
}

// InvalidParameterValue exception
func InvalidParameterValue(value, locator string) OWSException {
	return OWSException{
		ErrorMessage:  fmt.Sprintf("%s %s does not exist in this server. Please check the capabilities and reformulate your request", locator, value),
		LocatorCode:   locator,
		ExceptionCode: `InvalidParameterValue`,
	}
}

// VersionNegotiationFailed exception
func VersionNegotiationFailed(version string) OWSException {
	return OWSException{
		ErrorMessage:  fmt.Sprintf("%s is an invalid version number", version),
		ExceptionCode: `VersionNegotiationFailed`,
		LocatorCode:   "VERSION",
	}
}

// InvalidUpdateSequence exception
func InvalidUpdateSequence() OWSException {
	return OWSException{
		ExceptionCode: `InvalidUpdateSequence`,
	}
}

// OptionNotSupported exception
func OptionNotSupported() OWSException {
	return OWSException{
		ExceptionCode: `OptionNotSupported`,
	}
}

// NoApplicableCode exception
func NoApplicableCode(message string) OWSException {
	return OWSException{
		ErrorMessage:  message,
		ExceptionCode: `NoApplicableCode`,
	}
}
