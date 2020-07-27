package response

import (
	"encoding/xml"
	"regexp"

	"github.com/pdok/ogc-specifications/pkg/ows"
	"github.com/pdok/ogc-specifications/pkg/wms130/capabilities"
)

// Contains the WMS130 struct

//
const (
	getcapabilities = `GetCapabilities`
)

// Type and Version as constant
const (
	Service string = `WMS`
	Version string = `1.3.0`
)

// Type function needed for the interface
func (gc *GetCapabilities) Type() string {
	return getcapabilities
}

// Service function needed for the interface
func (gc *GetCapabilities) Service() string {
	return Service
}

// Version function needed for the interface
func (gc *GetCapabilities) Version() string {
	return Version
}

// Validate function of the wms130 spec
func (gc *GetCapabilities) Validate() ows.Exception {
	return nil
}

// BuildXML builds a GetCapabilities response object
func (gc *GetCapabilities) BuildXML() []byte {
	si, _ := xml.MarshalIndent(gc, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}

// GetCapabilities base struct
type GetCapabilities struct {
	XMLName    xml.Name `xml:"WMS_Capabilities"`
	Namespaces `yaml:"namespaces"`
	WMSService WMSService              `xml:"Service" yaml:"service"`
	Capability capabilities.Capability `xml:"Capability" yaml:"capability"`
}

// Namespaces struct containing the namespaces needed for the XML document
type Namespaces struct {
	XmlnsWMS           string `xml:"xmlns,attr" yaml:"wms"`                                              //http://www.opengis.net/wms
	XmlnsSLD           string `xml:"xmlns:sld,attr" yaml:"sld"`                                          //http://www.opengis.net/sld
	XmlnsXlink         string `xml:"xmlns:xlink,attr" yaml:"xlink"`                                      //http://www.w3.org/1999/xlink
	XmlnsXSI           string `xml:"xmlns:xsi,attr" yaml:"xsi"`                                          //http://www.w3.org/2001/XMLSchema-instance
	XmlnsInspireCommon string `xml:"xmlns:inspire_common,attr,omitempty" yaml:"inspirecommon,omitempty"` //http://inspire.ec.europa.eu/schemas/common/1.0
	XmlnsInspireVs     string `xml:"xmlns:inspire_vs,attr,omitempty" yaml:"inspirevs,omitempty"`         //http://inspire.ec.europa.eu/schemas/inspire_vs/1.0
	Version            string `xml:"version,attr" yaml:"version"`
	SchemaLocation     string `xml:"xsi:schemaLocation,attr" yaml:"schemalocation"`
}

// WMSService struct containing the base service information filled from the template
type WMSService struct {
	Name        string `xml:"Name" yaml:"name"`
	Title       string `xml:"Title" yaml:"title"`
	Abstract    string `xml:"Abstract" yaml:"abstract"`
	KeywordList struct {
		Keyword []string `xml:"Keyword" yaml:"keyword"`
	} `xml:"KeywordList" yaml:"keywordlist"`
	OnlineResource struct {
		Xlink *string `xml:"xmlns:xlink,attr" yaml:"xlink"`
		Type  *string `xml:"xlink:type,attr" yaml:"type"`
		Href  *string `xml:"xlink:href,attr" yaml:"href"`
	} `xml:"OnlineResource" yaml:"onlineresource"`
	ContactInformation struct {
		ContactPersonPrimary struct {
			ContactPerson       string `xml:"ContactPerson" yaml:"contactperson"`
			ContactOrganization string `xml:"ContactOrganization" yaml:"contactorganization"`
		} `xml:"ContactPersonPrimary" yaml:"contactpersonprimary"`
		ContactPosition string `xml:"ContactPosition" yaml:"contactposition"`
		ContactAddress  struct {
			AddressType     string `xml:"AddressType" yaml:"addresstype"`
			Address         string `xml:"Address" yaml:"address"`
			City            string `xml:"City" yaml:"city"`
			StateOrProvince string `xml:"StateOrProvince" yaml:"stateorprovince"`
			PostCode        string `xml:"PostCode" yaml:"postalcode"`
			Country         string `xml:"Country" yaml:"country"`
		} `xml:"ContactAddress" yaml:"contactaddress"`
		ContactVoiceTelephone        string `xml:"ContactVoiceTelephone" yaml:"contactvoicetelephone"`
		ContactFacsimileTelephone    string `xml:"ContactFacsimileTelephone" yaml:"contactfacsimiletelephone"`
		ContactElectronicMailAddress string `xml:"ContactElectronicMailAddress" yaml:"contactelectronicmailaddress"`
	} `xml:"ContactInformation"`
	Fees              string `xml:"Fees" yaml:"fees"`
	AccessConstraints string `xml:"AccessConstraints" yaml:"accessconstraints"`
	MaxWidth          string `xml:"MaxWidth" yaml:"maxwidth"`
	MaxHeight         string `xml:"MaxHeight" yaml:"maxheight"`
}
