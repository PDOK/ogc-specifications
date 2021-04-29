package request

import (
	"encoding/xml"
	"net/url"
	"regexp"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/common"
	"github.com/pdok/ogc-specifications/pkg/wfs200/capabilities"
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

// Validate returns GetCapabilities
func (gc *GetCapabilities) Validate(c capabilities.Capabilities) common.Exceptions {
	return nil
}

// ParseXML builds a GetCapabilities object based on a XML document
func (gc *GetCapabilities) ParseXML(doc []byte) common.Exception {
	var xmlattributes common.XMLAttribute
	if err := xml.Unmarshal(doc, &xmlattributes); err != nil {
		return common.NoApplicableCode("Could not process XML, is it XML?")
	}
	if err := xml.Unmarshal(doc, &gc); err != nil {
		return common.OperationNotSupported(err.Error()) //TODO Should be OperationParsingFailed
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

// ParseKVP builds a GetCapabilities object based on the available query parameters
func (gc *GetCapabilities) ParseKVP(query url.Values) common.Exceptions {
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
	XMLName xml.Name            `xml:"GetCapabilities" yaml:"getcapabilities"`
	Service string              `xml:"service,attr" yaml:"service"`
	Version string              `xml:"version,attr" yaml:"version"`
	Attr    common.XMLAttribute `xml:",attr"`
}
