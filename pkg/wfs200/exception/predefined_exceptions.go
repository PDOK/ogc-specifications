package exception

import "fmt"

// CannotLockAllFeatures exception
func CannotLockAllFeatures() WFSException {
	return WFSException{
		ExceptionCode: "CannotLockAllFeatures",
	}
}

// DuplicateStoredQueryIDValue exception
func DuplicateStoredQueryIDValue() WFSException {
	return WFSException{
		ExceptionCode: "DuplicateStoredQueryIDValue",
	}
}

// DuplicateStoredQueryParameterName exception
func DuplicateStoredQueryParameterName() WFSException {
	return WFSException{
		ExceptionCode: "DuplicateStoredQueryParameterName",
	}
}

// FeaturesNotLocked exception
func FeaturesNotLocked() WFSException {
	return WFSException{
		ExceptionCode: "FeaturesNotLocked",
	}
}

// InvalidLockID exception
func InvalidLockID() WFSException {
	return WFSException{
		ExceptionCode: "InvalidLockID",
	}
}

// InvalidValue exception
func InvalidValue(s ...string) WFSException {
	if len(s) == 1 {
		return WFSException{ExceptionText: fmt.Sprintf("The parameter: %s, contains a InvalidValue", s[0]),
			ExceptionCode: "InvalidValue",
			LocatorCode:   s[0]}
	}
	return WFSException{
		ExceptionCode: "InvalidValue",
	}
}

// LockHasExpired exception
func LockHasExpired() WFSException {
	return WFSException{
		ExceptionCode: "LockHasExpired",
	}
}

// OperationParsingFailed exception
func OperationParsingFailed(value, locator string) WFSException {
	return WFSException{
		ExceptionText: fmt.Sprintf("Failed to parse the operation, found: %s", value),
		LocatorCode:   locator,
		ExceptionCode: "OperationParsingFailed"}
}

// OperationProcessingFailed exception
func OperationProcessingFailed() WFSException {
	return WFSException{
		ExceptionCode: "OperationProcessingFailed",
	}
}

// ResponseCacheExpired exception
func ResponseCacheExpired() WFSException {
	return WFSException{
		ExceptionCode: "ResponseCacheExpired",
	}
}
