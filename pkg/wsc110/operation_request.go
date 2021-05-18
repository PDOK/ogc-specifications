package wsc110

import (
	"net/url"
)

type OperationRequest interface {
	Validate(Capabilities) Exceptions

	ParseQueryParameters(url.Values) Exceptions
	ToQueryParameters() url.Values

	ParseXML([]byte) Exceptions
	ToXML() []byte
}
