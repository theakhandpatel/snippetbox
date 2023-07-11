package validator

import (
	"strings"
	"unicode/utf8"
)

// Validator is a struct that validates fields.
type Validator struct {
	// FieldErrors is a map of field names to error messages.
	FieldErrors map[string]string
}

// Valid returns true if there are no field errors.
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

// AddFieldError adds an error message to the map of field errors.
func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}
	_, exists := v.FieldErrors[key]
	if !exists {
		v.FieldErrors[key] = message
	}
}

// CheckField checks a condition and adds an error message if the condition is not met.
func (v *Validator) CheckField(conditionMet bool, key, message string) {
	if !conditionMet {
		v.AddFieldError(key, message)
	}
}

// NotBlank returns true if the value is not blank.
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// MaxChars returns true if the value has a length of at most n characters.
func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// PermittedInt returns true if the value is one of the permitted values.
func PermittedInt(value int, permittedValues ...int) bool {
	for _, permittedValue := range permittedValues {
		if value == permittedValue {
			return true
		}
	}
	return false
}
