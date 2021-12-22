package validatorrule

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// MobileNumberRule ...
type MobileNumberRule struct{ *CommonRule }

// GetRule ...
func (*MobileNumberRule) GetRule() func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		// International Format: +66812345678
		// Local Format: 0812345678
		// Country Code +66
		// Mobile prefix: 8, 9 or 6
		// 8 digit user number: 12345678
		data := fl.Field().String()
		if data != "" {
			matched, err := regexp.Match(`^(0|\+66)(6|8|9){1}[0-9]{8}$`, []byte(data))
			if err != nil {
				return false
			}
			return matched
		}
		return true
	}
}

// NewMobileNumberRule ...
func NewMobileNumberRule() *MobileNumberRule {
	return &MobileNumberRule{&CommonRule{
		Field:     "mobile_number",
		FieldName: "Mobile Number",
		Message:   "{0} invalid format!",
	}}
}
