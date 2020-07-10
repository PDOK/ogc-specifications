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
	XMLName        xml.Name    `xml:"ows:ExceptionReport" yaml:"exceptionreport"`
	Ows            string      `xml:"xmlns:ows,attr,omitempty"`
	Xsi            string      `xml:"xmlns:xsi,attr,omitempty"`
	SchemaLocation string      `xml:"xsi:schemaLocation,attr,omitempty"`
	Version        string      `xml:"version,attr" yaml:"version"`
	Language       string      `xml:"xml:lang,attr,omitempty" yaml:"lang"`
	Exception      []Exception `xml:"Exception"`
}

// Report returns OWSExceptionReport
func (r OWSExceptionReport) Report(errors []Exception) []byte {
	r.SchemaLocation = `http://www.opengis.net/ows/1.1 http://schemas.opengis.net/ows/1.1.0/owsExceptionReport.xsd`
	r.Ows = `http://www.opengis.net/ows/1.1`
	r.Xsi = `http://www.w3.org/2001/XMLSchema-instance`
	r.Version = version
	r.Language = `en`
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
	XMLName       xml.Name `xml:"ows:Exception"`
	ExceptionText string   `xml:",chardata" yaml:"exception"`
	ExceptionCode string   `xml:"exceptionCode,attr" yaml:"exceptioncode"`
	LocatorCode   string   `xml:"locator,attr,omitempty" yaml:"locationcode"`
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
