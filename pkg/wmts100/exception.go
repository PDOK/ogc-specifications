package wmts100

import (
	"encoding/xml"

	"github.com/pdok/ogc-specifications/pkg/common"
	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

type exception struct {
	XMLName xml.Name `xml:"ows:Exception"`
	common.ExceptionDetails
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
