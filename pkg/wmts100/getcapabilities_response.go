package wmts100

import (
	"encoding/xml"
	"regexp"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

// Type function needed for the interface
func (gc *GetCapabilitiesResponse) Type() string {
	return getcapabilities
}

// Service function needed for the interface
func (gc *GetCapabilitiesResponse) Service() string {
	return Service
}

// Version function needed for the interface
func (gc *GetCapabilitiesResponse) Version() string {
	return Version
}

// Validate function of the wfs200 spec
func (gc *GetCapabilitiesResponse) Validate() wsc110.Exceptions {
	return nil
}

// ToXML builds a GetCapabilities response object
func (gc *GetCapabilitiesResponse) ToXML() []byte {
	si, _ := xml.MarshalIndent(gc, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}

// GetCapabilities base struct
type GetCapabilitiesResponse struct {
	XMLName               xml.Name `xml:"Capabilities"`
	Namespaces            `yaml:"namespaces"`
	ServiceIdentification ServiceIdentification `xml:"ows:ServiceIdentification" yaml:"serviceidentification"`
	Contents              Contents              `xml:"Contents" yaml:"contents"`
	ServiceMetadataURL    ServiceMetadataURL    `xml:"ServiceMetadataURL" yaml:"servicemetadataurl"`
}

// Namespaces struct containing the namespaces needed for the XML document
type Namespaces struct {
	Xmlns          string `xml:"xmlns,attr" yaml:"xmlns"`       //http://www.opengis.net/wmts/1.0
	XmlnsOws       string `xml:"xmlns:ows,attr" yaml:"common"`  //http://www.opengis.net/ows/1.1
	XmlnsXlink     string `xml:"xmlns:xlink,attr" yaml:"xlink"` //http://www.w3.org/1999/xlink
	XmlnsXSI       string `xml:"xmlns:xsi,attr" yaml:"xsi"`     //http://www.w3.org/2001/XMLSchema-instance
	XmlnsGml       string `xml:"xmlns:gml,attr" yaml:"gml"`     //http://www.opengis.net/gml
	Version        string `xml:"version,attr" yaml:"version"`
	SchemaLocation string `xml:"xsi:schemaLocation,attr" yaml:"schemalocation"`
}

// ServiceIdentification struct should only be fill by the "template" configuration wmts100.yaml
type ServiceIdentification struct {
	Title              string `xml:"ows:Title" yaml:"title"`
	Abstract           string `xml:"ows:Abstract" yaml:"abstract"`
	ServiceType        string `xml:"ows:ServiceType" yaml:"servicetype"`
	ServiceTypeVersion string `xml:"ows:ServiceTypeVersion" yaml:"servicetypeversion"`
	Fees               string `xml:"ows:Fees" yaml:"fees"`
	AccessConstraints  string `xml:"ows:AccessConstraints" yaml:"accessconstraints"`
}

// ServiceMetadataURL in struct for repeatability
type ServiceMetadataURL struct {
	Href string `xml:"xlink:href,attr" yaml:"href"`
}
