package wfs200

import (
	"fmt"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

// CannotLockAllFeatures exception
func CannotLockAllFeatures() wsc110.Exception {
	return exception{
		ExceptionCode: "CannotLockAllFeatures",
	}
}

// DuplicateStoredQueryIDValue exception
func DuplicateStoredQueryIDValue() wsc110.Exception {
	return exception{
		ExceptionCode: "DuplicateStoredQueryIDValue",
	}
}

// DuplicateStoredQueryParameterName exception
func DuplicateStoredQueryParameterName() wsc110.Exception {
	return exception{
		ExceptionCode: "DuplicateStoredQueryParameterName",
	}
}

// FeaturesNotLocked exception
func FeaturesNotLocked() wsc110.Exception {
	return exception{
		ExceptionCode: "FeaturesNotLocked",
	}
}

// InvalidLockID exception
func InvalidLockID() wsc110.Exception {
	return exception{
		ExceptionCode: "InvalidLockID",
	}
}

// InvalidValue exception
func InvalidValue(s ...string) wsc110.Exception {
	if len(s) == 1 {
		return exception{ExceptionText: fmt.Sprintf("The parameter: %s, contains a InvalidValue", s[0]),
			ExceptionCode: "InvalidValue",
			LocatorCode:   s[0]}
	}
	return exception{
		ExceptionCode: "InvalidValue",
	}
}

// LockHasExpired exception
func LockHasExpired() wsc110.Exception {
	return exception{
		ExceptionCode: "LockHasExpired",
	}
}

// OperationParsingFailed exception
func OperationParsingFailed(value, locator string) wsc110.Exception {
	return exception{
		ExceptionText: "Failed to parse the operation, found: " + value,
		LocatorCode:   locator,
		ExceptionCode: "OperationParsingFailed"}
}

// OperationProcessingFailed exception
func OperationProcessingFailed() wsc110.Exception {
	return exception{
		ExceptionCode: "OperationProcessingFailed",
	}
}

// ResponseCacheExpired exception
func ResponseCacheExpired() wsc110.Exception {
	return exception{
		ExceptionCode: "ResponseCacheExpired",
	}
}
