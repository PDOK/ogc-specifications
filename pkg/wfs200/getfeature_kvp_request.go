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

func (gfkvp *getFeatureKVPRequest) parseQueryParameters(query url.Values) wsc110.Exceptions {
	var exceptions wsc110.Exceptions
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
				gfkvp.standardPresentationParameters.startindex = &vp
			case COUNT:
				vp := v[0]
				gfkvp.standardPresentationParameters.count = &vp
			case OUTPUTFORMAT:
				vp := v[0]
				gfkvp.standardPresentationParameters.outputformat = &vp
			case RESULTTYPE:
				vp := v[0]
				gfkvp.standardPresentationParameters.resulttype = &vp
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
				gfkvp.commonKeywords.namespaces = &vp
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

func (gfkvp *getFeatureKVPRequest) parseGetFeatureRequest(gf GetFeatureRequest) wsc110.Exceptions {

	gfkvp.request = getfeature
	gfkvp.version = Version
	gfkvp.service = Service

	if gf.Startindex != nil {
		i := strconv.Itoa(*gf.Startindex)
		gfkvp.standardPresentationParameters.startindex = &i
	}

	if gf.Count != nil {
		i := strconv.Itoa(*gf.Count)
		gfkvp.standardPresentationParameters.count = &i
	}

	if gf.OutputFormat != nil {
		gfkvp.standardPresentationParameters.outputformat = gf.OutputFormat
	}

	if gf.ResultType != nil {
		gfkvp.standardPresentationParameters.resulttype = gf.ResultType
	}

	// gfkvp.standardResolveParameters.resolve = &vp
	// gfkvp.standardResolveParameters.resolvedepth = &vp
	// gfkvp.standardResolveParameters.resolvetimeout = &vp
	// gfkvp.commonKeywords.namespaces = &vp
	// gfkvp.adhocQueryKeywords.typenames = v[0]
	// gfkvp.adhocQueryKeywords.aliases = &vp
	// gfkvp.adhocQueryKeywords.srsname = &vp
	// gfkvp.adhocQueryKeywords.filter = &vp
	// gfkvp.adhocQueryKeywords.filter_language = &vp
	// gfkvp.adhocQueryKeywords.resourceid = &vp
	// gfkvp.adhocQueryKeywords.bbox = &vp
	// gfkvp.adhocQueryKeywords.sortby = &vp
	// gfkvp.storedQueryKeywords.storedqueryid = v[0]

	return nil
}

func (gfkvp *getFeatureKVPRequest) toQueryParameters() url.Values {
	return nil
}
