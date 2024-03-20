package controller

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	ErrorMsg string `json:"error"`
}

func failedResponse(errorMsg string) ErrorResponse {
	return ErrorResponse{
		ErrorMsg: errorMsg,
	}
}

func validationErrorParser(err error) (errMessage string) {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		var errNamespaces string
		for _, err := range errs {
			errNamespaces = fmt.Sprintf("%s, %s", errNamespaces, err.Namespace())
		}
		return "以下字段错误：" + errNamespaces
	}
	return "Validator异常"
}
