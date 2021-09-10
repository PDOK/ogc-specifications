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
	WMSCapabilities     `xml:"WMSCapabilities" yaml:"wmscapabilities"`
	OptionalConstraints `xml:"OptionalConstraints" yaml:"optionalconstraints"`
}

// WMSCapabilities base struct
type WMSCapabilities struct {
	Request              Request               `xml:"Request" yaml:"request"`
	Exception            ExceptionType         `xml:"Exception" yaml:"exception"`
	ExtendedCapabilities *ExtendedCapabilities `xml:"inspire_vs:ExtendedCapabilities" yaml:"extendedcapabilities"`
	Layer                []Layer               `xml:"Layer" yaml:"layer"`
}

// OptionalConstraints struct
type OptionalConstraints struct {
	LayerLimit int `xml:"LayerLimit,omitempty" yaml:"layerlimit,omitempty"`
	MaxWidth   int `xml:"MaxWidth,omitempty" yaml:"maxwidth,omitempty"`
	MaxHeight  int `xml:"MaxHeight,omitempty" yaml:"maxheight,omitempty"`
}

// Request struct with the different operations, should be filled from the template
type Request struct {
	GetCapabilities RequestType  `xml:"GetCapabilities" yaml:"getcapabilities"`
	GetMap          RequestType  `xml:"GetMap" yaml:"getmap"`
	GetFeatureInfo  *RequestType `xml:"GetFeatureInfo" yaml:"getfeatureinfo"`
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
	Opaque *string `xml:"opaque,attr" yaml:"opaque"`
	// no cascaded attr in Layer element, because we don't do cascaded services e.g. wms services "proxying" and/or combining other wms services
	//Cascaded                *string                  `xml:"cascaded,attr" yaml:"cascaded"`
	Name                    *string                  `xml:"Name" yaml:"name"`
	Title                   string                   `xml:"Title" yaml:"title"`
	Abstract                string                   `xml:"Abstract,omitempty" yaml:"abstract,omitempty"`
	KeywordList             *Keywords                `xml:"KeywordList" yaml:"keywordlist"`
	CRS                     []CRS                    `xml:"CRS" yaml:"crs"`
	EXGeographicBoundingBox *EXGeographicBoundingBox `xml:"EX_GeographicBoundingBox" yaml:"exgeographicboundingbox"`
	Dimension               []*Dimension             `xml:"Dimension" yaml:"dimension"`
	BoundingBox             []*LayerBoundingBox      `xml:"BoundingBox" yaml:"boundingbox"`
	AuthorityURL            *AuthorityURL            `xml:"AuthorityURL" yaml:"authorityurl"`
	Attribution             *Attribution             `xml:"Attribution,omitempty" yaml:"attribution"`
	Identifier              *Identifier              `xml:"Identifier" yaml:"identifier"`
	FeatureListURL          *FeatureListURL          `xml:"FeatureListURL,omitempty" yaml:"featurelisturl"`
	MetadataURL             []*MetadataURL           `xml:"MetadataURL" yaml:"metadataurl"`
	Style                   []*Style                 `xml:"Style" yaml:"style"`
	Layer                   []*Layer                 `xml:"Layer" yaml:"layer"`
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
	DCPType *DCPType `xml:"DCPType" yaml:"dcptype"`
}

// Identifier in struct for repeatability
type Identifier struct {
	Authority string `xml:"authority,attr" yaml:"authority"`
	Value     string `xml:",chardata" yaml:"value"`
}

// Attribution in struct for repeatability
type Attribution struct {
	Title          string         `xml:"Title" yaml:"title"`
	OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineresource"`
	LogoURL        struct {
		Format         *string        `xml:"Format" yaml:"format"`
		OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineresource"`
	} `xml:"LogoURL" yaml:"logourl"`
}

// Identifier in struct for repeatability
type FeatureListURL struct {
	Format         string         `xml:"Format" yaml:"format"`
	OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineresource"`
}

// MetadataURL in struct for repeatability
type MetadataURL struct {
	Type           *string        `xml:"type,attr" yaml:"type"`
	Format         *string        `xml:"Format" yaml:"format"`
	OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineresource"`
}

// AuthorityURL in struct for repeatability
type AuthorityURL struct {
	Name           string         `xml:"name,attr" yaml:"name"`
	OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineresource"`
}

// ExtendedCapabilities containing the inspire extendedcapabilities, when available
type ExtendedCapabilities struct {
	MetadataURL struct {
		URL       string `xml:"inspire_common:URL" yaml:"url"`
		MediaType string `xml:"inspire_common:MediaType" yaml:"mediatype"`
	} `xml:"inspire_common:MetadataUrl" yaml:"metadataurl"`
	SupportedLanguages struct {
		DefaultLanguage struct {
			Language string `xml:"inspire_common:Language" yaml:"language"`
		} `xml:"inspire_common:DefaultLanguage" yaml:"defaultlanguage"`
		SupportedLanguage *[]struct {
			Language string `xml:"inspire_common:Language" yaml:"language"`
		} `xml:"inspire_common:SupportedLanguage" yaml:"supportedlanguage"`
	} `xml:"inspire_common:SupportedLanguages" yaml:"supportedlanguages"`
	ResponseLanguage struct {
		Language string `xml:"inspire_common:Language" yaml:"language"`
	} `xml:"inspire_common:ResponseLanguage" yaml:"responselanguage"`
}

// EXGeographicBoundingBox in struct for repeatability
type EXGeographicBoundingBox struct {
	WestBoundLongitude float64 `xml:"westBoundLongitude" yaml:"westboundlongitude"`
	EastBoundLongitude float64 `xml:"eastBoundLongitude" yaml:"eastboundlongitude"`
	SouthBoundLatitude float64 `xml:"southBoundLatitude" yaml:"southboundlatitude"`
	NorthBoundLatitude float64 `xml:"northBoundLatitude" yaml:"northboundlatitude"`
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
	Name      string `xml:"Name" yaml:"name"`
	Title     string `xml:"Title" yaml:"title"`
	Abstract  string `xml:"Abstract,omitempty" yaml:"abstract"`
	LegendURL struct {
		Width          int            `xml:"width,attr" yaml:"width"`
		Height         int            `xml:"height,attr" yaml:"height"`
		Format         string         `xml:"Format" yaml:"format"`
		OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineresource"`
	} `xml:"LegendURL" yaml:"legendurl"`
	StyleSheetURL *struct {
		Format         string         `xml:"Format" yaml:"format"`
		OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineresource"`
	} `xml:"StyleSheetURL,omitempty" yaml:"stylesheeturl"`
}

// DCPType in struct for repeatability
type DCPType struct {
	HTTP struct {
		Get  *Method `xml:"Get" yaml:"get"`
		Post *Method `xml:"Post" yaml:"post"`
	} `xml:"HTTP" yaml:"http"`
}

// Method in struct for repeatability
type Method struct {
	OnlineResource OnlineResource `xml:"OnlineResource" yaml:"onlineresource"`
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
	Default      *string `xml:"default,attr,omitempty" yaml:"default"`
	NearestValue *string `xml:"nearestValue,attr,omitempty" yaml:"nearestvalue"`
	Value        *string `xml:",chardata" yaml:"value"`
}
