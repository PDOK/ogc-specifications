package wfs200

import "fmt"

// CannotLockAllFeatures exception
func CannotLockAllFeatures() Exception {
	return exception{
		ExceptionCode: "CannotLockAllFeatures",
	}
}

// DuplicateStoredQueryIDValue exception
func DuplicateStoredQueryIDValue() Exception {
	return exception{
		ExceptionCode: "DuplicateStoredQueryIDValue",
	}
}

// DuplicateStoredQueryParameterName exception
func DuplicateStoredQueryParameterName() Exception {
	return exception{
		ExceptionCode: "DuplicateStoredQueryParameterName",
	}
}

// FeaturesNotLocked exception
func FeaturesNotLocked() Exception {
	return exception{
		ExceptionCode: "FeaturesNotLocked",
	}
}

// InvalidLockID exception
func InvalidLockID() Exception {
	return exception{
		ExceptionCode: "InvalidLockID",
	}
}

// InvalidValue exception
func InvalidValue(s ...string) Exception {
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
func LockHasExpired() Exception {
	return exception{
		ExceptionCode: "LockHasExpired",
	}
}

// OperationParsingFailed exception
func OperationParsingFailed(value, locator string) Exception {
	return exception{
		ExceptionText: fmt.Sprintf("Failed to parse the operation, found: %s", value),
		LocatorCode:   locator,
		ExceptionCode: "OperationParsingFailed"}
}

// OperationProcessingFailed exception
func OperationProcessingFailed() Exception {
	return exception{
		ExceptionCode: "OperationProcessingFailed",
	}
}

// ResponseCacheExpired exception
func ResponseCacheExpired() Exception {
	return exception{
		ExceptionCode: "ResponseCacheExpired",
	}
}
