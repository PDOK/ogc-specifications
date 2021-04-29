package wms130

import (
	"encoding/xml"
	"github.com/pdok/ogc-specifications/pkg/common"
)

const (
	Service = `WMS`
	Version = `1.3.0`
)

type Exception common.Exception

type exception struct {
	ExceptionText string `xml:",chardata" yaml:"exception"`
	ExceptionCode string `xml:"code,attr" yaml:"code"`
	LocatorCode   string `xml:"locator,attr,omitempty" yaml:"locator,omitempty"`
}

type Exceptions []Exception

type WMSServiceExceptionReport struct {
	XMLName          xml.Name   `xml:"ServiceExceptionReport" yaml:"serviceexceptionreport"`
	Version          string     `xml:"version,attr" yaml:"version"`
	Xmlns            string     `xml:"xmlns,attr,omitempty"`
	Xsi              string     `xml:"xsi,attr,omitempty"`
	SchemaLocation   string     `xml:"schemaLocation,attr,omitempty"`
	ServiceException Exceptions `xml:"ServiceException"`
}

func (e Exceptions) ToReport() WMSServiceExceptionReport {
	r := WMSServiceExceptionReport{}
	r.Version = Version
	r.Xmlns = `http://www.opengis.net/ogc`
	r.Xsi = `http://www.w3.org/2001/XMLSchema-instance`
	r.SchemaLocation = `http://www.opengis.net/ogc http://schemas.opengis.net/wms/1.3.0/exceptions_1_3_0.xsd`
	r.ServiceException = e
	return r
}

func (r WMSServiceExceptionReport) ToBytes() []byte {
	si, _ := xml.MarshalIndent(r, "", " ")
	return append([]byte(xml.Header), si...)
}

func (e exception) Error() string {
	return e.ExceptionText
}

func (e exception) Code() string {
	return e.ExceptionCode
}

func (e exception) Locator() string {
	return e.LocatorCode
}
