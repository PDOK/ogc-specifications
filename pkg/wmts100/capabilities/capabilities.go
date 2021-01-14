package capabilities

// ParseXML func
func (c *Contents) ParseXML(doc []byte) error {
	return nil
}

// ParseYAMl func
func (c *Contents) ParseYAMl(doc []byte) error {
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
