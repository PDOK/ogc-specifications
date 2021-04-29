package response

import (
	"encoding/xml"
	"regexp"

	"github.com/pdok/ogc-specifications/pkg/common"
	"github.com/pdok/ogc-specifications/pkg/wfs200/capabilities"
)

//
const (
	getcapabilities = `GetCapabilities`
)

//
const (
	Service = `WFS`
	Version = `2.0.0`
)

// Contains the WFS200 struct

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

// Validate function of the wfs200 spec
func (gc *GetCapabilities) Validate() common.Exceptions {
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
	XMLName               xml.Name `xml:"wfs:WFS_Capabilities"`
	Namespaces            `yaml:"namespaces"`
	ServiceIdentification ServiceIdentification `xml:"common:ServiceIdentification" yaml:"serviceidentification"`
	ServiceProvider       ServiceProvider       `xml:"common:ServiceProvider" yaml:"serviceprovider"`
	capabilities.Capabilities
}

// Namespaces struct containing the namespaces needed for the XML document
type Namespaces struct {
	XmlnsGML           string `xml:"xmlns:gml,attr" yaml:"gml"`                                          //http://www.opengis.net/gml/3.2
	XmlnsWFS           string `xml:"xmlns:wfs,attr" yaml:"wfs"`                                          //http://www.opengis.net/wfs/2.0
	XmlnsOWS           string `xml:"xmlns:common,attr" yaml:"common"`                                    //http://www.opengis.net/ows/1.1
	XmlnsXlink         string `xml:"xmlns:xlink,attr" yaml:"xlink"`                                      //http://www.w3.org/1999/xlink
	XmlnsXSI           string `xml:"xmlns:xsi,attr" yaml:"xsi"`                                          //http://www.w3.org/2001/XMLSchema-instance
	XmlnsFes           string `xml:"xmlns:fes,attr" yaml:"fes"`                                          //http://www.opengis.net/fes/2.0
	XmlnsInspireCommon string `xml:"xmlns:inspire_common,attr,omitempty" yaml:"inspirecommon,omitempty"` //http://inspire.ec.europa.eu/schemas/common/1.0
	XmlnsInspireDls    string `xml:"xmlns:inspire_dls,attr,omitempty" yaml:"inspiredls,omitempty"`       //http://inspire.ec.europa.eu/schemas/inspire_dls/1.0
	XmlnsPrefix        string `xml:"xmlns:{{.Prefix}},attr" yaml:"prefix"`                               //namespace_uri placeholder
	Version            string `xml:"version,attr" yaml:"version"`
	SchemaLocation     string `xml:"xsi:schemaLocation,attr" yaml:"schemalocation"`
}

// ServiceIdentification struct should only be fill by the "template" configuration wfs200.yaml
type ServiceIdentification struct {
	XMLName     xml.Name         `xml:"common:ServiceIdentification"`
	Title       string           `xml:"common:Title" yaml:"title"`
	Abstract    string           `xml:"common:Abstract" yaml:"abstract"`
	Keywords    *common.Keywords `xml:"common:Keywords" yaml:"keywords"`
	ServiceType struct {
		Text      string `xml:",chardata" yaml:"text"`
		CodeSpace string `xml:"codeSpace,attr" yaml:"codespace"`
	} `xml:"common:ServiceType"`
	ServiceTypeVersion string `xml:"common:ServiceTypeVersion" yaml:"servicetypeversion"`
	Fees               string `xml:"common:Fees" yaml:"fees"`
	AccessConstraints  string `xml:"common:AccessConstraints" yaml:"accesscontraints"`
}

// ServiceProvider struct containing the provider/organization information should only be fill by the "template" configuration wfs200.yaml
type ServiceProvider struct {
	XMLName      xml.Name `xml:"common:ServiceProvider"`
	ProviderName string   `xml:"common:ProviderName" yaml:"providername"`
	ProviderSite struct {
		Type string `xml:"xlink:type,attr" yaml:"type"`
		Href string `xml:"xlink:href,attr" yaml:"href"`
	} `xml:"common:ProviderSite" yaml:"providersite"`
	ServiceContact struct {
		IndividualName string `xml:"common:IndividualName" yaml:"individualname"`
		PositionName   string `xml:"common:PositionName" yaml:"positionname"`
		ContactInfo    struct {
			Text  string `xml:",chardata"`
			Phone struct {
				Voice     string `xml:"common:Voice" yaml:"voice"`
				Facsimile string `xml:"common:Facsimile" yaml:"facsmile"`
			} `xml:"common:Phone" yaml:"phone"`
			Address struct {
				DeliveryPoint         string `xml:"common:DeliveryPoint" yaml:"deliverypoint"`
				City                  string `xml:"common:City" yaml:"city"`
				AdministrativeArea    string `xml:"common:AdministrativeArea" yaml:"administrativearea"`
				PostalCode            string `xml:"common:PostalCode" yaml:"postalcode"`
				Country               string `xml:"common:Country" yaml:"country"`
				ElectronicMailAddress string `xml:"common:ElectronicMailAddress" yaml:"electronicmailaddress"`
			} `xml:"common:Address" yaml:"address"`
			OnlineResource struct {
				Type string `xml:"xlink:type,attr" yaml:"type"`
				Href string `xml:"xlink:href,attr" yaml:"href"`
			} `xml:"common:OnlineResource" yaml:"onlineresource"`
			HoursOfService      string `xml:"common:HoursOfService" yaml:"hoursofservice"`
			ContactInstructions string `xml:"common:ContactInstructions" yaml:"contactinstructions"`
		} `xml:"common:ContactInfo" yaml:"contactinfo"`
		Role string `xml:"common:Role" yaml:"role"`
	} `xml:"common:ServiceContact" yaml:"servicecontact"`
}
