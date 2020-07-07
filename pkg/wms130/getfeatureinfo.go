package wms130

import (
	"encoding/xml"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/ows"
)

// <map_request_copy>
var getFeatureMandatoryGetMapParameters = []string{LAYERS, CRS, BBOX, WIDTH, HEIGHT}

//
const (
	getfeatureinfo = `GetFeatureInfo`

	// Mandatory
	QUERYLAYERS = `QUERY_LAYERS`
	X           = `X`
	Y           = `Y`

	// Optional
	INFOFORMAT   = `INFO_FORMAT`
	FEATURECOUNT = `FEATURE_COUNT`
)

var getFeatureInfoMandatoryParameters, getFeatureInfoOptionalParameters []string

func init() {
	getFeatureInfoMandatoryParameters = append(getMapMandatoryParameters, []string{QUERYLAYERS, X, Y}...)
	getFeatureInfoOptionalParameters = append(getMapOptionalParameters, []string{INFOFORMAT, FEATURECOUNT}...)
}

// Type returns GetFeatureInfo
func (gfi *GetFeatureInfo) Type() string {
	return getfeatureinfo
}

// ParseBody builds a GetFeatureInfo object based on the given body
func (gfi *GetFeatureInfo) ParseBody(body []byte) ows.Exception {
	var xmlattributes ows.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return ows.MissingParameterValue()
	}
	if err := xml.Unmarshal(body, &gfi); err != nil {
		return ows.MissingParameterValue("REQUEST")
	}
	var n []xml.Attr
	for _, a := range xmlattributes {
		switch strings.ToUpper(a.Name.Local) {
		case VERSION:
		case SERVICE:
		default:
			n = append(n, a)
		}
	}

	gfi.Attr = ows.StripDuplicateAttr(n)
	return nil
}

// ParseQuery builds a GetFeatureInfo object based on the available query parameters
func (gfi *GetFeatureInfo) ParseQuery(query url.Values) ows.Exception {
	return nil
}

// BuildQuery builds a new query string that will be proxied
func (gfi *GetFeatureInfo) BuildQuery() url.Values {
	querystring := make(map[string][]string)

	// base
	querystring[REQUEST] = []string{gfi.XMLName.Local}
	querystring[SERVICE] = []string{gfi.BaseRequest.Service}
	querystring[VERSION] = []string{gfi.BaseRequest.Version}

	for _, k := range getFeatureInfoMandatoryParameters {
		switch k {
		case LAYERS:
			querystring[LAYERS] = []string{gfi.StyledLayerDescriptor.getLayerQueryParameter()}
		case STYLES:
			querystring[STYLES] = []string{gfi.StyledLayerDescriptor.getStyleQueryParameter()}
		case CRS:
			querystring[CRS] = []string{gfi.CRS}
		case BBOX:
			querystring[BBOX] = []string{gfi.BoundingBox.BuildQueryString()}
		case WIDTH:
			querystring[WIDTH] = []string{strconv.Itoa(gfi.Output.Size.Width)}
		case HEIGHT:
			querystring[HEIGHT] = []string{strconv.Itoa(gfi.Output.Size.Height)}
		case FORMAT:
			querystring[FORMAT] = []string{gfi.Output.Format}
		case QUERYLAYERS:
			querystring[QUERYLAYERS] = []string{strings.Join(gfi.QueryLayers, ",")}
		case X:
			querystring[X] = []string{strconv.Itoa(gfi.X)}
		case Y:
			querystring[Y] = []string{strconv.Itoa(gfi.Y)}
		}
	}

	for _, k := range getFeatureInfoOptionalParameters {
		switch k {
		case TRANSPARENT:
			if gfi.Output.Transparent != nil {
				querystring[TRANSPARENT] = []string{*gfi.Output.Transparent}
			}
		case BGCOLOR:
			if gfi.Output.BGcolor != nil {
				querystring[BGCOLOR] = []string{*gfi.Output.BGcolor}
			}
		case EXCEPTIONS:
			if gfi.Exceptions != nil {
				querystring[EXCEPTIONS] = []string{*gfi.Exceptions}
			}
		}
	}

	return querystring
}

// BuildBody builds a 'new' XML document 'based' on the 'original' XML document
func (gfi *GetFeatureInfo) BuildBody() []byte {
	si, _ := xml.MarshalIndent(gfi, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}

// GetFeatureInfo struct with the needed parameters/attributes needed for making a GetFeatureInfo request
type GetFeatureInfo struct {
	XMLName xml.Name `xml:"GetFeatureInfo" validate:"required"`
	BaseRequest
	// <map_request_copy>
	StyledLayerDescriptor StyledLayerDescriptor `xml:"LayerDescriptor" validate:"required"`
	CRS                   string                `xml:"CRS" validate:"required"`
	BoundingBox           ows.BoundingBox       `xml:"BoundingBox" validate:"required"`
	Output                Output                `xml:"Output" validate:"required"`

	QueryLayers  []string `xml:"querylayers" validate:"required"`
	InfoFormat   *string  `xml:"infoformat"`
	FeatureCount *int     `xml:"featurecount"`
	X            int      `xml:"x" validate:"required"`
	Y            int      `xml:"y" validate:"required"`
	Exceptions   *string  `xml:"Exceptions"`
}
