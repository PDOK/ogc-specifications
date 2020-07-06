package ows

import (
	"encoding/xml"
	"fmt"
)

const (
	version = `1.0.0`
)

// ExceptionReport interface
type ExceptionReport interface {
	Report([]Exception) []byte
}

// OWSExceptionReport struct
type OWSExceptionReport struct {
	xmlname        xml.Name    `xml:"ExceptionReport"`
	ows            string      `xml:"xmlns:ows,attr"`
	xsi            string      `xml:"xmlns:xsi,attr"`
	schemaLocation string      `xml:"xsi:schemaLocation,attr"`
	version        string      `xml:"version,attr"`
	language       string      `xml:"xml:lang,attr"`
	Exception      []Exception `xml:"Exception"`
}

// Report returns OWSExceptionReport
func (r OWSExceptionReport) Report(errors []Exception) []byte {
	r.schemaLocation = `http://www.opengis.net/ows/1.1 http://schemas.opengis.net/ows/1.1.0/owsExceptionReport.xsd`
	r.ows = `http://www.opengis.net/ows/1.1`
	r.xsi = `http://www.w3.org/2001/XMLSchema-instance`
	r.version = version
	r.language = `en`
	r.Exception = errors

	si, _ := xml.MarshalIndent(r, "", " ")
	return append([]byte(xml.Header), si...)
}

// Exception interfact wraps the three variables:
// Error
// Code
// Locator
type Exception interface {
	Error() string
	Code() string
	Locator() string
}

// OWSException grouping the error message variables together
type OWSException struct {
	ExceptionText string `xml:",chardata"`
	ExceptionCode string `xml:"exceptionCode,attr"`
	LocatorCode   string `xml:"locator,attr,omitempty"`
}

// Error returns available ExceptionText
func (e OWSException) Error() string {
	return e.ExceptionText
}

// Code returns available ExceptionCode
func (e OWSException) Code() string {
	return e.ExceptionCode
}

// Locator returns available ExceptionCode
func (e OWSException) Locator() string {
	return e.LocatorCode
}

// OperationNotSupported exception
func OperationNotSupported(message string) OWSException {
	return OWSException{
		ExceptionText: fmt.Sprintf("This service does not know the operation: %s", message),
		ExceptionCode: `OperationNotSupported`,
		LocatorCode:   message,
	}
}

// MissingParameterValue exception
func MissingParameterValue(s ...string) OWSException {
	if len(s) >= 2 {
		return OWSException{ExceptionText: fmt.Sprintf("%s key got incorrect value: %s", s[0], s[1]), ExceptionCode: "MissingParameterValue", LocatorCode: s[0]}
	}
	if len(s) == 1 {
		return OWSException{ExceptionText: fmt.Sprintf("Missing key: %s", s[0]), ExceptionCode: "MissingParameterValue", LocatorCode: s[0]}
	}

	return OWSException{ExceptionText: `Could not determine REQUEST`, ExceptionCode: "MissingParameterValue", LocatorCode: "REQUEST"}
}

// InvalidParameterValue exception
func InvalidParameterValue(value, locator string) OWSException {
	return OWSException{
		ExceptionText: fmt.Sprintf("%s %s does not exist in this server. Please check the capabilities and reformulate your request", locator, value),
		LocatorCode:   value,
		ExceptionCode: `InvalidParameterValue`,
	}
}

// VersionNegotiationFailed exception
func VersionNegotiationFailed(version string) OWSException {
	return OWSException{
		ExceptionText: fmt.Sprintf("%s is an invalid version number", version),
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
		ExceptionText: message,
		ExceptionCode: `NoApplicableCode`,
	}
}
