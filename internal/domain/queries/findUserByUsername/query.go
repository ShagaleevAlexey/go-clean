package findUserByUsername

import (
	"github.com/ShagaleevAlexey/go-clean/internal/domain/db"
	"github.com/ShagaleevAlexey/go-clean/internal/domain/models"
	log "github.com/sirupsen/logrus"
)

type IFindUserByUsernameQuery interface {
	FindUserByUsername(*db.Session, string) (*models.User, error)
}

type FindUserByUsernameQuery struct {
}

func NewFindUserByUsernameQuery() *FindUserByUsernameQuery {
	return &FindUserByUsernameQuery{}
}

func (q *FindUserByUsernameQuery) FindUserByUsername(session *db.Session, u string) (*models.User, error) {
	var err error

	user := models.User{}
	if err = session.Gorm.Model(&models.User{}).Where("username = ?", u).First(&user).Error; err != nil {
		log.Errorf("<FindUserByUsernameQuery> %s", err)
		return nil, err
	}

	return &user, session.Gorm.Commit().Error
}
