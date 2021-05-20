package wms130

import (
	"encoding/xml"
	"net/url"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/utils"
)

// TODO move helper func
func sp(s string) *string {
	return &s
}

func ip(i int) *int {
	return &i
}

func bp(b bool) *bool {
	return &b
}

type identify struct {
	XMLName xml.Name
}

func ParsePostRequest(body []byte) (OperationRequest, Exceptions) {
	requestType, exception := IdentifyPostRequest(body)
	if exception != nil {
		return nil, exception
	}

	var request OperationRequest
	switch strings.ToLower(requestType) {
	case "getcapabilities":
		request = &GetCapabilitiesRequest{}
	case "getmap":
		request = &GetMapRequest{}
	case "getfeatureinfo":
		request = &GetFeatureInfoRequest{}
	default:
		return nil, OperationNotSupported(requestType).ToExceptions()
	}
	exception = request.ParseXML(body)
	return request, exception
}

func ParseGetRequest(queryParameters url.Values) (OperationRequest, Exceptions) {
	requestType, exception := IdentifyGetRequest(queryParameters)
	if exception != nil {
		return nil, exception
	}

	var request OperationRequest
	switch strings.ToLower(requestType) {
	case "getcapabilities":
		request = &GetCapabilitiesRequest{}
	case "getmap":
		request = &GetMapRequest{}
	case "getfeatureinfo":
		request = &GetFeatureInfoRequest{}
	default:
		return nil, OperationNotSupported(requestType).ToExceptions()
	}

	exception = request.ParseQueryParameters(queryParameters)
	return request, exception
}

func IdentifyPostRequest(doc []byte) (string, Exceptions) {
	var i identify

	if err := xml.Unmarshal(doc, &i); err != nil {
		return ``, Exceptions{MissingParameterValue()}
	} else {
		return i.XMLName.Local, nil
	}
}

func IdentifyGetRequest(q url.Values) (string, Exceptions) {

	query := utils.KeysToUpper(q)

	if query[REQUEST] != nil {
		return query[REQUEST][0], nil
	}

	return ``, Exceptions{MissingParameterValue()}
}
