package validations

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/iancoleman/strcase"
)

func GetValidationMessage(exception validator.ValidationErrors) map[string]interface{} {
	// errorData := make([]string, len(exception))
	errorData := map[string]interface{}{}
	for _, err := range exception {
		switch err.Tag() {
		case "required":
			errorData[strcase.ToSnake(err.Field())] = fmt.Sprintf("Kolom %s wajib diisi",
				strcase.ToSnake(err.Field()))
		case "email":
			errorData[strcase.ToSnake(err.Field())] = fmt.Sprintf("Kolom %s harus berisi email yang valid",
				strcase.ToSnake(err.Field()))
		case "min":
			errorData[strcase.ToSnake(err.Field())] = fmt.Sprintf("Kolom %s harus memiliki panjang minimal %s",
				strcase.ToSnake(err.Field()), err.Param())
		case "max":
			errorData[strcase.ToSnake(err.Field())] = fmt.Sprintf("Panjang kolom %s tidak boleh melebihi %s",
				strcase.ToSnake(err.Field()), err.Param())
		}
	}
	// errorMessage := strings.Join(errorData, ",")
	return errorData
}
