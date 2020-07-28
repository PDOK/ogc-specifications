package ows

// OperationResponse interface
type OperationResponse interface {
	Type() string
	Service() string
	Version() string
	Validate() Exceptions

	// ParseXML([]byte) Exceptions
	ParseYAML([]byte) Exception
	BuildXML() []byte
	BuildYAML() []byte
}
