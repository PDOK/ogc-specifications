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

// Type and Version as constant
const (
	Service string = `WCS`
	Version string = `2.0.1`
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

// Validate returns GetCapabilities
func (gc *GetCapabilities) Validate() ows.Exception {
	return nil
}

// ParseXML builds a GetCapabilities object based on a XML document
func (gc *GetCapabilities) ParseXML(body []byte) ows.Exception {
	var xmlattributes ows.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return ows.MissingParameterValue()
	}
	if err := xml.Unmarshal(body, &gc); err != nil {
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

	gc.Attr = ows.StripDuplicateAttr(n)
	return nil
}

// ParseKVP builds a GetCapabilities object based on the available query parameters
func (gc *GetCapabilities) ParseKVP(query url.Values) ows.Exception {
	for k, v := range query {
		switch strings.ToUpper(k) {
		case REQUEST:
			if strings.ToUpper(v[0]) == strings.ToUpper(getcapabilities) {
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
	XMLName xml.Name         `xml:"GetCapabilities" yaml:"getcapabilities" validate:"required"`
	Service string           `xml:"service,attr" yaml:"service" validate:"required,oneof=WCS wcs"`
	Version string           `xml:"version,attr" yaml:"version" validate:"eq=2.0.1"`
	Attr    ows.XMLAttribute `xml:",attr"`
}