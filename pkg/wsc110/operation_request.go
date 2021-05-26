package wsc110

import (
	"net/url"
)

// OperationRequest interface used by the wfs200 and wmts100 packages
type OperationRequest interface {
	Validate(Capabilities) []Exception

	ParseQueryParameters(url.Values) []Exception
	ToQueryParameters() url.Values

	ParseXML([]byte) []Exception
	ToXML() []byte
}
