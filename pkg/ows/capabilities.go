package ows

// Capabilities interface
// return an error, if needed, not a Exception
// because this isn't a true OWS object only a base
// from which GetCapabilities can build and OperationRequest
// and OperationResponse can be validated against
type Capabilities interface {
	ParseXML([]byte) error
	ParseYAMl([]byte) error
}
