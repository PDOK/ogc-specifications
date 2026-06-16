package wmts100

import (
	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

// ParseXML func
func (c *Contents) ParseXML(_ []byte) error {
	return nil
}

// ParseYAML func
func (c *Contents) ParseYAML(_ []byte) error {
	return nil
}

// Contents struct for the WMTS 1.0.0
type Contents struct {
	Layer         []Layer         `xml:"Layer" yaml:"layer"`
	TileMatrixSet []TileMatrixSet `xml:"TileMatrixSet" yaml:"tileMatrixSet"`
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
	WGS84BoundingBox  wsc110.WGS84BoundingBox `xml:"ows:WGS84BoundingBox" yaml:"wgs84BoundingBox"`
	Identifier        string                  `xml:"ows:Identifier" yaml:"identifier"`
	Metadata          *Metadata               `xml:"ows:Metadata,omitempty" yaml:"metadata"`
	Style             []Style                 `xml:"Style" yaml:"style"`
	Format            []string                `xml:"Format" yaml:"format"`
	InfoFormat        []string                `xml:"InfoFormat" yaml:"infoFormat"`
	TileMatrixSetLink []TileMatrixSetLink     `xml:"TileMatrixSetLink" yaml:"tileMatrixSetLink"`
	ResourceURL       []ResourceURL           `xml:"ResourceURL" yaml:"resourceUrl"`
}

type ResourceURL struct {
	Format       string `xml:"format,attr" yaml:"format"`
	ResourceType string `xml:"resourceType,attr" yaml:"resourceType"`
	Template     string `xml:"template,attr" yaml:"template"`
}

// Metadata  in struct for repeatability
type Metadata struct {
	Href string `xml:"xlink:href,attr,omitempty" yaml:"href"`
}

// Style in struct for repeatability
type Style struct {
	Title      *string          `xml:"ows:Title,omitempty" yaml:"title"`
	Abstract   *string          `xml:"ows:Abstract,omitempty" yaml:"abstract"`
	Keywords   *wsc110.Keywords `xml:"Keywords,omitempty" yaml:"keywords"`
	Identifier string           `xml:"ows:Identifier" yaml:"identifier"`
	LegendURL  []*LegendURL     `xml:"LegendURL,omitempty" yaml:"legendUrl"`
	IsDefault  *bool            `xml:"isDefault,attr,omitempty" yaml:"isDefault"`
}

// TileMatrixSetLink in struct for repeatability
type TileMatrixSetLink struct {
	TileMatrixSet string `xml:"TileMatrixSet" yaml:"tileMatrixSet"`
}

// TileMatrixSet in struct for repeatability
type TileMatrixSet struct {
	Identifier   string       `xml:"ows:Identifier" yaml:"identifier"`
	SupportedCRS string       `xml:"ows:SupportedCRS" yaml:"supportedCrs"`
	TileMatrix   []TileMatrix `xml:"TileMatrix" yaml:"tileMatrix"`
}

// TileMatrix in struct for repeatability
type TileMatrix struct {
	Identifier       string `xml:"ows:Identifier" yaml:"identifier"`
	ScaleDenominator string `xml:"ScaleDenominator" yaml:"scaleDenominator"`
	TopLeftCorner    string `xml:"TopLeftCorner" yaml:"topLeftCorner"`
	TileWidth        string `xml:"TileWidth" yaml:"tileWidth"`
	TileHeight       string `xml:"TileHeight" yaml:"tileHeight"`
	MatrixWidth      string `xml:"MatrixWidth" yaml:"matrixWidth"`
	MatrixHeight     string `xml:"MatrixHeight" yaml:"matrixHeight"`
}

// LegendURL in struct for optionality
type LegendURL struct {
	Format string `xml:"format,attr,omitempty" yaml:"format,omitempty"`
	Href   string `xml:"xlink:href,attr,omitempty" yaml:"href,omitempty"`
}
