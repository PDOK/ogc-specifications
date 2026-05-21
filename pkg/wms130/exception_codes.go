package wms130

import (
	"fmt"
	"github.com/pdok/ogc-specifications/pkg/common"
)

// InvalidFormat exception
func InvalidFormat(unknownformat string) exception {
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionText: fmt.Sprintf("The format: %s, is a invalid image format", unknownformat),
		ExceptionCode: `InvalidFormat`,
	}}
}

// InvalidCRS exception
func InvalidCRS(s ...string) exception {
	if len(s) == 1 {
		return exception{ExceptionDetails: common.ExceptionDetails{
			ExceptionText: fmt.Sprintf("CRS is not known by this service: %s", s[0]),
			ExceptionCode: `InvalidCRS`,
		}}
	}
	if len(s) == 2 {
		return exception{ExceptionDetails: common.ExceptionDetails{
			ExceptionText: fmt.Sprintf("The CRS: %s is not known by the layer: %s", s[0], s[1]),
			ExceptionCode: `InvalidCRS`,
		}}
	}
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionCode: `InvalidCRS`,
	}}
}

// LayerNotDefined exception
func LayerNotDefined(s ...string) exception {
	if len(s) == 1 {
		return exception{ExceptionDetails: common.ExceptionDetails{
			ExceptionText: fmt.Sprintf("The layer: %s is not known by the server", s[0]),
			ExceptionCode: `LayerNotDefined`,
		}}
	}
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionCode: `LayerNotDefined`,
	}}
}

// StyleNotDefined exception
func StyleNotDefined(s ...string) exception {
	if len(s) == 2 {
		return exception{ExceptionDetails: common.ExceptionDetails{
			ExceptionText: fmt.Sprintf("The style: %s is not known by the server for the layer: %s", s[0], s[1]),
			ExceptionCode: `StyleNotDefined`,
		}}
	}
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionText: `There is a one-to-one correspondence between the values in the LAYERS parameter and the values in the STYLES parameter. 
	Expecting an empty string for the STYLES like STYLES= or comma-separated list STYLES=,,, or using keyword default STYLES=default,default,...`,
		ExceptionCode: `StyleNotDefined`,
	}}
}

// LayerNotQueryable exception
func LayerNotQueryable(s ...string) exception {
	if len(s) == 1 {
		return exception{ExceptionDetails: common.ExceptionDetails{
			ExceptionText: fmt.Sprintf("Layer: %s, can not be queried", s[0]),
			ExceptionCode: `LayerNotQueryable`,
			LocatorCode:   s[0],
		}}
	}
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionCode: `LayerNotQueryable`,
	}}
}

// InvalidPoint exception
// i and j are strings so we can return none integer values in the exception
func InvalidPoint(i, j string) exception {
	// TODO provide giving WIDTH and HEIGHT values in Exception response
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionText: fmt.Sprintf("The parameters I and J are invalid, given: %s for I and %s for J", i, j),
		ExceptionCode: `InvalidPoint`,
	}}
}

// CurrentUpdateSequence exception
func CurrentUpdateSequence() exception {
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionCode: `CurrentUpdateSequence`,
	}}
}

// InvalidUpdateSequence exception
func InvalidUpdateSequence() exception {
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionCode: `InvalidUpdateSequence`,
	}}
}

// MissingDimensionValue exception
func MissingDimensionValue() exception {
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionCode: `MissingDimensionValue`,
	}}
}

// InvalidDimensionValue exception
func InvalidDimensionValue() exception {
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionCode: `InvalidDimensionValue`,
	}}
}

////////////////
////////////////

// MissingParameterValue exception
func MissingParameterValue(s ...string) exception {
	if len(s) >= 2 {
		return exception{ExceptionDetails: common.ExceptionDetails{ExceptionText: fmt.Sprintf("%s key got incorrect value: %s", s[0], s[1]), ExceptionCode: "MissingParameterValue", LocatorCode: s[0]}}
	}
	if len(s) == 1 {
		return exception{ExceptionDetails: common.ExceptionDetails{ExceptionText: fmt.Sprintf("Missing key: %s", s[0]), ExceptionCode: "MissingParameterValue", LocatorCode: s[0]}}
	}

	return exception{ExceptionDetails: common.ExceptionDetails{ExceptionText: `Could not determine REQUEST`, ExceptionCode: "MissingParameterValue", LocatorCode: "REQUEST"}}
}

// InvalidParameterValue exception
func InvalidParameterValue(value, locator string) exception {
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionText: fmt.Sprintf("%s contains a invalid value: %s", locator, value),
		LocatorCode:   value,
		ExceptionCode: `InvalidParameterValue`,
	}}
}

// NoApplicableCode exception
func NoApplicableCode(message string) exception {
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionText: message,
		ExceptionCode: `NoApplicableCode`,
	}}
}

// OperationNotSupported exception
func OperationNotSupported(message string) exception {
	return exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionText: fmt.Sprintf("This service does not know the operation: %s", message),
		ExceptionCode: `OperationNotSupported`,
		LocatorCode:   message,
	}}
}
