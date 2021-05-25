package wms130

import (
	"net/url"
	"strconv"
	"strings"
)

//GetFeatureInfoKVP struct
type getFeatureInfoKVPRequest struct {
	// Table 8 - The Parameters of a GetFeatureInfo request
	service string `yaml:"service,omitempty"`
	baseRequestKVP
	getMapKVPMandatory
	getFeatureInfoKVPMandatory
	getFeatureInfoKVPOptional
}

// ParseKVP builds a GetMapKVP object based on the available query parameters
func (gfikvp *getFeatureInfoKVPRequest) parseQueryParameters(query url.Values) Exceptions {
	var exceptions Exceptions
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, InvalidParameterValue(k, strings.Join(v, ",")))
		} else {
			switch strings.ToUpper(k) {
			case SERVICE:
				gfikvp.service = strings.ToUpper(v[0])
			case VERSION:
				gfikvp.baseRequestKVP.version = v[0]
			case REQUEST:
				gfikvp.baseRequestKVP.request = v[0]
			case LAYERS:
				gfikvp.getMapKVPMandatory.layers = v[0]
			case STYLES:
				gfikvp.getMapKVPMandatory.styles = v[0]
			case "CRS":
				gfikvp.getMapKVPMandatory.crs = v[0]
			case BBOX:
				gfikvp.getMapKVPMandatory.bbox = v[0]
			case WIDTH:
				gfikvp.getMapKVPMandatory.width = v[0]
			case HEIGHT:
				gfikvp.getMapKVPMandatory.height = v[0]
			case FORMAT:
				gfikvp.getMapKVPMandatory.format = v[0]
			case QUERYLAYERS:
				gfikvp.getFeatureInfoKVPMandatory.querylayers = v[0]
			case INFOFORMAT:
				gfikvp.getFeatureInfoKVPMandatory.infoformat = v[0]
			case I:
				gfikvp.getFeatureInfoKVPMandatory.i = v[0]
			case J:
				gfikvp.getFeatureInfoKVPMandatory.j = v[0]
			case FEATURECOUNT:
				gfikvp.getFeatureInfoKVPOptional.featurecount = &(v[0])
			case EXCEPTIONS:
				gfikvp.getFeatureInfoKVPOptional.exceptions = &(v[0])
			}
		}
	}

	if len(exceptions) > 0 {
		return exceptions
	}

	return nil
}

// ToQueryParameters builds a url.Values query from a GetMapKVP struct
func (gfikvp *getFeatureInfoKVPRequest) toQueryParameters() url.Values {
	query := make(map[string][]string)

	query[SERVICE] = []string{gfikvp.service}
	query[VERSION] = []string{gfikvp.version}
	query[REQUEST] = []string{gfikvp.request}
	query[LAYERS] = []string{gfikvp.layers}
	query[STYLES] = []string{gfikvp.styles}
	query["CRS"] = []string{gfikvp.crs}
	query[BBOX] = []string{gfikvp.bbox}
	query[WIDTH] = []string{gfikvp.width}
	query[HEIGHT] = []string{gfikvp.height}

	if gfikvp.format != `` {
		query[FORMAT] = []string{gfikvp.format}
	}

	query[QUERYLAYERS] = []string{gfikvp.querylayers}
	query[INFOFORMAT] = []string{gfikvp.infoformat}
	query[I] = []string{gfikvp.i}
	query[J] = []string{gfikvp.j}

	if gfikvp.featurecount != nil {
		query[FEATURECOUNT] = []string{*gfikvp.featurecount}
	}
	if gfikvp.exceptions != nil {
		query[EXCEPTIONS] = []string{*gfikvp.exceptions}
	}

	return query
}

// parseGetFeatureInfoRequest builds a getFeatureInfoKVPRequest object based on a GetFeatureInfoRequest struct
func (gfikvp *getFeatureInfoKVPRequest) parseGetFeatureInfoRequest(gfi GetFeatureInfoRequest) Exceptions {

	gfikvp.request = getfeatureinfo
	gfikvp.version = Version
	gfikvp.service = Service
	gfikvp.layers = gfi.StyledLayerDescriptor.getLayerKVPValue()
	gfikvp.styles = gfi.StyledLayerDescriptor.getStyleKVPValue()
	gfikvp.crs = gfi.CRS
	gfikvp.bbox = gfi.BoundingBox.ToQueryParameters()
	gfikvp.width = strconv.Itoa(gfi.Size.Width)
	gfikvp.height = strconv.Itoa(gfi.Size.Height)

	gfikvp.querylayers = strings.Join(gfi.QueryLayers, ",")
	gfikvp.infoformat = gfi.InfoFormat
	gfikvp.i = strconv.Itoa(gfi.I)
	gfikvp.j = strconv.Itoa(gfi.J)

	gfikvp.format = gfi.Format

	if gfi.FeatureCount != nil {
		fcp := strconv.Itoa(*gfi.FeatureCount)
		gfikvp.featurecount = &fcp
	}

	gfikvp.exceptions = gfi.Exceptions

	return nil
}

// GetFeatureInfoKVPMandatory struct containing the mandatory WMS request KVP
type getFeatureInfoKVPMandatory struct {
	querylayers string `yaml:"query_layers,omitempty"`
	infoformat  string `yaml:"info_format,omitempty"`
	i           string `yaml:"i,omitempty"`
	j           string `yaml:"j,omitempty"`
}

// GetFeatureInfoKVPOptional struct containing the optional WMS request KVP
type getFeatureInfoKVPOptional struct {
	featurecount *string `yaml:"feature_count,omitempty"`
	exceptions   *string `yaml:"exceptions,omitempty"`
}
