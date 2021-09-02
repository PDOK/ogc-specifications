package wcs201

// ParseXML func
func (c *Capabilities) ParseXML(doc []byte) error {
	return nil
}

// ParseYAML func
func (c *Capabilities) ParseYAML(doc []byte) error {
	return nil
}

// Capabilities struct
type Capabilities struct {
	OperationsMetadata OperationsMetadata `xml:"ows:OperationsMetadata" yaml:"operationsmetadata"`
	ServiceMetadata    ServiceMetadata    `xml:"wcs:ServiceMetadata" yaml:"servicemetadata"`
	Contents           Contents           `xml:"wcs:Contents" yaml:"contents"`
}

// OperationsMetadata struct for the WCS 2.0.1
type OperationsMetadata struct {
	Operation            []Operation           `xml:"ows:Operation" yaml:"operation"`
	ExtendedCapabilities *ExtendedCapabilities `xml:"ows:ExtendedCapabilities" yaml:"extendedcapabilities"`
}

// Operation in struct for repeatability
type Operation struct {
	Name string `xml:"name,attr" yaml:"name"`
	DCP  struct {
		HTTP struct {
			Get struct {
				Type string `xml:"xlink:type,attr" yaml:"type"`
				Href string `xml:"xlink:href,attr" yaml:"href"`
			} `xml:"ows:Get" yaml:"get"`
			Post *Post `xml:"ows:Post" yaml:"post"`
		} `xml:"ows:HTTP"  yaml:"http"`
	} `xml:"ows:DCP" yaml:"dcp"`
}

// Post in separated struct so to use it as a Pointer
type Post struct {
	Type       string `xml:"xlink:type,attr" yaml:"type"`
	Href       string `xml:"xlink:href,attr" yaml:"href"`
	Constraint struct {
		Name          string `xml:"name,attr" yaml:"name"`
		AllowedValues struct {
			Value []string `xml:"ows:Value" yaml:"value"`
		} `xml:"ows:AllowedValues" yaml:"allowedvalues"`
	} `xml:"ows:Constraint" yaml:"constraint"`
}

// ExtendedCapabilities struct for the WCS 2.0.1
type ExtendedCapabilities struct {
	ExtendedCapabilities struct {
		MetadataURL struct {
			URL       string `xml:"inspire_common:URL"`
			MediaType string `xml:"inspire_common:MediaType"`
		} `xml:"inspire_common:MetadataUrl"`
		SupportedLanguages struct {
			DefaultLanguage struct {
				Language string `xml:"inspire_common:Language"`
			} `xml:"inspire_common:DefaultLanguage"`
			SupportedLanguage *[]struct {
				Language string `xml:"inspire_common:Language" yaml:"language"`
			} `xml:"inspire_common:SupportedLanguage" yaml:"supportedlanguage"`
		} `xml:"inspire_common:SupportedLanguages"`
		ResponseLanguage struct {
			Language string `xml:"inspire_common:Language"`
		} `xml:"inspire_common:ResponseLanguage"`
		SpatialDataSetIdentifier struct {
			Code string `xml:"Code"`
		} `xml:"inspire_dls:SpatialDataSetIdentifier"`
	} `xml:"inspire_dls:ExtendedCapabilities"`
}

// ServiceMetadata struct for the WCS 2.0.1
type ServiceMetadata struct {
	FormatSupported []string `xml:"wcs:formatSupported"`
	Extension       struct {
		InterpolationMetadata struct {
			InterpolationSupported []string `xml:"int:InterpolationSupported"`
		} `xml:"int:InterpolationMetadata"`
		CrsMetadata struct {
			CrsSupported []string `xml:"crs:crsSupported"`
		} `xml:"crs:CrsMetadata"`
	} `xml:"wcs:Extension"`
}

// Contents in struct for repeatability
type Contents struct {
	CoverageSummary []CoverageSummary `xml:"wcs:CoverageSummary"`
}

// CoverageSummary in struct for repeatability
type CoverageSummary struct {
	CoverageID      string `xml:"wcs:CoverageId"`
	CoverageSubtype string `xml:"wcs:CoverageSubtype"`
}
