package wms130

import (
	"encoding/xml"
	"net/url"
	"regexp"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/utils"
)

// GetCapabilitiesRequest struct with the needed parameters/attributes needed for making a GetCapabilities request
type GetCapabilitiesRequest struct {
	XMLName xml.Name `xml:"GetCapabilities" yaml:"getcapabilities"`
	BaseRequest
}

// Validate returns GetCapabilities
func (g *GetCapabilitiesRequest) Validate(c Capabilities) Exceptions {
	var exceptions Exceptions
	return exceptions
}

// ParseXML builds a GetCapabilities object based on a XML document
func (g *GetCapabilitiesRequest) ParseXML(body []byte) Exceptions {
	var xmlattributes utils.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return Exceptions{MissingParameterValue()}
	}
	if err := xml.Unmarshal(body, &g); err != nil {
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

	g.Attr = utils.StripDuplicateAttr(n)
	return nil
}

// ParseQueryParameters builds a GetCapabilities object based on the available query parameters
func (g *GetCapabilitiesRequest) ParseQueryParameters(query url.Values) Exceptions {
	if len(query) == 0 {
		// When there are no query value we know that at least
		// the manadorty SERVICE and REQUEST parameter is missing.
		return Exceptions{MissingParameterValue(SERVICE), MissingParameterValue(REQUEST)}
	}

	gpv := getCapabilitiesParameterValueRequest{}
	if exception := gpv.parseQueryParameters(query); exception != nil {
		return exception
	}

	if exception := g.parseGetCapabilitiesParameterValueRequest(gpv); exception != nil {
		return exception
	}

	return nil
}

// parseGetCapabilitiesParameterValueRequest process the simple struct to a complex struct
func (g *GetCapabilitiesRequest) parseGetCapabilitiesParameterValueRequest(gpv getCapabilitiesParameterValueRequest) Exceptions {

	g.XMLName.Local = gpv.request
	g.BaseRequest.parseBaseParameterValueRequest(gpv.baseParameterValueRequest)
	return nil
}

// ToQueryParameters builds a new query string that will be proxied
func (g GetCapabilitiesRequest) ToQueryParameters() url.Values {
	gpv := getCapabilitiesParameterValueRequest{}
	gpv.parseGetCapabilitiesRequest(g)

	q := gpv.toQueryParameters()
	return q
}

// ToXML builds a 'new' XML document 'based' on the 'original' XML document
func (g GetCapabilitiesRequest) ToXML() []byte {
	si, _ := xml.MarshalIndent(&g, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}
