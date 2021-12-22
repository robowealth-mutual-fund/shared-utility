package validatorrule

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// SuitAnswerRule ...
type SuitAnswerRule struct{ *CommonRule }

// BoolSlice ..
type BoolSlice struct {
	BoolSlice []bool
}

// GetRule ...
func (*SuitAnswerRule) GetRule() func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		fieldName := fl.FieldName()
		answer := fl.Field().Interface()
		str := fmt.Sprintf("%v", answer)
		str = strings.ReplaceAll(str, " ", ",")
		boolSlice, err := ConvertToBoolSlice(str)

		if err != nil {
			return false
		}

		countAnswer := 0
		for _, boolVal := range boolSlice {
			if boolVal {
				countAnswer++
			}
		}

		if fieldName == "Answer4" {
			if countAnswer == 0 {
				return false
			}
		} else {
			if countAnswer == 0 || countAnswer > 1 {
				return false
			}
		}

		return true
	}
}

// ConvertToBoolSlice ..
func ConvertToBoolSlice(s string) ([]bool, error) {
	{
		a := &BoolSlice{}
		err := json.Unmarshal([]byte(`{"BoolSlice":`+s+"}"), a)
		return a.BoolSlice, err
	}
}

// NewSuitAnswerRule ...
func NewSuitAnswerRule() *SuitAnswerRule {
	return &SuitAnswerRule{&CommonRule{
		Field:     "suit_answer",
		FieldName: "Suit Answer",
		Message:   "{0} Need at least 1 and not more than 2 answers!",
	}}
}
