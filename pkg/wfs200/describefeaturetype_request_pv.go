package wfs200

import (
	"net/url"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

type describeFeatureTypeRequestParameterValue struct {
	service string `yaml:"service"`
	baseParameterValueRequest

	typeName     *string `yaml:"typeName"`     // [0..*]
	outputFormat *string `yaml:"outputFormat"` // default: "text/xml; subtype=gml/3.2"
}

func (dpv *describeFeatureTypeRequestParameterValue) parseQueryParameters(query url.Values) []wsc110.Exception {
	var exceptions []wsc110.Exception
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, wsc110.InvalidParameterValue(k, strings.Join(v, ",")))
		} else {
			switch strings.ToUpper(k) {
			case SERVICE:
				dpv.service = strings.ToUpper(v[0])
			case VERSION:
				dpv.baseParameterValueRequest.version = v[0]
			case REQUEST:
				dpv.baseParameterValueRequest.request = v[0]
			case TYPENAME:
				vp := v[0]
				dpv.typeName = &vp
			case OUTPUTFORMAT:
				vp := v[0]
				dpv.outputFormat = &vp
			}
		}
	}

	if len(exceptions) > 0 {
		return exceptions
	}

	return nil
}

func (dpv *describeFeatureTypeRequestParameterValue) parseDescribeFeatureTypeRequest(dft DescribeFeatureTypeRequest) []wsc110.Exception {
	dpv.request = describefeaturetype
	dpv.version = dft.Version
	dpv.service = dft.Service
	dpv.typeName = dft.TypeNames
	dpv.outputFormat = dft.OutputFormat
	return nil
}

func (dpv describeFeatureTypeRequestParameterValue) toQueryParameters() url.Values {
	querystring := make(map[string][]string)
	querystring[REQUEST] = []string{dpv.request}
	querystring[SERVICE] = []string{dpv.service}
	querystring[VERSION] = []string{dpv.version}
	if dpv.typeName != nil {
		querystring[TYPENAME] = []string{*dpv.typeName}
	}
	if dpv.outputFormat != nil {
		querystring[OUTPUTFORMAT] = []string{*dpv.outputFormat}
	}
	return querystring
}
