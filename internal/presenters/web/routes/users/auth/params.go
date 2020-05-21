package auth

import (
	"github.com/ShagaleevAlexey/go-clean/internal/presenters/web/models"
	"github.com/ShagaleevAlexey/go-clean/internal/platform/web"
	"github.com/ShagaleevAlexey/go-clean/internal/utils/validator"
)

type AuthUserParams struct {
	Body *models.RequestAuthUser `validate:"dive"`
}

func NewAuthUserParams(r *web.WebRequest) (*AuthUserParams, error){
	var err error
	params := &AuthUserParams{}

	params.Body, err = models.NewRequestAuthUserFromRequest(r)
	if err != nil {
		return nil, err
	}

	if err := validator.ValidateStruct(params); err != nil {
		return nil, err
	}

	return params, nil
}
