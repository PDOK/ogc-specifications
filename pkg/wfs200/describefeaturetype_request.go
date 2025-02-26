package wfs200

import (
	"encoding/xml"
	"net/url"

	"github.com/pdok/ogc-specifications/pkg/utils"
	"github.com/pdok/ogc-specifications/pkg/wsc110"

	"regexp"
	"strings"
)

const (
	TYPENAME = `TYPENAME` //NOTE: TYPENAME for Parameter Value encoding & typeNames for XML encoding
)

// Type returns DescribeFeatureType
func (d DescribeFeatureTypeRequest) Type() string {
	return describefeaturetype
}

// Validate returns GetCapabilities
func (d DescribeFeatureTypeRequest) Validate(c wsc110.Capabilities) []wsc110.Exception {
	return nil
}

// ParseXML builds a DescribeFeatureType object based on a XML document
func (d *DescribeFeatureTypeRequest) ParseXML(doc []byte) []wsc110.Exception {
	var xmlattributes utils.XMLAttribute
	if err := xml.Unmarshal(doc, &xmlattributes); err != nil {
		return wsc110.NoApplicableCode("Could not process XML, is it XML?").ToExceptions()
	}
	if err := xml.Unmarshal(doc, &d); err != nil {
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

	d.Attr = utils.StripDuplicateAttr(n)
	return nil
}

// ParseQueryParameters builds a DescribeFeatureType object based on the available query parameters
func (d *DescribeFeatureTypeRequest) ParseQueryParameters(query url.Values) []wsc110.Exception {
	if len(query) == 0 {
		// When there are no query value we know that at least
		// the manadorty VERSION parameter is missing.
		return wsc110.MissingParameterValue(VERSION).ToExceptions()
	}

	dpv := describeFeatureTypeRequestParameterValue{}

	if exceptions := dpv.parseQueryParameters(query); exceptions != nil {
		return exceptions
	}

	if exceptions := d.parsedescribeFeatureTypeRequestParameterValue(dpv); exceptions != nil {
		return exceptions
	}
	return nil
}

func (d *DescribeFeatureTypeRequest) parsedescribeFeatureTypeRequestParameterValue(dpv describeFeatureTypeRequestParameterValue) []wsc110.Exception {

	// Base
	d.XMLName.Local = describefeaturetype

	var br BaseRequest
	if exceptions := br.parseBaseParameterValueRequest(dpv.baseParameterValueRequest); exceptions != nil {
		return exceptions
	}
	d.BaseRequest = br

	d.TypeNames = dpv.typeName

	if dpv.outputFormat != nil {
		d.OutputFormat = dpv.outputFormat
	} else {
		s := gml32
		d.OutputFormat = &(s)
	}

	return nil
}

// ToQueryParameters  builds a new query string that will be proxied
func (d DescribeFeatureTypeRequest) ToQueryParameters() url.Values {
	dpv := describeFeatureTypeRequestParameterValue{}
	dpv.parseDescribeFeatureTypeRequest(d)

	q := dpv.toQueryParameters()
	return q
}

// ToXML builds a 'new' XML document 'based' on the 'original' XML document
func (d DescribeFeatureTypeRequest) ToXML() []byte {
	si, _ := xml.MarshalIndent(&d, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}

// DescribeFeatureTypeRequest struct with the needed parameters/attributes needed for making a DescribeFeatureType request
type DescribeFeatureTypeRequest struct {
	XMLName xml.Name `xml:"DescribeFeatureType" yaml:"describeFeatureType"`
	BaseRequest
	BaseDescribeFeatureTypeRequest
}

// BaseDescribeFeatureTypeRequest struct used by GetFeature
type BaseDescribeFeatureTypeRequest struct {
	OutputFormat *string `xml:"outputFormat,attr" yaml:"outputFormat"`
	TypeNames    *string `xml:"typeNames,attr" yaml:"typeNames"`
}
