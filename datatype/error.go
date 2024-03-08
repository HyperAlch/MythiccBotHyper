package datatype

import (
	"encoding/json"
	"strings"
)

type ApiError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewApiError(err error) (ApiError, bool) {
	if err == nil {
		return ApiError{}, false
	}

	errStr := err.Error()
	openB := strings.Index(errStr, "{")
	closeB := strings.Index(errStr, "}")

	if openB == -1 || closeB == -1 {
		return ApiError{}, false
	}
	apiError := ApiError{}
	jsonBytes := []byte(errStr[openB : closeB+1])

	err = json.Unmarshal(jsonBytes, &apiError)
	if err != nil {
		return ApiError{}, false
	}

	return apiError, true
}
