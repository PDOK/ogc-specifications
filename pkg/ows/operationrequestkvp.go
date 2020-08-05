package ows

import (
	"net/url"
)

// OperationRequestKVP interface
type OperationRequestKVP interface {
	Type() string
	Validate(Capability) Exceptions

	ParseKVP(url.Values) Exceptions
	ParseOperationsRequest(OperationRequest) Exceptions
	BuildKVP() url.Values
}
