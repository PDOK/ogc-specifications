package wms130

import (
	"net/url"
	"strconv"
	"strings"
)

//getFeatureInfoParameterValueRequest struct
type getFeatureInfoParameterValueRequest struct {
	// Table 8 - The Parameters of a GetFeatureInfo request
	service string `yaml:"service,omitempty"`
	baseParameterValueRequest
	getMapParameterValueMandatory
	getFeatureInfoParameterValueMandatory
	getFeatureInfoParameterValueOptional
}

// parseQueryParameters builds a getFeatureInfoParameterValueRequest object based on the available query parameters
func (ipv *getFeatureInfoParameterValueRequest) parseQueryParameters(query url.Values) Exceptions {
	var exceptions Exceptions
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, InvalidParameterValue(k, strings.Join(v, ",")))
		} else {
			switch strings.ToUpper(k) {
			case SERVICE:
				ipv.service = strings.ToUpper(v[0])
			case VERSION:
				ipv.baseParameterValueRequest.version = v[0]
			case REQUEST:
				ipv.baseParameterValueRequest.request = v[0]
			case LAYERS:
				ipv.getMapParameterValueMandatory.layers = v[0]
			case STYLES:
				ipv.getMapParameterValueMandatory.styles = v[0]
			case "CRS":
				ipv.getMapParameterValueMandatory.crs = v[0]
			case BBOX:
				ipv.getMapParameterValueMandatory.bbox = v[0]
			case WIDTH:
				ipv.getMapParameterValueMandatory.width = v[0]
			case HEIGHT:
				ipv.getMapParameterValueMandatory.height = v[0]
			case FORMAT:
				ipv.getMapParameterValueMandatory.format = v[0]
			case QUERYLAYERS:
				ipv.getFeatureInfoParameterValueMandatory.querylayers = v[0]
			case INFOFORMAT:
				ipv.getFeatureInfoParameterValueMandatory.infoformat = v[0]
			case I:
				ipv.getFeatureInfoParameterValueMandatory.i = v[0]
			case J:
				ipv.getFeatureInfoParameterValueMandatory.j = v[0]
			case FEATURECOUNT:
				ipv.getFeatureInfoParameterValueOptional.featurecount = &(v[0])
			case EXCEPTIONS:
				ipv.getFeatureInfoParameterValueOptional.exceptions = &(v[0])
			}
		}
	}

	if len(exceptions) > 0 {
		return exceptions
	}

	return nil
}

// ToQueryParameters builds a url.Values query from a getFeatureInfoParameterValueRequest struct
func (ipv getFeatureInfoParameterValueRequest) toQueryParameters() url.Values {
	query := make(map[string][]string)

	query[SERVICE] = []string{ipv.service}
	query[VERSION] = []string{ipv.version}
	query[REQUEST] = []string{ipv.request}
	query[LAYERS] = []string{ipv.layers}
	query[STYLES] = []string{ipv.styles}
	query["CRS"] = []string{ipv.crs}
	query[BBOX] = []string{ipv.bbox}
	query[WIDTH] = []string{ipv.width}
	query[HEIGHT] = []string{ipv.height}

	if ipv.format != `` {
		query[FORMAT] = []string{ipv.format}
	}

	query[QUERYLAYERS] = []string{ipv.querylayers}
	query[INFOFORMAT] = []string{ipv.infoformat}
	query[I] = []string{ipv.i}
	query[J] = []string{ipv.j}

	if ipv.featurecount != nil {
		query[FEATURECOUNT] = []string{*ipv.featurecount}
	}
	if ipv.exceptions != nil {
		query[EXCEPTIONS] = []string{*ipv.exceptions}
	}

	return query
}

// parseGetFeatureInfoRequest builds a getFeatureInfoParameterValueRequest object based on a GetFeatureInfoRequest struct
func (ipv *getFeatureInfoParameterValueRequest) parseGetFeatureInfoRequest(i GetFeatureInfoRequest) Exceptions {

	ipv.request = getfeatureinfo
	ipv.version = Version
	ipv.service = Service
	ipv.layers = i.StyledLayerDescriptor.getLayerParameterValue()
	ipv.styles = i.StyledLayerDescriptor.getStyleParameterValue()
	ipv.crs = i.CRS
	ipv.bbox = i.BoundingBox.ToQueryParameters()
	ipv.width = strconv.Itoa(i.Size.Width)
	ipv.height = strconv.Itoa(i.Size.Height)

	ipv.querylayers = strings.Join(i.QueryLayers, ",")
	ipv.infoformat = i.InfoFormat
	ipv.i = strconv.Itoa(i.I)
	ipv.j = strconv.Itoa(i.J)

	ipv.format = i.Format

	if i.FeatureCount != nil {
		fcp := strconv.Itoa(*i.FeatureCount)
		ipv.featurecount = &fcp
	}

	ipv.exceptions = i.Exceptions

	return nil
}

// GetFeatureInfoParameterValueMandatory struct containing the mandatory WMS request Parameter Value
type getFeatureInfoParameterValueMandatory struct {
	querylayers string `yaml:"query_layers,omitempty"`
	infoformat  string `yaml:"info_format,omitempty"`
	i           string `yaml:"i,omitempty"`
	j           string `yaml:"j,omitempty"`
}

// GetFeatureInfoParameterValueOptional struct containing the optional WMS request Parameter Value
type getFeatureInfoParameterValueOptional struct {
	featurecount *string `yaml:"feature_count,omitempty"`
	exceptions   *string `yaml:"exceptions,omitempty"`
}
