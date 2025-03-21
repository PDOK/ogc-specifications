package wcs201

import (
	"encoding/xml"
	"regexp"

	"github.com/pdok/ogc-specifications/pkg/wsc200"
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

// ToXML builds a GetCapabilities response object
func (gc GetCapabilitiesResponse) ToXML() []byte {
	si, _ := xml.MarshalIndent(gc, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}

// GetCapabilitiesResponse base struct
type GetCapabilitiesResponse struct {
	XMLName xml.Name `xml:"wcs:Capabilities" yaml:"wcsCapabilities"`
	Namespaces
	ServiceIdentification ServiceIdentification `xml:"ows:ServiceIdentification" yaml:"serviceIdentification"`
	ServiceProvider       ServiceProvider       `xml:"ows:ServiceProvider" yaml:"serviceProvider"`
	OperationsMetadata    OperationsMetadata    `xml:"ows:OperationsMetadata" yaml:"operationsMetadata"`
	ServiceMetadata       ServiceMetadata       `xml:"wcs:ServiceMetadata" yaml:"serviceMetadata"`
	Contents              Contents              `xml:"wcs:Contents" yaml:"contents"`
}

// Namespaces struct containing the namespaces needed for the XML document
type Namespaces struct {
	XmlnsWCS           string `xml:"xmlns:wcs,attr" yaml:"wcs"`                                //http://www.opengis.net/wcs/2.0
	XmlnsOWS           string `xml:"xmlns:ows,attr" yaml:"common"`                             //http://www.opengis.net/ows/1.1
	XmlnsOGC           string `xml:"xmlns:ogc,attr" yaml:"ogc"`                                //http://www.opengis.net/ogc
	XmlnsXSI           string `xml:"xmlns:xsi,attr" yaml:"xsi"`                                //http://www.w3.org/2001/XMLSchema-instance
	XmlnsXlink         string `xml:"xmlns:xlink,attr" yaml:"xlink"`                            //http://www.w3.org/1999/xlink
	XmlnsGML           string `xml:"xmlns:gml,attr" yaml:"gml"`                                //http://www.opengis.net/gml/3.2
	XmlnsGMLcov        string `xml:"xmlns:gmlcov,attr" yaml:"gmlcov"`                          //http://www.opengis.net/gmlcov/1.0
	XmlnsSWE           string `xml:"xmlns:swe,attr" yaml:"swe"`                                //http://www.opengis.net/swe/2.0
	XmlnsInspireCommon string `xml:"xmlns:inspire_common,attr,omitempty" yaml:"inspireCommon"` //http://inspire.ec.europa.eu/schemas/common/1.0
	XmlnsInspireDls    string `xml:"xmlns:inspire_dls,attr,omitempty" yaml:"inspireDls"`       //http://inspire.ec.europa.eu/schemas/inspire_dls/1.0
	XmlnsCrs           string `xml:"xmlns:crs,attr" yaml:"crs"`                                //http://www.opengis.net/wcs/crs/1.0
	XmlnsInt           string `xml:"xmlns:int,attr" yaml:"int"`                                //http://www.opengis.net/wcs/interpolation/1.0
	Version            string `xml:"version,attr" yaml:"version"`
	SchemaLocation     string `xml:"xsi:schemaLocation,attr" yaml:"schemaLocation"`
}

// ServiceIdentification struct should only be fill by the "template" configuration wcs201.yaml
type ServiceIdentification struct {
	Title       string           `xml:"ows:Title" yaml:"title"`
	Abstract    string           `xml:"ows:Abstract" yaml:"abstract"`
	Keywords    *wsc200.Keywords `xml:"ows:Keywords" yaml:"keywords"`
	ServiceType struct {
		Text      string `xml:",chardata" yaml:"text"`
		CodeSpace string `xml:"codeSpace,attr" yaml:"codespace"`
	} `xml:"ows:ServiceType" yaml:"servicetype"`
	ServiceTypeVersion []string `xml:"ows:ServiceTypeVersion" yaml:"serviceTypeVersion"`
	Profile            []string `xml:"ows:Profile" yaml:"profile"`
	Fees               string   `xml:"ows:Fees" yaml:"fees"`
	AccessConstraints  string   `xml:"ows:AccessConstraints" yaml:"accessConstraints"`
}

// ServiceProvider struct containing the provider/organization information should only be fill by the "template" configuration wcs201.yaml
type ServiceProvider struct {
	ProviderName string `xml:"ows:ProviderName" yaml:"providerName"`
	ProviderSite struct {
		Type string `xml:"xlink:type,attr" yaml:"type"`
		Href string `xml:"xlink:href,attr" yaml:"href"`
	} `xml:"ows:ProviderSite" yaml:"providersite"`
	ServiceContact struct {
		IndividualName string `xml:"ows:IndividualName" yaml:"individualName"`
		PositionName   string `xml:"ows:PositionName" yaml:"positionName"`
		ContactInfo    struct {
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
			OnlineResource *struct {
				Type string `xml:"xlink:type,attr,omitempty" yaml:"type"`
				Href string `xml:"xlink:href,attr,omitempty" yaml:"href"`
			} `xml:"ows:OnlineResource,omitempty" yaml:"onlineResource"`
			HoursOfService      string `xml:"ows:HoursOfService,omitempty" yaml:"hoursOfService"`
			ContactInstructions string `xml:"ows:ContactInstructions,omitempty" yaml:"contactInstructions"`
		} `xml:"ows:ContactInfo" yaml:"contactInfo"`
		Role string `xml:"ows:Role,omitempty" yaml:"role"`
	} `xml:"ows:ServiceContact" yaml:"serviceContact"`
}
