package wmts100

import (
	"encoding/xml"
	"regexp"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

// Type function needed for the interface
func (gc GetCapabilitiesResponse) Type() string {
	return getcapabilities
}

// Service function needed for the interface
func (gc GetCapabilitiesResponse) Service() string {
	return Service
}

// Version function needed for the interface
func (gc GetCapabilitiesResponse) Version() string {
	return Version
}

// Validate function of the wfs200 spec
func (gc GetCapabilitiesResponse) Validate() wsc110.Exceptions {
	return nil
}

// ToXML builds a GetCapabilities response object
func (gc GetCapabilitiesResponse) ToXML() []byte {
	si, _ := xml.MarshalIndent(gc, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}

// GetCapabilitiesResponse base struct
type GetCapabilitiesResponse struct {
	XMLName               xml.Name `xml:"Capabilities" yaml:"capabilities"`
	Namespaces            `yaml:"namespaces"`
	ServiceIdentification ServiceIdentification   `xml:"ows:ServiceIdentification" yaml:"serviceIdentification"`
	ServiceProvider       *wsc110.ServiceProvider `xml:"ows:ServiceProvider,omitempty" yaml:"serviceProvider"`
	OperationsMetadata    *OperationsMetadata     `xml:"ows:OperationsMetadata,omitempty" yaml:"operationsMetadata"`
	Contents              Contents                `xml:"Contents" yaml:"contents"`
	ServiceMetadataURL    *ServiceMetadataURL     `xml:"ServiceMetadataURL,omitempty" yaml:"serviceMetadataUrl"`
}

// Namespaces struct containing the namespaces needed for the XML document
type Namespaces struct {
	Xmlns          string `xml:"xmlns,attr" yaml:"xmlns"`       //http://www.opengis.net/wmts/1.0
	XmlnsOws       string `xml:"xmlns:ows,attr" yaml:"common"`  //http://www.opengis.net/ows/1.1
	XmlnsXlink     string `xml:"xmlns:xlink,attr" yaml:"xlink"` //http://www.w3.org/1999/xlink
	XmlnsXSI       string `xml:"xmlns:xsi,attr" yaml:"xsi"`     //http://www.w3.org/2001/XMLSchema-instance
	XmlnsGml       string `xml:"xmlns:gml,attr" yaml:"gml"`     //http://www.opengis.net/gml
	Version        string `xml:"version,attr" yaml:"version"`
	SchemaLocation string `xml:"xsi:schemaLocation,attr" yaml:"schemaLocation"`
}

type OperationsMetadata struct {
	XMLName   xml.Name    `xml:"ows:OperationsMetadata" yaml:"operationsMetadata"`
	Operation []Operation `xml:"ows:Operation" yaml:"operation"`
}

// Operation struct for the WFS 2.0.0
type Operation struct {
	Name string `xml:"name,attr" yaml:"name"`
	DCP  struct {
		HTTP struct {
			Get  *Method `xml:"ows:Get,omitempty" yaml:"get,omitempty"`
			Post *Method `xml:"ows:Post,omitempty" yaml:"post,omitempty"`
		} `xml:"ows:HTTP" yaml:"http"`
	} `xml:"ows:DCP" yaml:"dcp"`
	Parameter []struct {
		Name          string `xml:"name,attr" yaml:"name"`
		AllowedValues struct {
			Value []string `xml:"ows:Value" yaml:"value"`
		} `xml:"ows:AllowedValues" yaml:"allowedValues"`
	} `xml:"ows:Parameter" yaml:"parameter"`
	Constraints []struct {
		Name         string `xml:"name,attr" yaml:"name"`
		NoValues     string `xml:"ows:NoValues" yaml:"noValues"`
		DefaultValue string `xml:"ows:DefaultValue" yaml:"defaultValue"`
	} `xml:"ows:Constraint" yaml:"constraint"`
}

// Method in separated struct so to use it as a Pointer
type Method struct {
	Type       string `xml:"xlink:type,attr" yaml:"type"`
	Href       string `xml:"xlink:href,attr" yaml:"href"`
	Constraint []struct {
		Name          string `xml:"name,attr" yaml:"name"`
		AllowedValues struct {
			Value []string `xml:"ows:Value" yaml:"value"`
		} `xml:"ows:AllowedValues" yaml:"allowedValues"`
	} `xml:"ows:Constraint" yaml:"constraint"`
}

// ServiceIdentification struct should only be fill by the "template" configuration wmts100.yaml
type ServiceIdentification struct {
	Title              string           `xml:"ows:Title" yaml:"title"`
	Abstract           string           `xml:"ows:Abstract" yaml:"abstract"`
	Keywords           *wsc110.Keywords `xml:"ows:Keywords,omitempty" yaml:"keywords"`
	ServiceType        string           `xml:"ows:ServiceType" yaml:"serviceType"`
	ServiceTypeVersion string           `xml:"ows:ServiceTypeVersion" yaml:"serviceTypeVersion"`
	Fees               string           `xml:"ows:Fees" yaml:"fees"`
	AccessConstraints  string           `xml:"ows:AccessConstraints" yaml:"accessConstraints"`
}

// ServiceMetadataURL in struct for repeatability
type ServiceMetadataURL struct {
	Href string `xml:"xlink:href,attr" yaml:"href"`
}
