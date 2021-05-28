package wfs200

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

// getFeatureParameterValueRequest struct
type getFeatureParameterValueRequest struct {
	service string `yaml:"service,omitempty"`
	baseParameterValueRequest
	// Table 17 — Keywords for GetFeature KVP-encoding
	*commonKeywords
	*standardPresentationParameters
	*standardResolveParameters
	adhocQueryKeywords
	*storedQueryKeywords
}

// CommonKeywords struct
type commonKeywords struct {
	// Table 7
	namespaces *string `yaml:"namespaces,omitempty"`
	// VSP Vendor-specific parameters
}

// standardPresentationParameters struct
type standardPresentationParameters struct {
	// Table 5
	startindex   *string `yaml:"startindex,omitempty"`
	count        *string `yaml:"count,omitempty"`
	outputformat *string `yaml:"outputformat,omitempty"`
	resulttype   *string `yaml:"resulttype,omitempty"`
}

// standardResolveParameters struct
type standardResolveParameters struct {
	// Table 6
	resolve        *string `yaml:"resolve,omitempty"`
	resolvedepth   *string `yaml:"resolvedepth,omitempty"`
	resolvetimeout *string `yaml:"resolvetimeout,omitempty"`
}

// AdhocQueryKeywords struct
// NOTE filter, resourceid and bbox are mutually exclusive
type adhocQueryKeywords struct {
	// Table 8
	typenames string  `yaml:"typenames"`
	aliases   *string `yaml:"aliases,omitempty"`
	srsname   *string `yaml:"srsname,omitempty"`
	// Projection_clause not implemented
	filter         *string `yaml:"filter,omitempty"`
	filterlanguage *string `yaml:"filter_language,omitempty"`
	resourceid     *string `yaml:"resourceid,omitempty"`
	bbox           *string `yaml:"bbox,omitempty"`
	sortby         *string `yaml:"sortby,omitempty"`
}

// StoredQueryKeywords struct
type storedQueryKeywords struct {
	// Table 10
	storedqueryid string `yaml:"storedqueryid"`
	// storedquery_parameter not implemented
}

func (fpv *getFeatureParameterValueRequest) parseQueryParameters(query url.Values) []wsc110.Exception {
	var exceptions []wsc110.Exception
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, wsc110.InvalidParameterValue(k, strings.Join(v, ",")))
		} else {
			switch strings.ToUpper(k) {
			case SERVICE:
				fpv.service = strings.ToUpper(v[0])
			case VERSION:
				fpv.baseParameterValueRequest.version = v[0]
			case REQUEST:
				fpv.baseParameterValueRequest.request = v[0]
			case STARTINDEX:
				vp := v[0]
				if fpv.standardPresentationParameters != nil {
					fpv.standardPresentationParameters.startindex = &vp
				} else {
					spp := standardPresentationParameters{}
					spp.startindex = &vp
					fpv.standardPresentationParameters = &spp
				}
			case COUNT:
				vp := v[0]
				if fpv.standardPresentationParameters != nil {
					fpv.standardPresentationParameters.count = &vp
				} else {
					spp := standardPresentationParameters{}
					spp.count = &vp
					fpv.standardPresentationParameters = &spp
				}
			case OUTPUTFORMAT:
				vp := v[0]
				if fpv.standardPresentationParameters != nil {
					fpv.standardPresentationParameters.outputformat = &vp
				} else {
					spp := standardPresentationParameters{}
					spp.outputformat = &vp
					fpv.standardPresentationParameters = &spp
				}
			case RESULTTYPE:
				vp := v[0]
				if fpv.standardPresentationParameters != nil {
					fpv.standardPresentationParameters.resulttype = &vp
				} else {
					spp := standardPresentationParameters{}
					spp.resulttype = &vp
					fpv.standardPresentationParameters = &spp
				}
			case RESOLVE:
				vp := v[0]
				fpv.standardResolveParameters.resolve = &vp
			case RESOLVEDEPTH:
				vp := v[0]
				fpv.standardResolveParameters.resolvedepth = &vp
			case RESOLVETIMEOUT:
				vp := v[0]
				fpv.standardResolveParameters.resolvetimeout = &vp
			case NAMESPACES:
				vp := v[0]
				if fpv.commonKeywords != nil {
					fpv.commonKeywords.namespaces = &vp
				} else {
					ck := commonKeywords{}
					ck.namespaces = &vp
					fpv.commonKeywords = &ck
				}
			case TYPENAMES:
				fpv.adhocQueryKeywords.typenames = v[0]
			case ALIASES:
				vp := v[0]
				fpv.adhocQueryKeywords.aliases = &vp
			case SRSNAME:
				vp := v[0]
				fpv.adhocQueryKeywords.srsname = &vp
			case FILTER:
				vp := v[0]
				fpv.adhocQueryKeywords.filter = &vp
			case FILTERLANGUAGE:
				vp := v[0]
				fpv.adhocQueryKeywords.filterlanguage = &vp
			case RESOURCEID:
				vp := v[0]
				fpv.adhocQueryKeywords.resourceid = &vp
			case BBOX:
				vp := v[0]
				fpv.adhocQueryKeywords.bbox = &vp
			case SORTBY:
				vp := v[0]
				fpv.adhocQueryKeywords.sortby = &vp
			case STOREDQUERYID:
				fpv.storedQueryKeywords.storedqueryid = v[0]
			}
		}
	}

	if len(exceptions) > 0 {
		return exceptions
	}

	return nil
}

func (fpv *getFeatureParameterValueRequest) parseGetFeatureRequest(f GetFeatureRequest) []wsc110.Exception {

	fpv.request = getfeature
	fpv.version = Version
	fpv.service = Service

	if f.Startindex != nil {
		i := strconv.Itoa(*f.Startindex)
		if fpv.standardPresentationParameters != nil {
			fpv.standardPresentationParameters.startindex = &i
		} else {
			spp := standardPresentationParameters{}
			spp.startindex = &i
			fpv.standardPresentationParameters = &spp
		}
	}

	if f.Count != nil {
		i := strconv.Itoa(*f.Count)
		if fpv.standardPresentationParameters != nil {
			fpv.standardPresentationParameters.count = &i
		} else {
			spp := standardPresentationParameters{}
			spp.count = &i
			fpv.standardPresentationParameters = &spp
		}
	}

	if f.OutputFormat != nil {
		if fpv.standardPresentationParameters != nil {
			fpv.standardPresentationParameters.outputformat = f.OutputFormat
		} else {
			spp := standardPresentationParameters{}
			spp.outputformat = f.OutputFormat
			fpv.standardPresentationParameters = &spp
		}
	}

	if f.ResultType != nil {
		if fpv.standardPresentationParameters != nil {
			fpv.standardPresentationParameters.resulttype = f.ResultType
		} else {
			spp := standardPresentationParameters{}
			spp.resulttype = f.ResultType
			fpv.standardPresentationParameters = &spp
		}
	}

	if f.StandardResolveParameters != nil {
		if f.Resolve != nil {
			fpv.standardResolveParameters.resolve = f.Resolve
		}

		if f.ResolveDepth != nil {
			i := strconv.Itoa(*f.ResolveDepth)
			fpv.standardResolveParameters.resolvedepth = &i
		}

		if f.ResolveTimeout != nil {
			i := strconv.Itoa(*f.ResolveTimeout)
			fpv.standardResolveParameters.resolvetimeout = &i
		}
	}

	// TODO: extract namespaces from dataset specific and with inspire/ogc
	// fpv.commonKeywords.namespaces = &vp

	fpv.adhocQueryKeywords.typenames = f.Query.TypeNames

	// TODO
	// fpv.adhocQueryKeywords.aliases = &vp

	if f.Query.SrsName != nil {
		fpv.srsname = f.Query.SrsName
	}

	if f.Query.Filter != nil {
		if f.Query.Filter.ResourceID != nil {
			s := f.Query.Filter.ResourceID.toString()
			fpv.resourceid = &(s)
		} else if f.Query.Filter.BBOX != nil {
			s := f.Query.Filter.BBOX.MarshalText()
			fpv.bbox = &s
		} else {
			f := f.Query.Filter.toString()
			fpv.filter = &f
		}
	}

	// TODO
	// fpv.adhocQueryKeywords.filter_language = &vp

	if f.Query.SortBy != nil {
		s := f.Query.SortBy.toString()
		fpv.adhocQueryKeywords.sortby = &s
	}

	// TODO
	// fpv.storedQueryKeywords.storedqueryid = v[0]

	return nil
}

func (fpv getFeatureParameterValueRequest) toQueryParameters() url.Values {
	query := make(map[string][]string)

	query[SERVICE] = []string{fpv.service}
	query[VERSION] = []string{fpv.version}
	query[REQUEST] = []string{fpv.request}

	// commonKeywords
	if fpv.commonKeywords != nil {
		if fpv.namespaces != nil {
			query[NAMESPACES] = []string{*fpv.namespaces}
		}
	}

	// standardPresentationParameters
	if fpv.standardPresentationParameters != nil {
		if fpv.startindex != nil {
			query[STARTINDEX] = []string{*fpv.startindex}
		}
		if fpv.count != nil {
			query[COUNT] = []string{*fpv.count}
		}
		if fpv.outputformat != nil {
			query[OUTPUTFORMAT] = []string{*fpv.outputformat}
		}
		if fpv.resulttype != nil {
			query[RESULTTYPE] = []string{*fpv.resulttype}
		}
	}

	// standardResolveParameters
	if fpv.standardResolveParameters != nil {
		if fpv.resolve != nil {
			query[RESOLVE] = []string{*fpv.resolve}
		}
		if fpv.resolvedepth != nil {
			query[RESOLVEDEPTH] = []string{*fpv.resolvedepth}
		}
		if fpv.resolvetimeout != nil {
			query[RESOLVETIMEOUT] = []string{*fpv.resolvetimeout}
		}
	}

	// // adhocQueryKeywords
	query[TYPENAMES] = []string{fpv.typenames}

	if fpv.aliases != nil {
		query[ALIASES] = []string{*fpv.aliases}
	}

	if fpv.srsname != nil {
		query[SRSNAME] = []string{*fpv.srsname}
	}
	// // Projection_clause not implemented

	if fpv.filter != nil {
		query[FILTER] = []string{*fpv.filter}
	} else if fpv.resourceid != nil {
		query[RESOURCEID] = []string{*fpv.resourceid}
	} else if fpv.bbox != nil {
		query[BBOX] = []string{*fpv.bbox}
	}

	// filter_language *string `yaml:"filter_language,omitempty"`

	// sortby          *string `yaml:"sortby,omitempty"`
	if fpv.sortby != nil {
		query[SORTBY] = []string{*fpv.sortby}
	}

	// storedQueryKeywords
	if fpv.storedQueryKeywords != nil {
		query[STOREDQUERYID] = []string{fpv.storedqueryid}
	}

	return query
}
