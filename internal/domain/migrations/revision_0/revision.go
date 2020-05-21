package revision_0

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/gormigrate.v1"
	"time"
)

func Migration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID:       "0",
		Migrate:  upgrade,
		Rollback: downgrade,
	}
}

type user struct {
	Id        uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:current_timestamp"`
	Username  string    `gorm:"type:text;not null"`
	Password  string    `gorm:"type:text;not null"`
}

func (p *user) TableName() string {
	return "users"
}

func upgrade(db *gorm.DB) error {
	var err error

	err = db.AutoMigrate(&user{}).Error
	if err != nil {
		return err
	}

	return err
}

func downgrade(db *gorm.DB) error {
	var err error

	err = db.DropTable((&user{}).TableName()).Error
	if err != nil {
		return err
	}

	return err
}
