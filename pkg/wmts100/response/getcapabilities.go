package reponse

import "encoding/xml"

//
const (
	Service = `WMTS`
	Version = `1.0.0`
)

// Service function needed for the interface
func (wmts100 *Wmts100) Service() string {
	return Service
}

// Version function needed for the interface
func (wmts100 *Wmts100) Version() string {
	return Version
}

// Validate function of the wfs200 spec
func (wmts100 *Wmts100) Validate() bool {
	return false
}

// Wmts100 base struct
type Wmts100 struct {
	XMLName               xml.Name `xml:"Capabilities"`
	Namespaces            `yaml:"namespaces"`
	ServiceIdentification ServiceIdentification `xml:"ows:ServiceIdentification" yaml:"serviceidentification"`
	Contents              Contents              `xml:"Contents" yaml:"contents"`
	ServiceMetadataURL    ServiceMetadataURL    `xml:"ServiceMetadataURL" yaml:"servicemetadataurl"`
}

// Namespaces struct containing the namespaces needed for the XML document
type Namespaces struct {
	Xmlns          string `xml:"xmlns,attr" yaml:"xmlns"`       //http://www.opengis.net/wmts/1.0
	XmlnsOws       string `xml:"xmlns:ows,attr" yaml:"ows"`     //http://www.opengis.net/ows/1.1
	XmlnsXlink     string `xml:"xmlns:xlink,attr" yaml:"xlink"` //http://www.w3.org/1999/xlink
	XmlnsXSI       string `xml:"xmlns:xsi,attr" yaml:"xsi"`     //http://www.w3.org/2001/XMLSchema-instance
	XmlnsGml       string `xml:"xmlns:gml,attr" yaml:"gml"`     //http://www.opengis.net/gml
	Version        string `xml:"version,attr" yaml:"version"`
	SchemaLocation string `xml:"xsi:schemaLocation,attr" yaml:"schemalocation"`
}

// ServiceIdentification struct should only be fill by the "template" configuration wmts100.yaml
type ServiceIdentification struct {
	Title              string `xml:"ows:Title" yaml:"title"`
	Abstract           string `xml:"ows:Abstract" yaml:"abstract"`
	ServiceType        string `xml:"ows:ServiceType" yaml:"servicetype"`
	ServiceTypeVersion string `xml:"ows:ServiceTypeVersion" yaml:"servicetypeversion"`
	Fees               string `xml:"ows:Fees" yaml:"fees"`
	AccessConstraints  string `xml:"ows:AccessConstraints" yaml:"accessconstraints"`
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

// Layer in struct for repeatablity
type Layer struct {
	Title            string `xml:"ows:Title" yaml:"title"`
	Abstract         string `xml:"ows:Abstract" yaml:"abstract"`
	WGS84BoundingBox struct {
		LowerCorner string `xml:"ows:LowerCorner" yaml:"lowercorner"`
		UpperCorner string `xml:"ows:UpperCorner" yaml:"uppercorner"`
	} `xml:"ows:WGS84BoundingBox" yaml:"wgs84boundingbox"`
	Identifier string `xml:"ows:Identifier" yaml:"identifier"`
	Style      struct {
		Identifier string `xml:"ows:Identifier" yaml:"identifier"`
	} `xml:"Style" yaml:"style"`
	Format            string              `xml:"Format" yaml:"format"`
	TileMatrixSetLink []TileMatrixSetLink `xml:"TileMatrixSetLink" yaml:"tilematrixsetlink"`
	ResourceURL       struct {
		Format       string `xml:"format,attr" yaml:"format"`
		ResourceType string `xml:"resourceType,attr" yaml:"resourcetype"`
		Template     string `xml:"template,attr" yaml:"template"`
	} `xml:"ResourceURL" yaml:"resourceurl"`
}

// TileMatrixSetLink in struct for repeatablity
type TileMatrixSetLink struct {
	TileMatrixSet string `xml:"TileMatrixSet" yaml:"tilematrixset"`
}

// TileMatrixSet in struct for repeatablity
type TileMatrixSet struct {
	Identifier   string       `xml:"ows:Identifier" yaml:"identifier"`
	SupportedCRS string       `xml:"ows:SupportedCRS" yaml:"supportedcrs"`
	TileMatrix   []TileMatrix `xml:"TileMatrix" yaml:"tilematrix"`
}

// TileMatrix in struct for repeatablity
type TileMatrix struct {
	Identifier       string `xml:"ows:Identifier" yaml:"identifier"`
	ScaleDenominator string `xml:"ScaleDenominator" yaml:"scaledenominator"`
	TopLeftCorner    string `xml:"TopLeftCorner" yaml:"topleftcorner"`
	TileWidth        string `xml:"TileWidth" yaml:"tilewidth"`
	TileHeight       string `xml:"TileHeight" yaml:"tileheight"`
	MatrixWidth      string `xml:"MatrixWidth" yaml:"matrixwidth"`
	MatrixHeight     string `xml:"MatrixHeight" yaml:"matrixheight"`
}

// ServiceMetadataURL in struct for repeatablity
type ServiceMetadataURL struct {
	Href string `xml:"xlink:href,attr" yaml:"href"`
}
