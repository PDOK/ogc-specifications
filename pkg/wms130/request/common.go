package request

import (
	"net/url"

	"github.com/pdok/ogc-specifications/pkg/ows"
)

// Type and Version as constant
const (
	Service string = `WMS`
	Version string = `1.3.0`
)

// WMS 1.3.0 Tokens
const (
	SERVICE = `SERVICE`
	REQUEST = `REQUEST`
	VERSION = `VERSION`
)

// BaseRequest based on the SLD 1.1 spec 'containing' example implementation of a POST WMS 1.3.0 request
// http://schemas.opengis.net/sld/1.1//example_getmap.xml
// Note: not usable for GetCapabilities request regarding deviation of Optional/Mandatory parameters SERVICE and VERSION
type BaseRequest struct {
	Service string           `xml:"service,attr" yaml:"service,omitempty"`
	Version string           `xml:"version,attr" yaml:"version"`
	Attr    ows.XMLAttribute `xml:",attr"`
}

// ParseKVP builds a BaseRequest struct based on the given parameters
func (b *BaseRequest) ParseKVP(query url.Values) ows.Exception {
	if len(query[SERVICE]) > 0 {
		// Service is optional, because it's implicit for a GetMap/GetFeatureInfo request
		b.Service = query[SERVICE][0]
	}
	if len(query[VERSION]) > 0 {
		b.Version = query[VERSION][0]
	} else {
		// Version is mandatory
		return ows.MissingParameterValue(VERSION)
	}
	return nil
}

// Build builds a BaseRequest struct
func (b *BaseRequest) Build(service, version string) ows.Exception {
	if service != `` {
		// Service is optional, because it's implicit for a GetMap/GetFeatureInfo request
		b.Service = service
	}
	if version != `` {
		b.Version = version
	} else {
		// Version is mandatory
		return ows.MissingParameterValue(VERSION)
	}
	return nil
}
