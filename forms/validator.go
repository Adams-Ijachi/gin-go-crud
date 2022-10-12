package forms

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field string
	Msg   string
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"

	case "min":
		return "This field must be at least 10 characters"
	}
	return ""
}

func ValidateForm(err error) (bool, []ApiError) {

	var ve validator.ValidationErrors

	if ok := errors.As(err, &ve); ok {

		out := make([]ApiError, len(ve))

		for i, e := range ve {

			out[i] = ApiError{
				Field: e.Field(),
				Msg:   msgForTag(e.Tag()),
			}

		}

		return false, out
	}

	return true, nil
}
