package ows

// OperationResponse interface
type OperationResponse interface {
	Type() string
	Service() string
	Version() string
	Validate() Exceptions //TODO No sure about this one

	// ParseXML([]byte) Exceptions
	ParseYAML([]byte) Exception //TODO Maybe just return error
	BuildXML() []byte
	BuildYAML() []byte
}
