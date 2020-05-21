package users

import (
	"github.com/ShagaleevAlexey/go-clean/internal/config"
	"github.com/ShagaleevAlexey/go-clean/internal/domain/db"
	"github.com/ShagaleevAlexey/go-clean/internal/domain/queries"
)

type IUsersServices interface {

}

type UsersServices struct {
	db *db.DB
	appConf *config.AppConfig
	queries queries.IQueries
}

func NewUsersServices(db *db.DB, appConf *config.AppConfig, queries queries.IQueries) *UsersServices {
	return &UsersServices{
		db: db,
		appConf: appConf,
		queries: queries,
	}
}
