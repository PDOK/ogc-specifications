package wsc110

// Keywords in struct for repeatability
type Keywords struct {
	Keyword []string `xml:"ows:Keyword"`
	Type    *struct {
		Text      string  `xml:",chardata"`
		CodeSpace *string `xml:"codeSpace,attr,omitempty"`
	} `xml:"ows:Type"`
}
