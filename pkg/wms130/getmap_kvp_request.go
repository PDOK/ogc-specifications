package wms130

import (
	"net/url"
	"strconv"
	"strings"
)

//getMapKVPRequest struct
type getMapKVPRequest struct {
	// Table 8 - The Parameters of a GetMap request
	service string `yaml:"service,omitempty"`
	baseRequestKVP
	getMapKVPMandatory
	getMapKVPOptional
}

// parseQueryParameters builds a getMapKVPRequest object based on the available query parameters
func (gmkvp *getMapKVPRequest) parseQueryParameters(query url.Values) Exceptions {
	var exceptions Exceptions
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, InvalidParameterValue(k, strings.Join(v, ",")))
		} else {
			switch strings.ToUpper(k) {
			case SERVICE:
				gmkvp.service = strings.ToUpper(v[0])
			case VERSION:
				gmkvp.baseRequestKVP.version = v[0]
			case REQUEST:
				gmkvp.baseRequestKVP.request = v[0]
			case LAYERS:
				gmkvp.getMapKVPMandatory.layers = v[0]
			case STYLES:
				gmkvp.getMapKVPMandatory.styles = v[0]
			case "CRS":
				gmkvp.getMapKVPMandatory.crs = v[0]
			case BBOX:
				gmkvp.getMapKVPMandatory.bbox = v[0]
			case WIDTH:
				gmkvp.getMapKVPMandatory.width = v[0]
			case HEIGHT:
				gmkvp.getMapKVPMandatory.height = v[0]
			case FORMAT:
				gmkvp.getMapKVPMandatory.format = v[0]
			case TRANSPARENT:
				gmkvp.getMapKVPOptional.transparent = &(v[0])
			case BGCOLOR:
				gmkvp.getMapKVPOptional.bgcolor = &(v[0])
			case EXCEPTIONS:
				gmkvp.getMapKVPOptional.exceptions = &(v[0])
			}
		}
	}

	if len(exceptions) > 0 {
		return exceptions
	}

	return nil
}

// parseGetMapRequest builds a getMapKVPRequest object based on a GetMap struct
func (gmkvp *getMapKVPRequest) parseGetMapRequest(gm GetMapRequest) Exceptions {

	gmkvp.request = getmap
	gmkvp.version = Version
	gmkvp.service = Service
	gmkvp.layers = gm.StyledLayerDescriptor.getLayerKVPValue()
	gmkvp.styles = gm.StyledLayerDescriptor.getStyleKVPValue()
	gmkvp.crs = gm.CRS.String()
	gmkvp.bbox = gm.BoundingBox.BuildQueryParameters()
	gmkvp.width = strconv.Itoa(gm.Output.Size.Width)
	gmkvp.height = strconv.Itoa(gm.Output.Size.Height)
	gmkvp.format = gm.Output.Format

	if gm.Output.Transparent != nil {
		t := *gm.Output.Transparent
		tp := strconv.FormatBool(t)
		gmkvp.transparent = &tp
	}

	if gm.Output.BGcolor != nil {
		gmkvp.bgcolor = gm.Output.BGcolor
	}

	// TODO: something with Time & Elevation
	// gmkvp.Time = gm.Time
	// gmkvp.Elevation = gm.Elevation

	gmkvp.exceptions = gm.Exceptions

	return nil
}

// BuildOutput builds a Output struct from the KVP information
func (gmkvp *getMapKVPRequest) buildOutput() (Output, Exceptions) {
	output := Output{}

	h, err := strconv.Atoi(gmkvp.height)
	if err != nil {
		return output, InvalidParameterValue(HEIGHT, gmkvp.height).ToExceptions()
	}
	w, err := strconv.Atoi(gmkvp.width)
	if err != nil {
		return output, InvalidParameterValue(WIDTH, gmkvp.width).ToExceptions()
	}

	output.Size = Size{Height: h, Width: w}
	output.Format = gmkvp.format
	if b, err := strconv.ParseBool(*gmkvp.transparent); err == nil {
		output.Transparent = &b
	}
	output.BGcolor = gmkvp.bgcolor

	return output, nil
}

// StyledLayer struct
type styledLayer struct {
	layers string `yaml:"layers,omitempty"`
	styles string `yaml:"styles,omitempty"`
}

// BuildStyledLayerDescriptor builds a StyledLayerDescriptor struct from the KVP information
func (sl *styledLayer) buildStyledLayerDescriptor() (StyledLayerDescriptor, Exceptions) {
	var layers, styles []string
	if sl.layers != `` {
		layers = strings.Split(sl.layers, ",")
	}
	if sl.styles != `` {
		styles = strings.Split(sl.styles, ",")
	}

	sld, exceptions := buildStyledLayerDescriptor(layers, styles)
	if exceptions != nil {
		return sld, exceptions
	}

	return sld, nil
}

// toQueryParameters builds a url.Values query from a GetMapKVP struct
func (gmkvp *getMapKVPRequest) toQueryParameters() url.Values {
	query := make(map[string][]string)

	query[SERVICE] = []string{gmkvp.service}
	query[VERSION] = []string{gmkvp.version}
	query[REQUEST] = []string{gmkvp.request}
	query[LAYERS] = []string{gmkvp.layers}
	query[STYLES] = []string{gmkvp.styles}
	query["CRS"] = []string{gmkvp.crs}
	query[BBOX] = []string{gmkvp.bbox}
	query[WIDTH] = []string{gmkvp.width}
	query[HEIGHT] = []string{gmkvp.height}
	query[FORMAT] = []string{gmkvp.format}

	if gmkvp.transparent != nil {
		query[TRANSPARENT] = []string{*gmkvp.transparent}
	}
	if gmkvp.bgcolor != nil {
		query[BGCOLOR] = []string{*gmkvp.bgcolor}
	}
	if gmkvp.exceptions != nil {
		query[EXCEPTIONS] = []string{*gmkvp.exceptions}
	}

	return query
}

// GetMapKVPMandatory struct containing the mandatory WMS request KVP
type getMapKVPMandatory struct {
	styledLayer
	crs    string `yaml:"crs,omitempty"`
	bbox   string `yaml:"bbox,omitempty"`
	width  string `yaml:"width,omitempty"`
	height string `yaml:"height,omitempty"`
	format string `yaml:"format,omitempty"`
}

// GetMapKVPOptional struct containing the optional WMS request KVP
type getMapKVPOptional struct {
	transparent *string `yaml:"transparent,omitempty"`
	bgcolor     *string `yaml:"bgcolor,omitempty"`
	exceptions  *string `yaml:"exceptions,omitempty"`
	// TODO: something with Time & Elevation
	// Time        *string `yaml:"time,omitempty"`
	// Elevation   *string `yaml:"elevation,omitempty"`
}
