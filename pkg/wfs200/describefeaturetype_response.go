package wfs200

import "encoding/xml"

// Schema struct
// TODO build struct based on the capabilities and additional featureinfo
type Schema struct {
	XMLName                     xml.Name `xml:"schema" json:"schema"`
	Text                        string   `xml:",chardata" json:"text"`
	Xsd                         string   `xml:"xsd,attr" json:"xsd"`
	Digitaaltopografischbestand string   `xml:"digitaaltopografischbestand,attr" json:"digitaaltopografischbestandestand"`
	Gml                         string   `xml:"gml,attr" json:"gml"`
	Wfs                         string   `xml:"wfs,attr" json:"wfs"`
	ElementFormDefault          string   `xml:"elementFormDefault,attr" json:"elementFormDefault"`
	TargetNamespace             string   `xml:"targetNamespace,attr" json:"targetNamespace"`
	Import                      struct {
		Text           string `xml:",chardata" json:"text"`
		Namespace      string `xml:"namespace,attr" json:"namespace"`
		SchemaLocation string `xml:"schemaLocation,attr" json:"schemaLocation"`
	} `xml:"import" json:"import"`
	ComplexType []struct {
		Text           string `xml:",chardata" json:"text"`
		Name           string `xml:"name,attr" json:"name"`
		ComplexContent struct {
			Text      string `xml:",chardata" json:"text"`
			Extension struct {
				Text     string `xml:",chardata" json:"text"`
				Base     string `xml:"base,attr" json:"base"`
				Sequence struct {
					Text    string `xml:",chardata" json:"text"`
					Element []struct {
						Text      string `xml:",chardata" json:"text"`
						MaxOccurs string `xml:"maxOccurs,attr" json:"maxOccurs"`
						MinOccurs string `xml:"minOccurs,attr" json:"minOccurs"`
						Name      string `xml:"name,attr" json:"name"`
						Nillable  string `xml:"nillable,attr" json:"nillable"`
						Type      string `xml:"type,attr" json:"type"`
					} `xml:"element" json:"element"`
				} `xml:"sequence" json:"sequence"`
			} `xml:"extension" json:"extension"`
		} `xml:"complexContent" json:"complexContent"`
	} `xml:"complexType" json:"complexType"`
	Element []struct {
		Text              string `xml:",chardata" json:"text"`
		Name              string `xml:"name,attr" json:"name"`
		SubstitutionGroup string `xml:"substitutionGroup,attr" json:"substitutionGroup"`
		Type              string `xml:"type,attr" json:"type"`
	} `xml:"element" json:"element"`
}
