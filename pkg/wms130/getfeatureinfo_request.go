package wms130

import (
	"encoding/xml"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/utils"
)

// GetFeatureInfo
const (
	// Mandatory
	QUERYLAYERS = `QUERY_LAYERS`
	I           = `I`
	J           = `J`

	// Optional GetFeatureInfo Keys
	INFOFORMAT   = `INFO_FORMAT`
	FEATURECOUNT = `FEATURE_COUNT`
)

// GetFeatureInfoRequest struct with the needed parameters/attributes needed for making a GetFeatureInfo request
type GetFeatureInfoRequest struct {
	XMLName xml.Name `xml:"GetFeatureInfo" yaml:"getFeatureInfo"`
	BaseRequest

	// <map_request_copy>
	// These are the 'minimum' required GetMap parameters
	// needed in a GetFeatureInfo request
	StyledLayerDescriptor StyledLayerDescriptor `xml:"StyledLayerDescriptor" yaml:"styledLayerDescriptor"` // TODO layers is need styles is not!
	CRS                   string                `xml:"CRS" yaml:"crs"`
	BoundingBox           BoundingBox           `xml:"BoundingBox" yaml:"boundingBox"`
	// We skip the Output struct, because these are not required parameters
	Size   Size   `xml:"Size" yaml:"size"`
	Format string `xml:"Format,omitempty" yaml:"format,omitempty"`

	QueryLayers []string `xml:"QueryLayers" yaml:"queryLayers"`
	I           int      `xml:"I" yaml:"i"`
	J           int      `xml:"J" yaml:"j"`
	InfoFormat  string   `xml:"InfoFormat" yaml:"infoFormat" default:"text/plain"` // default text/plain

	// Optional Keys
	FeatureCount *int    `xml:"FeatureCount,omitempty" yaml:"featureCount,omitempty" default:"1"` // default 1
	Exceptions   *string `xml:"Exceptions" yaml:"exceptions"`
}

// Validate returns GetFeatureInfo
func (gfi *GetFeatureInfoRequest) Validate(c Capabilities) Exceptions {
	var exceptions Exceptions

	exceptions = append(exceptions, gfi.StyledLayerDescriptor.Validate(c)...)
	// exceptions = append(exceptions, gfi.Output.Validate(wmsCapabilities)...)

	return exceptions
}

// ParseXML builds a GetFeatureInfo object based on a XML document
// Note: the XML GetFeatureInfo body that is consumed is a interpretation.
// So we use the GetMap, that is a large part of this request, as a base
// with the additional GetFeatureInfo parameters.
func (gfi *GetFeatureInfoRequest) ParseXML(body []byte) Exceptions {
	var xmlAttributes utils.XMLAttribute
	if err := xml.Unmarshal(body, &xmlAttributes); err != nil {
		return Exceptions{MissingParameterValue()}
	}
	if err := xml.Unmarshal(body, &gfi); err != nil {
		return Exceptions{MissingParameterValue("REQUEST")}
	}
	var n []xml.Attr
	for _, a := range xmlAttributes {
		switch strings.ToUpper(a.Name.Local) {
		case VERSION:
		case SERVICE:
		default:
			n = append(n, a)
		}
	}

	gfi.Attr = utils.StripDuplicateAttr(n)
	return nil
}

// parseGetFeatureInfoRequestParameterValue process the simple struct to a complex struct
func (gfi *GetFeatureInfoRequest) parseGetFeatureInfoRequestParameterValue(ipv getFeatureInfoRequestParameterValue) Exceptions {

	var exceptions Exceptions

	gfi.XMLName.Local = getfeatureinfo
	gfi.BaseRequest.parseBaseParameterValueRequest(ipv.baseParameterValueRequest)

	sld, ex := ipv.buildStyledLayerDescriptor()
	if ex != nil {
		exceptions = append(exceptions, ex...)
	}
	gfi.StyledLayerDescriptor = sld

	gfi.CRS = ipv.crs

	var bbox BoundingBox
	if ex := bbox.parseString(ipv.bbox); ex != nil {
		exceptions = append(exceptions, ex...)
	}
	gfi.BoundingBox = bbox

	gfi.CRS = ipv.crs

	w, err := strconv.Atoi(ipv.width)
	if err != nil {
		exceptions = append(exceptions, Exceptions{MissingParameterValue(WIDTH, ipv.width)}...)
	}
	gfi.Size.Width = w

	h, err := strconv.Atoi(ipv.height)
	if err != nil {
		exceptions = append(exceptions, Exceptions{MissingParameterValue(HEIGHT, ipv.height)}...)
	}
	gfi.Size.Height = h

	gfi.QueryLayers = strings.Split(ipv.querylayers, ",")

	if exps := gfi.parseIJ(ipv.i, ipv.j); exps != nil {
		exceptions = append(exceptions, exps...)
	}

	gfi.InfoFormat = ipv.infoformat

	// Optional keys
	if ipv.featurecount != nil {
		fc, err := strconv.Atoi(*ipv.featurecount)
		if err != nil {
			exceptions = append(exceptions, NoApplicableCode("Unknown FEATURE_COUNT value"))
		}

		gfi.FeatureCount = &fc
	}

	if ipv.exceptions != nil {
		gfi.Exceptions = ipv.exceptions
	}

	if len(exceptions) > 0 {
		return exceptions
	}
	return nil
}

// ParseQueryParameters builds a GetFeatureInfo object based on the available query parameters
func (gfi *GetFeatureInfoRequest) ParseQueryParameters(query url.Values) Exceptions {
	if len(query) == 0 {
		// When there are no query value we know that at least
		// the manadorty VERSION and REQUEST parameter is missing.
		return Exceptions{MissingParameterValue(VERSION), MissingParameterValue(REQUEST)}
	}

	ipv := getFeatureInfoRequestParameterValue{}
	if exceptions := ipv.parseQueryParameters(query); len(exceptions) != 0 {
		return exceptions
	}

	if exceptions := gfi.parseGetFeatureInfoRequestParameterValue(ipv); len(exceptions) != 0 {
		return exceptions
	}

	return nil
}

// ToQueryParameters  builds a new query string that will be proxied
func (gfi *GetFeatureInfoRequest) ToQueryParameters() url.Values {
	ipv := getFeatureInfoRequestParameterValue{}
	ipv.parseGetFeatureInfoRequest(*gfi)

	q := ipv.toQueryParameters()
	return q
}

// ToXML builds a 'new' XML document 'based' on the 'original' XML document
// Note: this GetFeatureInfo XML body is a interpretation and there isn't a
// good/real OGC example request. So for now we use the GetMap, that is a large part
// of this request, as a base with the additional GetFeatureInfo parameters.
func (gfi *GetFeatureInfoRequest) ToXML() []byte {
	si, _ := xml.MarshalIndent(gfi, "", " ")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}

func (gfi *GetFeatureInfoRequest) parseIJ(i string, j string) Exceptions {
	ii, err := strconv.Atoi(i)
	if err != nil {
		return InvalidPoint(i, j).ToExceptions()
	}
	gfi.I = ii

	jj, err := strconv.Atoi(j)
	if err != nil {
		return InvalidPoint(i, j).ToExceptions()
	}
	gfi.J = jj

	return nil
}
