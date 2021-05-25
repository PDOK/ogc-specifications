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

// GetFeatureInfo struct with the needed parameters/attributes needed for making a GetFeatureInfo request
type GetFeatureInfoRequest struct {
	XMLName xml.Name `xml:"GetFeatureInfo" yaml:"getfeatureinfo"`
	BaseRequest

	// <map_request_copy>
	// These are the 'minimum' required GetMap parameters
	// needed in a GetFeatureInfo request
	StyledLayerDescriptor StyledLayerDescriptor `xml:"StyledLayerDescriptor" yaml:"styledlayerdescriptor"` //TODO layers is need styles is not!
	CRS                   string                `xml:"CRS" yaml:"crs"`
	BoundingBox           BoundingBox           `xml:"BoundingBox" yaml:"boundingbox"`
	// We skip the Output struct, because these are not required parameters
	Size   Size   `xml:"Size" yaml:"size"`
	Format string `xml:"Format,omitempty" yaml:"format,omitempty"`

	QueryLayers []string `xml:"QueryLayers" yaml:"querylayers"`
	I           int      `xml:"I" yaml:"i"`
	J           int      `xml:"J" yaml:"j"`
	InfoFormat  string   `xml:"InfoFormat" yaml:"infoformat" default:"text/plain"` // default text/plain

	// Optional Keys
	FeatureCount *int    `xml:"FeatureCount,omitempty" yaml:"featurecount,omitempty" default:"1"` // default 1
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
	var xmlattributes utils.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return Exceptions{MissingParameterValue()}
	}
	if err := xml.Unmarshal(body, &gfi); err != nil {
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

	gfi.Attr = utils.StripDuplicateAttr(n)
	return nil
}

// ParseOperationRequestKVP process the simple struct to a complex struct
func (gfi *GetFeatureInfoRequest) parseKVP(gfikvp getFeatureInfoKVPRequest) Exceptions {

	var exceptions Exceptions

	gfi.XMLName.Local = getfeatureinfo
	gfi.BaseRequest.parseKVP(gfikvp.baseRequestKVP)

	sld, ex := gfikvp.buildStyledLayerDescriptor()
	if ex != nil {
		exceptions = append(exceptions, ex...)
	}
	gfi.StyledLayerDescriptor = sld

	gfi.CRS = gfikvp.crs

	var bbox BoundingBox
	if ex := bbox.parseString(gfikvp.bbox); ex != nil {
		exceptions = append(exceptions, ex...)
	}
	gfi.BoundingBox = bbox

	gfi.CRS = gfikvp.crs

	w, err := strconv.Atoi(gfikvp.width)
	if err != nil {
		exceptions = append(exceptions, Exceptions{MissingParameterValue(WIDTH, gfikvp.width)}...)
	}
	gfi.Size.Width = w

	h, err := strconv.Atoi(gfikvp.height)
	if err != nil {
		exceptions = append(exceptions, Exceptions{MissingParameterValue(HEIGHT, gfikvp.height)}...)
	}
	gfi.Size.Height = h

	gfi.QueryLayers = strings.Split(gfikvp.querylayers, ",")

	if exps := gfi.parseIJ(gfikvp.i, gfikvp.j); exps != nil {
		exceptions = append(exceptions, exps...)
	}

	gfi.InfoFormat = gfikvp.infoformat

	// Optional keys
	if gfikvp.featurecount != nil {
		fc, err := strconv.Atoi(*gfikvp.featurecount)
		if err != nil {
			exceptions = append(exceptions, NoApplicableCode("Unknown FEATURE_COUNT value"))
		}

		gfi.FeatureCount = &fc
	}

	if gfikvp.exceptions != nil {
		gfi.Exceptions = gfikvp.exceptions
	}

	if len(exceptions) > 0 {
		return exceptions
	} else {
		return nil
	}
}

// ParseKVP builds a GetFeatureInfo object based on the available query parameters
func (gfi *GetFeatureInfoRequest) ParseQueryParameters(query url.Values) Exceptions {
	if len(query) == 0 {
		// When there are no query value we know that at least
		// the manadorty VERSION and REQUEST parameter is missing.
		return Exceptions{MissingParameterValue(VERSION), MissingParameterValue(REQUEST)}
	}

	gfikvp := getFeatureInfoKVPRequest{}
	if exceptions := gfikvp.parseQueryParameters(query); len(exceptions) != 0 {
		return exceptions
	}

	if exceptions := gfi.parseKVP(gfikvp); len(exceptions) != 0 {
		return exceptions
	}

	return nil
}

// ToQueryParameters  builds a new query string that will be proxied
func (gfi GetFeatureInfoRequest) ToQueryParameters() url.Values {
	gfikvp := getFeatureInfoKVPRequest{}
	gfikvp.parseGetFeatureInfoRequest(gfi)

	q := gfikvp.toQueryParameters()
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

func (gfi *GetFeatureInfoRequest) parseIJ(I, J string) Exceptions {
	i, err := strconv.Atoi(I)
	if err != nil {
		return InvalidPoint(I, J).ToExceptions()
	}
	gfi.I = i

	j, err := strconv.Atoi(J)
	if err != nil {
		return InvalidPoint(I, J).ToExceptions()
	}
	gfi.J = j

	return nil
}
