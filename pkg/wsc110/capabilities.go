package wsc110

type Capabilities interface {
	ParseXML([]byte) error
	ParseYAMl([]byte) error
}
