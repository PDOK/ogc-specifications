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
func (g GetCapabilitiesRequest) Type() string {
	return getcapabilities
}

// Validate returns GetCapabilities
func (g GetCapabilitiesRequest) Validate(_ wsc110.Capabilities) []wsc110.Exception {
	var exceptions []wsc110.Exception
	return exceptions
}

// ParseXML builds a GetCapabilities object based on a XML document
func (g *GetCapabilitiesRequest) ParseXML(doc []byte) []wsc110.Exception {
	var xmlAttributes utils.XMLAttribute
	if err := xml.Unmarshal(doc, &xmlAttributes); err != nil {
		return []wsc110.Exception{wsc110.NoApplicableCode("Could not process XML, is it XML?")}
	}
	if err := xml.Unmarshal(doc, &g); err != nil {
		return []wsc110.Exception{wsc110.OperationNotSupported(err.Error())} // TODO Should be OperationParsingFailed
	}
	var n []xml.Attr
	for _, a := range xmlAttributes {
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
func (g *GetCapabilitiesRequest) ParseQueryParameters(query url.Values) []wsc110.Exception {
	if len(query) == 0 {
		// When there are no query value we know that at least
		// the mandatory SERVICE and REQUEST parameter is missing.
		exceptions := wsc110.MissingParameterValue(SERVICE).ToExceptions()
		exceptions = append(exceptions, wsc110.MissingParameterValue(REQUEST))
		return exceptions
	}

	gpv := getCapabilitiesRequestParameterValue{}
	if exception := gpv.parseQueryParameters(query); exception != nil {
		return exception
	}

	if exception := g.parseGetCapabilitiesRequestParameterValue(gpv); exception != nil {
		return exception
	}

	return nil
}

// parseGetCapabilitiesRequestParameterValue process the simple struct to a complex struct
func (g *GetCapabilitiesRequest) parseGetCapabilitiesRequestParameterValue(gpv getCapabilitiesRequestParameterValue) []wsc110.Exception {
	g.XMLName.Local = gpv.request
	g.Service = gpv.service
	g.Version = gpv.version

	return nil
}

// ToQueryParameters builds a new query string that will be proxied
func (g GetCapabilitiesRequest) ToQueryParameters() url.Values {
	gpv := getCapabilitiesRequestParameterValue{}
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

// GetCapabilitiesRequest struct with the needed parameters/attributes needed for making a GetCapabilities request
type GetCapabilitiesRequest struct {
	XMLName xml.Name           `xml:"GetCapabilities" yaml:"getCapabilities"`
	Service string             `xml:"service,attr" yaml:"service"`
	Version string             `xml:"version,attr" yaml:"version"`
	Attr    utils.XMLAttribute `xml:",attr" yaml:"attr"`
}
