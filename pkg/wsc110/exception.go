package wsc110

import (
	"encoding/xml"

	"github.com/pdok/ogc-specifications/pkg/common"
)

// const (
// 	version = `1.0.0`
// )

// wsc110exception
type exception struct {
	XMLName       xml.Name `xml:"ows:Exception"`
	ExceptionText string   `xml:",chardata" yaml:"exception"`
	ExceptionCode string   `xml:"exceptionCode,attr" yaml:"exceptioncode"`
	LocatorCode   string   `xml:"locator,attr,omitempty" yaml:"locationcode"`
}

type Exceptions []common.Exception

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
