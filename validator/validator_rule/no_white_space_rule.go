package validatorrule

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// NoWhiteNoRule ...
type NoWhiteNoRule struct{ *CommonRule }

// GetRule ...
func (*NoWhiteNoRule) GetRule() func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		data := fl.Field().String()
		if data != "" {
			matched, err := regexp.Match(`^[\S]+$`, []byte(data))
			if err != nil {
				return false
			}
			return matched
		}
		return true
	}
}

// NewNoWhiteNoRule ...
func NewNoWhiteNoRule() *NoWhiteNoRule {
	return &NoWhiteNoRule{&CommonRule{
		Field:     "no_white_space",
		FieldName: "No White Space",
		Message:   "{0} invalid format!",
	}}
}
