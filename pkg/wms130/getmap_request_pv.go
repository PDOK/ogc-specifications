package wms130

import (
	"net/url"
	"strconv"
	"strings"
)

//getMapRequestParameterValue struct
type getMapRequestParameterValue struct {
	// Table 8 - The Parameters of a GetMap request
	service string `yaml:"service,omitempty"`
	baseParameterValueRequest
	getMapParameterValueMandatory
	getMapParameterValueOptional
}

// parseQueryParameters builds a getMapRequestParameterValue object based on the available query parameters
func (mpv *getMapRequestParameterValue) parseQueryParameters(query url.Values) Exceptions {
	var exceptions Exceptions
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, InvalidParameterValue(k, strings.Join(v, ",")))
		} else {
			switch strings.ToUpper(k) {
			case SERVICE:
				mpv.service = strings.ToUpper(v[0])
			case VERSION:
				mpv.baseParameterValueRequest.version = v[0]
			case REQUEST:
				mpv.baseParameterValueRequest.request = v[0]
			case LAYERS:
				mpv.getMapParameterValueMandatory.layers = v[0]
			case STYLES:
				mpv.getMapParameterValueMandatory.styles = v[0]
			case "CRS":
				mpv.getMapParameterValueMandatory.crs = v[0]
			case BBOX:
				mpv.getMapParameterValueMandatory.bbox = v[0]
			case WIDTH:
				mpv.getMapParameterValueMandatory.width = v[0]
			case HEIGHT:
				mpv.getMapParameterValueMandatory.height = v[0]
			case FORMAT:
				mpv.getMapParameterValueMandatory.format = v[0]
			case TRANSPARENT:
				mpv.getMapParameterValueOptional.transparent = &(v[0])
			case BGCOLOR:
				mpv.getMapParameterValueOptional.bgcolor = &(v[0])
			case EXCEPTIONS:
				mpv.getMapParameterValueOptional.exceptions = &(v[0])
			}
		}
	}
	if _, ok := query[VERSION]; !ok {
		exceptions = append(exceptions, MissingParameterValue(VERSION))
	}
	if _, ok := query[REQUEST]; !ok {
		exceptions = append(exceptions, MissingParameterValue(REQUEST))
	}
	if _, ok := query[LAYERS]; !ok {
		exceptions = append(exceptions, MissingParameterValue(LAYERS))
	}
	if _, ok := query[STYLES]; !ok {
		exceptions = append(exceptions, MissingParameterValue(STYLES))
	}
	if _, ok := query["CRS"]; !ok {
		exceptions = append(exceptions, MissingParameterValue("CRS"))
	}
	if _, ok := query[BBOX]; !ok {
		exceptions = append(exceptions, MissingParameterValue(BBOX))
	}
	if _, ok := query[WIDTH]; !ok {
		exceptions = append(exceptions, MissingParameterValue(WIDTH))
	}
	if _, ok := query[HEIGHT]; !ok {
		exceptions = append(exceptions, MissingParameterValue(HEIGHT))
	}
	if _, ok := query[FORMAT]; !ok {
		exceptions = append(exceptions, MissingParameterValue(FORMAT))
	}
	if len(exceptions) > 0 {
		return exceptions
	}
	return nil
}

// parseGetMapRequest builds a getMapRequestParameterValue object based on a GetMap struct
func (mpv *getMapRequestParameterValue) parseGetMapRequest(m GetMapRequest) Exceptions {

	mpv.request = getmap
	mpv.version = Version
	mpv.service = Service
	mpv.layers = m.StyledLayerDescriptor.getLayerParameterValue()
	mpv.styles = m.StyledLayerDescriptor.getStyleParameterValue()
	mpv.crs = m.CRS.String()
	mpv.bbox = m.BoundingBox.ToQueryParameters()
	mpv.width = strconv.Itoa(m.Output.Size.Width)
	mpv.height = strconv.Itoa(m.Output.Size.Height)
	mpv.format = m.Output.Format

	if m.Output.Transparent != nil {
		t := *m.Output.Transparent
		tp := strconv.FormatBool(t)
		mpv.transparent = &tp
	}

	if m.Output.BGcolor != nil {
		mpv.bgcolor = m.Output.BGcolor
	}

	// TODO: something with Time & Elevation
	// mpv.Time = m.Time
	// mpv.Elevation = m.Elevation

	mpv.exceptions = m.Exceptions

	return nil
}

// BuildOutput builds a Output struct from the getMapRequestParameterValue information
func (mpv *getMapRequestParameterValue) buildOutput() (Output, Exceptions) {
	output := Output{}

	h, err := strconv.Atoi(mpv.height)
	if err != nil {
		return output, InvalidParameterValue(mpv.height, HEIGHT).ToExceptions()
	}
	w, err := strconv.Atoi(mpv.width)
	if err != nil {
		return output, InvalidParameterValue(mpv.width, WIDTH).ToExceptions()
	}

	output.Size = Size{Height: h, Width: w}
	output.Format = mpv.format
	if mpv.transparent != nil {
		b, err := strconv.ParseBool(*mpv.transparent);
		if err != nil {
			return output, InvalidParameterValue(*mpv.transparent, TRANSPARENT).ToExceptions()
		}
		output.Transparent = &b
	}
	output.BGcolor = mpv.bgcolor

	return output, nil
}

// StyledLayer struct
type styledLayer struct {
	layers string `yaml:"layers,omitempty"`
	styles string `yaml:"styles,omitempty"`
}

// buildStyledLayerDescriptor builds a StyledLayerDescriptor struct from the parameter value information
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

// toQueryParameters builds a url.Values query from a getMapRequestParameterValue struct
func (mpv *getMapRequestParameterValue) toQueryParameters() url.Values {
	query := make(map[string][]string)

	query[SERVICE] = []string{mpv.service}
	query[VERSION] = []string{mpv.version}
	query[REQUEST] = []string{mpv.request}
	query[LAYERS] = []string{mpv.layers}
	query[STYLES] = []string{mpv.styles}
	query["CRS"] = []string{mpv.crs}
	query[BBOX] = []string{mpv.bbox}
	query[WIDTH] = []string{mpv.width}
	query[HEIGHT] = []string{mpv.height}
	query[FORMAT] = []string{mpv.format}

	if mpv.transparent != nil {
		query[TRANSPARENT] = []string{*mpv.transparent}
	}
	if mpv.bgcolor != nil {
		query[BGCOLOR] = []string{*mpv.bgcolor}
	}
	if mpv.exceptions != nil {
		query[EXCEPTIONS] = []string{*mpv.exceptions}
	}

	return query
}

// getMapParameterValueMandatory struct containing the mandatory WMS request Parameter Value
type getMapParameterValueMandatory struct {
	styledLayer
	crs    string `yaml:"crs,omitempty"`
	bbox   string `yaml:"bbox,omitempty"`
	width  string `yaml:"width,omitempty"`
	height string `yaml:"height,omitempty"`
	format string `yaml:"format,omitempty"`
}

// getMapParameterValueOptional struct containing the optional WMS request Parameter Value
type getMapParameterValueOptional struct {
	transparent *string `yaml:"transparent,omitempty"`
	bgcolor     *string `yaml:"bgcolor,omitempty"`
	exceptions  *string `yaml:"exceptions,omitempty"`
	// TODO: something with Time & Elevation
	// Time        *string `yaml:"time,omitempty"`
	// Elevation   *string `yaml:"elevation,omitempty"`
}
