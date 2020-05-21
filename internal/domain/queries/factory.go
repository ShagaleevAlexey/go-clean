package queries

import (
	"github.com/ShagaleevAlexey/go-clean/internal/domain/db"
	"github.com/ShagaleevAlexey/go-clean/internal/domain/models"
	"github.com/ShagaleevAlexey/go-clean/internal/domain/queries/findUserByUsername"
	"go.uber.org/dig"
)

type IQueries interface {
	findUserByUsername.IFindUserByUsernameQuery
}

type QueriesFactory struct {
	container *dig.Container
}

func NewQueriesFactory() *QueriesFactory {
	factory := &QueriesFactory{
		dig.New(),
	}

	factory.container.Provide(func() *findUserByUsername.FindUserByUsernameQuery {
		return findUserByUsername.NewFindUserByUsernameQuery()
	})

	return factory
}

func (qf *QueriesFactory) FindUserByUsername(session *db.Session, u string) (*models.User, error) {
	var (
		dbuser *models.User
		querryErr    error
	)
	err := qf.container.Invoke(func(q *findUserByUsername.FindUserByUsernameQuery) {
		dbuser, querryErr = q.FindUserByUsername(session, u)
	})
	if err != nil {
		return nil, err
	}

	return dbuser, querryErr
}
