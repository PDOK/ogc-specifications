package wfs200

import "github.com/pdok/ogc-specifications/pkg/ows"

// WFS 2.0.0 Tokens
const (
	SERVICE = `SERVICE`
	REQUEST = `REQUEST`
	VERSION = `VERSION`

	OUTPUTFORMAT = `OUTPUTFORMAT`
)

// BaseRequest based on Table 5 WFS2.0.0 spec
// Note: not usable for GetCapabilities request regarding deviation of Optional/Mandatory parameters SERVICE and VERSION
type BaseRequest struct {
	Service string           `xml:"service,attr" yaml:"service" validate:"oneof=WFS wfs"`
	Version string           `xml:"version,attr" yaml:"version" validate:"required,eq=2.0.0"`
	Attr    ows.XMLAttribute `xml:",attr"`
}
