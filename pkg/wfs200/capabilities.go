package wfs200

import (
	"encoding/xml"

	"github.com/pdok/ogc-specifications/pkg/wsc110"
)

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
	OperationsMetadata OperationsMetadata `xml:"ows:OperationsMetadata" yaml:"operationsMetadata"`
	FeatureTypeList    FeatureTypeList    `xml:"FeatureTypeList" yaml:"featureTypeList"`
	FilterCapabilities FilterCapabilities `xml:"fes:Filter_Capabilities" yaml:"filterCapabilities"`
}

// Method in separated struct so to use it as a Pointer
type Method struct {
	Type string `xml:"xlink:type,attr" yaml:"type"`
	Href string `xml:"xlink:href,attr" yaml:"href"`
}

// OperationsMetadata struct for the WFS 2.0.0
type OperationsMetadata struct {
	XMLName   xml.Name    `xml:"ows:OperationsMetadata" yaml:"operationsMetadata"`
	Operation []Operation `xml:"ows:Operation" yaml:"operation"`
	Parameter struct {
		Name          string         `xml:"name,attr" yaml:"name"`
		AllowedValues *AllowedValues `xml:"ows:AllowedValues" yaml:"allowedValues"`
	} `xml:"ows:Parameter" yaml:"parameter"`
	Constraint           []Constraint          `xml:"ows:Constraint" yaml:"constraint"`
	ExtendedCapabilities *ExtendedCapabilities `xml:"ows:ExtendedCapabilities" yaml:"extendedCapabilities"`
}

// Constraint struct for the WFS 2.0.0
type Constraint struct {
	Text          string         `xml:",chardata" yaml:"text"`
	Name          string         `xml:"name,attr" yaml:"name"`
	NoValues      *string        `xml:"ows:NoValues" yaml:"noValues"`
	DefaultValue  *string        `xml:"ows:DefaultValue" yaml:"defaultValue"`
	AllowedValues *AllowedValues `xml:"ows:AllowedValues" yaml:"allowedValues"`
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
	Name string `xml:"name,attr" yaml:"name"`
	DCP  struct {
		HTTP struct {
			Get  *Method `xml:"ows:Get,omitempty" yaml:"get,omitempty"`
			Post *Method `xml:"ows:Post,omitempty" yaml:"post,omitempty"`
		} `xml:"ows:HTTP" yaml:"http"`
	} `xml:"ows:DCP" yaml:"dcp"`
	Parameter []struct {
		Name          string `xml:"name,attr" yaml:"name"`
		AllowedValues struct {
			Value []string `xml:"ows:Value" yaml:"value"`
		} `xml:"ows:AllowedValues" yaml:"allowedValues"`
	} `xml:"ows:Parameter" yaml:"parameter"`
	Constraints []struct {
		Name         string `xml:"name,attr" yaml:"name"`
		NoValues     string `xml:"ows:NoValues" yaml:"noValues"`
		DefaultValue string `xml:"ows:DefaultValue" yaml:"defaultValue"`
	} `xml:"ows:Constraint" yaml:"constraint"`
}

// AllowedValues struct so it can be used as a pointer
type AllowedValues struct {
	Value []string `xml:"ows:Value" yaml:"value"`
}

// ExtendedCapabilities struct for the WFS 2.0.0
type ExtendedCapabilities struct {
	ExtendedCapabilities struct {
		Text        string `xml:",chardata" yaml:"text"`
		MetadataURL struct {
			URL       string `xml:"inspire_common:URL" yaml:"url"`
			MediaType string `xml:"inspire_common:MediaType" yaml:"mediaType"`
		} `xml:"inspire_common:MetadataUrl" yaml:"metadataUrl"`
		SupportedLanguages struct {
			DefaultLanguage struct {
				Language string `xml:"inspire_common:Language" yaml:"language"`
			} `xml:"inspire_common:DefaultLanguage" yaml:"defaultLanguage"`
			SupportedLanguage *[]struct {
				Language string `xml:"inspire_common:Language" yaml:"language"`
			} `xml:"inspire_common:SupportedLanguage" yaml:"supportedLanguage"`
		} `xml:"inspire_common:SupportedLanguages" yaml:"supportedLanguages"`
		ResponseLanguage struct {
			Language string `xml:"inspire_common:Language" yaml:"language"`
		} `xml:"inspire_common:ResponseLanguage" yaml:"responseLanguage"`
		SpatialDataSetIdentifier struct {
			Code string `xml:"inspire_common:Code" yaml:"code"`
		} `xml:"inspire_dls:SpatialDataSetIdentifier" yaml:"spatialDataSetIdentifier"`
	} `xml:"inspire_dls:ExtendedCapabilities" yaml:"extendedCapabilities"`
}

// FeatureTypeList struct for the WFS 2.0.0
type FeatureTypeList struct {
	XMLName     xml.Name      `xml:"FeatureTypeList" yaml:"featureTypeList"`
	FeatureType []FeatureType `xml:"FeatureType" yaml:"featureType"`
}

// FeatureType struct for the WFS 2.0.0
type FeatureType struct {
	Name          string             `xml:"Name" yaml:"name"`
	Title         string             `xml:"Title" yaml:"title"`
	Abstract      string             `xml:"Abstract" yaml:"abstract"`
	Keywords      *[]wsc110.Keywords `xml:"ows:Keywords" yaml:"keywords"`
	DefaultCRS    *CRS               `xml:"DefaultCRS" yaml:"defaultCrs"`
	OtherCRS      *[]CRS             `xml:"OtherCRS" yaml:"otherCrs"`
	OutputFormats struct {
		Format []string `xml:"Format" yaml:"format"`
	} `xml:"OutputFormats" yaml:"outputFormats"`
	WGS84BoundingBox *wsc110.WGS84BoundingBox `xml:"ows:WGS84BoundingBox" yaml:"wgs84BoundingBox"`
	MetadataURL      struct {
		Href string `xml:"xlink:href,attr" yaml:"href"`
	} `xml:"MetadataURL" yaml:"metadataUrl"`
}

// FilterCapabilities struct for the WFS 2.0.0
type FilterCapabilities struct {
	Conformance struct {
		Constraint []struct {
			Name         string `xml:"name,attr" yaml:"name"`
			NoValues     string `xml:"ows:NoValues" yaml:"noValues"`
			DefaultValue string `xml:"ows:DefaultValue" yaml:"defaultValue"`
		} `xml:"fes:Constraint" yaml:"constraint"`
	} `xml:"fes:Conformance" yaml:"conformance"`
	IDCapabilities struct {
		ResourceIdentifier struct {
			Name string `xml:"name,attr" yaml:"name"`
		} `xml:"fes:ResourceIdentifier" yaml:"resourceIdentifier"`
	} `xml:"fes:Id_Capabilities" yaml:"idCapabilities"`
	ScalarCapabilities struct {
		LogicalOperators    string `xml:"fes:LogicalOperators" yaml:"logicalOperators"`
		ComparisonOperators struct {
			ComparisonOperator []struct {
				Name string `xml:"name,attr"`
			} `xml:"fes:ComparisonOperator" yaml:"comparisonOperator"`
		} `xml:"fes:ComparisonOperators" yaml:"comparisonOperators"`
	} `xml:"fes:Scalar_Capabilities" yaml:"scalarCapabilities"`
	SpatialCapabilities struct {
		GeometryOperands struct {
			GeometryOperand []struct {
				Name string `xml:"name,attr" yaml:"name"`
			} `xml:"fes:GeometryOperand" yaml:"geometryOperand"`
		} `xml:"fes:GeometryOperands" yaml:"geometryOperands"`
		SpatialOperators struct {
			SpatialOperator []struct {
				Name string `xml:"name,attr" yaml:"name"`
			} `xml:"fes:SpatialOperator" yaml:"spatialOperator"`
		} `xml:"fes:SpatialOperators" yaml:"spatialOperators"`
	} `xml:"fes:Spatial_Capabilities" yaml:"spatialCapabilities"`
	// NO TemporalCapabilities!!!
	TemporalCapabilities *TemporalCapabilities `xml:"fes:Temporal_Capabilities" yaml:"temporalCapabilities"`
}

// TemporalCapabilities define but not used
type TemporalCapabilities struct {
	TemporalOperands struct {
		TemporalOperand []struct {
			Name string `xml:"name,attr" yaml:"name"`
		} `xml:"fes:TemporalOperand" yaml:"temporalOperand"`
	} `xml:"fes:TemporalOperands" yaml:"temporalOperands"`
	TemporalOperators struct {
		TemporalOperator []struct {
			Name string `xml:"name,attr,omitempty" yaml:"name,omitempty"`
		} `xml:"fes:TemporalOperator" yaml:"temporalOperator"`
	} `xml:"fes:TemporalOperators" yaml:"temporalOperators"`
}
