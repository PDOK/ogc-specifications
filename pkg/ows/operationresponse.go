package ows

// OperationResponse interface
type OperationResponse interface {
	Type() string
	Service() string
	Version() string
	Validate() Exception

	// ParseXML([]byte) Exception
	ParseYAML([]byte) Exception
	BuildXML() []byte
	BuildYAML() []byte
}
