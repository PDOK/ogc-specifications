package wfs200

import (
	"net/url"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/common"
)

//GetCapabilitiesKVP struct
type GetCapabilitiesKVP struct {
	// Table 8 - The Parameters of a GetMap request
	Service string `yaml:"service,omitempty"`
	BaseRequestKVP
}

// ParseKVP builds a GetCapabilities object based on the available query parameters
func (gckvp *GetCapabilitiesKVP) ParseKVP(query url.Values) common.Exceptions {
	var exceptions common.Exceptions
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, common.InvalidParameterValue(k, strings.Join(v, ",")))
		} else {
			switch strings.ToUpper(k) {
			case SERVICE:
				gckvp.Service = strings.ToUpper(v[0])
			case VERSION:
				gckvp.BaseRequestKVP.Version = v[0]
			case REQUEST:
				gckvp.BaseRequestKVP.Request = v[0]
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
//             SERVICE=WFS
// Optional:   VERSION=2.0.0
func (gckvp *GetCapabilitiesKVP) ParseOperationRequest(or common.OperationRequest) common.Exceptions {
	gc := or.(*GetCapabilitiesRequest)

	gckvp.Request = getcapabilities
	gckvp.Version = gc.Version
	gckvp.Service = gc.Service

	return nil
}

// BuildKVP builds a url.Values query from a GetMapKVP struct
func (gckvp *GetCapabilitiesKVP) BuildKVP() url.Values {
	query := make(map[string][]string)

	query[SERVICE] = []string{gckvp.Service}
	query[VERSION] = []string{gckvp.Version}
	query[REQUEST] = []string{gckvp.Request}

	return query
}
