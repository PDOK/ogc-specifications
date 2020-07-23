package validation

import (
	"github.com/pdok/ogc-specifications/pkg/wms130/request"
	"github.com/pdok/ogc-specifications/pkg/wms130/response"
)

// GetMapWrapper struct
type GetMapWrapper struct {
	getcapabilities *response.GetCapabilities
	getmap          *request.GetMap
}
