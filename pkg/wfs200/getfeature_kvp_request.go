package wfs200

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

// getFeatureKVPRequest struct
type getFeatureKVPRequest struct {
	service string `yaml:"service,omitempty"`
	baseRequestKVP
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
	filter          *string `yaml:"filter,omitempty"`
	filter_language *string `yaml:"filter_language,omitempty"`
	resourceid      *string `yaml:"resourceid,omitempty"`
	bbox            *string `yaml:"bbox,omitempty"`
	sortby          *string `yaml:"sortby,omitempty"`
}

// StoredQueryKeywords struct
type storedQueryKeywords struct {
	// Table 10
	storedqueryid string `yaml:"storedqueryid"`
	// storedquery_parameter not implemented
}

func (gfkvp *getFeatureKVPRequest) parseQueryParameters(query url.Values) []wsc110.Exception {
	var exceptions []wsc110.Exception
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, wsc110.InvalidParameterValue(k, strings.Join(v, ",")))
		} else {
			switch strings.ToUpper(k) {
			case SERVICE:
				gfkvp.service = strings.ToUpper(v[0])
			case VERSION:
				gfkvp.baseRequestKVP.version = v[0]
			case REQUEST:
				gfkvp.baseRequestKVP.request = v[0]
			case STARTINDEX:
				vp := v[0]
				if gfkvp.standardPresentationParameters != nil {
					gfkvp.standardPresentationParameters.startindex = &vp
				} else {
					spp := standardPresentationParameters{}
					spp.startindex = &vp
					gfkvp.standardPresentationParameters = &spp
				}
			case COUNT:
				vp := v[0]
				if gfkvp.standardPresentationParameters != nil {
					gfkvp.standardPresentationParameters.count = &vp
				} else {
					spp := standardPresentationParameters{}
					spp.count = &vp
					gfkvp.standardPresentationParameters = &spp
				}
			case OUTPUTFORMAT:
				vp := v[0]
				if gfkvp.standardPresentationParameters != nil {
					gfkvp.standardPresentationParameters.outputformat = &vp
				} else {
					spp := standardPresentationParameters{}
					spp.outputformat = &vp
					gfkvp.standardPresentationParameters = &spp
				}
			case RESULTTYPE:
				vp := v[0]
				if gfkvp.standardPresentationParameters != nil {
					gfkvp.standardPresentationParameters.resulttype = &vp
				} else {
					spp := standardPresentationParameters{}
					spp.resulttype = &vp
					gfkvp.standardPresentationParameters = &spp
				}
			case RESOLVE:
				vp := v[0]
				gfkvp.standardResolveParameters.resolve = &vp
			case RESOLVEDEPTH:
				vp := v[0]
				gfkvp.standardResolveParameters.resolvedepth = &vp
			case RESOLVETIMEOUT:
				vp := v[0]
				gfkvp.standardResolveParameters.resolvetimeout = &vp
			case NAMESPACES:
				vp := v[0]
				if gfkvp.commonKeywords != nil {
					gfkvp.commonKeywords.namespaces = &vp
				} else {
					ck := commonKeywords{}
					ck.namespaces = &vp
					gfkvp.commonKeywords = &ck
				}
			case TYPENAMES:
				gfkvp.adhocQueryKeywords.typenames = v[0]
			case ALIASES:
				vp := v[0]
				gfkvp.adhocQueryKeywords.aliases = &vp
			case SRSNAME:
				vp := v[0]
				gfkvp.adhocQueryKeywords.srsname = &vp
			case FILTER:
				vp := v[0]
				gfkvp.adhocQueryKeywords.filter = &vp
			case FILTERLANGUAGE:
				vp := v[0]
				gfkvp.adhocQueryKeywords.filter_language = &vp
			case RESOURCEID:
				vp := v[0]
				gfkvp.adhocQueryKeywords.resourceid = &vp
			case BBOX:
				vp := v[0]
				gfkvp.adhocQueryKeywords.bbox = &vp
			case SORTBY:
				vp := v[0]
				gfkvp.adhocQueryKeywords.sortby = &vp
			case STOREDQUERYID:
				gfkvp.storedQueryKeywords.storedqueryid = v[0]
			}
		}
	}

	if len(exceptions) > 0 {
		return exceptions
	}

	return nil
}

func (gfkvp *getFeatureKVPRequest) parseGetFeatureRequest(gf GetFeatureRequest) []wsc110.Exception {

	gfkvp.request = getfeature
	gfkvp.version = Version
	gfkvp.service = Service

	if gf.Startindex != nil {
		i := strconv.Itoa(*gf.Startindex)
		if gfkvp.standardPresentationParameters != nil {
			gfkvp.standardPresentationParameters.startindex = &i
		} else {
			spp := standardPresentationParameters{}
			spp.startindex = &i
			gfkvp.standardPresentationParameters = &spp
		}
	}

	if gf.Count != nil {
		i := strconv.Itoa(*gf.Count)
		if gfkvp.standardPresentationParameters != nil {
			gfkvp.standardPresentationParameters.count = &i
		} else {
			spp := standardPresentationParameters{}
			spp.count = &i
			gfkvp.standardPresentationParameters = &spp
		}
	}

	if gf.OutputFormat != nil {
		if gfkvp.standardPresentationParameters != nil {
			gfkvp.standardPresentationParameters.outputformat = gf.OutputFormat
		} else {
			spp := standardPresentationParameters{}
			spp.outputformat = gf.OutputFormat
			gfkvp.standardPresentationParameters = &spp
		}
	}

	if gf.ResultType != nil {
		if gfkvp.standardPresentationParameters != nil {
			gfkvp.standardPresentationParameters.resulttype = gf.ResultType
		} else {
			spp := standardPresentationParameters{}
			spp.resulttype = gf.ResultType
			gfkvp.standardPresentationParameters = &spp
		}
	}

	if gf.StandardResolveParameters != nil {
		if gf.Resolve != nil {
			gfkvp.standardResolveParameters.resolve = gf.Resolve
		}

		if gf.ResolveDepth != nil {
			i := strconv.Itoa(*gf.ResolveDepth)
			gfkvp.standardResolveParameters.resolvedepth = &i
		}

		if gf.ResolveTimeout != nil {
			i := strconv.Itoa(*gf.ResolveTimeout)
			gfkvp.standardResolveParameters.resolvetimeout = &i
		}
	}

	// TODO: extract namespaces from dataset specific and with inspire/ogc
	// gfkvp.commonKeywords.namespaces = &vp

	gfkvp.adhocQueryKeywords.typenames = gf.Query.TypeNames

	// TODO
	// gfkvp.adhocQueryKeywords.aliases = &vp

	if gf.Query.SrsName != nil {
		gfkvp.srsname = gf.Query.SrsName
	}

	if gf.Query.Filter != nil {
		if gf.Query.Filter.ResourceID != nil {
			s := gf.Query.Filter.ResourceID.toString()
			gfkvp.resourceid = &(s)
		} else if gf.Query.Filter.BBOX != nil {
			s := gf.Query.Filter.BBOX.MarshalText()
			gfkvp.bbox = &s
		} else {
			f := gf.Query.Filter.toString()
			gfkvp.filter = &f
		}
	}

	// TODO
	// gfkvp.adhocQueryKeywords.filter_language = &vp

	if gf.Query.SortBy != nil {
		s := gf.Query.SortBy.toString()
		gfkvp.adhocQueryKeywords.sortby = &s
	}

	// TODO
	// gfkvp.storedQueryKeywords.storedqueryid = v[0]

	return nil
}

func (gfkvp getFeatureKVPRequest) toQueryParameters() url.Values {
	query := make(map[string][]string)

	query[SERVICE] = []string{gfkvp.service}
	query[VERSION] = []string{gfkvp.version}
	query[REQUEST] = []string{gfkvp.request}

	// commonKeywords
	if gfkvp.commonKeywords != nil {
		if gfkvp.namespaces != nil {
			query[NAMESPACES] = []string{*gfkvp.namespaces}
		}
	}

	// standardPresentationParameters
	if gfkvp.standardPresentationParameters != nil {
		if gfkvp.startindex != nil {
			query[STARTINDEX] = []string{*gfkvp.startindex}
		}
		if gfkvp.count != nil {
			query[COUNT] = []string{*gfkvp.count}
		}
		if gfkvp.outputformat != nil {
			query[OUTPUTFORMAT] = []string{*gfkvp.outputformat}
		}
		if gfkvp.resulttype != nil {
			query[RESULTTYPE] = []string{*gfkvp.resulttype}
		}
	}

	// standardResolveParameters
	if gfkvp.standardResolveParameters != nil {
		if gfkvp.resolve != nil {
			query[RESOLVE] = []string{*gfkvp.resolve}
		}
		if gfkvp.resolvedepth != nil {
			query[RESOLVEDEPTH] = []string{*gfkvp.resolvedepth}
		}
		if gfkvp.resolvetimeout != nil {
			query[RESOLVETIMEOUT] = []string{*gfkvp.resolvetimeout}
		}
	}

	// // adhocQueryKeywords
	query[TYPENAMES] = []string{gfkvp.typenames}

	if gfkvp.aliases != nil {
		query[ALIASES] = []string{*gfkvp.aliases}
	}

	if gfkvp.srsname != nil {
		query[SRSNAME] = []string{*gfkvp.srsname}
	}
	// // Projection_clause not implemented

	if gfkvp.filter != nil {
		query[FILTER] = []string{*gfkvp.filter}
	} else if gfkvp.resourceid != nil {
		query[RESOURCEID] = []string{*gfkvp.resourceid}
	} else if gfkvp.bbox != nil {
		query[BBOX] = []string{*gfkvp.bbox}
	}

	// filter_language *string `yaml:"filter_language,omitempty"`

	// sortby          *string `yaml:"sortby,omitempty"`
	if gfkvp.sortby != nil {
		query[SORTBY] = []string{*gfkvp.sortby}
	}

	// storedQueryKeywords
	if gfkvp.storedQueryKeywords != nil {
		query[STOREDQUERYID] = []string{gfkvp.storedqueryid}
	}

	return query
}
