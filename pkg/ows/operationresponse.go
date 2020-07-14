package ows

// OperationResponse interface
type OperationResponse interface {
	Type() string
	Service() string
	Version() string
	Validate() bool

	// ParseXML([]byte) Exception
	ParseYAML([]byte) Exception
	BuildXML() []byte
	BuildYAML() []byte
}
