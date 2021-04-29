package common

import (
	"encoding/xml"
)

const (
	version = `1.0.0`
)

type Exception interface {
	Error() string
	Code() string
	Locator() string
}

type exception struct {
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
	Exception      Exceptions `xml:"Exception"`
}

func (e Exceptions) ToReport() ExceptionReport {
	// TODO validate all exceptions are of type common.exception
	r := ExceptionReport{}
	r.SchemaLocation = `http://www.opengis.net/ows/1.1 http://schemas.opengis.net/ows/1.1.0/owsExceptionReport.xsd`
	r.Ows = `http://www.opengis.net/ows/1.1`
	r.Xsi = `http://www.w3.org/2001/XMLSchema-instance`
	r.Version = version
	r.Language = `en`
	r.Exception = e
	return r
}

func (r ExceptionReport) ToBytes() []byte {
	si, _ := xml.MarshalIndent(r, "", " ")
	return append([]byte(xml.Header), si...)
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
