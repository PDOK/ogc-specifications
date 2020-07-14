package ows

import "net/url"

// OperationRequest interface
type OperationRequest interface {
	Type() string
	Validate() bool

	ParseXML([]byte) Exception
	ParseKVP(url.Values) Exception
	BuildXML() []byte
	BuildKVP() url.Values

	// TODO YAML support
	// ParseYAML([]byte) Exception
	// BuildYAML() []byte
}
