package wms130

import (
	"fmt"

	"github.com/pdok/ogc-specifications/pkg/common"
)

// InvalidFormat Exception
func InvalidFormat(unknownFormat string) Exception {
	return Exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionText: fmt.Sprintf("The format: %s, is a invalid image format", unknownFormat),
		ExceptionCode: `InvalidFormat`,
	}}
}

// InvalidCRS Exception
func InvalidCRS(s ...string) Exception {
	if len(s) == 1 {
		return Exception{ExceptionDetails: common.ExceptionDetails{
			ExceptionText: "CRS is not known by this service: " + s[0],
			ExceptionCode: `InvalidCRS`,
		}}
	}
	if len(s) == 2 {
		return Exception{ExceptionDetails: common.ExceptionDetails{
			ExceptionText: fmt.Sprintf("The CRS: %s is not known by the layer: %s", s[0], s[1]),
			ExceptionCode: `InvalidCRS`,
		}}
	}
	return Exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionCode: `InvalidCRS`,
	}}
}

// LayerNotDefined Exception
func LayerNotDefined(s ...string) Exception {
	if len(s) == 1 {
		return Exception{ExceptionDetails: common.ExceptionDetails{
			ExceptionText: fmt.Sprintf("The layer: %s is not known by the server", s[0]),
			ExceptionCode: `LayerNotDefined`,
		}}
	}
	return Exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionCode: `LayerNotDefined`,
	}}
}

// StyleNotDefined Exception
func StyleNotDefined(s ...string) Exception {
	if len(s) == 2 {
		return Exception{ExceptionDetails: common.ExceptionDetails{
			ExceptionText: fmt.Sprintf("The style: %s is not known by the server for the layer: %s", s[0], s[1]),
			ExceptionCode: `StyleNotDefined`,
		}}
	}
	return Exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionText: `There is a one-to-one correspondence between the values in the LAYERS parameter and the values in the STYLES parameter. 
	Expecting an empty string for the STYLES like STYLES= or comma-separated list STYLES=,,, or using keyword default STYLES=default,default,...`,
		ExceptionCode: `StyleNotDefined`,
	}}
}

// LayerNotQueryable Exception
func LayerNotQueryable(s ...string) Exception {
	if len(s) == 1 {
		return Exception{ExceptionDetails: common.ExceptionDetails{
			ExceptionText: fmt.Sprintf("Layer: %s, can not be queried", s[0]),
			ExceptionCode: `LayerNotQueryable`,
			LocatorCode:   s[0],
		}}
	}
	return Exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionCode: `LayerNotQueryable`,
	}}
}

// InvalidPoint Exception
// i and j are strings so we can return none integer values in the Exception
func InvalidPoint(i, j string) Exception {
	// TODO provide giving WIDTH and HEIGHT values in Exception response
	return Exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionText: fmt.Sprintf("The parameters I and J are invalid, given: %s for I and %s for J", i, j),
		ExceptionCode: `InvalidPoint`,
	}}
}

// CurrentUpdateSequence Exception
func CurrentUpdateSequence() Exception {
	return Exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionCode: `CurrentUpdateSequence`,
	}}
}

// InvalidUpdateSequence Exception
func InvalidUpdateSequence() Exception {
	return Exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionCode: `InvalidUpdateSequence`,
	}}
}

// MissingDimensionValue Exception
func MissingDimensionValue() Exception {
	return Exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionCode: `MissingDimensionValue`,
	}}
}

// InvalidDimensionValue Exception
func InvalidDimensionValue() Exception {
	return Exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionCode: `InvalidDimensionValue`,
	}}
}

////////////////
////////////////

// MissingParameterValue Exception
func MissingParameterValue(s ...string) Exception {
	if len(s) >= 2 {
		return Exception{ExceptionDetails: common.ExceptionDetails{ExceptionText: fmt.Sprintf("%s key got incorrect value: %s", s[0], s[1]), ExceptionCode: "MissingParameterValue", LocatorCode: s[0]}}
	}
	if len(s) == 1 {
		return Exception{ExceptionDetails: common.ExceptionDetails{ExceptionText: "Missing key: " + s[0], ExceptionCode: "MissingParameterValue", LocatorCode: s[0]}}
	}

	return Exception{ExceptionDetails: common.ExceptionDetails{ExceptionText: `Could not determine REQUEST`, ExceptionCode: "MissingParameterValue", LocatorCode: "REQUEST"}}
}

// InvalidParameterValue Exception
func InvalidParameterValue(value, locator string) Exception {
	return Exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionText: fmt.Sprintf("%s contains a invalid value: %s", locator, value),
		LocatorCode:   value,
		ExceptionCode: `InvalidParameterValue`,
	}}
}

// NoApplicableCode Exception
func NoApplicableCode(message string) Exception {
	return Exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionText: message,
		ExceptionCode: `NoApplicableCode`,
	}}
}

// OperationNotSupported Exception
func OperationNotSupported(message string) Exception {
	return Exception{ExceptionDetails: common.ExceptionDetails{
		ExceptionText: "This service does not know the operation: " + message,
		ExceptionCode: `OperationNotSupported`,
		LocatorCode:   message,
	}}
}
