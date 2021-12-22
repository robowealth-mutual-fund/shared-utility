package errors

type fieldError struct {
	FieldName   string
	Description string
}

// FormError ...
type FormError struct {
	Message string
	Fields  []*fieldError
}

// Error ...
func (e *FormError) Error() string {
	return e.Message
}

// AddErrorField ...
func (e *FormError) AddErrorField(fieldName string, description string) {
	e.Fields = append(e.Fields, &fieldError{
		FieldName:   fieldName,
		Description: description,
	})
}

// NewFormError ...
func NewFormError(message string) *FormError {
	err := &FormError{
		Message: message,
	}

	return err
}

// NewFormErrorWithFields ...
func NewFormErrorWithFields(message string, fields []*fieldError) *FormError {
	err := NewFormError(message)
	err.Fields = fields
	return err
}
