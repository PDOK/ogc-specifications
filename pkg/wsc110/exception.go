package wsc110

import (
	"encoding/xml"
)

type Exception interface {
	Error() string
	Code() string
	Locator() string
	ToExceptions() []Exception
}

// wsc110exception
type exception struct {
	XMLName       xml.Name `xml:"ows:Exception"`
	ExceptionText string   `xml:",chardata" yaml:"exception"`
	ExceptionCode string   `xml:"exceptionCode,attr" yaml:"exceptioncode"`
	LocatorCode   string   `xml:"locator,attr,omitempty" yaml:"locationcode"`
}

type Exceptions []Exception

func (e exception) ToExceptions() []Exception {
	return Exceptions{e}
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
