package wsc110

// Capabilities interface for the packages wfs200 and wmts100
type Capabilities interface {
	ParseXML([]byte) error
	ParseYAML([]byte) error
}
