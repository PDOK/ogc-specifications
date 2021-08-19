package wsc110

// Keywords in struct for repeatability
type Keywords struct {
	Keyword []string `xml:"ows:Keyword" yaml:"keyword"`
	Type    string   `xml:"ows:Type,omitempty" yaml:"type"`
}
