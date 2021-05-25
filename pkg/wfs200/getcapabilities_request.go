package wfs200

import (
	"encoding/xml"
	"net/url"
	"regexp"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/utils"
	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

// Contains the GetCapabilities struct and specific functions for building a GetCapabilities request

// Type returns GetCapabilities
func (gc *GetCapabilitiesRequest) Type() string {
	return getcapabilities
}

// Validate returns GetCapabilities
func (gc *GetCapabilitiesRequest) Validate(c wsc110.Capabilities) []wsc110.Exception {
	var exceptions []wsc110.Exception
	return exceptions
}

// ParseXML builds a GetCapabilities object based on a XML document
func (gc *GetCapabilitiesRequest) ParseXML(doc []byte) []wsc110.Exception {
	var xmlattributes utils.XMLAttribute
	if err := xml.Unmarshal(doc, &xmlattributes); err != nil {
		return []wsc110.Exception{wsc110.NoApplicableCode("Could not process XML, is it XML?")}
	}
	if err := xml.Unmarshal(doc, &gc); err != nil {
		return []wsc110.Exception{wsc110.OperationNotSupported(err.Error())} //TODO Should be OperationParsingFailed
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
func (gc *GetCapabilitiesRequest) ParseQueryParameters(query url.Values) []wsc110.Exception {
	if len(query) == 0 {
		// When there are no query value we know that at least
		// the manadorty SERVICE and REQUEST parameter is missing.
		exceptions := wsc110.MissingParameterValue(SERVICE).ToExceptions()
		exceptions = append(exceptions, wsc110.MissingParameterValue(REQUEST))
		return exceptions
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
func (gc *GetCapabilitiesRequest) parseKVPRequest(gckvp getCapabilitiesKVPRequest) []wsc110.Exception {
	gc.XMLName.Local = gckvp.request
	gc.Service = gckvp.service
	gc.Version = gckvp.version

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

// GetCapabilities struct with the needed parameters/attributes needed for making a GetCapabilities request
type GetCapabilitiesRequest struct {
	XMLName xml.Name           `xml:"GetCapabilities" yaml:"getcapabilities"`
	Service string             `xml:"service,attr" yaml:"service"`
	Version string             `xml:"version,attr" yaml:"version"`
	Attr    utils.XMLAttribute `xml:",attr"`
}
