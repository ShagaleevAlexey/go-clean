package models

import (
	"encoding/json"
	"github.com/ShagaleevAlexey/go-clean/internal/utils/validator"
)

type ResponseAuthToken struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
	ExpiresIn    int64  `json:"expires_in" validate:"required"`
}

func NewResponseAuthToken(at string, rt string, exp int64) (*ResponseAuthToken, error) {
	model := &ResponseAuthToken{
		AccessToken:  at,
		RefreshToken: rt,
		ExpiresIn:    exp,
	}

	if err := validator.ValidateStruct(model); err != nil {
		return nil, err
	}

	return model, nil
}

func (r *ResponseAuthToken) Marshal() ([]byte, error) {
	return json.Marshal(r)
}