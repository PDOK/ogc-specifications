package wsc110

// Keywords in struct for repeatability
type Keywords struct {
	Keyword []string `xml:"ows:Keyword" yaml:"keyword"`
	Type    *struct {
		Text      string  `xml:",chardata" yaml:"text"`
		CodeSpace *string `xml:"codeSpace,attr,omitempty" yaml:"codeSpace"`
	} `xml:"ows:Type" yaml:"type,omitempty"`
}
