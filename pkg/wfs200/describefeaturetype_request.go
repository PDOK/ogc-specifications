package wfs200

import (
	"encoding/xml"
	"net/url"

	"github.com/pdok/ogc-specifications/pkg/utils"
	"github.com/pdok/ogc-specifications/pkg/wsc110"

	"regexp"
	"strings"
)

//
const (
	TYPENAME = `TYPENAME` //NOTE: TYPENAME for KVP encoding & typeNames for XML encoding
)

// Type returns DescribeFeatureType
func (dft *DescribeFeatureTypeRequest) Type() string {
	return describefeaturetype
}

// Validate returns GetCapabilities
func (dft *DescribeFeatureTypeRequest) Validate(c Capabilities) []wsc110.Exception {
	return nil
}

// ParseXML builds a DescribeFeatureType object based on a XML document
func (dft *DescribeFeatureTypeRequest) ParseXML(doc []byte) []wsc110.Exception {
	var xmlattributes utils.XMLAttribute
	if err := xml.Unmarshal(doc, &xmlattributes); err != nil {
		return wsc110.NoApplicableCode("Could not process XML, is it XML?").ToExceptions()
	}
	if err := xml.Unmarshal(doc, &dft); err != nil {
		// TODO fix with pretty exception message
		return wsc110.NoApplicableCode(err.Error()).ToExceptions()
	}
	var n []xml.Attr
	for _, a := range xmlattributes {
		switch strings.ToUpper(a.Name.Local) {
		case VERSION:
		case SERVICE:
		case TYPENAME:
		default:
			n = append(n, a)
		}
	}

	dft.Attr = utils.StripDuplicateAttr(n)
	return nil
}

// ParseQueryParameters builds a DescribeFeatureType object based on the available query parameters
func (dft *DescribeFeatureTypeRequest) ParseQueryParameters(query url.Values) []wsc110.Exception {
	if len(query) == 0 {
		// When there are no query value we know that at least
		// the manadorty VERSION parameter is missing.
		return wsc110.MissingParameterValue(VERSION).ToExceptions()
	}

	dftkvp := describeFeatureTypeKVPRequest{}

	if exceptions := dftkvp.parseQueryParameters(query); exceptions != nil {
		return exceptions
	}

	if exceptions := dft.parseKVPRequest(dftkvp); exceptions != nil {
		return exceptions
	}
	return nil
}

func (dft *DescribeFeatureTypeRequest) parseKVPRequest(dftkvp describeFeatureTypeKVPRequest) []wsc110.Exception {

	// Base
	dft.XMLName.Local = describefeaturetype

	var br BaseRequest
	if exceptions := br.parseKVPRequest(dftkvp.baseRequestKVP); exceptions != nil {
		return exceptions
	}
	dft.BaseRequest = br

	dft.TypeName = dftkvp.typeName

	if dftkvp.outputFormat != nil {
		dft.OutputFormat = dftkvp.outputFormat
	} else {
		s := gml32
		dft.OutputFormat = &(s)
	}

	return nil
}

// ToQueryParameters  builds a new query string that will be proxied
func (dft DescribeFeatureTypeRequest) ToQueryParameters() url.Values {

	dftkvp := describeFeatureTypeKVPRequest{}
	dftkvp.parseDescribeFeatureTypeRequest(dft)

	q := dftkvp.toQueryParameters()
	return q
}

// ToXML builds a 'new' XML document 'based' on the 'original' XML document
func (dft *DescribeFeatureTypeRequest) ToXML() []byte {
	si, _ := xml.MarshalIndent(dft, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}

// DescribeFeatureType struct with the needed parameters/attributes needed for making a DescribeFeatureType request
type DescribeFeatureTypeRequest struct {
	XMLName xml.Name `xml:"DescribeFeatureType" yaml:"describefeaturetype"`
	BaseRequest
	BaseDescribeFeatureTypeRequest
}

// BaseDescribeFeatureTypeRequest struct used by GetFeature
type BaseDescribeFeatureTypeRequest struct {
	OutputFormat *string `xml:"outputFormat,attr" yaml:"outputformat"`
	TypeName     *string `xml:"typeNames,attr" yaml:"typenames"`
}
