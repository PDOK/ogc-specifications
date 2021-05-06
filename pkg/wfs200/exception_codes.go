package wfs200

import (
	"fmt"

	"github.com/pdok/ogc-specifications/pkg/common"
)

// CannotLockAllFeatures exception
func CannotLockAllFeatures() common.Exception {
	return exception{
		ExceptionCode: "CannotLockAllFeatures",
	}
}

// DuplicateStoredQueryIDValue exception
func DuplicateStoredQueryIDValue() common.Exception {
	return exception{
		ExceptionCode: "DuplicateStoredQueryIDValue",
	}
}

// DuplicateStoredQueryParameterName exception
func DuplicateStoredQueryParameterName() common.Exception {
	return exception{
		ExceptionCode: "DuplicateStoredQueryParameterName",
	}
}

// FeaturesNotLocked exception
func FeaturesNotLocked() common.Exception {
	return exception{
		ExceptionCode: "FeaturesNotLocked",
	}
}

// InvalidLockID exception
func InvalidLockID() common.Exception {
	return exception{
		ExceptionCode: "InvalidLockID",
	}
}

// InvalidValue exception
func InvalidValue(s ...string) common.Exception {
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
func LockHasExpired() common.Exception {
	return exception{
		ExceptionCode: "LockHasExpired",
	}
}

// OperationParsingFailed exception
func OperationParsingFailed(value, locator string) common.Exception {
	return exception{
		ExceptionText: fmt.Sprintf("Failed to parse the operation, found: %s", value),
		LocatorCode:   locator,
		ExceptionCode: "OperationParsingFailed"}
}

// OperationProcessingFailed exception
func OperationProcessingFailed() common.Exception {
	return exception{
		ExceptionCode: "OperationProcessingFailed",
	}
}

// ResponseCacheExpired exception
func ResponseCacheExpired() common.Exception {
	return exception{
		ExceptionCode: "ResponseCacheExpired",
	}
}
