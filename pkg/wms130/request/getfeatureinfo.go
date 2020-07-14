package request

import (
	"encoding/xml"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/ows"
	"github.com/pdok/ogc-specifications/pkg/utils"
	"github.com/pdok/ogc-specifications/pkg/wms130/exception"
)

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
	getFeatureInfoMandatoryParameters = []string{LAYERS, STYLES, CRS, BBOX, WIDTH, HEIGHT, FORMAT, QUERYLAYERS, I, J}
	getFeatureInfoOptionalParameters = []string{TRANSPARENT, BGCOLOR, EXCEPTIONS, INFOFORMAT, FEATURECOUNT}
}

// Type returns GetFeatureInfo
func (gfi *GetFeatureInfo) Type() string {
	return getfeatureinfo
}

// ParseXML builds a GetFeatureInfo object based on a XML document
// Note: the XML GetFeatureInfo body that is consumed is a interpretation.
// So we use the GetMap, that is a large part of this request, as a base
// with the additional GetFeatureInfo parameters.
func (gfi *GetFeatureInfo) ParseXML(body []byte) ows.Exception {
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

	if len(query) == 0 {
		// When there are no query value we know that at least
		// the manadorty VERSION parameter is missing.
		return ows.MissingParameterValue(VERSION)
	}

	q := utils.KeysToUpper(query)

	// Base
	if len(q[REQUEST]) > 0 {
		gfi.XMLName.Local = q[REQUEST][0]
	}

	var br BaseRequest
	if err := br.ParseQueryParameters(q); err != nil {
		return err
	}
	gfi.BaseRequest = br

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
				var bbox ows.BoundingBox
				var err ows.Exception
				if bbox, err = buildBoundingBox(query[k][0]); err != nil {
					return err
				}
				gfi.BoundingBox = bbox
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
					return exception.InvalidPoint(query[I][0], query[J][0])
				}
				gfi.I = i
			case J:
				i, err := strconv.Atoi(query[k][0])
				if err != nil {
					return exception.InvalidPoint(query[I][0], query[J][0])
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

// BuildXML builds a 'new' XML document 'based' on the 'original' XML document
// Note: this GetFeatureInfo XML body is a interpretation and there isn't a
// good/real OGC example request. So for now we use the GetMap, that is a large part
// of this request, as a base with the additional GetFeatureInfo parameters.
func (gfi *GetFeatureInfo) BuildXML() []byte {
	si, _ := xml.MarshalIndent(gfi, "", " ")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
	// return []byte(xml.Header + string(si))
}

// GetFeatureInfo struct with the needed parameters/attributes needed for making a GetFeatureInfo request
type GetFeatureInfo struct {
	XMLName xml.Name `xml:"GetFeatureInfo" yaml:"getfeatureinfo" validate:"required"`
	BaseRequest

	// <map_request_copy>
	// These are the 'minimum' required GetMap parameters
	// needed in a GetFeatureInfo request
	StyledLayerDescriptor StyledLayerDescriptor `xml:"StyledLayerDescriptor" yaml:"styledlayerdescriptor" validate:"required"`
	CRS                   string                `xml:"CRS" yaml:"crs" validate:"required"`
	BoundingBox           ows.BoundingBox       `xml:"BoundingBox" yaml:"boundingbox" validate:"required"`
	// We skip the OutPut struct, because these are not required parameters
	Size Size `xml:"Size" yaml:"size" validate:"required"`

	QueryLayers  []string `xml:"QueryLayers" yaml:"querylayers" validate:"required"`
	InfoFormat   *string  `xml:"InfoFormat" yaml:"infoformat"`
	FeatureCount *int     `xml:"FeatureCount" yaml:"featurecount"`
	I            int      `xml:"I" yaml:"i" validate:"required"`
	J            int      `xml:"J" yaml:"j" validate:"required"`
	Exceptions   *string  `xml:"Exceptions" yaml:"exceptions"`
}
