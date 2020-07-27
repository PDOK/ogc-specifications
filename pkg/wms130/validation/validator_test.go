package validation

import (
	"encoding/xml"
	"testing"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/pdok/ogc-specifications/pkg/ows"
	"github.com/pdok/ogc-specifications/pkg/wms130/capabilities"
	"github.com/pdok/ogc-specifications/pkg/wms130/request"
	"github.com/pdok/ogc-specifications/pkg/wms130/response"
)

func sp(s string) *string {
	return &s
}

func TestValidation(t *testing.T) {

	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	v := validator.New()
	en_translations.RegisterDefaultTranslations(v, trans)

	registerValidations(v)
	registerTranslations(v, &trans)

	getcapabilities := response.GetCapabilities{
		WMSService: response.WMSService{Name: "RiversRoadsAndHouses"},
		Capability: capabilities.Capability{
			Layer: []capabilities.Layer{
				{Name: sp(`Rivers`), CRS: []*string{sp(`EPSG:4326`)}},
				{Name: sp(`Roads`), CRS: []*string{sp(`EPSG:4326`)}},
				{Name: sp(`Houses`), CRS: []*string{sp(`EPSG:4326`)}},
			},
		},
	}

	var tests = []struct {
		Object request.GetMap
	}{
		0: {
			Object: request.GetMap{
				XMLName: xml.Name{Local: "GetMap"},
				BaseRequest: request.BaseRequest{
					Version: "1.3.0",
					Service: "WMS",
				},
				StyledLayerDescriptor: request.StyledLayerDescriptor{
					Version: "1.1.0",
					NamedLayer: []request.NamedLayer{
						{Name: "Rivers", NamedStyle: &request.NamedStyle{Name: "CenterLine"}},
						{Name: "Roads", NamedStyle: &request.NamedStyle{Name: "CenterLine"}},
						{Name: "Houses", NamedStyle: &request.NamedStyle{Name: "Outline"}},
					}},
				CRS: "EPSG:4326",
				BoundingBox: ows.BoundingBox{
					LowerCorner: [2]float64{-180.0, -90.0},
					UpperCorner: [2]float64{180.0, 90.0},
				},
				Output: request.Output{
					Size:        request.Size{Width: 1024, Height: 512},
					Format:      "image/jpeg",
					Transparent: sp("false")},
				Exceptions: sp("XML"),
			},
		},
	}

	for k, n := range tests {

		wr := GetMapWrapper{getcapabilities: &getcapabilities, getmap: &n.Object}

		err := v.Struct(wr)
		if err != nil {
			t.Errorf("test: %d, got: %s", k, (err.(validator.ValidationErrors)).Translate(trans))
		}

		err = v.Struct(n.Object)
		if err != nil {

			t.Errorf("test: %d, got: %s", k, (err.(validator.ValidationErrors)).Translate(trans))
		}
	}
}
