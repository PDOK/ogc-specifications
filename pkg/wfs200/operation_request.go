package wfs200

import (
	"github.com/pdok/ogc-specifications/pkg/wsc110"
	"net/url"
)

type OperationRequest interface {
	Validate(wsc110.Capabilities) wsc110.Exceptions

	ParseQueryParameters(url.Values) wsc110.Exceptions
	ToQueryParameters() url.Values

	ParseXML([]byte) wsc110.Exceptions
	ToXML() []byte
}
