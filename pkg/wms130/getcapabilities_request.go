package wms130

import (
	"encoding/xml"
	"net/url"
	"regexp"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/utils"
)

// GetCapabilities struct with the needed parameters/attributes needed for making a GetCapabilities request
type GetCapabilitiesRequest struct {
	XMLName xml.Name `xml:"GetCapabilities" yaml:"getcapabilities"`
	BaseRequest
}

// Validate returns GetCapabilities
func (gc *GetCapabilitiesRequest) Validate(c Capabilities) Exceptions {
	var exceptions Exceptions
	return exceptions
}

// ParseXML builds a GetCapabilities object based on a XML document
func (gc *GetCapabilitiesRequest) ParseXML(body []byte) Exceptions {
	var xmlattributes utils.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return Exceptions{MissingParameterValue()}
	}
	if err := xml.Unmarshal(body, &gc); err != nil {
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

	gc.Attr = utils.StripDuplicateAttr(n)
	return nil
}

// ParseQueryParameters builds a GetCapabilities object based on the available query parameters
func (gc *GetCapabilitiesRequest) ParseQueryParameters(query url.Values) Exceptions {
	if len(query) == 0 {
		// When there are no query value we know that at least
		// the manadorty SERVICE and REQUEST parameter is missing.
		return Exceptions{MissingParameterValue(SERVICE), MissingParameterValue(REQUEST)}
	}

	gckvp := getCapabilitiesKVPRequest{}
	if exception := gckvp.parseQueryParameters(query); exception != nil {
		return exception
	}

	if exception := gc.parseKVPRequest(gckvp); exception != nil {
		return exception
	}

	return nil
}

// ParseOperationRequestKVP process the simple struct to a complex struct
func (gc *GetCapabilitiesRequest) parseKVPRequest(gckvp getCapabilitiesKVPRequest) Exceptions {

	gc.XMLName.Local = gckvp.request
	gc.BaseRequest.parseKVPRequest(gckvp.baseRequestKVP)
	return nil
}

// ToQueryParameters builds a new query string that will be proxied
func (gc GetCapabilitiesRequest) ToQueryParameters() url.Values {
	gckvp := getCapabilitiesKVPRequest{}
	gckvp.parseGetCapabilitiesRequest(gc)

	q := gckvp.toQueryParameters()
	return q
}

// ToXML builds a 'new' XML document 'based' on the 'original' XML document
func (gc *GetCapabilitiesRequest) ToXML() []byte {
	si, _ := xml.MarshalIndent(gc, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}
