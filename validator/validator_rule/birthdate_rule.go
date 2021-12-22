package validatorrule

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// BirthdateRule ...
type BirthdateRule struct{ *CommonRule }

// GetRule ...
func (*BirthdateRule) GetRule() func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		data := fl.Field().String()
		if data != "" {
			matched, err := regexp.Match(`^\d{4}\-(0[0-9]|1[012])\-(0[0-9]|[12][0-9]|3[01])$`, []byte(data))
			if err != nil {
				return false
			}
			return matched
		}
		return true
	}
}

// NewBirthdateRule ...
func NewBirthdateRule() *BirthdateRule {
	return &BirthdateRule{&CommonRule{
		Field:     "birthdate",
		FieldName: "Birthdate",
		Message:   "{0} invalid format!",
	}}
}
