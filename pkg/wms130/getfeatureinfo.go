package wms130

import (
	"encoding/xml"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/common"
)

// GetFeatureInfo
const (
	getfeatureinfo = `GetFeatureInfo`

	// Mandatory
	QUERYLAYERS = `QUERY_LAYERS`
	I           = `I`
	J           = `J`

	// Optional GetFeatureInfo Keys
	INFOFORMAT   = `INFO_FORMAT`
	FEATURECOUNT = `FEATURE_COUNT`
)

// GetFeatureInfo struct with the needed parameters/attributes needed for making a GetFeatureInfo request
type GetFeatureInfo struct {
	XMLName xml.Name `xml:"GetFeatureInfo" yaml:"getfeatureinfo"`
	BaseRequest

	// <map_request_copy>
	// These are the 'minimum' required GetMap parameters
	// needed in a GetFeatureInfo request
	StyledLayerDescriptor StyledLayerDescriptor `xml:"StyledLayerDescriptor" yaml:"styledlayerdescriptor"` //TODO layers is need styles is not!
	CRS                   string                `xml:"CRS" yaml:"crs"`
	BoundingBox           common.BoundingBox    `xml:"BoundingBox" yaml:"boundingbox"`
	// We skip the Output struct, because these are not required parameters
	Size   Size   `xml:"Size" yaml:"size"`
	Format string `xml:"Format,omitempty" yaml:"format,omitempty"`

	QueryLayers  []string `xml:"QueryLayers" yaml:"querylayers"`
	I            int      `xml:"I" yaml:"i"`
	J            int      `xml:"J" yaml:"j"`
	InfoFormat   string   `xml:"InfoFormat" yaml:"infoformat" default:"text/plain"`                // default text/plain
	FeatureCount int      `xml:"FeatureCount,omitempty" yaml:"featurecount,omitempty" default:"1"` // default 1
	Exceptions   *string  `xml:"Exceptions" yaml:"exceptions"`
}

// Type returns GetFeatureInfo
func (gfi *GetFeatureInfo) Type() string {
	return getfeatureinfo
}

// Validate returns GetFeatureInfo
func (gfi *GetFeatureInfo) Validate(c common.Capabilities) common.Exceptions {
	var exceptions common.Exceptions

	wmsCapabilities := c.(Capabilities)

	exceptions = append(exceptions, gfi.StyledLayerDescriptor.Validate(wmsCapabilities)...)
	// exceptions = append(exceptions, gfi.Output.Validate(wmsCapabilities)...)

	return exceptions
}

// ParseXML builds a GetFeatureInfo object based on a XML document
// Note: the XML GetFeatureInfo body that is consumed is a interpretation.
// So we use the GetMap, that is a large part of this request, as a base
// with the additional GetFeatureInfo parameters.
func (gfi *GetFeatureInfo) ParseXML(body []byte) common.Exceptions {
	var xmlattributes common.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return common.Exceptions{common.MissingParameterValue()}
	}
	if err := xml.Unmarshal(body, &gfi); err != nil {
		return common.Exceptions{common.MissingParameterValue("REQUEST")}
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

	gfi.Attr = common.StripDuplicateAttr(n)
	return nil
}

// ParseOperationRequestKVP process the simple struct to a complex struct
func (gfi *GetFeatureInfo) ParseOperationRequestKVP(orkvp common.OperationRequestKVP) common.Exceptions {
	gfikvp := orkvp.(*GetFeatureInfoKVP)

	gfi.XMLName.Local = getfeatureinfo
	gfi.BaseRequest.Build(gfikvp.Service, gfikvp.Version)

	sld, ex := gfikvp.buildStyledLayerDescriptor()
	if ex != nil {
		return common.Exceptions{ex}
	}
	gfi.StyledLayerDescriptor = sld

	gfi.CRS = gfikvp.CRS

	var bbox common.BoundingBox
	if err := bbox.ParseString(gfikvp.Bbox); err != nil {
		return common.Exceptions{err}
	}
	gfi.BoundingBox = bbox

	gfi.CRS = gfikvp.CRS

	w, err := strconv.Atoi(gfikvp.Width)
	if err != nil {
		return common.Exceptions{common.MissingParameterValue(WIDTH, gfikvp.Width)}
	}
	gfi.Size.Width = w

	h, err := strconv.Atoi(gfikvp.Height)
	if err != nil {
		return common.Exceptions{common.MissingParameterValue(HEIGHT, gfikvp.Height)}
	}
	gfi.Size.Height = h

	gfi.QueryLayers = strings.Split(gfikvp.QueryLayers, ",")

	i, err := strconv.Atoi(gfikvp.I)
	if err != nil {
		return common.Exceptions{InvalidPoint(gfikvp.I, gfikvp.J)}
	}
	gfi.I = i

	j, err := strconv.Atoi(gfikvp.J)
	if err != nil {
		return common.Exceptions{InvalidPoint(gfikvp.I, gfikvp.J)}
	}
	gfi.J = j

	fc, err := strconv.Atoi(*gfikvp.FeatureCount)
	if err != nil {
		// TODO: ignore or a exception
		return common.Exceptions{common.NoApplicableCode("Unknown FeatureCount value")}
	}

	gfi.FeatureCount = fc
	gfi.InfoFormat = gfikvp.InfoFormat
	gfi.Exceptions = gfikvp.Exceptions

	return nil
}

// ParseKVP builds a GetFeatureInfo object based on the available query parameters
func (gfi *GetFeatureInfo) ParseKVP(query url.Values) common.Exceptions {
	if len(query) == 0 {
		// When there are no query value we know that at least
		// the manadorty VERSION and REQUEST parameter is missing.
		return common.Exceptions{common.MissingParameterValue(VERSION), common.MissingParameterValue(REQUEST)}
	}

	gfikvp := GetFeatureInfoKVP{}
	if err := gfikvp.ParseKVP(query); err != nil {
		return err
	}

	if err := gfi.ParseOperationRequestKVP(&gfikvp); err != nil {
		return err
	}

	return nil
}

// BuildKVP builds a new query string that will be proxied
func (gfi *GetFeatureInfo) BuildKVP() url.Values {
	gfikvp := GetFeatureInfoKVP{}
	gfikvp.ParseOperationRequest(gfi)

	kvp := gfikvp.BuildKVP()
	return kvp
}

// BuildXML builds a 'new' XML document 'based' on the 'original' XML document
// Note: this GetFeatureInfo XML body is a interpretation and there isn't a
// good/real OGC example request. So for now we use the GetMap, that is a large part
// of this request, as a base with the additional GetFeatureInfo parameters.
func (gfi *GetFeatureInfo) BuildXML() []byte {
	si, _ := xml.MarshalIndent(gfi, "", " ")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}
