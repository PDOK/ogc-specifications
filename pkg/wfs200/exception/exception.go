package exception

import (
	"encoding/xml"
)

const (
	Service = `WFS`
	Version = `2.0.0`
)

type WFSException struct {
	ExceptionText string `xml:",chardata" yaml:"exception"`
	ExceptionCode string `xml:"exceptionCode,attr" yaml:"exceptioncode"`
	LocatorCode   string `xml:"locator,attr,omitempty" yaml:"locatorcode,omitempty"`
}

type InvalidValueException WFSException

type WFSExceptions []WFSException

type WFSExceptionReport struct {
	XMLName        xml.Name      `xml:"ExceptionReport" yaml:"exceptionreport"`
	Ows            string        `xml:"xmlns:common,attr,omitempty"`
	Xsi            string        `xml:"xmlns:xsi,attr,omitempty"`
	SchemaLocation string        `xml:"xsi:schemaLocation,attr,omitempty"`
	Version        string        `xml:"version,attr" yaml:"version"`
	Language       string        `xml:"xml:lang,attr,omitempty" yaml:"lang,omitempty"`
	Exception      WFSExceptions `xml:"Exception"`
}

func (e WFSExceptions) ToReport() WFSExceptionReport {
	r := WFSExceptionReport{}
	r.Ows = `http://www.opengis.net/common/1.1`
	r.Xsi = `http://www.w3.org/2001/XMLSchema-instance`
	r.SchemaLocation = `http://www.opengis.net/common/1.1 http://schemas.opengis.net/common/1.1.0/owsExceptionReport.xsd`
	r.Version = Version
	r.Language = `en`
	r.Exception = e
	return r
}

func (r WFSExceptionReport) ToBytes() []byte {
	si, _ := xml.MarshalIndent(r, "", " ")
	return append([]byte(xml.Header), si...)
}

func (e WFSException) Error() string {
	return e.ExceptionText
}

func (e WFSException) Code() string {
	return e.ExceptionCode
}

func (e WFSException) Locator() string {
	return e.LocatorCode
}
