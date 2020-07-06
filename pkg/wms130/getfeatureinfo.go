package wms130

import (
	"net/url"

	"github.com/pdok/ogc-specifications/pkg/ows"
)

//
const (
	getfeatureinfo = `GetFeatureInfo`
)

// Type returns GetFeatureInfo
func (gfi *GetFeatureInfo) Type() string {
	return getfeatureinfo
}

// ParseBody builds a GetFeatureInfo object based on the given body
func (gfi *GetFeatureInfo) ParseBody(body []byte) ows.Exception {
	return nil
}

// ParseQuery builds a GetFeatureInfo object based on the available query parameters
func (gfi *GetFeatureInfo) ParseQuery(query url.Values) ows.Exception {
	return nil
}

// BuildQuery builds a new query string that will be proxied
func (gfi *GetFeatureInfo) BuildQuery() url.Values {
	querystring := make(map[string][]string)
	return querystring
}

// BuildBody builds a 'new' XML document 'based' on the 'original' XML document
func (gfi *GetFeatureInfo) BuildBody() []byte {
	return []byte(``)
}

// GetFeatureInfo struct with the needed parameters/attributes needed for making a GetFeatureInfo request
type GetFeatureInfo struct {
}

// Validate a GetFeatureInfo
func (gfi *GetFeatureInfo) Validate() ows.Exception {
	return nil
}
