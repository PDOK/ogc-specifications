package wsc110

import (
	"net/url"
)

type OperationRequest interface {
	Validate(Capabilities) []Exception

	ParseQueryParameters(url.Values) []Exception
	ToQueryParameters() url.Values

	ParseXML([]byte) []Exception
	ToXML() []byte
}
