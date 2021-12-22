package validatorrule

import (
	"strconv"

	"github.com/go-playground/validator/v10"
)

// IDCardRule ...
type IDCardRule struct{ *CommonRule }

// GetRule ...
func (*IDCardRule) GetRule() func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		citizenID := fl.Field().String()
		if citizenID != "" {
			if len(citizenID) == 13 {
				sum := 0
				for i := 0; i < 12; i++ {
					digit, _ := strconv.Atoi(string(citizenID[i]))
					sum += digit * (13 - i)
				}

				lastDigit, _ := strconv.Atoi(string(citizenID[12]))

				if (11-sum%11)%10 == lastDigit {
					return true
				}
			}

			return false
		}
		return true
	}
}

// NewIDCardRule ...
func NewIDCardRule() *IDCardRule {
	return &IDCardRule{&CommonRule{
		Field:     "id_card",
		FieldName: "Thai Citizen ID",
		Message:   "{0} invalid format!",
	}}
}
