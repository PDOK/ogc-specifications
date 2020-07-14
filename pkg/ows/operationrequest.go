package ows

import "net/url"

// OperationRequest interface
type OperationRequest interface {
	Type() string
	ParseXML([]byte) Exception
	ParseQuery(url.Values) Exception
	BuildQuery() url.Values
	BuildXML() []byte
	Validate() bool
}
