package models

import (
	"github.com/ShagaleevAlexey/go-clean/internal/platform/web"
	"github.com/ShagaleevAlexey/go-clean/internal/utils/validator"
)

type RequestAuthUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func NewRequestAuthUserFromRequest(r *web.WebRequest) (*RequestAuthUser, error) {
	model := &RequestAuthUser{}

	if err := r.DecodeBody(model); err != nil {
		return nil, err
	}

	if err := validator.ValidateStruct(model); err != nil {
		return nil, err
	}

	return model, nil
}
