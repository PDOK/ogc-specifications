package request

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/ows"
	"github.com/pdok/ogc-specifications/pkg/wms130/capabilities"
	"github.com/pdok/ogc-specifications/pkg/wms130/exception"
	"gopkg.in/yaml.v2"
)

//
const (
	getmap = `GetMap`

	// Mandatory
	LAYERS = `LAYERS`
	STYLES = `STYLES`
	CRS    = `CRS`
	BBOX   = `BBOX`
	WIDTH  = `WIDTH`
	HEIGHT = `HEIGHT`
	FORMAT = `FORMAT`

	//Optional
	TRANSPARENT = `TRANSPARENT`
	BGCOLOR     = `BGCOLOR`
	EXCEPTIONS  = `EXCEPTIONS` // defaults to XML

	// TODO: something with Time & Elevation
	// TIME        = `TIME`
	// ELEVATION   = `ELEVATION`
)

// Type returns GetMap
func (gm *GetMap) Type() string {
	return getmap
}

// Validate returns GetMap
func (gm *GetMap) Validate(c capabilities.Capability) ows.Exceptions {
	var exceptions ows.Exceptions

	exceptions = append(exceptions, gm.StyledLayerDescriptor.Validate(c)...)
	exceptions = append(exceptions, gm.Output.Validate(c)...)

	return exceptions
}

//GetMapKVP struct
type GetMapKVP struct {
	// Table 8 - The Parameters of a GetMap request
	Request     string  `yaml:"request,omitempty"`
	Version     string  `yaml:"version,omitempty"`
	Service     string  `yaml:"service,omitempty"`
	Layers      string  `yaml:"layers,omitempty"`
	Styles      string  `yaml:"styles,omitempty"`
	CRS         string  `yaml:"crs,omitempty"`
	Bbox        string  `yaml:"bbox,omitempty"`
	Width       string  `yaml:"width,omitempty"`
	Height      string  `yaml:"height,omitempty"`
	Format      string  `yaml:"format,omitempty"`
	Transparent *string `yaml:"transparent,omitempty"`
	BGColor     *string `yaml:"bgcolor,omitempty"`
	// TODO: something with Time & Elevation
	// Time        *string `yaml:"time,omitempty"`
	// Elevation   *string `yaml:"elevation,omitempty"`
	Exceptions *string `yaml:"exceptions,omitempty"`
}

// ParseKVP builds a GetMapKVP object based on the available query parameters
func (gmkvp *GetMapKVP) ParseKVP(query url.Values) ows.Exception {
	flatten := map[string]string{}
	for k, v := range query {
		if len(v) > 1 {
			// When there is more then one value
			// return a InvalidParameterValue Exception
			return ows.InvalidParameterValue(k, strings.Join(v, ","))
		}
		flatten[strings.ToLower(k)] = v[0]
	}

	y, _ := yaml.Marshal(&flatten)
	if err := yaml.Unmarshal(y, &gmkvp); err != nil {
		return ows.NoApplicableCode(`Could not read query parameters`)
	}

	return nil
}

// ParseGetMap builds a GetMapKVP object based on a GetMap struct
func (gmkvp *GetMapKVP) ParseGetMap(gm *GetMap) ows.Exception {
	gmkvp.Request = getmap
	gmkvp.Version = Version
	gmkvp.Service = Service
	gmkvp.Layers = gm.StyledLayerDescriptor.getLayerKVPValue()
	gmkvp.Styles = gm.StyledLayerDescriptor.getStyleKVPValue()
	gmkvp.CRS = gm.CRS
	gmkvp.Bbox = gm.BoundingBox.BuildKVP()
	gmkvp.Width = strconv.Itoa(gm.Output.Size.Width)
	gmkvp.Height = strconv.Itoa(gm.Output.Size.Height)
	gmkvp.Format = gm.Output.Format

	if gm.Output.Transparent != nil {
		t := *gm.Output.Transparent
		tp := strconv.FormatBool(t)
		gmkvp.Transparent = &tp
	}

	if gm.Output.BGcolor != nil {
		gmkvp.BGColor = gm.Output.BGcolor
	}

	// TODO: something with Time & Elevation
	// gmkvp.Time = gm.Time
	// gmkvp.Elevation = gm.Elevation

	gmkvp.Exceptions = gm.Exceptions

	return nil
}

// BuildOutput builds a Output struct from the KVP information
func (gmkvp *GetMapKVP) BuildOutput() (Output, ows.Exception) {
	output := Output{}

	h, err := strconv.Atoi(gmkvp.Height)
	if err != nil {
		return output, ows.InvalidParameterValue(HEIGHT, gmkvp.Height)
	}
	w, err := strconv.Atoi(gmkvp.Width)
	if err != nil {
		return output, ows.InvalidParameterValue(WIDTH, gmkvp.Width)
	}

	output.Size = Size{Height: h, Width: w}
	output.Format = gmkvp.Format
	if b, err := strconv.ParseBool(*gmkvp.Transparent); err == nil {
		output.Transparent = &b
	}
	output.BGcolor = gmkvp.BGColor

	return output, nil
}

// BuildStyledLayerDescriptor builds a StyledLayerDescriptor struct from the KVP information
func (gmkvp *GetMapKVP) BuildStyledLayerDescriptor() (StyledLayerDescriptor, ows.Exception) {
	var layers, styles []string
	if gmkvp.Layers != `` {
		layers = strings.Split(gmkvp.Layers, ",")
	}
	if gmkvp.Styles != `` {
		styles = strings.Split(gmkvp.Styles, ",")
	}

	sld, err := buildStyledLayerDescriptor(layers, styles)
	if err != nil {
		return sld, err
	}

	return sld, nil
}

// BuildKVP builds a url.Values query from a GetMapKVP struct
func (gmkvp *GetMapKVP) BuildKVP() url.Values {
	query := make(map[string][]string)

	fields := reflect.TypeOf(*gmkvp)
	values := reflect.ValueOf(*gmkvp)

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)

		switch value.Kind() {
		case reflect.String:
			v := value.String()
			query[strings.ToUpper(field.Name)] = []string{v}
		case reflect.Ptr:
			v := value.Elem()
			if v.IsValid() {
				query[strings.ToUpper(field.Name)] = []string{fmt.Sprintf("%v", v)}
			}
		}
	}

	return query
}

// ParseGetMapKVP process the simple struct to a complex struct
func (gm *GetMap) ParseGetMapKVP(gmkvp GetMapKVP) ows.Exception {
	gm.BaseRequest.Build(gmkvp.Service, gmkvp.Version)

	sld, err := gmkvp.BuildStyledLayerDescriptor()
	if err != nil {
		return err
	}
	gm.StyledLayerDescriptor = sld

	gm.CRS = gmkvp.CRS

	var bbox ows.BoundingBox
	if err := bbox.Build(gmkvp.Bbox); err != nil {
		return err
	}
	gm.BoundingBox = bbox

	output, err := gmkvp.BuildOutput()
	if err != nil {
		return err
	}
	gm.Output = output

	gm.Exceptions = gmkvp.Exceptions

	return nil
}

// ParseKVP builds a GetMap object based on the available query parameters
func (gm *GetMap) ParseKVP(query url.Values) ows.Exception {
	if len(query) == 0 {
		// When there are no query values we know that at least
		// the manadorty VERSION parameter is missing.
		return ows.MissingParameterValue(VERSION)
	}

	gmkvp := GetMapKVP{}
	if err := gmkvp.ParseKVP(query); err != nil {
		return err
	}

	if err := gm.ParseGetMapKVP(gmkvp); err != nil {
		return err
	}

	return nil
}

// ParseXML builds a GetMap object based on a XML document
func (gm *GetMap) ParseXML(body []byte) ows.Exception {
	var xmlattributes ows.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return ows.MissingParameterValue()
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
	gm.BaseRequest.Attr = ows.StripDuplicateAttr(n)
	return nil
}

// BuildKVP builds a new query string that will be proxied
func (gm *GetMap) BuildKVP() url.Values {
	gmkvp := GetMapKVP{}
	gmkvp.ParseGetMap(gm)

	query := gmkvp.BuildKVP()
	// query := map[string][]string{}

	return query
}

// BuildXML builds a 'new' XML document 'based' on the 'original' XML document
func (gm *GetMap) BuildXML() []byte {
	si, _ := xml.MarshalIndent(gm, "", " ")
	return append([]byte(xml.Header), si...)
}

func buildStyledLayerDescriptor(layers, styles []string) (StyledLayerDescriptor, ows.Exception) {
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
		return StyledLayerDescriptor{}, exception.StyleNotDefined()
	}

	return StyledLayerDescriptor{}, nil
}

// TODO maybe 'merge' both func in a single one with 2 outputs
// so their are 'in sync' ...?
func (sld *StyledLayerDescriptor) getLayerKVPValue() string {
	return strings.Join(sld.GetNamedLayers(), ",")
}

func (sld *StyledLayerDescriptor) getStyleKVPValue() string {
	return strings.Join(sld.GetNamedStyles(), ",")
}

// GetNamedLayers return an array of the Layer names
func (sld *StyledLayerDescriptor) GetNamedLayers() []string {
	layers := []string{}
	for _, l := range sld.NamedLayer {
		layers = append(layers, l.Name)
	}

	return layers
}

// GetNamedStyles return an array of the Layer names
func (sld *StyledLayerDescriptor) GetNamedStyles() []string {
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

// GetMap struct with the needed parameters/attributes needed for making a GetMap request
// Struct based on http://schemas.opengis.net/sld/1.1/example_getmap.xml
type GetMap struct {
	XMLName xml.Name `xml:"GetMap" yaml:"getmap" validate:"required"`
	BaseRequest
	StyledLayerDescriptor StyledLayerDescriptor `xml:"StyledLayerDescriptor" yaml:"styledlayerdescriptor" validate:"required"`
	CRS                   string                `xml:"CRS" yaml:"crs" validate:"required"`
	BoundingBox           ows.BoundingBox       `xml:"BoundingBox" yaml:"boundingbox" validate:"required"`
	Output                Output                `xml:"Output" yaml:"output" validate:"required"`
	Exceptions            *string               `xml:"Exceptions" yaml:"exceptions"`
	// TODO: something with Time & Elevation
	// Elevation             *[]Elevation          `xml:"Elevation" yaml:"elevation"`
	// Time                  *string               `xml:"Time" yaml:"time"`
}

// Validate validates the output parameters
func (output *Output) Validate(c capabilities.Capability) ows.Exceptions {
	var exceptions ows.Exceptions
	if output.Size.Width > c.MaxWidth {
		exceptions = append(exceptions, ows.NoApplicableCode(fmt.Sprintf("Image size out of range, WIDTH must be between 1 and %d pixels", c.MaxWidth)))
	}
	if output.Size.Height > c.MaxHeight {
		exceptions = append(exceptions, ows.NoApplicableCode(fmt.Sprintf("Image size out of range, HEIGHT must be between 1 and %d pixels", c.MaxHeight)))
	}

	for _, format := range c.WMSCapabilities.Request.GetMap.Format {
		found := false
		if format == output.Format {
			found = true
		}
		if !found {
			exceptions = append(exceptions, exception.InvalidFormat(output.Format))
		}
	}

	// Transparent is a bool so when it is parsed around in the application it is already valid
	// TODO: BGColor -> https://stackoverflow.com/questions/54197913/parse-hex-string-to-image-color

	return nil
}

// Output struct
type Output struct {
	Size        Size    `xml:"Size" yaml:"size" validate:"required"`
	Format      string  `xml:"Format" yaml:"format" validate:"required"`
	Transparent *bool   `xml:"Transparent" yaml:"transparent"`
	BGcolor     *string `xml:"BGcolor" yaml:"bgcolor"`
}

// Size struct
type Size struct {
	Width  int `xml:"Width" yaml:"width" validate:"required,min=1,max=5000"`
	Height int `xml:"Height" yaml:"height" validate:"required,min=1,max=5000"`
}

// StyledLayerDescriptor struct
type StyledLayerDescriptor struct {
	Version    string       `xml:"version,attr" yaml:"version" validate:"required,eq=1.1.0"`
	NamedLayer []NamedLayer `xml:"NamedLayer" yaml:"namedlayer" validate:"required"`
}

// Validate the StyledLayerDescriptor
func (sld *StyledLayerDescriptor) Validate(c capabilities.Capability) ows.Exceptions {
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

	var exceptions ows.Exceptions
	if len(unknownLayers) > 0 {
		for _, l := range unknownLayers {
			exceptions = append(exceptions, exception.LayerNotDefined(l))
		}
	}

	if len(unknownStyles) > 0 {
		for _, sld := range unknownStyles {
			exceptions = append(exceptions, exception.StyleNotDefined(sld.style, sld.layer))
		}
	}

	if len(exceptions) > 0 {
		return exceptions
	}

	return nil
}

// NamedLayer struct
type NamedLayer struct {
	Name       string      `xml:"Name" yaml:"name" validate:"required"`
	NamedStyle *NamedStyle `xml:"NamedStyle" yaml:"namedstyle"`
}

// NamedStyle contains the style name that needs be applied
type NamedStyle struct {
	Name string `xml:"Name" yaml:"name" validate:"required"`
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
