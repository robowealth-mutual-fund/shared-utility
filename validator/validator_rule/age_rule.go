package validatorrule

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

// AgeRule ...
type AgeRule struct{ *CommonRule }

// GetRule ...
func (*AgeRule) GetRule() func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		data := fl.Field().String()
		dateArr := strings.Split(data, "-")

		if dateArr[1] == "00" || dateArr[2] == "00" {
			data = fmt.Sprintf("%s-01-01", dateArr[0])
		}

		if data != "" {
			now := time.Now().UTC()
			date, err := time.Parse("2006-01-02", data)
			if err != nil {
				return false
			}
			check := date.AddDate(20, 0, 0).UTC()
			if check.After(now) {
				return false
			}
			return true
		}
		return true
	}
}

// NewAgeRule ...
func NewAgeRule() *AgeRule {
	return &AgeRule{&CommonRule{
		Field:     "age",
		FieldName: "Birthday",
		Message:   "{0} must be over 20 years old",
	}}
}
