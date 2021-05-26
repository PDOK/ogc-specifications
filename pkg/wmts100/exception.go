package wmts100

import (
	"encoding/xml"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

type exception struct {
	XMLName       xml.Name `xml:"ows:Exception"`
	ExceptionText string   `xml:",chardata" yaml:"exception"`
	ExceptionCode string   `xml:"exceptionCode,attr" yaml:"exceptioncode"`
	LocatorCode   string   `xml:"locator,attr,omitempty" yaml:"locationcode"`
}

// ToExceptions promotes a single exception to an array of one
func (e exception) ToExceptions() []wsc110.Exception {
	return []wsc110.Exception{e}
}

// Error returns available ExceptionText
func (e exception) Error() string {
	return e.ExceptionText
}

// Code returns available ExceptionCode
func (e exception) Code() string {
	return e.ExceptionCode
}

// Locator returns available ExceptionCode
func (e exception) Locator() string {
	return e.LocatorCode
}
