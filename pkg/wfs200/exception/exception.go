package exception

import (
	"encoding/xml"
	"fmt"

	"github.com/pdok/ogc-specifications/pkg/ows"
)

//
const (
	Service = `WFS`
	Version = `2.0.0`
)

// WFSExceptionReport struct
// TODO exception restucturing
type WFSExceptionReport struct {
	XMLName        xml.Name        `xml:"ExceptionReport" yaml:"exceptionreport"`
	Ows            string          `xml:"xmlns:ows,attr,omitempty"`
	Xsi            string          `xml:"xmlns:xsi,attr,omitempty"`
	SchemaLocation string          `xml:"xsi:schemaLocation,attr,omitempty"`
	Version        string          `xml:"version,attr" yaml:"version"`
	Language       string          `xml:"xml:lang,attr,omitempty" yaml:"lang,omitempty"`
	Exception      []ows.Exception `xml:"Exception"`
}

// Report returns WFSExceptionReport
func (r WFSExceptionReport) Report(errors []ows.Exception) []byte {
	r.Ows = `http://www.opengis.net/ows/1.1`
	r.Xsi = `http://www.w3.org/2001/XMLSchema-instance`
	r.SchemaLocation = `http://www.opengis.net/ows/1.1 http://schemas.opengis.net/ows/1.1.0/owsExceptionReport.xsd`
	r.Version = Version
	r.Language = `en`
	r.Exception = errors

	si, _ := xml.MarshalIndent(r, "", " ")
	return append([]byte(xml.Header), si...)
}

// WFSException struct
type WFSException struct {
	ExceptionText string `xml:",chardata" yaml:"exception"`
	ExceptionCode string `xml:"exceptionCode,attr" yaml:"exceptioncode"`
	LocatorCode   string `xml:"locator,attr,omitempty" yaml:"locatorcode,omitempty"`
	// ExceptionText string `xml:"ExceptionText"`
}

// Error returns available ExceptionText
func (e WFSException) Error() string {
	return e.ExceptionText
}

// Code returns available ErrorCode
func (e WFSException) Code() string {
	return e.ExceptionCode
}

// Locator returns available LocatorCode
func (e WFSException) Locator() string {
	return e.LocatorCode
}

// CannotLockAllFeatures exception
func CannotLockAllFeatures() WFSException {
	return WFSException{
		ExceptionCode: "CannotLockAllFeatures",
	}
}

// DuplicateStoredQueryIDValue exception
func DuplicateStoredQueryIDValue() WFSException {
	return WFSException{
		ExceptionCode: "DuplicateStoredQueryIDValue",
	}
}

// DuplicateStoredQueryParameterName exception
func DuplicateStoredQueryParameterName() WFSException {
	return WFSException{
		ExceptionCode: "DuplicateStoredQueryParameterName",
	}
}

// FeaturesNotLocked exception
func FeaturesNotLocked() WFSException {
	return WFSException{
		ExceptionCode: "FeaturesNotLocked",
	}
}

// InvalidLockID exception
func InvalidLockID() WFSException {
	return WFSException{
		ExceptionCode: "InvalidLockID",
	}
}

// InvalidValue exception
func InvalidValue(s ...string) WFSException {
	if len(s) == 1 {
		return WFSException{ExceptionText: fmt.Sprintf("The parameter: %s, contains a InvalidValue", s[0]),
			ExceptionCode: "InvalidValue",
			LocatorCode:   s[0]}
	}
	return WFSException{
		ExceptionCode: "InvalidValue",
	}
}

// LockHasExpired exception
func LockHasExpired() WFSException {
	return WFSException{
		ExceptionCode: "LockHasExpired",
	}
}

// OperationParsingFailed exception
func OperationParsingFailed(value, locator string) WFSException {
	return WFSException{
		ExceptionText: fmt.Sprintf("Failed to parse the operation, found: %s", value),
		LocatorCode:   locator,
		ExceptionCode: "OperationParsingFailed"}
}

// OperationProcessingFailed exception
func OperationProcessingFailed() WFSException {
	return WFSException{
		ExceptionCode: "OperationProcessingFailed",
	}
}

// ResponseCacheExpired exception
func ResponseCacheExpired() WFSException {
	return WFSException{
		ExceptionCode: "ResponseCacheExpired",
	}
}
