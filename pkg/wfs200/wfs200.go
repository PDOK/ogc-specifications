package wfs200

import (
	"encoding/xml"
)

//
const (
	Service = `WFS`
	Version = `2.0.0`
)

// Contains the WFS200 struct

// Service function needed for the interface
func (wfs200 *Wfs200) Service() string {
	return Service
}

// Version function needed for the interface
func (wfs200 *Wfs200) Version() string {
	return Version
}

// Validate function of the wfs200 spec
func (wfs200 *Wfs200) Validate() bool {
	return false
}

// Wfs200 base struct
type Wfs200 struct {
	XMLName               xml.Name `xml:"wfs:WFS_Capabilities"`
	Namespaces            `yaml:"namespaces"`
	ServiceIdentification ServiceIdentification `xml:"ows:ServiceIdentification" yaml:"serviceidentification"`
	ServiceProvider       ServiceProvider       `xml:"ows:ServiceProvider" yaml:"serviceprovider"`
	OperationsMetadata    OperationsMetadata    `xml:"ows:OperationsMetadata" yaml:"operationsmetadata"`
	FeatureTypeList       FeatureTypeList       `xml:"wfs:FeatureTypeList" yaml:"featuretypelist"`
	FilterCapabilities    FilterCapabilities    `xml:"fes:Filter_Capabilities" yaml:"filtercapabilities"`
}

// Namespaces struct containing the namespaces needed for the XML document
type Namespaces struct {
	XmlnsGML           string `xml:"xmlns:gml,attr" yaml:"gml"`                                //http://www.opengis.net/gml/3.2
	XmlnsWFS           string `xml:"xmlns:wfs,attr" yaml:"wfs"`                                //http://www.opengis.net/wfs/2.0
	XmlnsOWS           string `xml:"xmlns:ows,attr" yaml:"ows"`                                //http://www.opengis.net/ows/1.1
	XmlnsXlink         string `xml:"xmlns:xlink,attr" yaml:"xlink"`                            //http://www.w3.org/1999/xlink
	XmlnsXSI           string `xml:"xmlns:xsi,attr" yaml:"xsi"`                                //http://www.w3.org/2001/XMLSchema-instance
	XmlnsFes           string `xml:"xmlns:fes,attr" yaml:"fes"`                                //http://www.opengis.net/fes/2.0
	XmlnsInspireCommon string `xml:"xmlns:inspire_common,attr,omitempty" yaml:"inspirecommon"` //http://inspire.ec.europa.eu/schemas/common/1.0
	XmlnsInspireDls    string `xml:"xmlns:inspire_dls,attr,omitempty" yaml:"inspiredls"`       //http://inspire.ec.europa.eu/schemas/inspire_dls/1.0
	XmlnsPrefix        string `xml:"xmlns:{{.Prefix}},attr" yaml:"prefix"`                     //namespace_uri placeholder
	Version            string `xml:"version,attr" yaml:"version"`
	SchemaLocation     string `xml:"xsi:schemaLocation,attr" yaml:"schemalocation"`
}

// ServiceIdentification struct should only be fill by the "template" configuration wfs200.yaml
type ServiceIdentification struct {
	XMLName     xml.Name `xml:"ows:ServiceIdentification"`
	Title       string   `xml:"ows:Title" yaml:"title"`
	Abstract    string   `xml:"ows:Abstract" yaml:"abstract"`
	Keywords    Keywords `xml:"ows:Keywords" yaml:"keywords"`
	ServiceType struct {
		Text      string `xml:",chardata" yaml:"text"`
		CodeSpace string `xml:"codeSpace,attr" yaml:"codespace"`
	} `xml:"ows:ServiceType"`
	ServiceTypeVersion string `xml:"ows:ServiceTypeVersion" yaml:"servicetypeversion"`
	Fees               string `xml:"ows:Fees" yaml:"fees"`
	AccessConstraints  string `xml:"ows:AccessConstraints" yaml:"accesscontraints"`
}

// ServiceProvider struct containing the provider/organization information should only be fill by the "template" configuration wfs200.yaml
type ServiceProvider struct {
	XMLName      xml.Name `xml:"ows:ServiceProvider"`
	ProviderName string   `xml:"ows:ProviderName" yaml:"providername"`
	ProviderSite struct {
		Type string `xml:"xlink:type,attr" yaml:"type"`
		Href string `xml:"xlink:href,attr" yaml:"href"`
	} `xml:"ows:ProviderSite" yaml:"providersite"`
	ServiceContact struct {
		IndividualName string `xml:"ows:IndividualName" yaml:"individualname"`
		PositionName   string `xml:"ows:PositionName" yaml:"positionname"`
		ContactInfo    struct {
			Text  string `xml:",chardata"`
			Phone struct {
				Voice     string `xml:"ows:Voice" yaml:"voice"`
				Facsimile string `xml:"ows:Facsimile" yaml:"facsmile"`
			} `xml:"ows:Phone" yaml:"phone"`
			Address struct {
				DeliveryPoint         string `xml:"ows:DeliveryPoint" yaml:"deliverypoint"`
				City                  string `xml:"ows:City" yaml:"city"`
				AdministrativeArea    string `xml:"ows:AdministrativeArea" yaml:"administrativearea"`
				PostalCode            string `xml:"ows:PostalCode" yaml:"postalcode"`
				Country               string `xml:"ows:Country" yaml:"country"`
				ElectronicMailAddress string `xml:"ows:ElectronicMailAddress" yaml:"electronicmailaddress"`
			} `xml:"ows:Address" yaml:"address"`
			OnlineResource struct {
				Type string `xml:"xlink:type,attr" yaml:"type"`
				Href string `xml:"xlink:href,attr" yaml:"href"`
			} `xml:"ows:OnlineResource" yaml:"onlineresource"`
			HoursOfService      string `xml:"ows:HoursOfService" yaml:"hoursofservice"`
			ContactInstructions string `xml:"ows:ContactInstructions" yaml:"contactinstructions"`
		} `xml:"ows:ContactInfo" yaml:"contactinfo"`
		Role string `xml:"ows:Role" yaml:"role"`
	} `xml:"ows:ServiceContact" yaml:"servicecontact"`
}

// Method in separated struct so to use it as a Pointer
type Method struct {
	Type string `xml:"xlink:type,attr" yaml:"type"`
	Href string `xml:"xlink:href,attr" yaml:"href"`
}

// OperationsMetadata struct for the WFS 2.0.0
type OperationsMetadata struct {
	XMLName   xml.Name    `xml:"ows:OperationsMetadata"`
	Operation []Operation `xml:"ows:Operation"`
	Parameter struct {
		Name          string `xml:"name,attr" yaml:"name"`
		AllowedValues struct {
			Value []string `xml:"ows:Value" yaml:"value"`
		} `xml:"ows:AllowedValues" yaml:"allowedvalues"`
	} `xml:"ows:Parameter" yaml:"parameter"`
	Constraint           []Constraint          `xml:"ows:Constraint" yaml:"constraint"`
	ExtendedCapabilities *ExtendedCapabilities `xml:"ows:ExtendedCapabilities" yaml:"extendedcapabilities"`
}

// Constraint struct for the WFS 2.0.0
type Constraint struct {
	Text          string         `xml:",chardata"`
	Name          string         `xml:"name,attr" yaml:"name"`
	NoValues      *string        `xml:"ows:NoValues" yaml:"novalue"`
	DefaultValue  *string        `xml:"ows:DefaultValue" yaml:"defaultvalue"`
	AllowedValues *AllowedValues `xml:"ows:AllowedValues" yaml:"allowedvalues"`
}

// Operation struct for the WFS 2.0.0
type Operation struct {
	Name string `xml:"name,attr"`
	DCP  struct {
		HTTP struct {
			Get  *Method `xml:"ows:Get,omitempty" yaml:"get"`
			Post *Method `xml:"ows:Post,omitempty" yaml:"post"`
		} `xml:"ows:HTTP" yaml:"http"`
	} `xml:"ows:DCP" yaml:"dcp"`
	Parameter []struct {
		Name          string `xml:"name,attr"`
		AllowedValues struct {
			Value []string `xml:"ows:Value"`
		} `xml:"ows:AllowedValues"`
	} `xml:"ows:Parameter"`
}

// AllowedValues struct so it can be used as a pointer
type AllowedValues struct {
	Value []string `xml:"ows:Value" yaml:"value"`
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
	Name          string    `xml:"wfs:Name" yaml:"name"`
	Title         string    `xml:"wfs:Title" yaml:"title"`
	Abstract      string    `xml:"wfs:Abstract" yaml:"abstract"`
	Keywords      *Keywords `xml:"ows:Keywords" yaml:"keywords"`
	DefaultCRS    *string   `xml:"wfs:DefaultCRS" yaml:"defaultcrs"`
	OtherCRS      *[]string `xml:"wfs:OtherCRS" yaml:"othercrs"`
	OutputFormats struct {
		Format []string `xml:"wfs:Format" yaml:"format"`
	} `xml:"wfs:OutputFormats" yaml:"outputformats"`
	WGS84BoundingBox struct {
		Dimensions  string `xml:"dimensions,attr" yaml:"dimensions"`
		LowerCorner string `xml:"ows:LowerCorner" yaml:"lowercorner"`
		UpperCorner string `xml:"ows:UpperCorner" yaml:"uppercorner"`
	} `xml:"ows:WGS84BoundingBox" yaml:"wgs84boundingbox"`
	MetadataURL struct {
		Href string `xml:"xlink:href,attr" yaml:"href"`
	} `xml:"wfs:MetadataURL" yaml:"metadataurl"`
}

// FilterCapabilities struct for the WFS 2.0.0
type FilterCapabilities struct {
	Conformance struct {
		Constraint []struct {
			Name         string `xml:"name,attr" yaml:"name"`
			NoValues     string `xml:"ows:NoValues" yaml:"novalues"`
			DefaultValue string `xml:"ows:DefaultValue" yaml:"defaultvalue"`
		} `xml:"fes:Constraint" yaml:"constraint"`
	} `xml:"fes:Conformance" yaml:"conformance"`
	IDCapabilities struct {
		ResourceIdentifier struct {
			Name string `xml:"name,attr" yaml:"name" `
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
			Name string `xml:"name,attr,omitempty" yaml:"name"`
		} `xml:"fes:TemporalOperator" yaml:"temporaloperator"`
	} `xml:"fes:TemporalOperators" yaml:"temporaloperators"`
}

// Keywords in struct for repeatablity
type Keywords struct {
	Keyword []string `xml:"ows:Keyword" yaml:"keyword"`
}
