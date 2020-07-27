package validation

import (
	"github.com/pdok/ogc-specifications/pkg/wms130/capabilities"
	"github.com/pdok/ogc-specifications/pkg/wms130/request"
)

// GetMapWrapper struct
type GetMapWrapper struct {
	capabilities *capabilities.Capability
	getmap       *request.GetMap
}
