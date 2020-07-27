package capabilities

import "github.com/pdok/ogc-specifications/pkg/ows"

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
	KeywordList             *ows.Keywords            `xml:"KeywordList" yaml:"keywordlist"`
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
	WestBoundLongitude float64 `xml:"westBoundLongitude" yaml:"westboundlongitude"`
	EastBoundLongitude float64 `xml:"eastBoundLongitude" yaml:"eastboundlongitude"`
	SouthBoundLatitude float64 `xml:"southBoundLatitude" yaml:"southboundlatitude"`
	NorthBoundLatitude float64 `xml:"northBoundLatitude" yaml:"northboundlatitude"`
}

// BoundingBox in struct for repeatablity
type BoundingBox struct {
	CRS  string  `xml:"CRS,attr" yaml:"crs"`
	Minx float64 `xml:"minx,attr" yaml:"minx"`
	Miny float64 `xml:"miny,attr" yaml:"miny"`
	Maxx float64 `xml:"maxx,attr" yaml:"maxx"`
	Maxy float64 `xml:"maxy,attr" yaml:"maxy"`
}

// Style in struct for repeatablity
type Style struct {
	Name      string `xml:"Name" yaml:"name"`
	Title     string `xml:"Title" yaml:"title"`
	LegendURL struct {
		Width          int            `xml:"width,attr" yaml:"width"`
		Height         int            `xml:"height,attr" yaml:"height"`
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

// OnlineResource in struct for repeatablity
type OnlineResource struct {
	Xlink *string `xml:"xmlns:xlink,attr" yaml:"xlink"`
	Type  *string `xml:"xlink:type,attr" yaml:"type"`
	Href  *string `xml:"xlink:href,attr" yaml:"href"`
}