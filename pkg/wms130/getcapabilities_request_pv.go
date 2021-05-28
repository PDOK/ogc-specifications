package wms130

import (
	"net/url"
	"strings"
)

//getCapabilitiesParameterValueRequest struct
type getCapabilitiesParameterValueRequest struct {
	// Table 8 - The Parameters of a GetMap request
	service string `yaml:"service,omitempty"`
	baseParameterValueRequest
}

// ParseQueryParameters builds a GetCapabilities object based on the available query parameters
func (gpv *getCapabilitiesParameterValueRequest) parseQueryParameters(query url.Values) Exceptions {
	var exceptions Exceptions
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, InvalidParameterValue(k, strings.Join(v, ",")))
		} else {
			switch strings.ToUpper(k) {
			case SERVICE:
				gpv.service = strings.ToUpper(v[0])
			case VERSION:
				gpv.baseParameterValueRequest.version = v[0]
			case REQUEST:
				gpv.baseParameterValueRequest.request = v[0]
			}
		}
	}

	if len(exceptions) > 0 {
		return exceptions
	}

	return nil
}

// ParseOperationRequest builds a getCapabilitiesParameterValueRequest object based on a GetCapabilities struct
// This is a 'dummy' implementation, because for a GetCapabilities request it will always be
// Mandatory:  REQUEST=GetCapabilities
//             SERVICE=WMS
// Optional:   VERSION=1.3.0
func (gpv *getCapabilitiesParameterValueRequest) parseGetCapabilitiesRequest(g GetCapabilitiesRequest) Exceptions {
	gpv.request = getcapabilities
	gpv.version = g.Version
	gpv.service = g.Service

	return nil
}

// toQueryParameters builds a url.Values query from a getCapabilitiesParameterValueRequest struct
func (gpv getCapabilitiesParameterValueRequest) toQueryParameters() url.Values {
	query := make(map[string][]string)

	query[SERVICE] = []string{gpv.service}
	query[VERSION] = []string{gpv.version}
	query[REQUEST] = []string{gpv.request}

	return query
}
