package wms130

import "github.com/pdok/ogc-specifications/pkg/ows"

// WMS 1.3.0 Tokens
const (
	SERVICE = `SERVICE`
	REQUEST = `REQUEST`
	VERSION = `VERSION`
)

// WMSbaseparameters array of base WMS parameters
var WMSbaseparameters = []string{SERVICE, REQUEST, VERSION}

// BaseRequest based on the SLD 1.1 spec 'containing' example implementation of a POST WMS 1.3.0 request
// http://schemas.opengis.net/sld/1.1//example_getmap.xml
// Note: not usable for GetCapabilities request regarding deviation of Optional/Mandatory parameters SERVICE and VERSION
type BaseRequest struct {
	Service string           `xml:"service,attr" yaml:"service" validate:"oneof=WMS wms"`
	Version string           `xml:"version,attr" yaml:"version" validate:"required,eq=1.3.0"`
	Attr    ows.XMLAttribute `xml:",attr"`
}
