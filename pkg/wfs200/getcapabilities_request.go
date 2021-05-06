package wfs200

import (
	"encoding/xml"
	"net/url"
	"regexp"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/common"
	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

// Contains the GetCapabilities struct and specific functions for building a GetCapabilities request

// Type returns GetCapabilities
func (gc *GetCapabilitiesRequest) Type() string {
	return getcapabilities
}

// Validate returns GetCapabilities
func (gc *GetCapabilitiesRequest) Validate(c common.Capabilities) common.Exceptions {
	var exceptions common.Exceptions
	return exceptions
}

// ParseXML builds a GetCapabilities object based on a XML document
func (gc *GetCapabilitiesRequest) ParseXML(doc []byte) common.Exceptions {
	var xmlattributes common.XMLAttribute
	if err := xml.Unmarshal(doc, &xmlattributes); err != nil {
		return common.Exceptions{wsc110.NoApplicableCode("Could not process XML, is it XML?")}
	}
	if err := xml.Unmarshal(doc, &gc); err != nil {
		return common.Exceptions{wsc110.OperationNotSupported(err.Error())} //TODO Should be OperationParsingFailed
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
func (gc *GetCapabilitiesRequest) ParseKVP(query url.Values) common.Exceptions {
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

// ParseOperationRequestKVP process the simple struct to a complex struct
func (gc *GetCapabilitiesRequest) ParseOperationRequestKVP(orkvp common.OperationRequestKVP) common.Exceptions {
	gckvp := orkvp.(*GetCapabilitiesKVP)

	gc.XMLName.Local = gckvp.Request
	gc.Service = gckvp.Service
	gc.Version = gckvp.Version

	return nil
}

// BuildKVP builds a new query string that will be proxied
func (gc *GetCapabilitiesRequest) BuildKVP() url.Values {
	querystring := make(map[string][]string)
	querystring[REQUEST] = []string{gc.XMLName.Local}
	querystring[SERVICE] = []string{gc.Service}
	querystring[VERSION] = []string{gc.Version}

	return querystring
}

// BuildXML builds a 'new' XML document 'based' on the 'original' XML document
func (gc *GetCapabilitiesRequest) BuildXML() []byte {
	si, _ := xml.MarshalIndent(gc, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}

// GetCapabilities struct with the needed parameters/attributes needed for making a GetCapabilities request
type GetCapabilitiesRequest struct {
	XMLName xml.Name            `xml:"GetCapabilities" yaml:"getcapabilities"`
	Service string              `xml:"service,attr" yaml:"service"`
	Version string              `xml:"version,attr" yaml:"version"`
	Attr    common.XMLAttribute `xml:",attr"`
}
