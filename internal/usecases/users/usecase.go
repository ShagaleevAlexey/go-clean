package users

import (
	"github.com/ShagaleevAlexey/go-clean/internal/config"
	"github.com/ShagaleevAlexey/go-clean/internal/domain/db"
	"github.com/ShagaleevAlexey/go-clean/internal/domain/queries"
	"github.com/ShagaleevAlexey/go-clean/internal/services/users"
)

type UsersUseCases struct {
	db       *db.DB
	config   *config.AppConfig
	queries  queries.IQueries
	services users.IUsersServices
}

func NewUsersUseCases(db *db.DB, appConf *config.AppConfig, queries queries.IQueries, services users.IUsersServices) *UsersUseCases {
	return &UsersUseCases{
		db:       db,
		config:   appConf,
		queries:  queries,
		services: services,
	}
}
