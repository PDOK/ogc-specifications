package wfs200

import (
	"net/url"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

//getCapabilitiesParameterValueRequest struct
type getCapabilitiesParameterValueRequest struct {
	service string `yaml:"service,omitempty"`
	baseParameterValueRequest
}

// ParseQueryParameters builds a GetCapabilities object based on the available query parameters
func (gpv *getCapabilitiesParameterValueRequest) parseQueryParameters(query url.Values) []wsc110.Exception {
	var exceptions []wsc110.Exception
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, wsc110.InvalidParameterValue(k, strings.Join(v, ",")))
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

// parseGetCapabilitiesRequest builds a getCapabilitiesParameterValueRequest object based on a GetCapabilities struct
// This is a 'dummy' implementation, because for a GetCapabilities request it will always be
// Mandatory:  REQUEST=GetCapabilities
//             SERVICE=WFS
// Optional:   VERSION=2.0.0
func (gpv *getCapabilitiesParameterValueRequest) parseGetCapabilitiesRequest(gc GetCapabilitiesRequest) []wsc110.Exception {
	gpv.request = getcapabilities
	gpv.version = gc.Version
	gpv.service = gc.Service

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
