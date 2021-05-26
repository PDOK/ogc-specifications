package wmts100

import (
	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

// ParseXML func
func (c *Contents) ParseXML(doc []byte) error {
	return nil
}

// ParseYAML func
func (c *Contents) ParseYAML(doc []byte) error {
	return nil
}

// Contents struct for the WMTS 1.0.0
type Contents struct {
	Layer         []Layer         `xml:"Layer" yaml:"layer"`
	TileMatrixSet []TileMatrixSet `xml:"TileMatrixSet" yaml:"tilematrixset"`
}

// GetTilematrixsets helper function for collecting the provided TileMatrixSets, so th base can be cleanup for unused TileMatrixSets
func (c Contents) GetTilematrixsets() map[string]bool {
	tilematrixsets := make(map[string]bool)
	for _, l := range c.Layer {
		for _, t := range l.TileMatrixSetLink {
			tilematrixsets[t.TileMatrixSet] = true
		}
	}
	return tilematrixsets
}

// Layer in struct for repeatability
type Layer struct {
	Title             string                  `xml:"ows:Title" yaml:"title"`
	Abstract          string                  `xml:"ows:Abstract" yaml:"abstract"`
	WGS84BoundingBox  wsc110.WGS84BoundingBox `xml:"ows:WGS84BoundingBox" yaml:"wgs84boundingbox"`
	Identifier        string                  `xml:"ows:Identifier" yaml:"identifier"`
	Metadata          *Metadata               `xml:"ows:Metadata,omitempty" yaml:"metadata"`
	Style             []Style                 `xml:"Style" yaml:"style"`
	Format            string                  `xml:"Format" yaml:"format"`
	TileMatrixSetLink []TileMatrixSetLink     `xml:"TileMatrixSetLink" yaml:"tilematrixsetlink"`
	ResourceURL       struct {
		Format       string `xml:"format,attr" yaml:"format"`
		ResourceType string `xml:"resourceType,attr" yaml:"resourcetype"`
		Template     string `xml:"template,attr" yaml:"template"`
	} `xml:"ResourceURL" yaml:"resourceurl"`
}

// Metadata  in struct for repeatability
type Metadata struct {
	Href string `xml:"xlink:href,attr,omitempty" yaml:"href"`
}

// Style in struct for repeatability
type Style struct {
	Identifier string           `xml:"ows:Identifier" yaml:"identifier"`
	Title      *string          `xml:"ows:Title,omitempty" yaml:"title"`
	Abstract   *string          `xml:"ows:Abstract,omitempty" yaml:"abstract"`
	Keywords   *wsc110.Keywords `xml:"Keywords,omitempty" yaml:"keywords"`
	LegendURL  []*LegendURL     `xml:"LegendURL,omitempty" yaml:"legendurl"`
	IsDefault  *bool            `xml:"isDefault,attr,omitempty" yaml:"isdefault"`
}

// TileMatrixSetLink in struct for repeatability
type TileMatrixSetLink struct {
	TileMatrixSet string `xml:"TileMatrixSet" yaml:"tilematrixset"`
}

// TileMatrixSet in struct for repeatability
type TileMatrixSet struct {
	Identifier   string       `xml:"ows:Identifier" yaml:"identifier"`
	SupportedCRS string       `xml:"ows:SupportedCRS" yaml:"supportedcrs"`
	TileMatrix   []TileMatrix `xml:"TileMatrix" yaml:"tilematrix"`
}

// TileMatrix in struct for repeatability
type TileMatrix struct {
	Identifier       string `xml:"ows:Identifier" yaml:"identifier"`
	ScaleDenominator string `xml:"ScaleDenominator" yaml:"scaledenominator"`
	TopLeftCorner    string `xml:"TopLeftCorner" yaml:"topleftcorner"`
	TileWidth        string `xml:"TileWidth" yaml:"tilewidth"`
	TileHeight       string `xml:"TileHeight" yaml:"tileheight"`
	MatrixWidth      string `xml:"MatrixWidth" yaml:"matrixwidth"`
	MatrixHeight     string `xml:"MatrixHeight" yaml:"matrixheight"`
}

// LegendURL in struct for optionality
type LegendURL struct {
	Format string `xml:"format,attr,omitempty" yaml:"format,omitempty"`
	Href   string `xml:"xlink:href,attr,omitempty" yaml:"href,omitempty"`
}
