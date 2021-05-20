package wms130

import (
	"net/url"
	"strconv"
	"strings"
)

//GetFeatureInfoKVP struct
type GetFeatureInfoKVP struct {
	// Table 8 - The Parameters of a GetFeatureInfo request
	Service string `yaml:"service,omitempty"`
	BaseRequestKVP
	GetMapKVPMandatory
	GetFeatureInfoKVPMandatory
	GetFeatureInfoKVPOptional
}

// GetFeatureInfoKVPMandatory struct containing the mandatory WMS request KVP
type GetFeatureInfoKVPMandatory struct {
	QueryLayers string `yaml:"query_layers,omitempty"`
	InfoFormat  string `yaml:"info_format,omitempty"`
	I           string `yaml:"i,omitempty"`
	J           string `yaml:"j,omitempty"`
}

// GetFeatureInfoKVPOptional struct containing the optional WMS request KVP
type GetFeatureInfoKVPOptional struct {
	FeatureCount *string `yaml:"feature_count,omitempty"`
	Exceptions   *string `yaml:"exceptions,omitempty"`
}

// ParseKVP builds a GetMapKVP object based on the available query parameters
func (gfikvp *GetFeatureInfoKVP) ParseQueryParameters(query url.Values) Exceptions {
	var exceptions Exceptions
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, InvalidParameterValue(k, strings.Join(v, ",")))
		} else {
			switch strings.ToUpper(k) {
			case SERVICE:
				gfikvp.Service = strings.ToUpper(v[0])
			case VERSION:
				gfikvp.BaseRequestKVP.Version = v[0]
			case REQUEST:
				gfikvp.BaseRequestKVP.Request = v[0]
			case LAYERS:
				gfikvp.GetMapKVPMandatory.Layers = v[0]
			case STYLES:
				gfikvp.GetMapKVPMandatory.Styles = v[0]
			case "CRS":
				gfikvp.GetMapKVPMandatory.CRS = v[0]
			case BBOX:
				gfikvp.GetMapKVPMandatory.Bbox = v[0]
			case WIDTH:
				gfikvp.GetMapKVPMandatory.Width = v[0]
			case HEIGHT:
				gfikvp.GetMapKVPMandatory.Height = v[0]
			case FORMAT:
				gfikvp.GetMapKVPMandatory.Format = v[0]
			case QUERYLAYERS:
				gfikvp.GetFeatureInfoKVPMandatory.QueryLayers = v[0]
			case INFOFORMAT:
				gfikvp.GetFeatureInfoKVPMandatory.InfoFormat = v[0]
			case I:
				gfikvp.GetFeatureInfoKVPMandatory.I = v[0]
			case J:
				gfikvp.GetFeatureInfoKVPMandatory.J = v[0]
			case FEATURECOUNT:
				gfikvp.GetFeatureInfoKVPOptional.FeatureCount = &(v[0])
			case EXCEPTIONS:
				gfikvp.GetFeatureInfoKVPOptional.Exceptions = &(v[0])
			}
		}
	}

	if len(exceptions) > 0 {
		return exceptions
	}

	return nil
}

// BuildKVP builds a url.Values query from a GetMapKVP struct
func (gfikvp *GetFeatureInfoKVP) ToQueryParameters() url.Values {
	query := make(map[string][]string)

	query[SERVICE] = []string{gfikvp.Service}
	query[VERSION] = []string{gfikvp.Version}
	query[REQUEST] = []string{gfikvp.Request}
	query[LAYERS] = []string{gfikvp.Layers}
	query[STYLES] = []string{gfikvp.Styles}
	query["CRS"] = []string{gfikvp.CRS}
	query[BBOX] = []string{gfikvp.Bbox}
	query[WIDTH] = []string{gfikvp.Width}
	query[HEIGHT] = []string{gfikvp.Height}

	if gfikvp.Format != `` {
		query[FORMAT] = []string{gfikvp.Format}
	}

	query[QUERYLAYERS] = []string{gfikvp.QueryLayers}
	query[INFOFORMAT] = []string{gfikvp.InfoFormat}
	query[I] = []string{gfikvp.I}
	query[J] = []string{gfikvp.J}

	if gfikvp.FeatureCount != nil {
		query[FEATURECOUNT] = []string{*gfikvp.FeatureCount}
	}
	if gfikvp.Exceptions != nil {
		query[EXCEPTIONS] = []string{*gfikvp.Exceptions}
	}

	return query
}

// ParseOperationRequest builds a GetFeatureInfoKVP object based on a GetFeatureInfo struct
func (gfirkvp *GetFeatureInfoKVP) ParseOperationRequest(or OperationRequest) Exceptions {
	gfi := or.(*GetFeatureInfoRequest)
	gfikvp := gfirkvp
	gfikvp.Request = getfeatureinfo
	gfikvp.Version = Version
	gfikvp.Service = Service
	gfikvp.Layers = gfi.StyledLayerDescriptor.getLayerKVPValue()
	gfikvp.Styles = gfi.StyledLayerDescriptor.getStyleKVPValue()
	gfikvp.CRS = gfi.CRS
	gfikvp.Bbox = gfi.BoundingBox.BuildQueryParameters()
	gfikvp.Width = strconv.Itoa(gfi.Size.Width)
	gfikvp.Height = strconv.Itoa(gfi.Size.Height)

	gfikvp.QueryLayers = strings.Join(gfi.QueryLayers, ",")
	gfikvp.InfoFormat = gfi.InfoFormat
	gfikvp.I = strconv.Itoa(gfi.I)
	gfikvp.J = strconv.Itoa(gfi.J)

	gfikvp.Format = gfi.Format

	if gfi.FeatureCount != nil {
		fcp := strconv.Itoa(*gfi.FeatureCount)
		gfikvp.FeatureCount = &fcp
	}

	gfikvp.Exceptions = gfi.Exceptions

	return nil
}
