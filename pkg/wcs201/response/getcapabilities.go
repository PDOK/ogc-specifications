package wcs201

import (
	"encoding/xml"
	"regexp"
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
func (gc *GetCapabilities) Validate() bool {
	return false
}

// BuildXML builds a GetCapabilities response object
func (gc *GetCapabilities) BuildXML() []byte {
	si, _ := xml.MarshalIndent(gc, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}

// GetCapabilities base struct
type GetCapabilities struct {
	XMLName               xml.Name `xml:"wcs:Capabilities"`
	Namespaces            `yaml:"namespaces"`
	ServiceIdentification ServiceIdentification `xml:"ows:ServiceIdentification" yaml:"serviceidentification"`
	ServiceProvider       ServiceProvider       `xml:"ows:ServiceProvider" yaml:"serviceprovider"`
	OperationsMetadata    OperationsMetadata    `xml:"ows:OperationsMetadata" yaml:"operationsmetadata"`
	ServiceMetadata       ServiceMetadata       `xml:"wcs:ServiceMetadata" yaml:"servicemetadata"`
	Contents              Contents              `xml:"wcs:Contents" yaml:"contents"`
}

// Namespaces struct containing the namespaces needed for the XML document
type Namespaces struct {
	XmlnsWCS           string `xml:"xmlns:wcs,attr" yaml:"wcs"`                                //http://www.opengis.net/wcs/2.0
	XmlnsOWS           string `xml:"xmlns:ows,attr" yaml:"ows"`                                //http://www.opengis.net/ows/1.1
	XmlnsOGC           string `xml:"xmlns:ogc,attr" yaml:"ogc"`                                //http://www.opengis.net/ogc
	XmlnsXSI           string `xml:"xmlns:xsi,attr" yaml:"xsi"`                                //http://www.w3.org/2001/XMLSchema-instance
	XmlnsXlink         string `xml:"xmlns:xlink,attr" yaml:"xlink"`                            //http://www.w3.org/1999/xlink
	XmlnsGML           string `xml:"xmlns:gml,attr" yaml:"gml"`                                //http://www.opengis.net/gml/3.2
	XmlnsGMLcov        string `xml:"xmlns:gmlcov,attr" yaml:"gmlcov"`                          //http://www.opengis.net/gmlcov/1.0
	XmlnsSWE           string `xml:"xmlns:swe,attr" yaml:"swe"`                                //http://www.opengis.net/swe/2.0
	XmlnsInspireCommon string `xml:"xmlns:inspire_common,attr,omitempty" yaml:"inspirecommon"` //http://inspire.ec.europa.eu/schemas/common/1.0
	XmlnsInspireDls    string `xml:"xmlns:inspire_dls,attr,omitempty" yaml:"inspiredls"`       //http://inspire.ec.europa.eu/schemas/inspire_dls/1.0
	XmlnsCrs           string `xml:"xmlns:crs,attr" yaml:"crs"`                                //http://www.opengis.net/wcs/crs/1.0
	XmlnsInt           string `xml:"xmlns:int,attr" yaml:"int"`                                //http://www.opengis.net/wcs/interpolation/1.0
	Version            string `xml:"version,attr" yaml:"version"`
	SchemaLocation     string `xml:"xsi:schemaLocation,attr" yaml:"schemalocation"`
}

// ServiceIdentification struct should only be fill by the "template" configuration wcs201.yaml
type ServiceIdentification struct {
	Title    string `xml:"ows:Title" yaml:"title"`
	Abstract string `xml:"ows:Abstract" yaml:"abstract"`
	Keywords struct {
		Keyword []string `xml:"ows:Keyword" yaml:"keyword"`
	} `xml:"ows:Keywords" yaml:"keywords"`
	ServiceType struct {
		Text      string `xml:",chardata" yaml:"text"`
		CodeSpace string `xml:"codeSpace,attr" yaml:"codespace"`
	} `xml:"ows:ServiceType" yaml:"servicetype"`
	ServiceTypeVersion []string `xml:"ows:ServiceTypeVersion" yaml:"servicetypeversion"`
	Profile            []string `xml:"ows:Profile" yaml:"profile"`
	Fees               string   `xml:"ows:Fees" yaml:"fees"`
	AccessConstraints  string   `xml:"ows:AccessConstraints" yaml:"accessconstraints"`
}

// ServiceProvider struct containing the provider/organization information should only be fill by the "template" configuration wcs201.yaml
type ServiceProvider struct {
	ProviderName string `xml:"ows:ProviderName" yaml:"providername"`
	ProviderSite struct {
		Type string `xml:"xlink:type,attr" yaml:"type"`
		Href string `xml:"xlink:href,attr" yaml:"href"`
	} `xml:"ows:ProviderSite" yaml:"providersite"`
	ServiceContact struct {
		IndividualName string `xml:"ows:IndividualName" yaml:"individualname"`
		PositionName   string `xml:"ows:PositionName" yaml:"positionname"`
		ContactInfo    struct {
			Phone struct {
				Voice     string `xml:"ows:Voice" yaml:"voice"`
				Facsimile string `xml:"ows:Facsimile" yaml:"facsimile"`
			} `xml:"ows:Phone" yaml:"phone"`
			Address struct {
				DeliveryPoint         string `xml:"ows:DeliveryPoint" yaml:"deliverypoint"`
				City                  string `xml:"ows:City" yaml:"city"`
				AdministrativeArea    string `xml:"ows:AdministrativeArea" yaml:"administrativearea"`
				PostalCode            string `xml:"ows:PostalCode" yaml:"postalcode"`
				Country               string `xml:"ows:Country" yaml:"country"`
				ElectronicMailAddress string `xml:"ows:ElectronicMailAddress" yaml:"electronicmailaddress"`
			} `xml:"ows:Address" yaml:"address"`
			OnlineResource struct {
				Type string `xml:"xlink:type,attr" yaml:"type"`
				Href string `xml:"xlink:href,attr" yaml:"href"`
			} `xml:"ows:OnlineResource" yaml:"onlineresource"`
			HoursOfService      string `xml:"ows:HoursOfService" yaml:"hoursofservice"`
			ContactInstructions string `xml:"ows:ContactInstructions" yaml:"contactinstructions"`
		} `xml:"ows:ContactInfo" yaml:"contactinfo"`
		Role string `xml:"ows:Role" yaml:"role"`
	} `xml:"ows:ServiceContact" yaml:"servicecontact"`
}

// OperationsMetadata struct for the WCS 2.0.1
type OperationsMetadata struct {
	Operation            []Operation           `xml:"ows:Operation" yaml:"operation"`
	ExtendedCapabilities *ExtendedCapabilities `xml:"ows:ExtendedCapabilities" yaml:"extendedcapabilities"`
}

// Operation in struct for repeatablity
type Operation struct {
	Name string `xml:"name,attr" yaml:"name"`
	DCP  struct {
		HTTP struct {
			Get struct {
				Type string `xml:"xlink:type,attr" yaml:"type"`
				Href string `xml:"xlink:href,attr" yaml:"href"`
			} `xml:"ows:Get" yaml:"get"`
			Post *Post `xml:"ows:Post" yaml:"post"`
		} `xml:"ows:HTTP"  yaml:"http"`
	} `xml:"ows:DCP" yaml:"dcp"`
}

// Post in separated struct so to use it as a Pointer
type Post struct {
	Type       string `xml:"xlink:type,attr" yaml:"type"`
	Href       string `xml:"xlink:href,attr" yaml:"href"`
	Constraint struct {
		Name          string `xml:"name,attr" yaml:"name"`
		AllowedValues struct {
			Value []string `xml:"ows:Value" yaml:"value"`
		} `xml:"ows:AllowedValues" yaml:"allowedvalues"`
	} `xml:"ows:Constraint" yaml:"constraint"`
}

// ExtendedCapabilities struct for the WCS 2.0.1
type ExtendedCapabilities struct {
	ExtendedCapabilities struct {
		MetadataURL struct {
			Type      string `xml:"xsi:type,attr"`
			URL       string `xml:"inspire_common:URL"`
			MediaType string `xml:"inspire_common:MediaType"`
		} `xml:"inspire_common:MetadataUrl"`
		SupportedLanguages struct {
			DefaultLanguage struct {
				Language string `xml:"inspire_common:Language"`
			} `xml:"inspire_common:DefaultLanguage"`
		} `xml:"inspire_common:SupportedLanguages"`
		ResponseLanguage struct {
			Language string `xml:"inspire_common:Language"`
		} `xml:"inspire_common:ResponseLanguage"`
		SpatialDataSetIdentifier struct {
			Code string `xml:"Code"`
		} `xml:"inspire_dls:SpatialDataSetIdentifier"`
	} `xml:"inspire_dls:ExtendedCapabilities"`
}

// ServiceMetadata struct for the WCS 2.0.1
type ServiceMetadata struct {
	FormatSupported []string `xml:"wcs:formatSupported"`
	Extension       struct {
		InterpolationMetadata struct {
			InterpolationSupported []string `xml:"int:InterpolationSupported"`
		} `xml:"int:InterpolationMetadata"`
		CrsMetadata struct {
			CrsSupported []string `xml:"crs:crsSupported"`
		} `xml:"crs:CrsMetadata"`
	} `xml:"wcs:Extension"`
}

// Contents in struct for repeatablity
type Contents struct {
	CoverageSummary []CoverageSummary `xml:"wcs:CoverageSummary"`
}

// CoverageSummary in struct for repeatablity
type CoverageSummary struct {
	CoverageID      string `xml:"wcs:CoverageId"`
	CoverageSubtype string `xml:"wcs:CoverageSubtype"`
}
