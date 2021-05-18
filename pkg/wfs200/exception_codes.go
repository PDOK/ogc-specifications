package wfs200

import (
	"fmt"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

// CannotLockAllFeatures exception
func CannotLockAllFeatures() wsc110.Exception {
	return wsc110.Exception{
		ExceptionCode: "CannotLockAllFeatures",
	}
}

// DuplicateStoredQueryIDValue exception
func DuplicateStoredQueryIDValue() wsc110.Exception {
	return wsc110.Exception{
		ExceptionCode: "DuplicateStoredQueryIDValue",
	}
}

// DuplicateStoredQueryParameterName exception
func DuplicateStoredQueryParameterName() wsc110.Exception {
	return wsc110.Exception{
		ExceptionCode: "DuplicateStoredQueryParameterName",
	}
}

// FeaturesNotLocked exception
func FeaturesNotLocked() wsc110.Exception {
	return wsc110.Exception{
		ExceptionCode: "FeaturesNotLocked",
	}
}

// InvalidLockID exception
func InvalidLockID() wsc110.Exception {
	return wsc110.Exception{
		ExceptionCode: "InvalidLockID",
	}
}

// InvalidValue exception
func InvalidValue(s ...string) wsc110.Exception {
	if len(s) == 1 {
		return wsc110.Exception{ExceptionText: fmt.Sprintf("The parameter: %s, contains a InvalidValue", s[0]),
			ExceptionCode: "InvalidValue",
			LocatorCode:   s[0]}
	}
	return wsc110.Exception{
		ExceptionCode: "InvalidValue",
	}
}

// LockHasExpired exception
func LockHasExpired() wsc110.Exception {
	return wsc110.Exception{
		ExceptionCode: "LockHasExpired",
	}
}

// OperationParsingFailed exception
func OperationParsingFailed(value, locator string) wsc110.Exception {
	return wsc110.Exception{
		ExceptionText: fmt.Sprintf("Failed to parse the operation, found: %s", value),
		LocatorCode:   locator,
		ExceptionCode: "OperationParsingFailed"}
}

// OperationProcessingFailed exception
func OperationProcessingFailed() wsc110.Exception {
	return wsc110.Exception{
		ExceptionCode: "OperationProcessingFailed",
	}
}

// ResponseCacheExpired exception
func ResponseCacheExpired() wsc110.Exception {
	return wsc110.Exception{
		ExceptionCode: "ResponseCacheExpired",
	}
}
