package auth

import (
	"github.com/ShagaleevAlexey/go-clean/internal/presenters/web/models"
	"github.com/ShagaleevAlexey/go-clean/internal/platform/web"
)

func NewWebResponseWithResponseAuthTokens(status int, m *models.ResponseAuthToken) (*web.WebResponse, error) {
	webResponse := &web.WebResponse{}

	err := webResponse.WriteStruct(status, web.HeaderContentJson, m)
	if err != nil {
		return nil, err
	}

	return webResponse, nil
}

