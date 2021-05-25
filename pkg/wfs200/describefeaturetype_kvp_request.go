package wfs200

import (
	"net/url"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

type describeFeatureTypeKVPRequest struct {
	service string `yaml:"service"`
	baseRequestKVP

	typeName     *string `yaml:"typename"`     // [0..*]
	outputFormat *string `yaml:"outputformat"` // default: "text/xml; subtype=gml/3.2"
}

func (dftkvp *describeFeatureTypeKVPRequest) parseQueryParameters(query url.Values) []wsc110.Exception {
	var exceptions []wsc110.Exception
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, wsc110.InvalidParameterValue(k, strings.Join(v, ",")))
		} else {
			switch strings.ToUpper(k) {
			case SERVICE:
				dftkvp.service = strings.ToUpper(v[0])
			case VERSION:
				dftkvp.baseRequestKVP.version = v[0]
			case REQUEST:
				dftkvp.baseRequestKVP.request = v[0]
			case TYPENAME:
				vp := v[0]
				dftkvp.typeName = &vp
			case OUTPUTFORMAT:
				vp := v[0]
				dftkvp.outputFormat = &vp
			}
		}
	}

	if len(exceptions) > 0 {
		return exceptions
	}

	return nil
}

func (dftkvp *describeFeatureTypeKVPRequest) parseDescribeFeatureTypeRequest(dft DescribeFeatureTypeRequest) []wsc110.Exception {
	dftkvp.request = describefeaturetype
	dftkvp.version = dft.Version
	dftkvp.service = dft.Service
	dftkvp.typeName = dft.TypeName
	dftkvp.outputFormat = dft.OutputFormat
	return nil
}

func (dftkvp describeFeatureTypeKVPRequest) toQueryParameters() url.Values {
	return nil
}
