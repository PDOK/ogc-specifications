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
	StyledLayerDescriptor StyledLayerDescriptor `xml:"StyledLayerDescriptor" yaml:"styledLayerDescriptor"` //TODO layers is need styles is not!
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
func (i *GetFeatureInfoRequest) Validate(c Capabilities) Exceptions {
	var exceptions Exceptions

	exceptions = append(exceptions, i.StyledLayerDescriptor.Validate(c)...)
	// exceptions = append(exceptions, i.Output.Validate(wmsCapabilities)...)

	return exceptions
}

// ParseXML builds a GetFeatureInfo object based on a XML document
// Note: the XML GetFeatureInfo body that is consumed is a interpretation.
// So we use the GetMap, that is a large part of this request, as a base
// with the additional GetFeatureInfo parameters.
func (i *GetFeatureInfoRequest) ParseXML(body []byte) Exceptions {
	var xmlattributes utils.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return Exceptions{MissingParameterValue()}
	}
	if err := xml.Unmarshal(body, &i); err != nil {
		return Exceptions{MissingParameterValue("REQUEST")}
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

	i.Attr = utils.StripDuplicateAttr(n)
	return nil
}

// parsegetFeatureInfoRequestParameterValue process the simple struct to a complex struct
func (i *GetFeatureInfoRequest) parsegetFeatureInfoRequestParameterValue(ipv getFeatureInfoRequestParameterValue) Exceptions {

	var exceptions Exceptions

	i.XMLName.Local = getfeatureinfo
	i.BaseRequest.parseBaseParameterValueRequest(ipv.baseParameterValueRequest)

	sld, ex := ipv.buildStyledLayerDescriptor()
	if ex != nil {
		exceptions = append(exceptions, ex...)
	}
	i.StyledLayerDescriptor = sld

	i.CRS = ipv.crs

	var bbox BoundingBox
	if ex := bbox.parseString(ipv.bbox); ex != nil {
		exceptions = append(exceptions, ex...)
	}
	i.BoundingBox = bbox

	i.CRS = ipv.crs

	w, err := strconv.Atoi(ipv.width)
	if err != nil {
		exceptions = append(exceptions, Exceptions{MissingParameterValue(WIDTH, ipv.width)}...)
	}
	i.Size.Width = w

	h, err := strconv.Atoi(ipv.height)
	if err != nil {
		exceptions = append(exceptions, Exceptions{MissingParameterValue(HEIGHT, ipv.height)}...)
	}
	i.Size.Height = h

	i.QueryLayers = strings.Split(ipv.querylayers, ",")

	if exps := i.parseIJ(ipv.i, ipv.j); exps != nil {
		exceptions = append(exceptions, exps...)
	}

	i.InfoFormat = ipv.infoformat

	// Optional keys
	if ipv.featurecount != nil {
		fc, err := strconv.Atoi(*ipv.featurecount)
		if err != nil {
			exceptions = append(exceptions, NoApplicableCode("Unknown FEATURE_COUNT value"))
		}

		i.FeatureCount = &fc
	}

	if ipv.exceptions != nil {
		i.Exceptions = ipv.exceptions
	}

	if len(exceptions) > 0 {
		return exceptions
	}
	return nil
}

// ParseQueryParameters builds a GetFeatureInfo object based on the available query parameters
func (i *GetFeatureInfoRequest) ParseQueryParameters(query url.Values) Exceptions {
	if len(query) == 0 {
		// When there are no query value we know that at least
		// the manadorty VERSION and REQUEST parameter is missing.
		return Exceptions{MissingParameterValue(VERSION), MissingParameterValue(REQUEST)}
	}

	ipv := getFeatureInfoRequestParameterValue{}
	if exceptions := ipv.parseQueryParameters(query); len(exceptions) != 0 {
		return exceptions
	}

	if exceptions := i.parsegetFeatureInfoRequestParameterValue(ipv); len(exceptions) != 0 {
		return exceptions
	}

	return nil
}

// ToQueryParameters  builds a new query string that will be proxied
func (i GetFeatureInfoRequest) ToQueryParameters() url.Values {
	ipv := getFeatureInfoRequestParameterValue{}
	ipv.parseGetFeatureInfoRequest(i)

	q := ipv.toQueryParameters()
	return q
}

// ToXML builds a 'new' XML document 'based' on the 'original' XML document
// Note: this GetFeatureInfo XML body is a interpretation and there isn't a
// good/real OGC example request. So for now we use the GetMap, that is a large part
// of this request, as a base with the additional GetFeatureInfo parameters.
func (i GetFeatureInfoRequest) ToXML() []byte {
	si, _ := xml.MarshalIndent(i, "", " ")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}

func (i *GetFeatureInfoRequest) parseIJ(I, J string) Exceptions {
	ii, err := strconv.Atoi(I)
	if err != nil {
		return InvalidPoint(I, J).ToExceptions()
	}
	i.I = ii

	j, err := strconv.Atoi(J)
	if err != nil {
		return InvalidPoint(I, J).ToExceptions()
	}
	i.J = j

	return nil
}
