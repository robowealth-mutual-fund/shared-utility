package validator

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	vr "github.com/robowealth-mutual-fund/shared-utility/validator/validator_rule"

	errs "github.com/robowealth-mutual-fund/shared-utility/errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// CustomValidator ...
type CustomValidator struct {
	Validator *validator.Validate
	Trans     ut.Translator
}

// Configure ...
func (cv *CustomValidator) Configure() {
	// Setup validator and initial language
	v := validator.New()

	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(v, trans)

	// Assign the Validate and Trans
	cv.Validator = v
	cv.Trans = trans

	// Start register a custom rules here ...
	cv.RegisterRule(vr.NewMobileNumberRule())
	cv.RegisterRule(vr.NewIDCardRule())
	cv.RegisterRule(vr.NewSuitAnswerRule())
	cv.RegisterRule(vr.NewNoWhiteNoRule())
	cv.RegisterRule(vr.NewBirthdateRule())
	cv.RegisterRule(vr.NewImage64BitRule())
	cv.RegisterRule(vr.NewTokenRule())
	cv.RegisterRule(vr.NewAgeRule())
}

// RegisterRule ...
func (cv *CustomValidator) RegisterRule(rule vr.Rule) {
	vr.RegisterValidationRule(rule, cv.Validator, cv.Trans)
}

// Validate ...
func (cv *CustomValidator) Validate(structRule interface{}) error {
	if err := cv.Validator.Struct(structRule); err != nil {
		formError := errs.NewFormError("Wrong Input")

		if err != nil {
			for _, e := range err.(validator.ValidationErrors) {

				jsonFieldName := e.Field()
				if field, ok := reflect.TypeOf(structRule).Elem().FieldByName(e.Field()); ok {
					if jsonTag, ok := field.Tag.Lookup("json"); ok {
						jsonFieldName = strings.Split(jsonTag, ",")[0]
					}
				}

				param := e.Param()
				if param != "" {
					param = "=" + param
				}

				errorMsg := fmt.Sprintf("ERROR_%s%s: %s", strings.ToUpper(e.Tag()), param, e.Translate(cv.Trans))

				log.Println(jsonFieldName, ">>>", errorMsg)

				formError.AddErrorField(jsonFieldName, errorMsg)
			}
		}

		return formError
	}

	return nil
}

// NewCustomValidator ...
func NewCustomValidator() *CustomValidator {
	var cv = &CustomValidator{}
	cv.Configure()
	return cv
}
