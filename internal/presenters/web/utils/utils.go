package utils

import (
	"encoding/json"
	"github.com/ShagaleevAlexey/go-clean/internal/platform/web"
	"github.com/ShagaleevAlexey/go-clean/internal/usecases"
)

type WebError struct {
	Message string `json:"message"`
	Code int32 `json:"code"`
}

func NewWebError(message string, code int32) *WebError {
	return &WebError{
		Message: message,
		Code:    code,
	}
}

func (we *WebError) Marshal() ([]byte, error) {
	return json.Marshal(we)
}

func NewWebResponseWithError(status int, err error) *web.WebResponse {
	body, _ := NewWebError(err.Error(), 0).Marshal()
	if body == nil {
		body = []byte("Unexpected error")
	}

	return web.NewWebResponse(status, web.HeaderContentJson, body)
}

func NewWebResponseWithAppError(status int, err *usecases.AppError) *web.WebResponse {
	return web.NewWebResponse(status, web.HeaderContentJson, err.Errorb())
}
