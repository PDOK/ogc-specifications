package validation

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func registerTranslations(v *validator.Validate, trans *ut.Translator) {

	v.RegisterTranslation(
		"epsg",
		*trans,
		func(ut ut.Translator) error {
			return ut.Add("epsg", "{0} has a invalid value: {1}", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("epsg", fe.Field(), fmt.Sprintf("%v", fe.Value()))
			return t
		},
	)

	v.RegisterTranslation("version", *trans,
		func(ut ut.Translator) error {
			return ut.Add("version", "{0} has a invalid value: {1}", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("version", fe.Field(), fmt.Sprintf("%v", fe.Value()))
			return t
		})

	v.RegisterTranslation("required", *trans,
		func(ut ut.Translator) error {
			return ut.Add("required", "{0} is missing", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		})

	v.RegisterTranslation("knownlayer", *trans,
		func(ut ut.Translator) error {
			return ut.Add("knownlayer", "{0} has a invalid value {1}", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("knownlayer", fe.Field(), fmt.Sprintf("%v", fe.Value()))
			return t
		})
}
