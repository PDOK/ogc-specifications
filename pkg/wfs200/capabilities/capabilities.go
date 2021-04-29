package capabilities

import (
	"encoding/xml"

	"github.com/pdok/ogc-specifications/pkg/common"
)

// ParseXML func
func (c *Capabilities) ParseXML(doc []byte) error {
	return nil
}

// ParseYAMl func
func (c *Capabilities) ParseYAMl(doc []byte) error {
	return nil
}

// Capabilities struct
type Capabilities struct {
	OperationsMetadata OperationsMetadata `xml:"common:OperationsMetadata" yaml:"operationsmetadata"`
	FeatureTypeList    FeatureTypeList    `xml:"wfs:FeatureTypeList" yaml:"featuretypelist"`
	FilterCapabilities FilterCapabilities `xml:"fes:Filter_Capabilities" yaml:"filtercapabilities"`
}

// Method in separated struct so to use it as a Pointer
type Method struct {
	Type string `xml:"xlink:type,attr" yaml:"type"`
	Href string `xml:"xlink:href,attr" yaml:"href"`
}

// OperationsMetadata struct for the WFS 2.0.0
type OperationsMetadata struct {
	XMLName   xml.Name    `xml:"common:OperationsMetadata"`
	Operation []Operation `xml:"common:Operation"`
	Parameter struct {
		Name          string `xml:"name,attr" yaml:"name"`
		AllowedValues struct {
			Value []string `xml:"common:Value" yaml:"value"`
		} `xml:"common:AllowedValues" yaml:"allowedvalues"`
	} `xml:"common:Parameter" yaml:"parameter"`
	Constraint           []Constraint          `xml:"common:Constraint" yaml:"constraint"`
	ExtendedCapabilities *ExtendedCapabilities `xml:"common:ExtendedCapabilities" yaml:"extendedcapabilities"`
}

// Constraint struct for the WFS 2.0.0
type Constraint struct {
	Text          string         `xml:",chardata"`
	Name          string         `xml:"name,attr" yaml:"name"`
	NoValues      *string        `xml:"common:NoValues" yaml:"novalues"`
	DefaultValue  *string        `xml:"common:DefaultValue" yaml:"defaultvalue"`
	AllowedValues *AllowedValues `xml:"common:AllowedValues" yaml:"allowedvalues"`
}

// Operation struct for the WFS 2.0.0
type Operation struct {
	Name string `xml:"name,attr"`
	DCP  struct {
		HTTP struct {
			Get  *Method `xml:"common:Get,omitempty" yaml:"get,omitempty"`
			Post *Method `xml:"common:Post,omitempty" yaml:"post,omitempty"`
		} `xml:"common:HTTP" yaml:"http"`
	} `xml:"common:DCP" yaml:"dcp"`
	Parameter []struct {
		Name          string `xml:"name,attr"`
		AllowedValues struct {
			Value []string `xml:"common:Value"`
		} `xml:"common:AllowedValues"`
	} `xml:"common:Parameter"`
}

// AllowedValues struct so it can be used as a pointer
type AllowedValues struct {
	Value []string `xml:"common:Value" yaml:"value"`
}

// ExtendedCapabilities struct for the WFS 2.0.0
type ExtendedCapabilities struct {
	ExtendedCapabilities struct {
		Text        string `xml:",chardata"`
		MetadataURL struct {
			Type      string `xml:"xsi:type,attr" yaml:"type"`
			URL       string `xml:"inspire_common:URL" yaml:"url"`
			MediaType string `xml:"inspire_common:MediaType" yaml:"mediatype"`
		} `xml:"inspire_common:MetadataUrl" yaml:"metadataurl"`
		SupportedLanguages struct {
			DefaultLanguage struct {
				Language string `xml:"inspire_common:Language" yaml:"language"`
			} `xml:"inspire_common:DefaultLanguage" yaml:"defaultlanguage"`
		} `xml:"inspire_common:SupportedLanguages" yaml:"supportedlanguages"`
		ResponseLanguage struct {
			Language string `xml:"inspire_common:Language" yaml:"language"`
		} `xml:"inspire_common:ResponseLanguage" yaml:"responselanguage"`
		SpatialDataSetIdentifier struct {
			Code string `xml:"inspire_common:Code" yaml:"code"`
		} `xml:"inspire_dls:SpatialDataSetIdentifier" yaml:"spatialdatasetidentifier"`
	} `xml:"inspire_dls:ExtendedCapabilities" yaml:"extendedcapabilities"`
}

// FeatureTypeList struct for the WFS 2.0.0
type FeatureTypeList struct {
	XMLName     xml.Name      `xml:"wfs:FeatureTypeList"`
	FeatureType []FeatureType `xml:"wfs:FeatureType" yaml:"featuretype"`
}

// FeatureType struct for the WFS 2.0.0
type FeatureType struct {
	Name          string           `xml:"wfs:Name" yaml:"name"`
	Title         string           `xml:"wfs:Title" yaml:"title"`
	Abstract      string           `xml:"wfs:Abstract" yaml:"abstract"`
	Keywords      *common.Keywords `xml:"common:Keywords" yaml:"keywords"`
	DefaultCRS    *common.CRS      `xml:"wfs:DefaultCRS" yaml:"defaultcrs"`
	OtherCRS      *[]common.CRS    `xml:"wfs:OtherCRS" yaml:"othercrs"`
	OutputFormats struct {
		Format []string `xml:"wfs:Format" yaml:"format"`
	} `xml:"wfs:OutputFormats" yaml:"outputformats"`
	WGS84BoundingBox common.BoundingBox `xml:"common:WGS84BoundingBox" yaml:"wgs84boundingbox"`
	MetadataURL      struct {
		Href string `xml:"xlink:href,attr" yaml:"href"`
	} `xml:"wfs:MetadataURL" yaml:"metadataurl"`
}

// FilterCapabilities struct for the WFS 2.0.0
type FilterCapabilities struct {
	Conformance struct {
		Constraint []struct {
			Name         string `xml:"name,attr" yaml:"name"`
			NoValues     string `xml:"common:NoValues" yaml:"novalues"`
			DefaultValue string `xml:"common:DefaultValue" yaml:"defaultvalue"`
		} `xml:"fes:Constraint" yaml:"constraint"`
	} `xml:"fes:Conformance" yaml:"conformance"`
	IDCapabilities struct {
		ResourceIdentifier struct {
			Name string `xml:"name,attr" yaml:"name"`
		} `xml:"fes:ResourceIdentifier" yaml:"resourceidentifier"`
	} `xml:"fes:Id_Capabilities" yaml:"idcapabilities"`
	ScalarCapabilities struct {
		LogicalOperators    string `xml:"fes:LogicalOperators" yaml:"logicaloperators"`
		ComparisonOperators struct {
			ComparisonOperator []struct {
				Name string `xml:"name,attr"`
			} `xml:"fes:ComparisonOperator" yaml:"comparisonoperator"`
		} `xml:"fes:ComparisonOperators" yaml:"comparisonoperators"`
	} `xml:"fes:Scalar_Capabilities" yaml:"scalarcapabilities"`
	SpatialCapabilities struct {
		GeometryOperands struct {
			GeometryOperand []struct {
				Name string `xml:"name,attr"`
			} `xml:"fes:GeometryOperand"`
		} `xml:"fes:GeometryOperands"`
		SpatialOperators struct {
			SpatialOperator []struct {
				Name string `xml:"name,attr"`
			} `xml:"fes:SpatialOperator"`
		} `xml:"fes:SpatialOperators"`
	} `xml:"fes:Spatial_Capabilities"`
	// NO TemporalCapabilities!!!
	TemporalCapabilities *TemporalCapabilities `xml:"fes:Temporal_Capabilities" yaml:"temporalcapabilities"`
}

// TemporalCapabilities define but not used
type TemporalCapabilities struct {
	TemporalOperands struct {
		TemporalOperand []struct {
			Name string `xml:"name,attr" yaml:"name"`
		} `xml:"fes:TemporalOperand" yaml:"temporaloperand"`
	} `xml:"fes:TemporalOperands" yaml:"temporaloperands"`
	TemporalOperators struct {
		TemporalOperator []struct {
			Name string `xml:"name,attr,omitempty" yaml:"name,omitempty"`
		} `xml:"fes:TemporalOperator" yaml:"temporaloperator"`
	} `xml:"fes:TemporalOperators" yaml:"temporaloperators"`
}
