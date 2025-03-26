package wfs200

import (
	"encoding/xml"
	"regexp"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

// Contains the WFS200 struct

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
func (gc *GetCapabilitiesResponse) Validate() []wsc110.Exception {
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
	XMLName               xml.Name `xml:"WFS_Capabilities"`
	Namespaces            `yaml:"namespaces"`
	ServiceIdentification ServiceIdentification `xml:"ows:ServiceIdentification" yaml:"serviceIdentification"`
	ServiceProvider       ServiceProvider       `xml:"ows:ServiceProvider" yaml:"serviceProvider"`
	Capabilities          `yaml:"capabilities"`
}

// Namespaces struct containing the namespaces needed for the XML document
type Namespaces struct {
	XmlnsGML           string `xml:"xmlns:gml,attr" yaml:"gml"`                                          //http://www.opengis.net/gml/3.2
	XmlnsWFS           string `xml:"xmlns,attr" yaml:"wfs"`                                              //http://www.opengis.net/wfs/2.0
	XmlnsOWS           string `xml:"xmlns:ows,attr" yaml:"common"`                                       //http://www.opengis.net/ows/1.1
	XmlnsXlink         string `xml:"xmlns:xlink,attr" yaml:"xlink"`                                      //http://www.w3.org/1999/xlink
	XmlnsXSI           string `xml:"xmlns:xsi,attr" yaml:"xsi"`                                          //http://www.w3.org/2001/XMLSchema-instance
	XmlnsFes           string `xml:"xmlns:fes,attr" yaml:"fes"`                                          //http://www.opengis.net/fes/2.0
	XmlnsInspireCommon string `xml:"xmlns:inspire_common,attr,omitempty" yaml:"inspireCommon,omitempty"` //http://inspire.ec.europa.eu/schemas/common/1.0
	XmlnsInspireDls    string `xml:"xmlns:inspire_dls,attr,omitempty" yaml:"inspireDls,omitempty"`       //http://inspire.ec.europa.eu/schemas/inspire_dls/1.0
	XmlnsPrefix        string `xml:"xmlns:{{.Prefix}},attr" yaml:"prefix"`                               //namespace_uri placeholder
	Version            string `xml:"version,attr" yaml:"version"`
	SchemaLocation     string `xml:"xsi:schemaLocation,attr" yaml:"schemaLocation"`
}

// ServiceIdentification struct should only be fill by the "template" configuration wfs200.yaml
type ServiceIdentification struct {
	XMLName     xml.Name         `xml:"ows:ServiceIdentification"`
	Title       string           `xml:"ows:Title" yaml:"title"`
	Abstract    string           `xml:"ows:Abstract" yaml:"abstract"`
	Keywords    *wsc110.Keywords `xml:"ows:Keywords" yaml:"keywords"`
	ServiceType struct {
		Text      string `xml:",chardata" yaml:"text"`
		CodeSpace string `xml:"codeSpace,attr" yaml:"codeSpace"`
	} `xml:"ows:ServiceType" yaml:"serviceType"`
	ServiceTypeVersion string `xml:"ows:ServiceTypeVersion" yaml:"serviceTypeVersion"`
	Fees               string `xml:"ows:Fees" yaml:"fees"`
	AccessConstraints  string `xml:"ows:AccessConstraints" yaml:"accessConstraints"`
}

// ServiceProvider struct containing the provider/organization information should only be fill by the "template" configuration wfs200.yaml
type ServiceProvider struct {
	XMLName      xml.Name `xml:"ows:ServiceProvider"`
	ProviderName string   `xml:"ows:ProviderName" yaml:"providerName"`
	ProviderSite struct {
		Type string `xml:"xlink:type,attr" yaml:"type"`
		Href string `xml:"xlink:href,attr" yaml:"href"`
	} `xml:"ows:ProviderSite" yaml:"providerSite"`
	ServiceContact struct {
		IndividualName string `xml:"ows:IndividualName" yaml:"individualName"`
		PositionName   string `xml:"ows:PositionName" yaml:"positionName"`
		ContactInfo    struct {
			Text  string `xml:",chardata"`
			Phone struct {
				Voice     string `xml:"ows:Voice" yaml:"voice"`
				Facsimile string `xml:"ows:Facsimile" yaml:"facsimile"`
			} `xml:"ows:Phone" yaml:"phone"`
			Address struct {
				DeliveryPoint         string `xml:"ows:DeliveryPoint" yaml:"deliveryPoint"`
				City                  string `xml:"ows:City" yaml:"city"`
				AdministrativeArea    string `xml:"ows:AdministrativeArea" yaml:"administrativeArea"`
				PostalCode            string `xml:"ows:PostalCode" yaml:"postalCode"`
				Country               string `xml:"ows:Country" yaml:"country"`
				ElectronicMailAddress string `xml:"ows:ElectronicMailAddress" yaml:"electronicMailAddress"`
			} `xml:"ows:Address" yaml:"address"`
			OnlineResource struct {
				Type string `xml:"xlink:type,attr" yaml:"type"`
				Href string `xml:"xlink:href,attr" yaml:"href"`
			} `xml:"ows:OnlineResource" yaml:"onlineResource"`
			HoursOfService      string `xml:"ows:HoursOfService" yaml:"hoursOfService"`
			ContactInstructions string `xml:"ows:ContactInstructions" yaml:"contactInstructions"`
		} `xml:"ows:ContactInfo" yaml:"contactInfo"`
		Role string `xml:"ows:Role" yaml:"role"`
	} `xml:"ows:ServiceContact" yaml:"serviceContact"`
}
