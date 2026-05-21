package wms130

import (
	"encoding/xml"
	"log"

	"gopkg.in/yaml.v3"
)

// ParseXML func
func (c *Capabilities) ParseXML(doc []byte) error {
	if err := xml.Unmarshal(doc, c); err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	return nil
}

// ParseYAML func
func (c *Capabilities) ParseYAML(doc []byte) error {
	if err := yaml.Unmarshal(doc, c); err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	return nil
}

// Capabilities struct needed for keeping all constraints and capabilities together
type Capabilities struct {
	WMSCapabilities      `xml:"WMSCapabilities" yaml:"wmsCapabilities"`
	*OptionalConstraints `xml:"OptionalConstraints" yaml:"optionalConstraints,omitempty"`
}

// WMSCapabilities base struct
type WMSCapabilities struct {
	Request              Request               `xml:"Request" yaml:"request"`
	Exception            ExceptionType         `xml:"Exception" yaml:"exception"`
	ExtendedCapabilities *ExtendedCapabilities `xml:"inspire_vs:ExtendedCapabilities" yaml:"extendedCapabilities,omitempty"`
	Layer                []Layer               `xml:"Layer" yaml:"layer"`
}

// OptionalConstraints struct
type OptionalConstraints struct {
	LayerLimit int `xml:"LayerLimit,omitempty" yaml:"layerLimit,omitempty"`
	MaxWidth   int `xml:"MaxWidth,omitempty" yaml:"maxWidth,omitempty"`
	MaxHeight  int `xml:"MaxHeight,omitempty" yaml:"maxHeight,omitempty"`
}

// Request struct with the different operations, should be filled from the template
type Request struct {
	GetCapabilities RequestType  `xml:"GetCapabilities" yaml:"getCapabilities"`
	GetMap          RequestType  `xml:"GetMap" yaml:"getMap"`
	GetFeatureInfo  *RequestType `xml:"GetFeatureInfo" yaml:"getFeatureInfo"`
}

// ExceptionType struct containing the different available exceptions, should be filled from the template
// default: XML
// other commonly used: BLANK, INIMAGE and JSON
type ExceptionType struct {
	Format []string `xml:"Format" yaml:"format"`
}

// Layer contains the WMS 1.3.0 layer configuration
type Layer struct {
	Queryable *int `xml:"queryable,attr" yaml:"queryable"`
	// layer has a full/complete map coverage
	Opaque *string `xml:"opaque,attr" yaml:"opaque,omitempty"`
	// no cascaded attr in Layer element, because we don't do cascaded services e.g. wms services "proxying" and/or combining other wms services
	//Cascaded                *string                  `xml:"cascaded,attr" yaml:"cascaded"`
	Name                    *string                  `xml:"Name" yaml:"name,omitempty"`
	Title                   string                   `xml:"Title" yaml:"title"`
	Abstract                *string                  `xml:"Abstract,omitempty" yaml:"abstract,omitempty"`
	KeywordList             *Keywords                `xml:"KeywordList" yaml:"keywordList,omitempty"`
	CRS                     []CRS                    `xml:"CRS" yaml:"crs,omitempty"`
	EXGeographicBoundingBox *EXGeographicBoundingBox `xml:"EX_GeographicBoundingBox" yaml:"exGeographicBoundingBox,omitempty"`
	BoundingBox             []*LayerBoundingBox      `xml:"BoundingBox" yaml:"boundingBox,omitempty"`
	Dimension               []*Dimension             `xml:"Dimension" yaml:"dimension,omitempty"`
	Attribution             *Attribution             `xml:"Attribution,omitempty" yaml:"attribution,omitempty"`
	AuthorityURL            *AuthorityURL            `xml:"AuthorityURL" yaml:"authorityUrl,omitempty"`
	Identifier              *Identifier              `xml:"Identifier" yaml:"identifier,omitempty"`
	MetadataURL             []*MetadataURL           `xml:"MetadataURL" yaml:"metadataUrl,omitempty"`
	DataURL                 *URL                     `xml:"DataURL,omitempty" yaml:"dataUrl,omitempty"`
	FeatureListURL          *URL                     `xml:"FeatureListURL,omitempty" yaml:"featureListUrl,omitempty"`
	Style                   []*Style                 `xml:"Style" yaml:"style,omitempty"`
	MinScaleDenominator     *float64                 `xml:"MinScaleDenominator,omitempty" yaml:"minScaleDenominator,omitempty"`
	MaxScaleDenominator     *float64                 `xml:"MaxScaleDenominator,omitempty" yaml:"maxScaleDenominator,omitempty"`
	Layer                   []*Layer                 `xml:"Layer" yaml:"layer,omitempty"`
}

// StyleDefined checks if the style that is defined is available for the requested layer
func (c *Capabilities) StyleDefined(layername, stylename string) bool {
	defined := false
	for _, layer := range c.Layer {
		defined = layer.styleDefined(layername, stylename)
		if defined {
			return defined
		}
	}

	return defined
}

// styleDefined checks if the style is defined
func (l *Layer) styleDefined(layername, stylename string) bool {
	if l.Name != nil && *l.Name == layername {
		if l.Style != nil {
			for _, sld := range l.Style {
				if sld.Name == stylename {
					return true
				}
			}
		}
	}

	// Not found top level, so we go deeper
	defined := false
	if l.Layer != nil {
		for _, nestedlayer := range l.Layer {
			defined = nestedlayer.styleDefined(layername, stylename)
			if defined {
				return defined
			}
		}
	}

	return defined
}

// GetLayerNames returns the available layers as []string
func (c *Capabilities) GetLayerNames() []string {
	var layers []string

	for _, l := range c.Layer {
		if l.Name != nil {
			layers = append(layers, *l.Name)
		}
		if l.Layer != nil {
			for _, n := range l.Layer {
				u := n.getLayerNames()
				layers = append(layers, u...)
			}
		}
	}

	return layers
}

// getLayerNames returns the available layers as []string
func (l *Layer) getLayerNames() []string {
	var layers []string

	layers = append(layers, *l.Name)
	if l.Layer != nil {
		for _, n := range l.Layer {
			u := n.getLayerNames()
			layers = append(layers, u...)
		}
	}

	return layers
}

func (l *Layer) findLayer(layername string) *Layer {
	if *l.Name == layername {
		return l
	}
	if l.Layer != nil {
		for _, n := range l.Layer {
			u := n.findLayer(layername)
			if u != nil {
				if *u.Name == layername {
					return u
				}
			}
		}
	}
	return nil
}

// GetLayer returns the Layer Capabilities from the Capabilities document.
// when the requested Layer is not found a exception is thrown.
func (c *Capabilities) GetLayer(layername string) (Layer, Exceptions) {
	var layer Layer

	found := false
	for _, l := range c.GetLayerNames() {
		if l == layername {
			found = true
		}
	}

	if !found {
		return layer, Exceptions{LayerNotDefined(layername)}
	}

	for _, l := range c.Layer {
		if l.Name != nil && *l.Name == layername {
			layer = l
			break
		}
		if l.Layer != nil {
			for _, n := range l.Layer {
				u := n.findLayer(layername)
				if u != nil {
					return *u, nil
				}
			}
		}
	}

	return layer, nil
}

// RequestType containing the formats and DCPTypes available
type RequestType struct {
	Format  []string `xml:"Format" yaml:"format"`
	DCPType *DCPType `xml:"DCPType" yaml:"dcpType"`
}

// Identifier in struct for repeatability
type Identifier struct {
	Authority string `xml:"authority,attr" yaml:"authority"`
	Value     string `xml:",chardata" yaml:"value"`
}

// Attribution in struct for repeatability
type Attribution struct {
	Title          *string         `xml:"Title" yaml:"title"`
	OnlineResource *OnlineResource `xml:"OnlineResource" yaml:"onlineResource"`
	LogoURL        *LogoURL        `xml:"LogoURL" yaml:"logoUrl"`
}

// LogoURL in struct for repeatability
type LogoURL struct {
	Format         *string        `xml:"Format" yaml:"format"`
	OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineResource"`
}

// URL in struct for repeatability
type URL struct {
	Format         *string        `xml:"Format" yaml:"format"`
	OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineResource"`
}

// MetadataURL in struct for repeatability
type MetadataURL struct {
	Type           *string        `xml:"type,attr" yaml:"type"`
	Format         *string        `xml:"Format" yaml:"format"`
	OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineResource"`
}

// AuthorityURL in struct for repeatability
type AuthorityURL struct {
	Name           string         `xml:"name,attr" yaml:"name"`
	OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineResource"`
}

// ExtendedCapabilities containing the inspire extendedcapabilities, when available
type ExtendedCapabilities struct {
	MetadataURL        ExtendedMetadataURL `xml:"inspire_common:MetadataUrl" yaml:"metadataUrl"`
	SupportedLanguages SupportedLanguages  `xml:"inspire_common:SupportedLanguages" yaml:"supportedLanguages"`
	ResponseLanguage   Language            `xml:"inspire_common:ResponseLanguage" yaml:"responseLanguage"`
}

// ExtendedMetadataURL struct for the WMS 1.3.0
type ExtendedMetadataURL struct {
	URL       string `xml:"inspire_common:URL" yaml:"url"`
	MediaType string `xml:"inspire_common:MediaType" yaml:"mediaType"`
}

// SupportedLanguages struct for the WMS 1.3.0
type SupportedLanguages struct {
	DefaultLanguage   Language    `xml:"inspire_common:DefaultLanguage" yaml:"defaultLanguage"`
	SupportedLanguage *[]Language `xml:"inspire_common:SupportedLanguage" yaml:"supportedLanguage"`
}

// Language struct for the WMS 1.3.0
type Language struct {
	Language string `xml:"inspire_common:Language" yaml:"language"`
}

// EXGeographicBoundingBox in struct for repeatability
type EXGeographicBoundingBox struct {
	WestBoundLongitude float64 `xml:"westBoundLongitude" yaml:"westBoundLongitude"`
	EastBoundLongitude float64 `xml:"eastBoundLongitude" yaml:"eastBoundLongitude"`
	SouthBoundLatitude float64 `xml:"southBoundLatitude" yaml:"southBoundLatitude"`
	NorthBoundLatitude float64 `xml:"northBoundLatitude" yaml:"northBoundLatitude"`
}

// LayerBoundingBox in struct for repeatability
type LayerBoundingBox struct {
	CRS  string  `xml:"CRS,attr" yaml:"crs"`
	Minx float64 `xml:"minx,attr" yaml:"minx"`
	Miny float64 `xml:"miny,attr" yaml:"miny"`
	Maxx float64 `xml:"maxx,attr" yaml:"maxx"`
	Maxy float64 `xml:"maxy,attr" yaml:"maxy"`
	Resx float64 `xml:"resx,attr,omitempty" yaml:"resx,omitempty"`
	Resy float64 `xml:"resy,attr,omitempty" yaml:"resy,omitempty"`
}

// Style in struct for repeatability
type Style struct {
	Name          string         `xml:"Name" yaml:"name"`
	Title         string         `xml:"Title" yaml:"title"`
	Abstract      *string        `xml:"Abstract,omitempty" yaml:"abstract,omitempty"`
	LegendURL     *LegendURL     `xml:"LegendURL" yaml:"legendUrl"`
	StyleSheetURL *StyleSheetURL `xml:"StyleSheetURL,omitempty" yaml:"styleSheetUrl,omitempty"`
}

// LegendURL struct for the WMS 1.3.0
type LegendURL struct {
	Width          int            `xml:"width,attr" yaml:"width"`
	Height         int            `xml:"height,attr" yaml:"height"`
	Format         string         `xml:"Format" yaml:"format"`
	OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineResource"`
}

// StyleSheetURL struct for the WMS 1.3.0
type StyleSheetURL struct {
	Format         string         `xml:"Format" yaml:"format"`
	OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineResource"`
}

// DCPType in struct for repeatability
type DCPType struct {
	HTTP struct {
		Get  Method  `xml:"Get" yaml:"get"`
		Post *Method `xml:"Post" yaml:"post"`
	} `xml:"HTTP" yaml:"http"`
}

// Method in struct for repeatability
type Method struct {
	OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineResource"`
}

// OnlineResource in struct for repeatability
type OnlineResource struct {
	Xlink *string `xml:"xmlns:xlink,attr" yaml:"xlink"`
	Type  *string `xml:"xlink:type,attr" yaml:"type"`
	Href  *string `xml:"xlink:href,attr" yaml:"href"`
}

type Dimension struct {
	Name         *string `xml:"name,attr" yaml:"name"`
	Units        *string `xml:"units,attr" yaml:"units"`
	Default      *string `xml:"default,attr,omitempty" yaml:"default,omitempty"`
	NearestValue *string `xml:"nearestValue,attr,omitempty" yaml:"nearestValue,omitempty"`
	Value        *string `xml:",chardata" yaml:"value"`
}
