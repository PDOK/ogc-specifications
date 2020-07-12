package response

import (
	"encoding/xml"
)

// Contains the WMS130 struct

// Type and Version as constant
const (
	Service string = `WMS`
	Version string = `1.3.0`
)

// Service function needed for the interface
func (wms130 *Wms130) Service() string {
	return Service
}

// Version function needed for the interface
func (wms130 *Wms130) Version() string {
	return Version
}

// Validate function of the wms130 spec
func (wms130 *Wms130) Validate() bool {
	return false
}

// Wms130 base struct
type Wms130 struct {
	XMLName    xml.Name `xml:"WMS_Capabilities"`
	Namespaces `yaml:"namespaces"`
	WMSService WMSService `xml:"Service" yaml:"service"`
	Capability Capability `xml:"Capability" yaml:"capability"`
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
	OnlineResource     OnlineResource `xml:"OnlineResource" yaml:"onlineresource"`
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

// Capability base struct
type Capability struct {
	Request              Request               `xml:"Request" yaml:"request"`
	Exception            Exception             `xml:"Exception" yaml:"exception"`
	ExtendedCapabilities *ExtendedCapabilities `xml:"inspire_vs:ExtendedCapabilities" yaml:"extendedcapabilities"`
	Layer                []Layer               `xml:"Layer" yaml:"layer"`
}

// Request struct with the different operations, should be filled from the template
type Request struct {
	GetCapabilities RequestType  `xml:"GetCapabilities" yaml:"getcapabilities"`
	GetMap          RequestType  `xml:"GetMap" yaml:"getmap"`
	GetFeatureInfo  *RequestType `xml:"GetFeatureInfo" yaml:"getfeatureinfo"`
}

// Exception struct containing the different available exceptions, should be filled from the template
type Exception struct {
	Format []string `xml:"Format" yaml:"format"`
}

// Layer contains the WMS 1.3.0 layer configuration
type Layer struct {
	Queryable *string `xml:"queryable,attr" yaml:"queryable"`
	// layer has a full/complete map coverage
	Opaque *string `xml:"opaque,attr" yaml:"opaque"`
	// no cascaded attr in Layer element, because we don't do cascaded services e.g. wms services "proxying" and/or combining other wms services
	//Cascaded                *string                  `xml:"cascaded,attr" yaml:"cascaded"`
	Name                    *string                  `xml:"Name" yaml:"name"`
	Title                   string                   `xml:"Title" yaml:"title"`
	Abstract                string                   `xml:"Abstract" yaml:"abstract"`
	KeywordList             *KeywordList             `xml:"KeywordList" yaml:"keywordlist"`
	CRS                     []*string                `xml:"CRS" yaml:"crs"`
	EXGeographicBoundingBox *EXGeographicBoundingBox `xml:"EX_GeographicBoundingBox" yaml:"exgeographicboundingbox"`
	BoundingBox             []*BoundingBox           `xml:"BoundingBox" yaml:"boundingbox"`
	AuthorityURL            *AuthorityURL            `xml:"AuthorityURL" yaml:"authorityurl"`
	Identifier              *Identifier              `xml:"Identifier" yaml:"identifier"`
	MetadataURL             []*MetadataURL           `xml:"MetadataURL" yaml:"metadataurl"`
	Style                   []*Style                 `xml:"Style" yaml:"style"`
	Layer                   []*Layer                 `xml:"Layer" yaml:"layer"`
}

// RequestType containing the formats and DCPTypes available
type RequestType struct {
	Format  []string `xml:"Format" yaml:"format"`
	DCPType DCPType  `xml:"DCPType" yaml:"dcptype"`
}

// Identifier in struct for repeatablity
type Identifier struct {
	Authority string `xml:"authority,attr" yaml:"authority"`
	Value     string `xml:",chardata" yaml:"value"`
}

// KeywordList in struct for repeatablity
type KeywordList struct {
	Keyword []string `xml:"Keyword" yaml:"keyword"`
}

// OnlineResource in struct for repeatablity
type OnlineResource struct {
	Xlink *string `xml:"xmlns:xlink,attr" yaml:"xlink"`
	Type  *string `xml:"xlink:type,attr" yaml:"type"`
	Href  *string `xml:"xlink:href,attr" yaml:"href"`
}

// MetadataURL in struct for repeatablity
type MetadataURL struct {
	Type           *string        `xml:"type,attr" yaml:"type"`
	Format         *string        `xml:"Format" yaml:"format"`
	OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineresource"`
}

// AuthorityURL in struct for repeatablity
type AuthorityURL struct {
	Name           string         `xml:"name,attr" yaml:"name"`
	OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineresource"`
}

// ExtendedCapabilities containing the inspire extendedcapbilities, when available
type ExtendedCapabilities struct {
	MetadataURL struct {
		Type      string `xml:"xsi:type,attr" yaml:"type"`
		URL       string `xml:"inspire_common:URL" yaml:"url"`
		MediaType string `xml:"inspire_common:MediaType" yaml:"mediatype"`
	} `xml:"inspire_common:MetadataUrl" yaml:"metadataurl"`
	SupportedLanguages struct {
		DefaultLanguage struct {
			Language string `xml:"inspire_common:Language" yaml:"language"`
		} `xml:"inspire_common:DefaultLanguage" yaml:"defaultlanguage"`
	} `xml:"inspire_common:SupportedLanguages" yaml:"supportedlanguages"`
	ResponseLanguage struct {
		Language string `xml:"inspire_common:Language" yaml:"language"`
	} `xml:"inspire_common:ResponseLanguage" yaml:"responselanguage"`
}

// EXGeographicBoundingBox in struct for repeatablity
type EXGeographicBoundingBox struct {
	WestBoundLongitude string `xml:"westBoundLongitude" yaml:"westboundlongitude"`
	EastBoundLongitude string `xml:"eastBoundLongitude" yaml:"eastboundlongitude"`
	SouthBoundLatitude string `xml:"southBoundLatitude" yaml:"southboundlatitude"`
	NorthBoundLatitude string `xml:"northBoundLatitude" yaml:"northboundlatitude"`
}

// BoundingBox in struct for repeatablity
type BoundingBox struct {
	CRS  string `xml:"CRS,attr" yaml:"crs"`
	Minx string `xml:"minx,attr" yaml:"minx"`
	Miny string `xml:"miny,attr" yaml:"miny"`
	Maxx string `xml:"maxx,attr" yaml:"maxx"`
	Maxy string `xml:"maxy,attr" yaml:"maxy"`
}

// Style in struct for repeatablity
type Style struct {
	Name      string `xml:"Name" yaml:"name"`
	Title     string `xml:"Title" yaml:"title"`
	LegendURL struct {
		Width          string         `xml:"width,attr" yaml:"width"`
		Height         string         `xml:"height,attr" yaml:"height"`
		Format         string         `xml:"Format" yaml:"format"`
		OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineresource"`
	} `xml:"LegendURL" yaml:"legendurl"`
}

// DCPType in struct for repeatablity
type DCPType struct {
	HTTP struct {
		Get  *Method `xml:"Get" yaml:"get"`
		Post *Method `xml:"Post" yaml:"post"`
	} `xml:"HTTP" yaml:"http"`
}

// Method in struct for repeatablity
type Method struct {
	OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineresource"`
}
