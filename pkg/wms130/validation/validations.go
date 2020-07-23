package validation

import (
	"github.com/go-playground/validator/v10"
	ows "github.com/pdok/ogc-specifications/pkg/ows"
)

func registerValidations(v *validator.Validate) {

	v.RegisterStructValidation(BboxValidator, ows.BoundingBox{})
	v.RegisterStructValidation(GetMapValidation, GetMapWrapper{})
}

// BboxValidator implements validator.CustomTypeFunc
func BboxValidator(sl validator.StructLevel) {

	bbox := sl.Current().Interface().(ows.BoundingBox)

	if bbox.LowerCorner[0] >= bbox.UpperCorner[0] {
		sl.ReportError(bbox, "lowerCorner", "LowerCorner", "bbox", `bbox`)
	}

	if bbox.LowerCorner[1] >= bbox.UpperCorner[1] {
		sl.ReportError(bbox, "lowerCorner", "LowerCorner", "bbox", `bbox`)
	}

}

// GetMapValidation structlevel
func GetMapValidation(sl validator.StructLevel) {

	wr := sl.Current().Interface().(GetMapWrapper)

	// layer
	for _, sldname := range wr.getmap.StyledLayerDescriptor.GetNamedLayers() {
		known := false
		for _, l := range wr.getcapabilities.Capability.Layer {
			if *l.Name == sldname {
				known = true
			}
		}
		if !known {
			sl.ReportError(wr.getmap.StyledLayerDescriptor.NamedLayer, "namedlayer", "NamedLayer", "knownlayer", sldname)
		}
	}
}
