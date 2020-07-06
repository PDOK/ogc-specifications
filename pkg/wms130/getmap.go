package wms130

import (
	"encoding/xml"
	"net/url"
	"strconv"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/ows"
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
	TIME        = `TIME`
	ELEVATION   = `ELEVATION`
)

var wmsmandatoryparameters = []string{LAYERS, STYLES, CRS, BBOX, WIDTH, HEIGHT, FORMAT}
var wmsoptionalparameters = []string{TRANSPARENT, BGCOLOR, EXCEPTIONS, TIME, ELEVATION}

// Type returns GetMap
func (gm *GetMap) Type() string {
	return getmap
}

// ParseBody builds a GetMap object based on the given body
func (gm *GetMap) ParseBody(body []byte) ows.Exception {
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

// ParseQuery builds a GetMap object based on the available query parameters
func (gm *GetMap) ParseQuery(query url.Values) ows.Exception {
	// Base
	for _, k := range WMSbaseparameters {
		if len(query[k]) > 0 {
			switch k {
			case REQUEST:
				if strings.ToUpper(query[k][0]) == strings.ToUpper(getmap) {
					gm.XMLName.Local = getmap
				}
			case SERVICE:
				gm.BaseRequest.Service = strings.ToUpper(query[k][0])
			case VERSION:
				gm.BaseRequest.Version = strings.ToUpper(query[k][0])
			}
		}
	}

	// WMS mandatory parameters

	var styles, layers []string

	if len(query[STYLES]) > 0 {
		styles = strings.Split(query[STYLES][0], ",")
	}
	if len(query[LAYERS]) > 0 {
		layers = strings.Split(query[LAYERS][0], ",")
	}

	sld, err := buildStyledLayerDescriptor(layers, styles)
	if err == nil {
		gm.StyledLayerDescriptor = sld
	} else {
		return err
	}

	if len(query[CRS]) > 0 {
		gm.CRS = query[CRS][0]
	}
	if len(query[BBOX]) > 0 {
		gm.BoundingBox = buildBoundingBox(query[BBOX][0])
	}
	if len(query[WIDTH]) > 0 {
		i, _ := strconv.Atoi(query[WIDTH][0])
		gm.Output.Size.Width = i
	}
	if len(query[HEIGHT]) > 0 {
		i, _ := strconv.Atoi(query[HEIGHT][0])
		gm.Output.Size.Height = i
	}
	if len(query[FORMAT]) > 0 {
		gm.Output.Format = query[FORMAT][0]
	}

	// WMS optional parameters
	for _, k := range wmsoptionalparameters {
		if len(query[k]) > 0 {
			switch k {
			case TRANSPARENT:
				gm.Output.Transparent = &query[k][0]
			case BGCOLOR:
				gm.Output.BGcolor = &query[k][0]
			case EXCEPTIONS:
				gm.Exceptions = &query[k][0]
				// case TIME:
				// No Time implementation (for now...)
				// Time format in ccyy-mm-ddThh:mm:ss.sssZ but also need support for time ranges
				// see: OGC 06-042 (WMS 1.3.0 spec)
				// case ELEVATION:
				// skip for now, same 'issue' as with the TIME
			}
		}
	}

	return nil
}

func buildBoundingBox(boundingbox string) ows.BoundingBox {
	bbox := strings.Split(boundingbox, ",")

	// check if all 'strings' are parsable to float64
	for _, crd := range bbox {
		_, err := strconv.ParseFloat(crd, 64)
		if err != nil {
			return ows.BoundingBox{}
		}
	}

	if len(bbox) == 4 {
		lcx, _ := strconv.ParseFloat(bbox[0], 64)
		lcy, _ := strconv.ParseFloat(bbox[1], 64)
		ucx, _ := strconv.ParseFloat(bbox[2], 64)
		ucy, _ := strconv.ParseFloat(bbox[3], 64)
		return ows.BoundingBox{LowerCorner: [2]float64{lcx, lcy},
			UpperCorner: [2]float64{ucx, ucy}}
	}
	return ows.BoundingBox{}
}

func buildStyledLayerDescriptor(layers, styles []string) (StyledLayerDescriptor, ows.Exception) {
	// Because the LAYERS & STYLES parameters are intertwined we proces as follows:
	// 1. cnt(STYLE) == 0 -> Added LAYERS
	// 2. cnt(LAYERS) == 0 -> Added no LAYERS (and no STYLES)
	// 3. cnt(LAYERS) == cnt(STYLES) -> merge LAYERS STYLES
	// 4. cnt(LAYERS) != cnt(STYLES) -> raise error Style not defined/Styles do not correspond with layers
	//    normally when 4 would occure this sould be done in the validate step... but,..
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
		return StyledLayerDescriptor{}, StyleNotDefined()
	}

	return StyledLayerDescriptor{}, nil
}

// BuildQuery builds a new query string that will be proxied
func (gm *GetMap) BuildQuery() url.Values {
	querystring := make(map[string][]string)

	// base
	querystring[REQUEST] = []string{gm.XMLName.Local}
	querystring[SERVICE] = []string{gm.BaseRequest.Service}
	querystring[VERSION] = []string{gm.BaseRequest.Version}

	for _, k := range wmsmandatoryparameters {
		switch k {
		case LAYERS:
			querystring[LAYERS] = []string{gm.StyledLayerDescriptor.getLayerQueryParameter()}
		case STYLES:
			querystring[STYLES] = []string{gm.StyledLayerDescriptor.getStyleQueryParameter()}
		case CRS:
			querystring[CRS] = []string{gm.CRS}
		case BBOX:
			querystring[BBOX] = []string{gm.BoundingBox.BuildQueryString()}
		case WIDTH:
			querystring[WIDTH] = []string{strconv.Itoa(gm.Output.Size.Width)}
		case HEIGHT:
			querystring[HEIGHT] = []string{strconv.Itoa(gm.Output.Size.Height)}
		case FORMAT:
			querystring[FORMAT] = []string{gm.Output.Format}
		}
	}

	for _, k := range wmsoptionalparameters {
		switch k {
		case TRANSPARENT:
			if gm.Output.Transparent != nil {
				querystring[TRANSPARENT] = []string{*gm.Output.Transparent}
			}
		case BGCOLOR:
			if gm.Output.BGcolor != nil {
				querystring[BGCOLOR] = []string{*gm.Output.BGcolor}
			}
		case EXCEPTIONS:
			if gm.Exceptions != nil {
				querystring[EXCEPTIONS] = []string{*gm.Exceptions}
			}
			// case TIME:
			// case ELEVATION:
		}
	}

	return querystring
}

// BuildBody builds a 'new' XML document 'based' on the 'original' XML document
func (gm *GetMap) BuildBody() []byte {
	si, _ := xml.MarshalIndent(gm, "", " ")
	return append([]byte(xml.Header), si...)
}

// GetMap struct with the needed parameters/attributes needed for making a GetMap request
// Struct based on http://schemas.opengis.net/sld/1.1//example_getmap.xml
type GetMap struct {
	XMLName xml.Name `xml:"GetMap" validate:"required"`
	BaseRequest
	StyledLayerDescriptor StyledLayerDescriptor `xml:"StyledLayerDescriptor" validate:"required"`
	CRS                   string                `xml:"CRS" validate:"required,epsg"`
	BoundingBox           ows.BoundingBox       `xml:"BoundingBox" validate:"required"`
	Output                Output                `xml:"Output" validate:"required"`
	Exceptions            *string               `xml:"Exceptions"`
	Elevation             *[]Elevation          `xml:"Elevation"`
	Time                  *string               `xml:"Time"`
}

// Output struct
type Output struct {
	Size        Size    `xml:"Size" validate:"required"`
	Format      string  `xml:"Format" validate:"required"`
	Transparent *string `xml:"Transparent"`
	BGcolor     *string `xml:"BGcolor"`
}

// Size struct
type Size struct {
	Width  int `xml:"Width" validate:"required,min=1,max=5000"`
	Height int `xml:"Height" validate:"required,min=1,max=5000"`
}

// StyledLayerDescriptor struct
type StyledLayerDescriptor struct {
	Version    string       `xml:"version,attr" validate:"required"`
	NamedLayer []NamedLayer `xml:"NamedLayer" validate:"required"`
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

// NamedLayer struct
type NamedLayer struct {
	Name       string      `xml:"Name" validate:"required"`
	NamedStyle *NamedStyle `xml:"NamedStyle"`
}

// NamedStyle contains the style name that needs be applied
type NamedStyle struct {
	Name string `xml:"Name" validate:"required"`
}

// Elevation struct for GetMap requests
// The extent string declares what value(s) along the Dimension axis are appropriate for the corresponding layer.
// The extent string has the syntax shown in Table C.2.
type Elevation struct {
	Value    float64 `xml:"Value"`
	Interval struct {
		Min float64 `xml:"Min"`
		Max float64 `xml:"Max"`
	} `xml:"Interval"`
}

// Validate a GetMap
func (gm *GetMap) Validate() ows.Exception {
	return nil
}
