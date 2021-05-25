package wfs200

import "net/url"

type OperationRequestKVP interface {
	ParseQueryParameters(url.Values) Exceptions
	ParseOperationRequest(OperationRequest) Exceptions
	ToQueryParameters() url.Values
}
