package wms130

import (
	"encoding/xml"
	"github.com/pdok/ogc-specifications/pkg/common"
)

// exception
type exception struct {
	common.ExceptionDetails
}

// Exceptions is a array of the Exception interface
type Exceptions []exception

// ServiceExceptionReport struct
type ServiceExceptionReport struct {
	XMLName          xml.Name   `xml:"ServiceExceptionReport" yaml:"serviceExceptionReport"`
	Version          string     `xml:"version,attr" yaml:"version"`
	Xmlns            string     `xml:"xmlns,attr,omitempty" yaml:"xmlns"`
	Xsi              string     `xml:"xsi,attr,omitempty" yaml:"xsi"`
	SchemaLocation   string     `xml:"schemaLocation,attr,omitempty" yaml:"schemaLocation"`
	ServiceException Exceptions `xml:"ServiceException" yaml:"serviceException"`
}

// ToReport builds a ServiceExceptionReport from an array of Exceptions
func (e Exceptions) ToReport() ServiceExceptionReport {
	r := ServiceExceptionReport{}
	r.SchemaLocation = `http://www.opengis.net/ogc http://schemas.opengis.net/wms/1.3.0/exceptions_1_3_0.xsd`
	r.Xmlns = `http://www.opengis.net/ogc`
	r.Xsi = `http://www.w3.org/2001/XMLSchema-instance`
	r.Version = Version
	r.ServiceException = e
	return r
}

// ToBytes makes from a ServiceExceptionReport a []byte
func (r ServiceExceptionReport) ToBytes() []byte {
	si, _ := xml.MarshalIndent(r, "", " ")
	return append([]byte(xml.Header), si...)
}

// ToExceptions promotes a single exception to an array of one
func (e exception) ToExceptions() Exceptions {
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
