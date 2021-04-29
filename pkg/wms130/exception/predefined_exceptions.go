package exception

import "fmt"

// InvalidFormat exception
func InvalidFormat(unknownformat string) WMSException {
	return WMSException{
		ExceptionText: fmt.Sprintf("The format: %s, is a invalid image format", unknownformat),
		ExceptionCode: `InvalidFormat`,
	}
}

// InvalidCRS exception
func InvalidCRS(s ...string) WMSException {
	if len(s) == 1 {
		return WMSException{
			ExceptionText: fmt.Sprintf("CRS is not known by this service: %s", s[0]),
			ExceptionCode: `InvalidCRS`,
		}
	}
	if len(s) == 2 {
		return WMSException{
			ExceptionText: fmt.Sprintf("The CRS: %s is not known by the layer: %s", s[0], s[1]),
			ExceptionCode: `InvalidCRS`,
		}
	}
	return WMSException{
		ExceptionCode: `InvalidCRS`,
	}
}

// LayerNotDefined exception
func LayerNotDefined(s ...string) WMSException {
	if len(s) == 1 {
		return WMSException{
			ExceptionText: fmt.Sprintf("The layer: %s is not known by the server", s[0]),
			ExceptionCode: `LayerNotDefined`,
		}
	}
	return WMSException{
		ExceptionCode: `LayerNotDefined`,
	}
}

// StyleNotDefined exception
func StyleNotDefined(s ...string) WMSException {
	if len(s) == 2 {
		return WMSException{
			ExceptionText: fmt.Sprintf("The style: %s is not known by the server for the layer: %s", s[0], s[1]),
			ExceptionCode: `StyleNotDefined`,
		}
	}
	return WMSException{
		ExceptionText: `There is a one-to-one correspondence between the values in the LAYERS parameter and the values in the STYLES parameter. 
	Expecting an empty string for the STYLES like STYLES= or comma-separated list STYLES=,,, or using keyword default STYLES=default,default,...`,
		ExceptionCode: `StyleNotDefined`,
	}
}

// LayerNotQueryable exception
func LayerNotQueryable(s ...string) WMSException {
	if len(s) == 1 {
		return WMSException{
			ExceptionText: fmt.Sprintf("Layer: %s, can not be queried", s[0]),
			ExceptionCode: `LayerNotQueryable`,
			LocatorCode:   s[0],
		}
	}
	return WMSException{
		ExceptionCode: `LayerNotQueryable`,
	}
}

// InvalidPoint exception
// i and j are strings so we can return none integer values in the exception
func InvalidPoint(i, j string) WMSException {
	// TODO provide giving WIDTH and HEIGHT values in Exception response
	return WMSException{
		ExceptionText: fmt.Sprintf("The parameters I and J are invalid, given: %s for I and %s for J", i, j),
		ExceptionCode: `InvalidPoint`,
	}
}

// CurrentUpdateSequence exception
func CurrentUpdateSequence() WMSException {
	return WMSException{
		ExceptionCode: `CurrentUpdateSequence`,
	}
}

// InvalidUpdateSequence exception
func InvalidUpdateSequence() WMSException {
	return WMSException{
		ExceptionCode: `InvalidUpdateSequence`,
	}
}

// MissingDimensionValue exception
func MissingDimensionValue() WMSException {
	return WMSException{
		ExceptionCode: `MissingDimensionValue`,
	}
}

// InvalidDimensionValue exception
func InvalidDimensionValue() WMSException {
	return WMSException{
		ExceptionCode: `InvalidDimensionValue`,
	}
}
