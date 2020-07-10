package wfs200

import (
	"encoding/xml"
	"net/url"

	"github.com/pdok/ogc-specifications/pkg/ows"
	"github.com/pdok/ogc-specifications/pkg/utils"

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

// ParseBody builds a DescribeFeatureType object based on the given body
func (dft *DescribeFeatureType) ParseBody(body []byte) ows.Exception {
	var xmlattributes ows.XMLAttribute
	if err := xml.Unmarshal(body, &xmlattributes); err != nil {
		return ows.NoApplicableCode("Could not process XML, is it XML?")
	}
	if err := xml.Unmarshal(body, &dft); err != nil {
		return ows.OperationNotSupported(err.Error())
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

	dft.Attr = ows.StripDuplicateAttr(n)
	return nil
}

// ParseQuery builds a DescribeFeatureType object based on the available query parameters
func (dft *DescribeFeatureType) ParseQuery(query url.Values) ows.Exception {

	q := utils.KeysToUpper(query)

	var br BaseRequest
	if err := br.ParseQueryParameters(q); err != nil {
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

// BuildQuery builds a new query string that will be proxied
func (dft *DescribeFeatureType) BuildQuery() url.Values {
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

// BuildBody builds a 'new' XML document 'based' on the 'original' XML document
func (dft *DescribeFeatureType) BuildBody() []byte {
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
