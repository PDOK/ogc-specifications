package ows

import "net/url"

// OperationRequest interface
type OperationRequest interface {
	Type() string
	Validate(Capability) Exceptions

	ParseXML([]byte) Exceptions
	ParseKVP(url.Values) Exceptions
	BuildXML() []byte
	BuildKVP() url.Values

	// TODO YAML support
	// ParseYAML([]byte) Exceptions
	// BuildYAML() []byte
}
