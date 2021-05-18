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
	requestType, error := IdentifyPostRequest(body)
	if error != nil {
		return nil, error
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
	error = request.ParseXML(body)
	return request, error
}

func ParseGetRequest(queryParameters url.Values) (OperationRequest, Exceptions) {
	requestType, error := IdentifyGetRequest(queryParameters)
	if error != nil {
		return nil, error
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

	error = request.ParseQueryParameters(queryParameters)
	return request, error
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
