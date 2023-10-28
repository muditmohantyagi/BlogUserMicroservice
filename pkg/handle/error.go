package handle

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Ivalid email"
	case "min":
		return "Should be greater " + fe.Param()
	case "max":
		return "Should be greater than " + fe.Param()
	case "eqfield":
		return "Passwords are not equal,Value: " + fe.Value().(string)

	default:
		return "UnHandeledMessage:" + fe.Error()
	}

}
func Error(err error) []ErrorMsg {
	var ve validator.ValidationErrors
	var OutPut []ErrorMsg
	if errors.As(err, &ve) {

		for _, fe := range ve {
			custome_message := ErrorMsg{fe.Field(), getErrorMsg(fe)}
			OutPut = append(OutPut, custome_message)
		}
	} else {
		custome_message := ErrorMsg{"Validation", err.Error()}
		OutPut = append(OutPut, custome_message)
	}
	return OutPut
}
