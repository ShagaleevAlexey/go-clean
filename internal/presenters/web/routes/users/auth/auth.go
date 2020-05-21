package auth

import (
	"context"
	"github.com/ShagaleevAlexey/go-clean/internal/presenters/web/models"
	"github.com/ShagaleevAlexey/go-clean/internal/platform/web"
	"github.com/ShagaleevAlexey/go-clean/internal/presenters/web/utils"
	"github.com/ShagaleevAlexey/go-clean/internal/usecases"
	"github.com/ShagaleevAlexey/go-clean/internal/usecases/users"
	"github.com/ShagaleevAlexey/go-clean/internal/usecases/users/access_token"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func NewUCAuthUserModelFromParams(params *AuthUserParams) *users.UCAuthUserModel {
	return users.NewUCAuthUserModel(params.Body.Username, params.Body.Password)
}

type UseCaseFuncAuthUser func(ctx context.Context, input *users.UCAuthUserModel) (*access_token.AuthTokens, *usecases.AppError)

func MountAuthUser(r web.Router, f UseCaseFuncAuthUser) {
	r.Mount(http.MethodPost, "/auth", func(r *web.WebRequest) *web.WebResponse {
		return HandleAuthUser(r, f)
	})
}

func HandleAuthUser(request *web.WebRequest, f UseCaseFuncAuthUser) *web.WebResponse {
	var err error

	params, err := NewAuthUserParams(request)
	if err != nil {
		log.Errorf("<api> failed mapping request: %s", err)
		return utils.NewWebResponseWithError(http.StatusInternalServerError, err)
	}
	ucParams := NewUCAuthUserModelFromParams(params)

	tokens, appErr := f(request.Context(), ucParams)
	if appErr != nil {
		log.Errorf("<api> error: %s", appErr)
		var httpStatus = http.StatusInternalServerError

		switch appErr {
		case usecases.ErrUserNotFound:
			httpStatus = http.StatusNotFound
		}

		return utils.NewWebResponseWithAppError(httpStatus, appErr)
	}

	response, err := models.NewResponseAuthToken(tokens.AccessToken, tokens.RefreshToken, tokens.Expiration)
	if err != nil {
		log.Errorf("<api> failed mapping response: %s", err)
		return utils.NewWebResponseWithError(http.StatusInternalServerError, err)
	}

	webResponse, err := NewWebResponseWithResponseAuthTokens(http.StatusOK, response)
	if err != nil {
		log.Errorf("<api> failed write response: %s", err)
		return utils.NewWebResponseWithError(http.StatusInternalServerError, err)
	}
	return webResponse
}
