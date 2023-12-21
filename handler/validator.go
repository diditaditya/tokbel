package handler

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func ValidateStruct(s any) *map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	result := map[string]string{}
	for _, err := range err.(validator.ValidationErrors) {
		// fmt.Println("================= VALIDATION ERROR ===========")
		// fmt.Println("namespace: ", err.Namespace())
		// fmt.Println("field: ", err.Field())
		// fmt.Println("struct namespace: ", err.StructNamespace())
		// fmt.Println("struct field: ", err.StructField())
		// fmt.Println("tag: ", err.Tag())
		// fmt.Println("actual tag: ", err.ActualTag())
		// fmt.Println("kind: ", err.Kind())
		// fmt.Println("type: ", err.Type())
		// fmt.Println("value: ", err.Value())
		// fmt.Println("param: ", err.Param())
		// fmt.Println()
		errMsg := fmt.Sprintf("invalid %s", err.Field())
		if err.Tag() == "required" {
			errMsg = fmt.Sprintf("%s is required", err.Field())
		}
		result[err.Field()] = errMsg
	}

	return &result
}
