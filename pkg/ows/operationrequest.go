package ows

import "net/url"

// OperationRequest interface
type OperationRequest interface {
	Type() string
	ParseBody([]byte) Exception
	ParseQuery(url.Values) Exception
	BuildQuery() url.Values
	BuildBody() []byte
	Validate() bool
}
