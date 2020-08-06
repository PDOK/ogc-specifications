package request

import (
	"encoding/xml"
	"net/url"
	"regexp"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/ows"
)

//
const (
	getcapabilities = `GetCapabilities`
)

// Type returns GetCapabilities
func (gc *GetCapabilities) Type() string {
	return getcapabilities
}

// Validate returns GetCapabilities
func (gc *GetCapabilities) Validate(c ows.Capability) ows.Exceptions {
	var exceptions ows.Exceptions
	return exceptions
}

// ParseXML builds a GetCapabilities object based on a XML document
func (gc *GetCapabilities) ParseXML(body []byte) ows.Exceptions {
	var xmlattributes ows.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return ows.Exceptions{ows.MissingParameterValue()}
	}
	if err := xml.Unmarshal(body, &gc); err != nil {
		return ows.Exceptions{ows.MissingParameterValue("REQUEST")}
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

	gc.Attr = ows.StripDuplicateAttr(n)
	return nil
}

// ParseKVP builds a GetCapabilities object based on the available query parameters
func (gc *GetCapabilities) ParseKVP(query url.Values) ows.Exceptions {
	if len(query) == 0 {
		// When there are no query value we know that at least
		// the manadorty SERVICE and REQUEST parameter is missing.
		return ows.Exceptions{ows.MissingParameterValue(SERVICE), ows.MissingParameterValue(REQUEST)}
	}

	gckvp := GetCapabilitiesKVP{}
	if err := gckvp.ParseKVP(query); err != nil {
		return err
	}

	if err := gc.ParseOperationRequestKVP(&gckvp); err != nil {
		return err
	}

	return nil
}

// ParseOperationRequestKVP process the simple struct to a complex struct
func (gc *GetCapabilities) ParseOperationRequestKVP(orkvp ows.OperationRequestKVP) ows.Exceptions {
	gckvp := orkvp.(*GetCapabilitiesKVP)

	gc.XMLName.Local = gckvp.Request
	gc.BaseRequest.Build(gckvp.Service, gckvp.Version)
	return nil
}

// BuildKVP builds a new query string that will be proxied
func (gc *GetCapabilities) BuildKVP() url.Values {
	gckvp := GetCapabilitiesKVP{}
	gckvp.ParseOperationRequest(gc)

	kvp := gckvp.BuildKVP()
	return kvp
}

// BuildXML builds a 'new' XML document 'based' on the 'original' XML document
func (gc *GetCapabilities) BuildXML() []byte {
	si, _ := xml.MarshalIndent(gc, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}

// GetCapabilities struct with the needed parameters/attributes needed for making a GetCapabilities request
type GetCapabilities struct {
	XMLName xml.Name `xml:"GetCapabilities" yaml:"getcapabilities"`
	BaseRequest
}
