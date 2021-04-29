package wfs200

import (
	"net/url"

	"github.com/pdok/ogc-specifications/pkg/common"
)

type DescribeFeatureTypeKVP struct {
}

func (dftkvp *DescribeFeatureTypeKVP) ParseKVP(query url.Values) common.Exceptions {
	return nil
}

func (dftkvp *DescribeFeatureTypeKVP) ParseOperationRequest(or common.OperationRequest) common.Exceptions {
	return nil
}

func (dftkvp *DescribeFeatureTypeKVP) BuildKVP() url.Values {
	return nil
}
