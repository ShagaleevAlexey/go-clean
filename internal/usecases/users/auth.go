package users

import (
	"context"
	"github.com/ShagaleevAlexey/go-clean/internal/usecases"
	"github.com/ShagaleevAlexey/go-clean/internal/usecases/users/access_token"
	"github.com/ShagaleevAlexey/go-clean/internal/usecases/users/passwds"
)

type UCAuthUserModel struct {
	Username string
	Password string
}

func NewUCAuthUserModel(username string, password string) *UCAuthUserModel {
	return &UCAuthUserModel{
		Username: username,
		Password: password,
	}
}

func (uc *UsersUseCases) AuthUser(ctx context.Context, input *UCAuthUserModel) (*access_token.AuthTokens, *usecases.AppError) {
	session := uc.db.CreateSession()
	defer session.DefferWithRollback()
	var err error

	user, err := uc.queries.FindUserByUsername(session, input.Username)
	if err != nil {
		return nil, usecases.ErrUserNotFound
	}

	err = passwds.MatchPasswordHash(user.Password, input.Password)
	if err != nil {
		return nil, usecases.ErrUserUnauthorizeError
	}

	authTokens, err := access_token.GenerateTokens(user.Id, []byte(uc.config.JwtSignature))
	if err != nil {
		return nil, usecases.TokenGenerationError
	}

	return authTokens, nil
}
