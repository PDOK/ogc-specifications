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
	I           = `I`
	J           = `J`

	// Optional
	INFOFORMAT   = `INFO_FORMAT`
	FEATURECOUNT = `FEATURE_COUNT`
)

var getFeatureInfoMandatoryParameters, getFeatureInfoOptionalParameters []string

func init() {
	getFeatureInfoMandatoryParameters = append(getMapMandatoryParameters, []string{QUERYLAYERS, I, J}...)
	getFeatureInfoOptionalParameters = append(getMapOptionalParameters, []string{INFOFORMAT, FEATURECOUNT}...)
}

// Type returns GetFeatureInfo
func (gfi *GetFeatureInfo) Type() string {
	return getfeatureinfo
}

// ParseBody builds a GetFeatureInfo object based on the given body
// Note: the XML GetFeatureInfo body that is consumed is a interpretation.
// So we use the GetMap, that is a large part of this request, as a base
// with the additional GetFeatureInfo parameters.
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
	// Base
	for _, k := range WMSbaseparameters {
		if len(query[k]) > 0 {
			switch k {
			case REQUEST:
				if strings.ToUpper(query[k][0]) == strings.ToUpper(getfeatureinfo) {
					gfi.XMLName.Local = getfeatureinfo
				}
			case SERVICE:
				gfi.BaseRequest.Service = strings.ToUpper(query[k][0])
			case VERSION:
				gfi.BaseRequest.Version = strings.ToUpper(query[k][0])
			}
		}
	}

	var styles, layers []string

	// GetFeatureInfo mandatory parameters
	for _, k := range getFeatureInfoMandatoryParameters {
		if len(query[k]) > 0 {
			switch k {
			case LAYERS:
				layers = strings.Split(query[k][0], ",")
			case STYLES:
				styles = strings.Split(query[k][0], ",")
			case CRS:
				gfi.CRS = query[k][0]
			case BBOX:
				gfi.BoundingBox = buildBoundingBox(query[k][0])
			case WIDTH:
				i, err := strconv.Atoi(query[k][0])
				if err != nil {
					return ows.MissingParameterValue(WIDTH, query[k][0])
				}
				gfi.Size.Width = i
			case HEIGHT:
				i, err := strconv.Atoi(query[k][0])
				if err != nil {
					// TODO: ignore or a exception
					return ows.MissingParameterValue(HEIGHT, query[k][0])
				}
				gfi.Size.Height = i
			case QUERYLAYERS:
				gfi.QueryLayers = strings.Split(query[k][0], ",")
			case I:
				i, err := strconv.Atoi(query[k][0])
				if err != nil {
					return InvalidPoint(query[I][0], query[J][0])
				}
				gfi.I = i
			case J:
				i, err := strconv.Atoi(query[k][0])
				if err != nil {
					return InvalidPoint(query[I][0], query[J][0])
				}
				gfi.J = i
			}
		}
	}

	sld, err := buildStyledLayerDescriptor(layers, styles)
	if err == nil {
		gfi.StyledLayerDescriptor = sld
	} else {
		return err
	}

	// GetFeatureInfo optional parameters
	for _, k := range getFeatureInfoOptionalParameters {
		if len(query[k]) > 0 {
			switch k {
			case INFOFORMAT:
				gfi.InfoFormat = &query[k][0]
			case FEATURECOUNT:
				i, err := strconv.Atoi(query[k][0])
				if err != nil {
					// TODO: ignore or a exception
				}
				gfi.FeatureCount = &i
			case EXCEPTIONS:
				gfi.Exceptions = &query[k][0]
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
			querystring[WIDTH] = []string{strconv.Itoa(gfi.Size.Width)}
		case HEIGHT:
			querystring[HEIGHT] = []string{strconv.Itoa(gfi.Size.Height)}
		case QUERYLAYERS:
			querystring[QUERYLAYERS] = []string{strings.Join(gfi.QueryLayers, ",")}
		case I:
			querystring[J] = []string{strconv.Itoa(gfi.I)}
		case J:
			querystring[I] = []string{strconv.Itoa(gfi.J)}
		}
	}

	for _, k := range getFeatureInfoOptionalParameters {
		switch k {
		case INFOFORMAT:
			if gfi.InfoFormat != nil {
				querystring[INFOFORMAT] = []string{*gfi.InfoFormat}
			}
		case FEATURECOUNT:
			if gfi.FeatureCount != nil {
				querystring[FEATURECOUNT] = []string{strconv.Itoa(*gfi.FeatureCount)}
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
// Note: this GetFeatureInfo XML body is a interpretation and there isn't a
// good/real OGC exmaple request. So for now we use the GetMap, that is a large part
// of this request, as a base with the additional GetFeatureInfo parameters.
func (gfi *GetFeatureInfo) BuildBody() []byte {
	si, _ := xml.MarshalIndent(gfi, "", " ")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
	// return []byte(xml.Header + string(si))
}

// GetFeatureInfo struct with the needed parameters/attributes needed for making a GetFeatureInfo request
type GetFeatureInfo struct {
	XMLName xml.Name `xml:"GetFeatureInfo" validate:"required"`
	BaseRequest

	// <map_request_copy>
	// These are the 'minimum' required GetMap parameters
	// needed in a GetFeatureInfo request
	StyledLayerDescriptor StyledLayerDescriptor `xml:"StyledLayerDescriptor" validate:"required"`
	CRS                   string                `xml:"CRS" validate:"required"`
	BoundingBox           ows.BoundingBox       `xml:"BoundingBox" validate:"required"`
	// We skip the OutPut struct, because these are not required parameters
	Size Size `xml:"Size" validate:"required"`

	QueryLayers  []string `xml:"QueryLayers" validate:"required"`
	InfoFormat   *string  `xml:"InfoFormat"`
	FeatureCount *int     `xml:"FeatureCount"`
	I            int      `xml:"I" validate:"required"`
	J            int      `xml:"J" validate:"required"`
	Exceptions   *string  `xml:"Exceptions"`
}
