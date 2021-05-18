package wsc110

type Capabilities interface {
	ParseXML([]byte) error
	ParseYAML([]byte) error
}
