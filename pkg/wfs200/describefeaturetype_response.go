package wfs200

import "encoding/xml"

// Schema struct
// TODO build struct based on the capabilities and additional featureinfo
type Schema struct {
	XMLName                     xml.Name `xml:"schema"`
	Text                        string   `xml:",chardata"`
	Xsd                         string   `xml:"xsd,attr"`
	Digitaaltopografischbestand string   `xml:"digitaaltopografischbestand,attr"`
	Gml                         string   `xml:"gml,attr"`
	Wfs                         string   `xml:"wfs,attr"`
	ElementFormDefault          string   `xml:"elementFormDefault,attr"`
	TargetNamespace             string   `xml:"targetNamespace,attr"`
	Import                      struct {
		Text           string `xml:",chardata"`
		Namespace      string `xml:"namespace,attr"`
		SchemaLocation string `xml:"schemaLocation,attr"`
	} `xml:"import"`
	ComplexType []struct {
		Text           string `xml:",chardata"`
		Name           string `xml:"name,attr"`
		ComplexContent struct {
			Text      string `xml:",chardata"`
			Extension struct {
				Text     string `xml:",chardata"`
				Base     string `xml:"base,attr"`
				Sequence struct {
					Text    string `xml:",chardata"`
					Element []struct {
						Text      string `xml:",chardata"`
						MaxOccurs string `xml:"maxOccurs,attr"`
						MinOccurs string `xml:"minOccurs,attr"`
						Name      string `xml:"name,attr"`
						Nillable  string `xml:"nillable,attr"`
						Type      string `xml:"type,attr"`
					} `xml:"element"`
				} `xml:"sequence"`
			} `xml:"extension"`
		} `xml:"complexContent"`
	} `xml:"complexType"`
	Element []struct {
		Text              string `xml:",chardata"`
		Name              string `xml:"name,attr"`
		SubstitutionGroup string `xml:"substitutionGroup,attr"`
		Type              string `xml:"type,attr"`
	} `xml:"element"`
}
