package wcs201

// ParseXML func
func (c *Capabilities) ParseXML(_ []byte) error {
	return nil
}

// ParseYAML func
func (c *Capabilities) ParseYAML(_ []byte) error {
	return nil
}

// Capabilities struct
type Capabilities struct {
	OperationsMetadata OperationsMetadata `xml:"ows:OperationsMetadata" yaml:"operationsMetadata"`
	ServiceMetadata    ServiceMetadata    `xml:"wcs:ServiceMetadata" yaml:"serviceMetadata"`
	Contents           Contents           `xml:"wcs:Contents" yaml:"contents"`
}

// OperationsMetadata struct for the WCS 2.0.1
type OperationsMetadata struct {
	Operation            []Operation           `xml:"ows:Operation" yaml:"operation"`
	ExtendedCapabilities *ExtendedCapabilities `xml:"ows:ExtendedCapabilities" yaml:"extendedCapabilities"`
}

// Operation in struct for repeatability
type Operation struct {
	Name string `xml:"name,attr" yaml:"name"`
	DCP  DCP    `xml:"ows:DCP" yaml:"dcp"`
}

// DCP struct for the WCS 2.0.1
type DCP struct {
	HTTP HTTP `xml:"ows:HTTP"  yaml:"http"`
}

// HTTP struct for the WCS 2.0.1
type HTTP struct {
	Get  Get   `xml:"ows:Get" yaml:"get"`
	Post *Post `xml:"ows:Post" yaml:"post"`
}

// Get struct for the WCS 2.0.1
type Get struct {
	Type string `xml:"xlink:type,attr" yaml:"type"`
	Href string `xml:"xlink:href,attr" yaml:"href"`
}

// Post in separated struct so to use it as a Pointer
type Post struct {
	Type       string     `xml:"xlink:type,attr" yaml:"type"`
	Href       string     `xml:"xlink:href,attr" yaml:"href"`
	Constraint Constraint `xml:"ows:Constraint" yaml:"constraint"`
}

// Constraint struct for the WCS 2.0.1
type Constraint struct {
	Name          string        `xml:"name,attr" yaml:"name"`
	AllowedValues AllowedValues `xml:"ows:AllowedValues" yaml:"allowedValues"`
}

// AllowedValues struct for the WCS 2.0.1
type AllowedValues struct {
	Value []string `xml:"ows:Value" yaml:"value"`
}

// ExtendedCapabilities struct for the WCS 2.0.1
type ExtendedCapabilities struct {
	ExtendedCapabilities NestedExtendedCapabilities `xml:"inspire_dls:ExtendedCapabilities" yaml:"extendedCapabilities"`
}

// NestedExtendedCapabilities struct for the WCS 2.0.1
type NestedExtendedCapabilities struct {
	MetadataURL              MetadataURL              `xml:"inspire_common:MetadataUrl"`
	ResponseLanguage         Language                 `xml:"inspire_common:ResponseLanguage" yaml:"responseLanguage"`
	SpatialDataSetIdentifier SpatialDataSetIdentifier `xml:"inspire_dls:SpatialDataSetIdentifier" yaml:"spatialDataSetIdentifier"`
}

// MetadataURL struct { struct for the WCS 2.0.1
type MetadataURL struct {
	URL       string `xml:"inspire_common:URL" yaml:"url"`
	MediaType string `xml:"inspire_common:MediaType" yaml:"mediaType"`
}

// SupportedLanguages struct for the struct for the WCS 2.0.1
type SupportedLanguages struct {
	DefaultLanguage   Language    `xml:"inspire_common:DefaultLanguage" yaml:"defaultLanguage"`
	SupportedLanguage *[]Language `xml:"inspire_common:SupportedLanguage" yaml:"supportedLanguage"`
}

// Language struct for the WCS 2.0.1
type Language struct {
	Language string `xml:"inspire_common:Language" yaml:"language"`
}

// SpatialDataSetIdentifier struct for the WCS 2.0.1
type SpatialDataSetIdentifier struct {
	Code string `xml:"Code" yaml:"code"`
}

// ServiceMetadata struct for the WCS 2.0.1
type ServiceMetadata struct {
	FormatSupported []string  `xml:"wcs:formatSupported" yaml:"formatSupported"`
	Extension       Extension `xml:"wcs:Extension" yaml:"extension"`
}

// Extension struct for the WCS 2.0.1
type Extension struct {
	InterpolationMetadata InterpolationMetadata `xml:"int:InterpolationMetadata" yaml:"interpolationMetadata"`
	CrsMetadata           CrsMetadata           `xml:"crs:CrsMetadata" yaml:"crsMetadata"`
}

// InterpolationMetadata struct for the WCS 2.0.1
type InterpolationMetadata struct {
	InterpolationSupported []string `xml:"int:InterpolationSupported" yaml:"interpolationSupported"`
}

// CrsMetadata struct for the WCS 2.0.1
type CrsMetadata struct {
	CrsSupported []string `xml:"crs:crsSupported" yaml:"crsSupported"`
}

// Contents in struct for repeatability
type Contents struct {
	CoverageSummary []CoverageSummary `xml:"wcs:CoverageSummary" yaml:"coverageSummary"`
}

// CoverageSummary in struct for repeatability
type CoverageSummary struct {
	CoverageID      string `xml:"wcs:CoverageId" yaml:"coverageId"`
	CoverageSubtype string `xml:"wcs:CoverageSubtype" yaml:"coverageSubtype"`
}
