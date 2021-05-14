package wsc110

import (
	"encoding/xml"
)

// wsc110exception
type Exception struct {
	XMLName       xml.Name `xml:"ows:Exception"`
	ExceptionText string   `xml:",chardata" yaml:"exception"`
	ExceptionCode string   `xml:"exceptionCode,attr" yaml:"exceptioncode"`
	LocatorCode   string   `xml:"locator,attr,omitempty" yaml:"locationcode"`
}

type Exceptions []Exception

func (e Exception) ToExceptions() Exceptions {
	return Exceptions{e}
}

// Error returns available ExceptionText
func (e Exception) Error() string {
	return e.ExceptionText
}

// Code returns available ExceptionCode
func (e Exception) Code() string {
	return e.ExceptionCode
}

// Locator returns available ExceptionCode
func (e Exception) Locator() string {
	return e.LocatorCode
}
