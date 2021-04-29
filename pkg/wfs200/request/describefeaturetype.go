package request

import (
	"encoding/xml"
	"net/url"

	"github.com/pdok/ogc-specifications/pkg/common"
	"github.com/pdok/ogc-specifications/pkg/utils"
	"github.com/pdok/ogc-specifications/pkg/wfs200/capabilities"

	"regexp"
	"strings"
)

//
const (
	TYPENAME            = `TYPENAME` //NOTE: TYPENAME for KVP encoding & typeNames for XML encoding
	describefeaturetype = `DescribeFeatureType`
)

// Type returns DescribeFeatureType
func (dft *DescribeFeatureType) Type() string {
	return describefeaturetype
}

// Validate returns GetCapabilities
func (dft *DescribeFeatureType) Validate(c capabilities.Capabilities) common.Exceptions {
	return nil
}

// ParseXML builds a DescribeFeatureType object based on a XML document
func (dft *DescribeFeatureType) ParseXML(doc []byte) common.Exception {
	var xmlattributes common.XMLAttribute
	if err := xml.Unmarshal(doc, &xmlattributes); err != nil {
		return common.NoApplicableCode("Could not process XML, is it XML?")
	}
	if err := xml.Unmarshal(doc, &dft); err != nil {
		return common.OperationNotSupported(err.Error())
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

	dft.Attr = common.StripDuplicateAttr(n)
	return nil
}

// ParseKVP builds a DescribeFeatureType object based on the available query parameters
func (dft *DescribeFeatureType) ParseKVP(query url.Values) common.Exceptions {
	if len(query) == 0 {
		// When there are no query value we know that at least
		// the manadorty VERSION parameter is missing.
		return common.Exceptions{common.MissingParameterValue(VERSION)}
	}

	q := utils.KeysToUpper(query)

	var br BaseRequest
	if err := br.ParseKVP(q); err != nil {
		return err
	}
	dft.BaseRequest = br

	for k, v := range query {
		switch strings.ToUpper(k) {
		case REQUEST:
			if strings.ToUpper(v[0]) == strings.ToUpper(describefeaturetype) {
				dft.XMLName.Local = describefeaturetype
			}
		case TYPENAME:
			dft.BaseDescribeFeatureTypeRequest.TypeName = &v[0] //TODO maybe process as a comma separated list
		case OUTPUTFORMAT:
			// TODO nothing for now always assume the default text/xml; subtype=gml/3.2
		}
	}

	return nil
}

// BuildKVP builds a new query string that will be proxied
func (dft *DescribeFeatureType) BuildKVP() url.Values {
	querystring := make(map[string][]string)
	querystring[REQUEST] = []string{dft.XMLName.Local}
	querystring[SERVICE] = []string{dft.BaseRequest.Service}
	querystring[VERSION] = []string{dft.BaseRequest.Version}
	if dft.BaseDescribeFeatureTypeRequest.TypeName != nil {
		querystring[TYPENAME] = []string{*dft.BaseDescribeFeatureTypeRequest.TypeName}
	}
	if dft.BaseDescribeFeatureTypeRequest.OutputFormat != nil {
		querystring[OUTPUTFORMAT] = []string{*dft.BaseDescribeFeatureTypeRequest.OutputFormat}
	}
	return querystring
}

// BuildXML builds a 'new' XML document 'based' on the 'original' XML document
func (dft *DescribeFeatureType) BuildXML() []byte {
	si, _ := xml.MarshalIndent(dft, "", "")
	re := regexp.MustCompile(`><.*>`)
	return []byte(xml.Header + re.ReplaceAllString(string(si), "/>"))
}

// DescribeFeatureType struct with the needed parameters/attributes needed for making a DescribeFeatureType request
type DescribeFeatureType struct {
	XMLName xml.Name `xml:"DescribeFeatureType" yaml:"describefeaturetype"`
	BaseRequest
	BaseDescribeFeatureTypeRequest
}

// BaseDescribeFeatureTypeRequest struct used by GetFeature
type BaseDescribeFeatureTypeRequest struct {
	OutputFormat *string `xml:"outputFormat,attr" yaml:"outputformat"`
	TypeName     *string `xml:"typeNames,attr" yaml:"typenames"`
}
