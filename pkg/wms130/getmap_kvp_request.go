package wms130

import (
	"net/url"
	"strconv"
	"strings"
)

//GetMapKVP struct
type GetMapKVP struct {
	// Table 8 - The Parameters of a GetMap request
	Service string `yaml:"service,omitempty"`
	BaseRequestKVP
	GetMapKVPMandatory
	GetMapKVPOptional
}

// StyledLayer struct
type StyledLayer struct {
	Layers string `yaml:"layers,omitempty"`
	Styles string `yaml:"styles,omitempty"`
}

// GetMapKVPMandatory struct containing the mandatory WMS request KVP
type GetMapKVPMandatory struct {
	StyledLayer
	CRS    string `yaml:"crs,omitempty"`
	Bbox   string `yaml:"bbox,omitempty"`
	Width  string `yaml:"width,omitempty"`
	Height string `yaml:"height,omitempty"`
	Format string `yaml:"format,omitempty"`
}

// GetMapKVPOptional struct containing the optional WMS request KVP
type GetMapKVPOptional struct {
	Transparent *string `yaml:"transparent,omitempty"`
	BGColor     *string `yaml:"bgcolor,omitempty"`
	Exceptions  *string `yaml:"exceptions,omitempty"`
	// TODO: something with Time & Elevation
	// Time        *string `yaml:"time,omitempty"`
	// Elevation   *string `yaml:"elevation,omitempty"`
}

// ParseKVP builds a GetMapKVP object based on the available query parameters
func (gmkvp *GetMapKVP) ParseQueryParameters(query url.Values) Exceptions {
	var exceptions Exceptions
	for k, v := range query {
		if len(v) != 1 {
			exceptions = append(exceptions, InvalidParameterValue(k, strings.Join(v, ",")))
		} else {
			switch strings.ToUpper(k) {
			case SERVICE:
				gmkvp.Service = strings.ToUpper(v[0])
			case VERSION:
				gmkvp.BaseRequestKVP.Version = v[0]
			case REQUEST:
				gmkvp.BaseRequestKVP.Request = v[0]
			case LAYERS:
				gmkvp.GetMapKVPMandatory.Layers = v[0]
			case STYLES:
				gmkvp.GetMapKVPMandatory.Styles = v[0]
			case "CRS":
				gmkvp.GetMapKVPMandatory.CRS = v[0]
			case BBOX:
				gmkvp.GetMapKVPMandatory.Bbox = v[0]
			case WIDTH:
				gmkvp.GetMapKVPMandatory.Width = v[0]
			case HEIGHT:
				gmkvp.GetMapKVPMandatory.Height = v[0]
			case FORMAT:
				gmkvp.GetMapKVPMandatory.Format = v[0]
			case TRANSPARENT:
				gmkvp.GetMapKVPOptional.Transparent = &(v[0])
			case BGCOLOR:
				gmkvp.GetMapKVPOptional.BGColor = &(v[0])
			case EXCEPTIONS:
				gmkvp.GetMapKVPOptional.Exceptions = &(v[0])
			}
		}
	}

	if len(exceptions) > 0 {
		return exceptions
	}

	return nil
}

// ParseOperationRequest builds a GetMapKVP object based on a GetMap struct
func (gmkvp *GetMapKVP) ParseOperationRequest(or OperationRequest) Exceptions {
	gm := or.(*GetMapRequest)

	gmkvp.Request = getmap
	gmkvp.Version = Version
	gmkvp.Service = Service
	gmkvp.Layers = gm.StyledLayerDescriptor.getLayerKVPValue()
	gmkvp.Styles = gm.StyledLayerDescriptor.getStyleKVPValue()
	gmkvp.CRS = gm.CRS.String()
	gmkvp.Bbox = gm.BoundingBox.BuildQueryParameters()
	gmkvp.Width = strconv.Itoa(gm.Output.Size.Width)
	gmkvp.Height = strconv.Itoa(gm.Output.Size.Height)
	gmkvp.Format = gm.Output.Format

	if gm.Output.Transparent != nil {
		t := *gm.Output.Transparent
		tp := strconv.FormatBool(t)
		gmkvp.Transparent = &tp
	}

	if gm.Output.BGcolor != nil {
		gmkvp.BGColor = gm.Output.BGcolor
	}

	// TODO: something with Time & Elevation
	// gmkvp.Time = gm.Time
	// gmkvp.Elevation = gm.Elevation

	gmkvp.Exceptions = gm.Exceptions

	return nil
}

// BuildOutput builds a Output struct from the KVP information
func (gmkvp *GetMapKVP) buildOutput() (Output, Exceptions) {
	output := Output{}

	h, err := strconv.Atoi(gmkvp.Height)
	if err != nil {
		return output, InvalidParameterValue(HEIGHT, gmkvp.Height).ToExceptions()
	}
	w, err := strconv.Atoi(gmkvp.Width)
	if err != nil {
		return output, InvalidParameterValue(WIDTH, gmkvp.Width).ToExceptions()
	}

	output.Size = Size{Height: h, Width: w}
	output.Format = gmkvp.Format
	if b, err := strconv.ParseBool(*gmkvp.Transparent); err == nil {
		output.Transparent = &b
	}
	output.BGcolor = gmkvp.BGColor

	return output, nil
}

// BuildStyledLayerDescriptor builds a StyledLayerDescriptor struct from the KVP information
func (sl *StyledLayer) buildStyledLayerDescriptor() (StyledLayerDescriptor, Exceptions) {
	var layers, styles []string
	if sl.Layers != `` {
		layers = strings.Split(sl.Layers, ",")
	}
	if sl.Styles != `` {
		styles = strings.Split(sl.Styles, ",")
	}

	sld, exceptions := buildStyledLayerDescriptor(layers, styles)
	if exceptions != nil {
		return sld, exceptions
	}

	return sld, nil
}

// BuildKVP builds a url.Values query from a GetMapKVP struct
func (gmkvp *GetMapKVP) ToQueryParameters() url.Values {
	query := make(map[string][]string)

	query[SERVICE] = []string{gmkvp.Service}
	query[VERSION] = []string{gmkvp.Version}
	query[REQUEST] = []string{gmkvp.Request}
	query[LAYERS] = []string{gmkvp.Layers}
	query[STYLES] = []string{gmkvp.Styles}
	query["CRS"] = []string{gmkvp.CRS}
	query[BBOX] = []string{gmkvp.Bbox}
	query[WIDTH] = []string{gmkvp.Width}
	query[HEIGHT] = []string{gmkvp.Height}
	query[FORMAT] = []string{gmkvp.Format}

	if gmkvp.Transparent != nil {
		query[TRANSPARENT] = []string{*gmkvp.Transparent}
	}
	if gmkvp.BGColor != nil {
		query[BGCOLOR] = []string{*gmkvp.BGColor}
	}
	if gmkvp.Exceptions != nil {
		query[EXCEPTIONS] = []string{*gmkvp.Exceptions}
	}

	return query
}
