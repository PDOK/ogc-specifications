package wms130

import (
	"net/url"
)

// OperationRequest interface used by the wms130 packages
type OperationRequest interface {
	Validate(Capabilities) Exceptions

	ParseQueryParameters(url.Values) Exceptions
	ToQueryParameters() url.Values

	ParseXML([]byte) Exceptions
	ToXML() []byte
}
