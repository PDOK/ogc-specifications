package wms130

import (
	"encoding/xml"
	"fmt"

	"github.com/pdok/ogc-specifications/pkg/ows"
)

// WMSServiceExceptionReport struct
// TODO exception restucturing
type WMSServiceExceptionReport struct {
	XMLName          xml.Name        `xml:"ServiceExceptionReport"`
	Version          string          `xml:"version,attr"`
	Xmlns            string          `xml:"xmlns,attr"`
	Xsi              string          `xml:"xsi,attr"`
	SchemaLocation   string          `xml:"schemaLocation,attr"`
	ServiceException []ows.Exception `xml:"ServiceException"`
}

// Report returns WMSServiceExceptionReport
func (r WMSServiceExceptionReport) Report(errors []ows.Exception) []byte {
	r.ServiceException = errors
	si, _ := xml.MarshalIndent(r, "", " ")
	return append([]byte(xml.Header), si...)
}

// WMSException grouping the error message variables together
type WMSException struct {
	ExceptionText string `xml:",chardata"`
	ErrorCode     string `xml:"code,attr"`
	LocatorCode   string `xml:"locator,attr"`
}

// Error returns available ExceptionText
func (e WMSException) Error() string {
	return e.ExceptionText
}

// Code returns available ErrorCode
func (e WMSException) Code() string {
	return e.ErrorCode
}

// Locator returns available LocatorCode
func (e WMSException) Locator() string {
	return e.LocatorCode
}

// InvalidFormat exception
func InvalidFormat() WMSException {
	return WMSException{
		ErrorCode: `InvalidFormat`,
	}
}

// InvalidCRS exception
func InvalidCRS(s ...string) WMSException {
	if len(s) == 1 {
		return WMSException{
			ExceptionText: s[0],
			ErrorCode:     `InvalidCRS`,
		}
	}
	return WMSException{
		ErrorCode: `InvalidCRS`,
	}
}

// LayerNotDefined exception
func LayerNotDefined(undefinedlayer string) WMSException {
	return WMSException{
		ExceptionText: fmt.Sprintf("The layer: %s is not known by the server", undefinedlayer),
		ErrorCode:     `LayerNotDefined`,
	}
}

// StyleNotDefined exception
func StyleNotDefined() WMSException {
	return WMSException{
		ExceptionText: `There is a one-to-one correspondence between the values in the LAYERS parameter and the values in the STYLES parameter. 
	Expecting an empty string for the STYLES like STYLES= or comma-separated list STYLES=,,, or using keyword default STYLES=default,default,...`,
		ErrorCode: `StyleNotDefined`,
	}
}

// LayerNotQueryable exception
func LayerNotQueryable() WMSException {
	return WMSException{
		ErrorCode: `LayerNotQueryable`,
	}
}

// InvalidPoint exception
func InvalidPoint() WMSException {
	return WMSException{
		ErrorCode: `InvalidPoint`,
	}
}

// CurrentUpdateSequence exception
func CurrentUpdateSequence() WMSException {
	return WMSException{
		ErrorCode: `CurrentUpdateSequence`,
	}
}

// InvalidUpdateSequence exception
func InvalidUpdateSequence() WMSException {
	return WMSException{
		ErrorCode: `InvalidUpdateSequence`,
	}
}

// MissingDimensionValue exception
func MissingDimensionValue() WMSException {
	return WMSException{
		ErrorCode: `MissingDimensionValue`,
	}
}

// InvalidDimensionValue exception
func InvalidDimensionValue() WMSException {
	return WMSException{
		ErrorCode: `InvalidDimensionValue`,
	}
}

// OperationNotSupported exception -> available in OWS Exceptions
func OperationNotSupported() WMSException {
	// TODO Use the error.OperationNotSupported instead of this one
	return WMSException{}
}
