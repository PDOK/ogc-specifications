// Package wsc110
package wsc110

import (
	"encoding/xml"
	"net/url"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/utils"
)

const (
	SERVICE = `SERVICE`
	REQUEST = `REQUEST`
	VERSION = `VERSION`
)

// GetCapabilitiesRequest contains the data structure as described in Fig.2 and
// Specified in Table 3. of OGC 06-121r3
type GetCapabilitiesRequest struct {
	XMLName        xml.Name           `xml:"GetCapabilities"`
	Service        string             `xml:"service,attr"`
	AcceptVersions AcceptVersions     `xml:"AcceptVersions"`
	Sections       Sections           `xml:"Sections"`
	AcceptFormats  AcceptFormats      `xml:"AcceptFormats"`
	UpdateSequence *string            `xml:"updateSequence,attr"`
	Attr           utils.XMLAttribute `xml:",attr"`
}

type AcceptVersions struct {
	Version []string `xml:"Version"`
}

func (a AcceptVersions) String() []string {
	return a.Version
}

type Sections struct {
	Section []string `xml:"Section"`
}

type AcceptFormats struct {
	OutputFormat []string `xml:"OutputFormat"`
}

// Marshal

// ToXML builds a 'new' XML document 'based' on the 'original' XML document
func (gc GetCapabilitiesRequest) ToXML() []byte {
	return utils.ToXML(gc)
}

// ToQueryParameters builds a new query string that will be proxied
func (gc GetCapabilitiesRequest) ToQueryParameters() url.Values {
	querystring := make(map[string][]string)
	querystring[REQUEST] = []string{gc.XMLName.Local}
	querystring[SERVICE] = []string{gc.Service}
	querystring[VERSION] = gc.AcceptVersions.String()

	return querystring
}

// ParseXML builds a GetCapabilities object based on a XML document
func (g *GetCapabilitiesRequest) ParseXML(doc []byte) []Exception {
	var xmlattributes utils.XMLAttribute
	if err := xml.Unmarshal(doc, &xmlattributes); err != nil {
		return []Exception{NoApplicableCode("Could not process XML, is it XML?")}
	}
	if err := xml.Unmarshal(doc, &g); err != nil {
		return []Exception{OperationNotSupported(err.Error())} //TODO Should be OperationParsingFailed
	}
	var n []xml.Attr
	for _, a := range xmlattributes {
		switch strings.ToUpper(a.Name.Local) {
		case VERSION:
		case SERVICE:
		case `UPDATESEQUENCE`:
		default:
			n = append(n, a)
		}
	}

	g.Attr = utils.StripDuplicateAttr(n)
	return nil
}

// ParseQueryParameters builds a GetCapabilities object based on the available query parameters
func (g *GetCapabilitiesRequest) ParseQueryParameters(query url.Values) []Exception {
	if len(query) == 0 {
		// When there are no query value we know that at least
		// the manadorty SERVICE and REQUEST parameter is missing.
		exceptions := MissingParameterValue(SERVICE).ToExceptions()
		exceptions = append(exceptions, MissingParameterValue(REQUEST))
		return exceptions
	}

	// gpv := getCapabilitiesRequestParameterValue{}
	// if exception := gpv.parseQueryParameters(query); exception != nil {
	// 	return exception
	// }

	// if exception := g.parsegetCapabilitiesRequestParameterValue(gpv); exception != nil {
	// 	return exception
	// }

	return nil
}
