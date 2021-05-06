package common

type Exception interface {
	Error() string
	Code() string
	Locator() string
}

type Exceptions []Exception

type ExceptionReport interface {
	ToBytes() []byte
}
