package wfs200

import (
	"encoding/xml"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

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
	OperationsMetadata *OperationsMetadata `xml:"ows:OperationsMetadata" yaml:"operationsMetadata,omitempty"`
	FeatureTypeList    FeatureTypeList     `xml:"FeatureTypeList" yaml:"featureTypeList"`
	FilterCapabilities *FilterCapabilities `xml:"fes:Filter_Capabilities" yaml:"filterCapabilities,omitempty"`
}

// Method in separated struct so to use it as a Pointer
type Method struct {
	Type string `xml:"xlink:type,attr" yaml:"type"`
	Href string `xml:"xlink:href,attr" yaml:"href"`
}

// OperationsMetadata struct for the WFS 2.0.0
type OperationsMetadata struct {
	XMLName              xml.Name              `xml:"ows:OperationsMetadata" yaml:"-"`
	Operation            []Operation           `xml:"ows:Operation" yaml:"operation,omitempty"`
	Parameter            *Parameter            `xml:"ows:Parameter" yaml:"parameter,omitempty"`
	Constraint           []Constraint          `xml:"ows:Constraint" yaml:"constraint,omitempty"`
	ExtendedCapabilities *ExtendedCapabilities `xml:"ows:ExtendedCapabilities" yaml:"extendedCapabilities,omitempty"`
}

// Parameter struct for the WFS 2.0.0
type Parameter struct {
	Name          string         `xml:"name,attr" yaml:"name"`
	AllowedValues *AllowedValues `xml:"ows:AllowedValues" yaml:"allowedValues"`
}

// Constraint struct for the WFS 2.0.0
type Constraint struct {
	Text          string         `xml:",chardata" yaml:"text"`
	Name          string         `xml:"name,attr" yaml:"name"`
	NoValues      *string        `xml:"ows:NoValues" yaml:"noValues"`
	DefaultValue  *string        `xml:"ows:DefaultValue" yaml:"defaultValue"`
	AllowedValues *AllowedValues `xml:"ows:AllowedValues" yaml:"allowedValues"`
}

// ValueConstraint struct for the WFS 2.0.0
type ValueConstraint struct {
	Name         string `xml:"name,attr" yaml:"name"`
	NoValues     string `xml:"ows:NoValues" yaml:"noValues"`
	DefaultValue string `xml:"ows:DefaultValue" yaml:"defaultValue"`
}

// when AllowedValues are defined, NoValues should not be present and vice versa
func (c Constraint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type constraint Constraint // prevent recursion
	x := constraint(c)
	if x.AllowedValues != nil {
		x.NoValues = nil
	} else {
		s := ""
		x.NoValues = &s
	}
	return e.EncodeElement(x, start)
}

// Operation struct for the WFS 2.0.0
type Operation struct {
	Name        string            `xml:"name,attr" yaml:"name"`
	DCP         DCP               `xml:"ows:DCP" yaml:"dcp"`
	Parameter   []Parameter       `xml:"ows:Parameter" yaml:"parameter"`
	Constraints []ValueConstraint `xml:"ows:Constraint" yaml:"constraint"`
}

// DCP struct for the WFS 2.0.0
type DCP struct {
	HTTP HTTP `xml:"ows:HTTP" yaml:"http"`
}

// HTTP struct for the WFS 2.0.0
type HTTP struct {
	Get  *Method `xml:"ows:Get,omitempty" yaml:"get,omitempty"`
	Post *Method `xml:"ows:Post,omitempty" yaml:"post,omitempty"`
}

// AllowedValues struct so it can be used as a pointer
type AllowedValues struct {
	Value []string `xml:"ows:Value" yaml:"value"`
}

// ExtendedCapabilities struct for the WFS 2.0.0
type ExtendedCapabilities struct {
	ExtendedCapabilities NestedExtendedCapabilities `xml:"inspire_dls:ExtendedCapabilities" yaml:"extendedCapabilities"`
}

// NestedExtendedCapabilities struct for the WFS 2.0.0
type NestedExtendedCapabilities struct {
	Text                     *string                  `xml:",chardata" yaml:"text,omitempty"`
	MetadataURL              MetadataURL              `xml:"inspire_common:MetadataUrl" yaml:"metadataUrl"`
	SupportedLanguages       SupportedLanguages       `xml:"inspire_common:SupportedLanguages" yaml:"supportedLanguages"`
	ResponseLanguage         Language                 `xml:"inspire_common:ResponseLanguage" yaml:"responseLanguage"`
	SpatialDataSetIdentifier SpatialDataSetIdentifier `xml:"inspire_dls:SpatialDataSetIdentifier" yaml:"spatialDataSetIdentifier"`
}

// MetadataURL struct for the WFS 2.0.0
type MetadataURL struct {
	URL       string `xml:"inspire_common:URL" yaml:"url"`
	MediaType string `xml:"inspire_common:MediaType" yaml:"mediaType"`
}

// SupportedLanguages struct for the WFS 2.0.0
type SupportedLanguages struct {
	DefaultLanguage   Language    `xml:"inspire_common:DefaultLanguage" yaml:"defaultLanguage"`
	SupportedLanguage *[]Language `xml:"inspire_common:SupportedLanguage" yaml:"supportedLanguage,omitempty"`
}

// Language struct for the WFS 2.0.0
type Language struct {
	Language string `xml:"inspire_common:Language" yaml:"language"`
}

// SpatialDataSetIdentifier struct for the WFS 2.0.0
type SpatialDataSetIdentifier struct {
	Code string `xml:"inspire_common:Code" yaml:"code"`
}

// FeatureTypeList struct for the WFS 2.0.0
type FeatureTypeList struct {
	XMLName     xml.Name      `xml:"FeatureTypeList" yaml:"-"`
	FeatureType []FeatureType `xml:"FeatureType" yaml:"featureType"`
}

// FeatureType struct for the WFS 2.0.0
type FeatureType struct {
	Name             string                   `xml:"Name" yaml:"name"`
	Title            string                   `xml:"Title" yaml:"title"`
	Abstract         string                   `xml:"Abstract" yaml:"abstract"`
	Keywords         *[]wsc110.Keywords       `xml:"ows:Keywords" yaml:"keywords"`
	DefaultCRS       *CRS                     `xml:"DefaultCRS" yaml:"defaultCrs"`
	OtherCRS         []*CRS                   `xml:"OtherCRS" yaml:"otherCrs"`
	OutputFormats    *OutputFormats           `xml:"OutputFormats" yaml:"outputFormats,omitempty"`
	WGS84BoundingBox *wsc110.WGS84BoundingBox `xml:"ows:WGS84BoundingBox" yaml:"wgs84BoundingBox,omitempty"`
	MetadataURL      MetadataHref             `xml:"MetadataURL" yaml:"metadataUrl"`
}

// OutputFormats struct for the WFS 2.0.0
type OutputFormats struct {
	Format []string `xml:"Format" yaml:"format"`
}

// MetadataHref struct for the WFS 2.0.0
type MetadataHref struct {
	Href string `xml:"xlink:href,attr" yaml:"href"`
}

// FilterCapabilities struct for the WFS 2.0.0
type FilterCapabilities struct {
	Conformance         Conformance         `xml:"fes:Conformance" yaml:"conformance"`
	IDCapabilities      IDCapabilities      `xml:"fes:Id_Capabilities" yaml:"idCapabilities"`
	ScalarCapabilities  ScalarCapabilities  `xml:"fes:Scalar_Capabilities" yaml:"scalarCapabilities"`
	SpatialCapabilities SpatialCapabilities `xml:"fes:Spatial_Capabilities" yaml:"spatialCapabilities"`
	// NO TemporalCapabilities!!!
	TemporalCapabilities *TemporalCapabilities `xml:"fes:Temporal_Capabilities" yaml:"temporalCapabilities"`
}

// Conformance struct for the WFS 2.0.0
type Conformance struct {
	Constraint []ValueConstraint `xml:"fes:Constraint" yaml:"constraint"`
}

// IDCapabilities struct for the WFS 2.0.0
type IDCapabilities struct {
	ResourceIdentifier ResourceIdentifier `xml:"fes:ResourceIdentifier" yaml:"resourceIdentifier"`
}

// ScalarCapabilities struct for the WFS 2.0.0
type ScalarCapabilities struct {
	LogicalOperators    string              `xml:"fes:LogicalOperators" yaml:"logicalOperators"`
	ComparisonOperators ComparisonOperators `xml:"fes:ComparisonOperators" yaml:"comparisonOperators"`
}

// ComparisonOperators struct for the WFS 2.0.0
type ComparisonOperators struct {
	ComparisonOperator []ComparisonOperatorName `xml:"fes:ComparisonOperator" yaml:"comparisonOperator"`
}

// ComparisonOperatorName struct for the WFS 2.0.0
type ComparisonOperatorName struct {
	Name string `xml:"name,attr"`
}

// SpatialCapabilities struct for the WFS 2.0.0
type SpatialCapabilities struct {
	GeometryOperands GeometryOperands `xml:"fes:GeometryOperands" yaml:"geometryOperands"`
	SpatialOperators SpatialOperators `xml:"fes:SpatialOperators" yaml:"spatialOperators"`
}

// GeometryOperands struct for the WFS 2.0.0
type GeometryOperands struct {
	GeometryOperand []GeometryOperandName `xml:"fes:GeometryOperand" yaml:"geometryOperand"`
}

// GeometryOperandName struct for the WFS 2.0.0
type GeometryOperandName struct {
	Name string `xml:"name,attr" yaml:"name"`
}

// SpatialOperators struct for the WFS 2.0.0
type SpatialOperators struct {
	SpatialOperator []SpatialOperatorName `xml:"fes:SpatialOperator" yaml:"spatialOperator"`
}

// SpatialOperatorName struct for the WFS 2.0.0
type SpatialOperatorName struct {
	Name string `xml:"name,attr" yaml:"name"`
}

// ResourceIdentifier struct for the WFS 2.0.0
type ResourceIdentifier struct {
	Name string `xml:"name,attr" yaml:"name"`
}

// TemporalCapabilities define but not used
type TemporalCapabilities struct {
	TemporalOperands  TemporalOperands  `xml:"fes:TemporalOperands" yaml:"temporalOperands"`
	TemporalOperators TemporalOperators `xml:"fes:TemporalOperators" yaml:"temporalOperators"`
}

// TemporalOperands struct for the WFS 2.0.0
type TemporalOperands struct {
	TemporalOperand []TemporalOperand `xml:"fes:TemporalOperand" yaml:"temporalOperand"`
}

// TemporalOperand struct for the WFS 2.0.0
type TemporalOperand struct {
	Name string `xml:"name,attr" yaml:"name"`
}

// TemporalOperators  struct for the WFS 2.0.0
type TemporalOperators struct {
	TemporalOperator []TemporalOperator `xml:"fes:TemporalOperator" yaml:"temporalOperator"`
}

// TemporalOperator  struct for the WFS 2.0.0
type TemporalOperator struct {
	Name string `xml:"name,attr,omitempty" yaml:"name,omitempty"`
}
