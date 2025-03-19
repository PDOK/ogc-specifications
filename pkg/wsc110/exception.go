package wsc110

import (
	"encoding/xml"
)

// Exception interface
type Exception interface {
	Error() string
	Code() string
	Locator() string
	ToExceptions() []Exception
}

// exception
type exception struct {
	XMLName       xml.Name `xml:"ows:Exception"`
	ExceptionText string   `xml:",chardata" yaml:"exception"`
	ExceptionCode string   `xml:"exceptionCode,attr" yaml:"exceptionCode"`
	LocatorCode   string   `xml:"locator,attr,omitempty" yaml:"locatorCode"`
}

// ExceptionReport struct
type ExceptionReport struct {
	XMLName        xml.Name   `xml:"ows:ExceptionReport" yaml:"exceptionReport"`
	Ows            string     `xml:"xmlns:ows,attr,omitempty" yaml:"ows"`
	Xsi            string     `xml:"xmlns:xsi,attr,omitempty" yaml:"xsi"`
	SchemaLocation string     `xml:"xsi:schemaLocation,attr,omitempty" yaml:"schemaLocation"`
	Version        string     `xml:"version,attr" yaml:"version"`
	Language       string     `xml:"xml:lang,attr,omitempty" yaml:"lang,omitempty"`
	Exception      Exceptions `xml:"ows:Exception" yaml:"exception"`
}

// Exceptions is a array of the Exception interface
type Exceptions []Exception

// ToReport builds a ExceptionReport from an array of Exceptions
func (e Exceptions) ToReport(version string) ExceptionReport {
	r := ExceptionReport{}
	r.SchemaLocation = `http://www.opengis.net/ows/1.1 http://schemas.opengis.net/ows/1.1.0/owsExceptionReport.xsd`
	r.Ows = `http://www.opengis.net/ows/1.1`
	r.Xsi = `http://www.w3.org/2001/XMLSchema-instance`
	r.Version = version
	r.Language = `en`
	r.Exception = e
	return r
}

// ToBytes makes from a ExceptionReport a []byte
func (r ExceptionReport) ToBytes() []byte {
	si, _ := xml.MarshalIndent(r, "", " ")
	return append([]byte(xml.Header), si...)
}

// ToExceptions promotes a single exception to an array of one
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
