package wsc110

import (
	"encoding/xml"
)

const (
	Version string = `1.1.0`
)

// wsc110exception
type Exception struct {
	XMLName       xml.Name `xml:"ows:Exception"`
	ExceptionText string   `xml:",chardata" yaml:"exception"`
	ExceptionCode string   `xml:"exceptionCode,attr" yaml:"exceptioncode"`
	LocatorCode   string   `xml:"locator,attr,omitempty" yaml:"locationcode"`
}

type Exceptions []Exception

type ExceptionReport struct {
	XMLName        xml.Name   `xml:"ows:ExceptionReport" yaml:"exceptionreport"`
	Ows            string     `xml:"xmlns:ows,attr,omitempty"`
	Xsi            string     `xml:"xmlns:xsi,attr,omitempty"`
	SchemaLocation string     `xml:"xsi:schemaLocation,attr,omitempty"`
	Version        string     `xml:"version,attr" yaml:"version"`
	Language       string     `xml:"xml:lang,attr,omitempty" yaml:"lang,omitempty"`
	Exception      Exceptions `xml:"ows:Exception"`
}

func (e Exceptions) ToReport() ExceptionReport {
	r := ExceptionReport{}
	r.SchemaLocation = `http://www.opengis.net/ows/1.1 http://schemas.opengis.net/ows/1.1.0/owsExceptionReport.xsd`
	r.Ows = `http://www.opengis.net/ows/1.1`
	r.Xsi = `http://www.w3.org/2001/XMLSchema-instance`
	r.Version = Version
	r.Language = `en`
	r.Exception = e
	return r
}

func (r ExceptionReport) ToBytes() []byte {
	si, _ := xml.MarshalIndent(r, "", " ")
	return append([]byte(xml.Header), si...)
}

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
