package wcs201

import (
	"encoding/xml"
	"net/url"
	"regexp"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/utils"
	"github.com/pdok/ogc-specifications/pkg/wsc200"
)

// WCS 2.0.1 Tokens
const (
	SERVICE = `SERVICE`
	REQUEST = `REQUEST`
	VERSION = `VERSION`
)

// Type returns GetCapabilities
func (gc *GetCapabilitiesRequest) Type() string {
	return getcapabilities
}

// Validate validates the GetCapabilities struct
func (gc *GetCapabilitiesRequest) Validate(c Capabilities) []wsc200.Exception {
	return nil
}

// ParseXML builds a GetCapabilities object based on a XML document
func (gc *GetCapabilitiesRequest) ParseXML(body []byte) []wsc200.Exception {
	var xmlattributes utils.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return wsc200.MissingParameterValue().ToExceptions()
	}
	if err := xml.Unmarshal(body, &gc); err != nil {
		return wsc200.MissingParameterValue(REQUEST).ToExceptions()
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

// QueryParameters builds a GetCapabilities object based on the available query parameters
func (gc *GetCapabilitiesRequest) QueryParameters(query url.Values) []wsc200.Exception {
	for k, v := range query {
		switch strings.ToUpper(k) {
		case REQUEST:
			if strings.EqualFold(v[0], getcapabilities) {
				gc.XMLName.Local = getcapabilities
			}
		case SERVICE:
			gc.Service = strings.ToUpper(v[0])
		case VERSION:
			gc.Version = strings.ToUpper(v[0])
		}
	}
	return nil
}

// ToQueryParameters builds a new query string that will be proxied
func (gc *GetCapabilitiesRequest) ToQueryParameters() url.Values {
	querystring := make(map[string][]string)
	querystring[REQUEST] = []string{gc.XMLName.Local}
	querystring[SERVICE] = []string{gc.Service}
	querystring[VERSION] = []string{gc.Version}

	return querystring
}

// ToXML builds a 'new' XML document 'based' on the 'original' XML document
func (gc GetCapabilitiesRequest) ToXML() []byte {
	si, _ := xml.MarshalIndent(gc, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}

// GetCapabilitiesRequest struct with the needed parameters/attributes needed for making a GetCapabilities request
type GetCapabilitiesRequest struct {
	XMLName xml.Name           `xml:"GetCapabilities" yaml:"getcapabilities"`
	Service string             `xml:"service,attr" yaml:"service"`
	Version string             `xml:"version,attr" yaml:"version"`
	Attr    utils.XMLAttribute `xml:",attr"`
}
