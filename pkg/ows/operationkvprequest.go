package ows

import "net/url"

// OperationKVPRequest interface
type OperationKVPRequest interface {
	Type() string
	Validate(Capability) Exceptions

	ParseKVP(url.Values) Exceptions
	BuildKVP() url.Values
}
