package wms130

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/utils"
)

// Mandatory GetMap Keys
const (
	LAYERS = `LAYERS`
	STYLES = `STYLES`
	BBOX   = `BBOX`
	WIDTH  = `WIDTH`
	HEIGHT = `HEIGHT`
	FORMAT = `FORMAT`
)

// Optional GetMap Keys
const (
	TRANSPARENT = `TRANSPARENT`
	BGCOLOR     = `BGCOLOR`
	EXCEPTIONS  = `EXCEPTIONS` // defaults to XML

	// TODO: something with Time & Elevation
	// TIME        = `TIME`
	// ELEVATION   = `ELEVATION`
)

// GetMap struct with the needed parameters/attributes needed for making a GetMap request
// Struct based on http://schemas.opengis.net/sld/1.1/example_getmap.xml
type GetMapRequest struct {
	XMLName xml.Name `xml:"GetMap" yaml:"getmap"`
	BaseRequest
	StyledLayerDescriptor StyledLayerDescriptor `xml:"StyledLayerDescriptor" yaml:"styledlayerdescriptor"`
	CRS                   CRS                   `xml:"CRS" yaml:"crs"`
	BoundingBox           BoundingBox           `xml:"BoundingBox" yaml:"boundingbox"`
	Output                Output                `xml:"Output" yaml:"output"`
	Exceptions            *string               `xml:"Exceptions" yaml:"exceptions"`
	// TODO: something with Time & Elevation
	// Elevation             *[]Elevation          `xml:"Elevation" yaml:"elevation"`
	// Time                  *string               `xml:"Time" yaml:"time"`BuildKVP
}

// Validate returns GetMap
func (gm *GetMapRequest) Validate(c Capabilities) Exceptions {
	var exceptions Exceptions

	exceptions = append(exceptions, gm.StyledLayerDescriptor.Validate(c)...)
	exceptions = append(exceptions, gm.Output.Validate(c)...)

	for _, sld := range gm.StyledLayerDescriptor.NamedLayer {
		layer, layerexception := c.GetLayer(sld.Name)
		if layerexception != nil {
			exceptions = append(exceptions, layerexception...)
		}
		if CRSException := checkCRS(gm.CRS, layer.CRS); CRSException != nil {
			exceptions = append(exceptions, InvalidCRS(gm.CRS.String(), sld.Name))
		}
	}

	return exceptions
}

// ParseOperationRequestKVP process the simple struct to a complex struct
func (gm *GetMapRequest) parseKVP(gmkvp getMapKVPRequest) Exceptions {
	gm.XMLName.Local = getmap
	gm.BaseRequest.parseKVP(gmkvp.baseRequestKVP)

	sld, exceptions := gmkvp.buildStyledLayerDescriptor()
	if exceptions != nil {
		return exceptions
	}
	gm.StyledLayerDescriptor = sld

	var crs CRS
	crs.parseString(gmkvp.crs)
	gm.CRS = crs

	var bbox BoundingBox
	if exceptions := bbox.parseString(gmkvp.bbox); exceptions != nil {
		return exceptions
	}
	gm.BoundingBox = bbox

	output, exceptions := gmkvp.buildOutput()
	if exceptions != nil {
		return exceptions
	}
	gm.Output = output

	gm.Exceptions = gmkvp.exceptions

	return nil
}

// ParseKVP builds a GetMap object based on the available query parameters
func (gm *GetMapRequest) ParseQueryParameters(query url.Values) Exceptions {
	if len(query) == 0 {
		// When there are no query values we know that at least
		// the manadorty VERSION and REQUEST parameter is missing.
		return Exceptions{MissingParameterValue(VERSION), MissingParameterValue(REQUEST)}
	}

	gmkvp := getMapKVPRequest{}
	if exceptions := gmkvp.parseQueryParameters(query); exceptions != nil {
		return exceptions
	}

	if exceptions := gm.parseKVP(gmkvp); exceptions != nil {
		return exceptions
	}

	return nil
}

// ParseXML builds a GetMap object based on a XML document
func (gm *GetMapRequest) ParseXML(body []byte) Exceptions {
	var xmlattributes utils.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return Exceptions{MissingParameterValue()}
	}
	xml.Unmarshal(body, &gm) //When object can be Unmarshalled -> XMLAttributes, it can be Unmarshalled -> GetMap
	var n []xml.Attr
	for _, a := range xmlattributes {
		switch strings.ToUpper(a.Name.Local) {
		case VERSION:
		default:
			n = append(n, a)
		}
	}
	gm.BaseRequest.Attr = utils.StripDuplicateAttr(n)
	return nil
}

// BuildKVP builds a new query string that will be proxied
func (gm GetMapRequest) ToQueryParameters() url.Values {
	gmkvp := getMapKVPRequest{}
	gmkvp.parseGetMapRequest(gm)

	q := gmkvp.toQueryParameters()
	return q
}

// BuildXML builds a 'new' XML document 'based' on the 'original' XML document
func (gm *GetMapRequest) ToXML() []byte {
	si, _ := xml.MarshalIndent(gm, "", " ")
	return append([]byte(xml.Header), si...)
}

// Output struct
type Output struct {
	Size        Size    `xml:"Size" yaml:"size"`
	Format      string  `xml:"Format" yaml:"format"`
	Transparent *bool   `xml:"Transparent" yaml:"transparent"`
	BGcolor     *string `xml:"BGcolor" yaml:"bgcolor"`
}

// Validate validates the output parameters
func (output *Output) Validate(c Capabilities) Exceptions {
	exceptions := Exceptions{}
	if c.MaxWidth > 0 && output.Size.Width > c.MaxWidth {
		exceptions = append(exceptions, NoApplicableCode(fmt.Sprintf("Image size out of range, WIDTH must be between 1 and %d pixels", c.MaxWidth)))
	}
	if c.MaxHeight > 0 && output.Size.Height > c.MaxHeight {
		exceptions = append(exceptions, NoApplicableCode(fmt.Sprintf("Image size out of range, HEIGHT must be between 1 and %d pixels", c.MaxHeight)))
	}

	for _, format := range c.WMSCapabilities.Request.GetMap.Format {
		found := false
		if format == output.Format {
			found = true
		}
		if !found {
			exceptions = append(exceptions, InvalidFormat(output.Format))
		}
	}

	if exceptions != nil {
		return exceptions
	}

	// Transparent is a bool so when it is parsed around in the application it is already valid
	// TODO: BGColor -> https://stackoverflow.com/questions/54197913/parse-hex-string-to-image-color

	return nil
}

// Size struct
type Size struct {
	Width  int `xml:"Width" yaml:"width"`
	Height int `xml:"Height" yaml:"height"`
}

// StyledLayerDescriptor struct
type StyledLayerDescriptor struct {
	Version    string       `xml:"version,attr" yaml:"version"`
	NamedLayer []NamedLayer `xml:"NamedLayer" yaml:"namedlayer"`
}

// Validate the StyledLayerDescriptor
func (sld StyledLayerDescriptor) Validate(c Capabilities) Exceptions {
	var unknownLayers []string
	var unknownStyles []struct{ layer, style string }

	for _, namedLayer := range sld.NamedLayer {
		found := false
		for _, c := range c.GetLayerNames() {
			if namedLayer.Name == c {
				found = true
			}
		}
		if !found {
			unknownLayers = append(unknownLayers, namedLayer.Name)
		}

		if namedLayer.NamedStyle != nil {
			if namedLayer.NamedStyle.Name != `` {
				if !c.StyleDefined(namedLayer.Name, namedLayer.NamedStyle.Name) {
					unknownStyles = append(unknownStyles, struct{ layer, style string }{namedLayer.Name, namedLayer.NamedStyle.Name})
				}
			}
		}
	}

	var exceptions Exceptions
	if len(unknownLayers) > 0 {
		for _, l := range unknownLayers {
			exceptions = append(exceptions, LayerNotDefined(l))
		}
	}

	if len(unknownStyles) > 0 {
		for _, sld := range unknownStyles {
			exceptions = append(exceptions, StyleNotDefined(sld.style, sld.layer))
		}
	}

	if len(exceptions) > 0 {
		return exceptions
	}

	return nil
}

// TODO maybe 'merge' both func in a single one with 2 outputs
// so their are 'in sync' ...?
func (sld *StyledLayerDescriptor) getLayerKVPValue() string {
	return strings.Join(sld.getNamedLayers(), ",")
}

func (sld *StyledLayerDescriptor) getStyleKVPValue() string {
	return strings.Join(sld.getNamedStyles(), ",")
}

// GetNamedLayers return an array of the Layer names
func (sld *StyledLayerDescriptor) getNamedLayers() []string {
	layers := []string{}
	for _, l := range sld.NamedLayer {
		layers = append(layers, l.Name)
	}
	return layers
}

// GetNamedStyles return an array of the Layer names
func (sld *StyledLayerDescriptor) getNamedStyles() []string {
	styles := []string{}
	for _, l := range sld.NamedLayer {
		if l.Name != "" {
			if l.NamedStyle != nil {
				styles = append(styles, l.NamedStyle.Name)
			} else {
				styles = append(styles, "")
			}
		}
	}
	return styles
}

// NamedLayer struct
type NamedLayer struct {
	Name       string      `xml:"Name" yaml:"name"`
	NamedStyle *NamedStyle `xml:"NamedStyle" yaml:"namedstyle"`
}

// NamedStyle contains the style name that needs be applied
type NamedStyle struct {
	Name string `xml:"Name" yaml:"name"`
}

// Elevation struct for GetMap requests
// The extent string declares what value(s) along the Dimension axis are appropriate for the corresponding layer.
// The extent string has the syntax shown in Table C.2.
type Elevation struct {
	Value    float64 `xml:"Value" yaml:"value"`
	Interval struct {
		Min float64 `xml:"Min" yaml:"min"`
		Max float64 `xml:"Max" yaml:"max"`
	} `xml:"Interval" yaml:"interval"`
}

func buildStyledLayerDescriptor(layers, styles []string) (StyledLayerDescriptor, Exceptions) {
	// Because the LAYERS & STYLES parameters are intertwined we process as follows:
	// 1. cnt(STYLE) == 0 -> Added LAYERS
	// 2. cnt(LAYERS) == 0 -> Added no LAYERS (and no STYLES)
	// 3. cnt(LAYERS) == cnt(STYLES) -> merge LAYERS STYLES
	// 4. cnt(LAYERS) != cnt(STYLES) -> raise error Style not defined/Styles do not correspond with layers
	//    normally when 4 would occur this could be done in the validate step... but,..
	//    with the serialization -> struct it would become a valid object (yes!?.. YES!)
	//    That is because POST xml and GET KVP handle this 'different' (at least not in the same way...)
	//    When 3 is hit the validation at the Validation step wil resolve this

	// 1.
	if len(styles) == 0 {
		var sld StyledLayerDescriptor
		for _, layer := range layers {
			sld.NamedLayer = append(sld.NamedLayer, NamedLayer{Name: layer})
		}
		sld.Version = "1.1.0"
		return sld, nil
		// 2.
	} else if len(layers) == 0 {
		// do nothing
		// will be resolved during validation

		// 3.
	} else if len(layers) == len(styles) {
		var sld StyledLayerDescriptor
		for k, layer := range layers {
			sld.NamedLayer = append(sld.NamedLayer, NamedLayer{Name: layer, NamedStyle: &NamedStyle{Name: styles[k]}})
		}
		sld.Version = "1.1.0"
		return sld, nil
		// 4.
	} else if len(layers) != len(styles) {
		return StyledLayerDescriptor{}, StyleNotDefined().ToExceptions()
	}

	return StyledLayerDescriptor{}, nil
}

// checkCRS against a given list of CRS
func checkCRS(crs CRS, definedCrs []CRS) Exceptions {
	for _, defined := range definedCrs {
		if defined == crs {
			return nil
		}
	}
	return InvalidCRS(crs.String()).ToExceptions()
}
