package apiaierrors

import (
	"errors"
	"fmt"
)

type ApiAiRequestError struct {
	Message string
	Code    int
}

func (apiAiRequestError *ApiAiRequestError) Throw() error {
	return errors.New(fmt.Sprintf("%d : %s", apiAiRequestError.Code, apiAiRequestError.Message))
}
