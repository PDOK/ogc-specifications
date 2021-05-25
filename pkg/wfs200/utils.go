package wfs200

import (
	"encoding/xml"
	"github.com/pdok/ogc-specifications/pkg/wsc110"
	"net/url"
	"strings"

	"github.com/pdok/ogc-specifications/pkg/utils"
)

// TODO move helper func
func sp(s string) *string {
	return &s
}

type identify struct {
	XMLName xml.Name
}

func ParsePostRequest(body []byte) (OperationRequest, wsc110.Exceptions) {
	requestType, error := IdentifyPostRequest(body)
	if error != nil {
		return nil, error
	}

	request, error := parseRequestType(requestType)
	if error != nil {
		return nil, error
	}

	error = request.ParseXML(body)
	return request, error
}

func ParseGetRequest(queryParameters url.Values) (OperationRequest, wsc110.Exceptions) {
	requestType, error := IdentifyGetRequest(queryParameters)
	if error != nil {
		return nil, error
	}

	request, error := parseRequestType(requestType)
	if error != nil {
		return nil, error
	}

	error = request.ParseQueryParameters(queryParameters)
	return request, error
}

func IdentifyPostRequest(doc []byte) (string, wsc110.Exceptions) {
	var i identify

	if err := xml.Unmarshal(doc, &i); err != nil {
		return ``, nil
	} else {
		return i.XMLName.Local, nil
	}
}

func IdentifyGetRequest(q url.Values) (string, wsc110.Exceptions) {

	query := utils.KeysToUpper(q)

	if query[REQUEST] != nil {
		return query[REQUEST][0], nil
	}

	return ``, wsc110.MissingParameterValue(REQUEST).ToExceptions()
}

func parseRequestType(requestType string) (OperationRequest, wsc110.Exceptions) {
	var request OperationRequest
	switch strings.ToLower(requestType) {
	case "getcapabilities":
		request = &GetCapabilitiesRequest{}
	case "describefeaturetype":
		request = &DescribeFeatureTypeRequest{}
	case "getfeature":
		request = &GetFeatureRequest{}
	default:
		return nil, wsc110.OperationNotSupported(requestType).ToExceptions()
	}
	return request, nil
}
