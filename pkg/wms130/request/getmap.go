package request

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/ows"
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

// ParseQuery builds a GetMapKVP object based on the available query parameters
func (gmkvp *GetMapKVP) ParseQuery(query url.Values) ows.Exception {
	flatten := map[string]string{}
	for k, v := range query {
		if len(v) > 1 {
			// When there are is more then one value
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
	gmkvp.Layers = gm.StyledLayerDescriptor.getLayerQueryParameter()
	gmkvp.Styles = gm.StyledLayerDescriptor.getStyleQueryParameter()
	gmkvp.CRS = gm.CRS
	gmkvp.Bbox = gm.BoundingBox.BuildQueryString()
	gmkvp.Width = strconv.Itoa(gm.Output.Size.Width)
	gmkvp.Height = strconv.Itoa(gm.Output.Size.Height)
	gmkvp.Format = gm.Output.Format
	gmkvp.Transparent = gm.Output.Transparent
	gmkvp.BGColor = gm.Output.BGcolor
	// TODO: something with Time & Elevation
	// gmkvp.Time = gm.Time
	// gmkvp.Elevation = gm.Elevation
	gmkvp.Exceptions = gm.Exceptions

	return nil
}

// ParseKVP process the simple struct to a complex struct
func (gm *GetMap) ParseKVP(gmkvp GetMapKVP) ows.Exception {

	gm.BaseRequest.Build(gmkvp.Service, gmkvp.Version)

	var layers, styles []string
	if gmkvp.Layers != `` {
		layers = strings.Split(gmkvp.Layers, ",")
	}
	if gmkvp.Styles != `` {
		styles = strings.Split(gmkvp.Styles, ",")
	}

	sld, err := buildStyledLayerDescriptor(layers, styles)
	if err != nil {
		return err
	}
	gm.StyledLayerDescriptor = sld

	gm.CRS = gmkvp.CRS

	bbox, err := buildBoundingBox(gmkvp.Bbox)
	if err != nil {
		return err
	}
	gm.BoundingBox = bbox

	output, err := buildOutput(gmkvp.Height, gmkvp.Width, gmkvp.Format, gmkvp.Transparent, gmkvp.BGColor)
	if err != nil {
		return err
	}
	gm.Output = output

	gm.Exceptions = gmkvp.Exceptions

	return nil
}

// ParseQuery builds a GetMap object based on the available query parameters
func (gm *GetMap) ParseQuery(query url.Values) ows.Exception {

	if len(query) == 0 {
		// When there are no query value we know that at least
		// the manadorty VERSION parameter is missing.
		return ows.MissingParameterValue(VERSION)
	}

	gmkvp := GetMapKVP{}
	if err := gmkvp.ParseQuery(query); err != nil {
		return err
	}

	if err := gm.ParseKVP(gmkvp); err != nil {
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

// BuildQuery builds a url.Values query from a GetMapKVP struct
func (gmkvp *GetMapKVP) BuildQuery() url.Values {
	query := make(map[string][]string)

	fields := reflect.TypeOf(*gmkvp)
	values := reflect.ValueOf(*gmkvp)

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)
		// fmt.Print("Type:", field.Type, ",", field.Name, "=", value, "\n")

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

// BuildQuery builds a new query string that will be proxied
func (gm *GetMap) BuildQuery() url.Values {

	gmkvp := GetMapKVP{}
	gmkvp.ParseGetMap(gm)

	query := gmkvp.BuildQuery()

	return query
}

// BuildXML builds a 'new' XML document 'based' on the 'original' XML document
func (gm *GetMap) BuildXML() []byte {
	si, _ := xml.MarshalIndent(gm, "", " ")
	return append([]byte(xml.Header), si...)
}

func buildOutput(height, width, format string, transparent, bgcolor *string) (Output, ows.Exception) {

	output := Output{}

	h, err := strconv.Atoi(height)
	if err != nil {
		return output, ows.InvalidParameterValue(HEIGHT, height)
	}
	w, err := strconv.Atoi(width)
	if err != nil {
		return output, ows.InvalidParameterValue(WIDTH, width)
	}

	output.Size = Size{Height: h, Width: w}
	output.Format = format
	output.Transparent = transparent
	output.BGcolor = bgcolor

	return output, nil
}

func buildBoundingBox(boundingbox string) (ows.BoundingBox, ows.Exception) {
	result := strings.Split(boundingbox, ",")
	var lx, ly, ux, uy float64
	var err error

	if len(result) < 4 {
		return ows.BoundingBox{}, ows.InvalidParameterValue(boundingbox, BBOX)
	}

	if len(result) == 4 || len(result) == 5 {
		if lx, err = strconv.ParseFloat(result[0], 64); err != nil {
			return ows.BoundingBox{}, ows.InvalidParameterValue(boundingbox, BBOX)
		}
		if ly, err = strconv.ParseFloat(result[1], 64); err != nil {
			return ows.BoundingBox{}, ows.InvalidParameterValue(boundingbox, BBOX)
		}
		if ux, err = strconv.ParseFloat(result[2], 64); err != nil {
			return ows.BoundingBox{}, ows.InvalidParameterValue(boundingbox, BBOX)
		}
		if uy, err = strconv.ParseFloat(result[3], 64); err != nil {
			return ows.BoundingBox{}, ows.InvalidParameterValue(boundingbox, BBOX)
		}
	}

	if len(result) == 5 {
		return ows.BoundingBox{LowerCorner: [2]float64{lx, ly},
			UpperCorner: [2]float64{ux, uy}, Crs: result[4]}, nil
	}

	return ows.BoundingBox{LowerCorner: [2]float64{lx, ly},
		UpperCorner: [2]float64{ux, uy}}, nil
}

func buildStyledLayerDescriptor(layers, styles []string) (StyledLayerDescriptor, ows.Exception) {
	// Because the LAYERS & STYLES parameters are intertwined we process as follows:
	// 1. cnt(STYLE) == 0 -> Added LAYERS
	// 2. cnt(LAYERS) == 0 -> Added no LAYERS (and no STYLES)
	// 3. cnt(LAYERS) == cnt(STYLES) -> merge LAYERS STYLES
	// 4. cnt(LAYERS) != cnt(STYLES) -> raise error Style not defined/Styles do not correspond with layers
	//    normally when 4 would occure this could be done in the validate step... but,..
	//    with the serialisation -> struct it would become a valid object (yes!?.. YES!)
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
// so their are 'insync' ...?
func (sld *StyledLayerDescriptor) getLayerQueryParameter() string {
	queryvalue := ""
	for p, l := range sld.NamedLayer {
		queryvalue = queryvalue + l.Name
		if p < len(sld.NamedLayer)-1 {
			queryvalue = queryvalue + ","
		}
	}
	return queryvalue
}

func (sld *StyledLayerDescriptor) getStyleQueryParameter() string {
	queryvalue := ""
	for p, l := range sld.NamedLayer {
		if l.Name != "" {
			if l.NamedStyle != nil {
				queryvalue = queryvalue + l.NamedStyle.Name
			}
			if p < len(sld.NamedLayer)-1 {
				queryvalue = queryvalue + ","
			}
		}
	}
	return queryvalue
}

// GetMap struct with the needed parameters/attributes needed for making a GetMap request
// Struct based on http://schemas.opengis.net/sld/1.1//example_getmap.xml
type GetMap struct {
	XMLName xml.Name `xml:"GetMap" yaml:"getmap" validate:"required"`
	BaseRequest
	StyledLayerDescriptor StyledLayerDescriptor `xml:"StyledLayerDescriptor" yaml:"styledlayerdescriptor" validate:"required"`
	CRS                   string                `xml:"CRS" yaml:"crs" validate:"required"`
	BoundingBox           ows.BoundingBox       `xml:"BoundingBox" yaml:"boundingbox" validate:"required"`
	Output                Output                `xml:"Output" yaml:"output" validate:"required"`
	Exceptions            *string               `xml:"Exceptions" yaml:"exceptions"`
	// TODO: something with Time & Elevation
	// Elevation             *[]Elevation          `xml:"Elevation" yaml:"elavation"`
	// Time                  *string               `xml:"Time" yaml:"time"`
}

// Output struct
type Output struct {
	Size        Size    `xml:"Size" yaml:"size" validate:"required"`
	Format      string  `xml:"Format" yaml:"format" validate:"required"`
	Transparent *string `xml:"Transparent" yaml:"transparent"`
	BGcolor     *string `xml:"BGcolor" yaml:"bgcolor"`
}

// Size struct
type Size struct {
	Width  int `xml:"Width" yaml:"width" validate:"required,min=1,max=5000"`
	Height int `xml:"Height" yaml:"height" validate:"required,min=1,max=5000"`
}

// StyledLayerDescriptor struct
type StyledLayerDescriptor struct {
	Version    string       `xml:"version,attr" yaml:"version" validate:"required"`
	NamedLayer []NamedLayer `xml:"NamedLayer" yaml:"namedlayer" validate:"required"`
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
