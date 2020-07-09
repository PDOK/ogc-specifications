package wfs200

import (
	"encoding/xml"
	"net/url"
	"regexp"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/ows"
)

// Type and Version as constant
const (
	getcapabilities = `GetCapabilities`
)

// Contains the GetCapabilities struct and specific functions for building a GetCapabilities request

// Type returns GetCapabilities
func (gc *GetCapabilities) Type() string {
	return getcapabilities
}

// ParseBody builds a GetCapabilities object based on the given body
func (gc *GetCapabilities) ParseBody(body []byte) ows.Exception {
	var xmlattributes ows.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return ows.NoApplicableCode("Could not process XML, is it XML?")
	}
	if err := xml.Unmarshal(body, &gc); err != nil {
		return ows.OperationNotSupported(err.Error()) //TODO Should be OperationParsingFailed
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

// ParseQuery builds a GetCapabilities object based on the available query parameters
func (gc *GetCapabilities) ParseQuery(query url.Values) ows.Exception {
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

// BuildQuery builds a new query string that will be proxied
func (gc *GetCapabilities) BuildQuery() url.Values {
	querystring := make(map[string][]string)
	querystring[REQUEST] = []string{gc.XMLName.Local}
	querystring[SERVICE] = []string{gc.Service}
	querystring[VERSION] = []string{gc.Version}

	return querystring
}

// BuildBody builds a 'new' XML document 'based' on the 'original' XML document
func (gc *GetCapabilities) BuildBody() []byte {
	si, _ := xml.MarshalIndent(gc, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}

// GetCapabilities struct with the needed parameters/attributes needed for making a GetCapabilities request
type GetCapabilities struct {
	XMLName xml.Name         `xml:"GetCapabilities" yaml:"getcapabilities"`
	Service string           `xml:"service,attr" yaml:"service" validate:"required,oneof=WFS wfs"`
	Version string           `xml:"version,attr" yaml:"version" validate:"eq=2.0.0"`
	Attr    ows.XMLAttribute `xml:",attr"`
}
