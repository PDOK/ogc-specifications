package wmts100

import "github.com/pdok/ogc-specifications/pkg/wsc110"

// TileOutOfRange exception
func TileOutOfRange() wsc110.Exception {
	return exception{
		ExceptionCode: "TileOutOfRange",
		ExceptionText: "TileRow or TileCol out of rangeName",
		LocatorCode:   "", // TODO parse the right parameter TileRow or TileCol
	}
}

// The four values listed below are copied
// from Table 8 in subclause 7.4.1 of OWS Common [OGC 06-121r3].

func OperationNotSupported(message string) wsc110.Exception {
	return wsc110.OperationNotSupported(message)
}

func MissingParameterValue(s ...string) wsc110.Exception {
	return wsc110.MissingParameterValue(s...)
}

func InvalidParameterValue(value, locator string) wsc110.Exception {
	return wsc110.InvalidParameterValue(value, locator)
}

func NoApplicableCode(message string) wsc110.Exception {
	return wsc110.NoApplicableCode(message)
}
