package wms130

import (
	"encoding/xml"

	"github.com/pdok/ogc-specifications/pkg/common"
)

//
const (
	getcapabilities = `GetCapabilities`
)

// GetCapabilities struct with the needed parameters/attributes needed for making a GetCapabilities request
type GetCapabilities struct {
	XMLName xml.Name `xml:"GetCapabilities" yaml:"getcapabilities"`
	BaseRequest
}

// Type returns GetCapabilities
func (gc *GetCapabilities) Type() string {
	return getcapabilities
}

// Validate returns GetCapabilities
func (gc *GetCapabilities) Validate(c common.Capabilities) common.Exceptions {
	var exceptions common.Exceptions
	return exceptions
}
