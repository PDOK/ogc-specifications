package wcs201

import (
	"encoding/xml"
	"net/url"
	"regexp"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/common"
	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

// WCS 2.0.1 Tokens
const (
	SERVICE = `SERVICE`
	REQUEST = `REQUEST`
	VERSION = `VERSION`
)

// Type returns GetCapabilities
func (gc *GetCapabilities) Type() string {
	return getcapabilities
}

// Validate validates the GetCapabilities struct
func (gc *GetCapabilities) Validate(c Capabilities) common.Exceptions {
	return nil
}

// ParseXML builds a GetCapabilities object based on a XML document
func (gc *GetCapabilities) ParseXML(body []byte) common.Exceptions {
	var xmlattributes common.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return common.Exceptions{wsc110.MissingParameterValue()}
	}
	if err := xml.Unmarshal(body, &gc); err != nil {
		return common.Exceptions{wsc110.MissingParameterValue("REQUEST")}
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

	gc.Attr = common.StripDuplicateAttr(n)
	return nil
}

// QueryParameters builds a GetCapabilities object based on the available query parameters
func (gc *GetCapabilities) QueryParameters(query url.Values) common.Exceptions {
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

// BuildKVP builds a new query string that will be proxied
func (gc *GetCapabilities) BuildKVP() url.Values {
	querystring := make(map[string][]string)
	querystring[REQUEST] = []string{gc.XMLName.Local}
	querystring[SERVICE] = []string{gc.Service}
	querystring[VERSION] = []string{gc.Version}

	return querystring
}

// BuildXML builds a 'new' XML document 'based' on the 'original' XML document
func (gc *GetCapabilities) BuildXML() []byte {
	si, _ := xml.MarshalIndent(gc, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}

// GetCapabilities struct with the needed parameters/attributes needed for making a GetCapabilities request
type GetCapabilities struct {
	XMLName xml.Name            `xml:"GetCapabilities" yaml:"getcapabilities"`
	Service string              `xml:"service,attr" yaml:"service"`
	Version string              `xml:"version,attr" yaml:"version"`
	Attr    common.XMLAttribute `xml:",attr"`
}
