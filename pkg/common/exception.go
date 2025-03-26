package common

type ExceptionDetails struct {
	ExceptionText string `xml:",chardata" yaml:"exception"`
	ExceptionCode string `xml:"code,attr" yaml:"exceptionCode"`
	LocatorCode   string `xml:"locator,attr,omitempty" yaml:"locatorCode"`
}
