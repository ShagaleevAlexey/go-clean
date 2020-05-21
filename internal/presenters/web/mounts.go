package web

import (
	"context"
	"github.com/ShagaleevAlexey/go-clean/internal/platform/web"
	"github.com/ShagaleevAlexey/go-clean/internal/presenters/web/routes/users/auth"
	"github.com/ShagaleevAlexey/go-clean/internal/usecases"
	"github.com/ShagaleevAlexey/go-clean/internal/usecases/users"
	"github.com/ShagaleevAlexey/go-clean/internal/usecases/users/access_token"
)

type IUsersUsecase interface {
	AuthUser(ctx context.Context, input *users.UCAuthUserModel) (*access_token.AuthTokens, *usecases.AppError)
}

func Mounts(r web.Router, usersUsecase IUsersUsecase)  {
	r.Route("/users", func(r web.Router) {
		auth.MountAuthUser(r, usersUsecase.AuthUser)
	})
}
