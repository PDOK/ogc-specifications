package wfs200

import (
	"net/url"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

type DescribeFeatureTypeKVP struct {
}

func (dftkvp *DescribeFeatureTypeKVP) ParseKVP(query url.Values) wsc110.Exceptions {
	return nil
}

func (dftkvp *DescribeFeatureTypeKVP) ParseOperationRequest(or wsc110.OperationRequest) wsc110.Exceptions {
	return nil
}

func (dftkvp *DescribeFeatureTypeKVP) BuildKVP() url.Values {
	return nil
}
