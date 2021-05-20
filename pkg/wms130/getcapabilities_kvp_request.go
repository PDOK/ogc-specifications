package wms130

import (
	"net/url"
	"strings"
)

//getCapabilitiesKVPRequest struct
type getCapabilitiesKVPRequest struct {
	// Table 8 - The Parameters of a GetMap request
	service string `yaml:"service,omitempty"`
	baseRequestKVP
}

// ParseQueryParameters builds a GetCapabilities object based on the available query parameters
func (gckvp *getCapabilitiesKVPRequest) parseQueryParameters(query url.Values) Exceptions {
	var exceptions Exceptions
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, InvalidParameterValue(k, strings.Join(v, ",")))
		} else {
			switch strings.ToUpper(k) {
			case SERVICE:
				gckvp.service = strings.ToUpper(v[0])
			case VERSION:
				gckvp.baseRequestKVP.version = v[0]
			case REQUEST:
				gckvp.baseRequestKVP.request = v[0]
			}
		}
	}

	if len(exceptions) > 0 {
		return exceptions
	}

	return nil
}

// ParseOperationRequest builds a GetCapabilitiesKVP object based on a GetCapabilities struct
// This is a 'dummy' implementation, because for a GetCapabilities request it will always be
// Mandatory:  REQUEST=GetCapabilities
//             SERVICE=WMS
// Optional:   VERSION=1.3.0
func (gckvp *getCapabilitiesKVPRequest) parseGetCapabilitiesRequest(gc GetCapabilitiesRequest) Exceptions {
	gckvp.request = getcapabilities
	gckvp.version = gc.Version
	gckvp.service = gc.Service

	return nil
}

// BuildKVP builds a url.Values query from a GetMapKVP struct
func (gckvp *getCapabilitiesKVPRequest) toQueryParameters() url.Values {
	query := make(map[string][]string)

	query[SERVICE] = []string{gckvp.service}
	query[VERSION] = []string{gckvp.version}
	query[REQUEST] = []string{gckvp.request}

	return query
}
