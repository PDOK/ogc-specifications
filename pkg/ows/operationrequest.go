package ows

import "net/url"

// OperationRequest interface
type OperationRequest interface {
	Type() string
	Validate(Capabilities) Exceptions

	ParseXML([]byte) Exceptions
	ParseKVP(url.Values) Exceptions
	BuildXML() []byte
	BuildKVP() url.Values
	ParseOperationRequestKVP(OperationRequestKVP) Exceptions

	// TODO YAML support
	// ParseYAML([]byte) Exceptions
	// BuildYAML() []byte
}
