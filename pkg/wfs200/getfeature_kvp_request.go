package wfs200

import (
	"net/url"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/common"
	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

// GetFeatureKVP struct
type GetFeatureKVP struct {
	Service string `yaml:"service,omitempty"`
	BaseRequestKVP
	// Table 17 — Keywords for GetFeature KVP-encoding
	*CommonKeywords
	*StandardPresentationParameters
	*StandardResolveParameters
	AdhocQueryKeywords
	*StoredQueryKeywords
}

// CommonKeywords struct
type CommonKeywords struct {
	// Table 7
	Namespaces *string `yaml:"namespaces,omitempty"`
	// VSP Vendor-specific parameters
}

// StandardPresentationParameters struct
type StandardPresentationParameters struct {
	// Table 5
	StartIndex   *string `yaml:"startindex,omitempty"`
	Count        *string `yaml:"count,omitempty"`
	OutputFormat *string `yaml:"outputformat,omitempty"`
	ResultType   *string `yaml:"resulttype,omitempty"`
}

// StandardResolveParameters struct
type StandardResolveParameters struct {
	// Table 6
	Resolve        *string `yaml:"resolve,omitempty"`
	ResolveDepth   *string `yaml:"resolvedepth,omitempty"`
	ResolveTimeout *string `yaml:"resolvetimeout,omitempty"`
}

// AdhocQueryKeywords struct
type AdhocQueryKeywords struct {
	// Table 8
	TypeNames string  `yaml:"typenames"`
	Aliases   *string `yaml:"aliases,omitempty"`
	SrsName   *string `yaml:"srsname,omitempty"`
	// Projection_clause not implemented
	Filter          *string `yaml:"filter,omitempty"`
	Filter_Language *string `yaml:"filter_language,omitempty"`
	ResourceId      *string `yaml:"resourceid,omitempty"`
	Bbox            *string `yaml:"bbox,omitempty"`
	SortBy          *string `yaml:"sortby,omitempty"`
}

// StoredQueryKeywords struct
type StoredQueryKeywords struct {
	// Table 10
	StoredQueryId string `yaml:"storedqueryid"`
	// storedquery_parameter not implemented
}

func (gfkvp *GetFeatureKVP) ParseKVP(query url.Values) wsc110.Exceptions {
	var exceptions wsc110.Exceptions
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, wsc110.InvalidParameterValue(k, strings.Join(v, ",")))
		} else {
			switch strings.ToUpper(k) {
			case SERVICE:
				gfkvp.Service = strings.ToUpper(v[0])
			case VERSION:
				gfkvp.BaseRequestKVP.Version = v[0]
			case REQUEST:
				gfkvp.BaseRequestKVP.Request = v[0]
			case STARTINDEX:
				vp := v[0]
				gfkvp.StandardPresentationParameters.StartIndex = &vp
			case COUNT:
				vp := v[0]
				gfkvp.StandardPresentationParameters.Count = &vp
			case OUTPUTFORMAT:
				vp := v[0]
				gfkvp.StandardPresentationParameters.OutputFormat = &vp
			case RESULTTYPE:
				vp := v[0]
				gfkvp.StandardPresentationParameters.ResultType = &vp
			case RESOLVE:
				vp := v[0]
				gfkvp.StandardResolveParameters.Resolve = &vp
			case RESOLVEDEPTH:
				vp := v[0]
				gfkvp.StandardResolveParameters.ResolveDepth = &vp
			case RESOLVETIMEOUT:
				vp := v[0]
				gfkvp.StandardResolveParameters.ResolveTimeout = &vp
			case NAMESPACES:
				vp := v[0]
				gfkvp.CommonKeywords.Namespaces = &vp
			case TYPENAMES:
				gfkvp.AdhocQueryKeywords.TypeNames = v[0]
			case ALIASES:
				vp := v[0]
				gfkvp.AdhocQueryKeywords.Aliases = &vp
			case SRSNAME:
				vp := v[0]
				gfkvp.AdhocQueryKeywords.SrsName = &vp
			case FILTER:
				vp := v[0]
				gfkvp.AdhocQueryKeywords.Filter = &vp
			case FILTERLANGUAGE:
				vp := v[0]
				gfkvp.AdhocQueryKeywords.Filter_Language = &vp
			case RESOURCEID:
				vp := v[0]
				gfkvp.AdhocQueryKeywords.ResourceId = &vp
			case BBOX:
				vp := v[0]
				gfkvp.AdhocQueryKeywords.Bbox = &vp
			case SORTBY:
				vp := v[0]
				gfkvp.AdhocQueryKeywords.SortBy = &vp
			case STOREDQUERYID:
				gfkvp.StoredQueryKeywords.StoredQueryId = v[0]
			}
		}
	}

	if len(exceptions) > 0 {
		return exceptions
	}

	return nil
}

func (gfkvp *GetFeatureKVP) ParseOperationRequest(or common.OperationRequest) wsc110.Exceptions {
	return nil
}

func (gfkvp *GetFeatureKVP) BuildKVP() url.Values {
	return nil
}
