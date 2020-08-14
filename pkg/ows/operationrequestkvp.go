package ows

import (
	"net/url"
)

// OperationRequestKVP interface
// This interface is a layer in front of a OperationRequest struct
// to translate KVP to OperationRequest structs and generating
// OperationRequestKVP struct from OperationRequests
type OperationRequestKVP interface {
	ParseKVP(url.Values) Exceptions
	ParseOperationRequest(OperationRequest) Exceptions
	BuildKVP() url.Values
}
