package migrations

import (
	"github.com/ShagaleevAlexey/go-clean/internal/domain/db"
	"github.com/ShagaleevAlexey/go-clean/internal/domain/migrations/revision_0"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gormigrate.v1"
)

type MigrationService struct {
	DB *db.DB
}

func InitMigrationService(db *db.DB) *MigrationService {
	return &MigrationService{
		DB: db,
	}
}

var migrationOptions = &gormigrate.Options{
	TableName:                 "migrations",
	IDColumnName:              "id",
	IDColumnSize:              255,
	UseTransaction:            true,
	ValidateUnknownMigrations: false,
}

func (ms *MigrationService) UpgradeProcess() {
	defer log.Info("Migration was successfully")

	ms.DB.Gorm.LogMode(true)

	revisions := make([]*gormigrate.Migration, 0)

	for _, f := range revFuncs {
		revisions = append(revisions, f())
	}

	m := gormigrate.New(ms.DB.Gorm, migrationOptions, revisions)

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migration: %v", err)
	}
}

func (ms *MigrationService) RollbackProcess(id string) {
	defer log.Info("Rollback was successfully")

	revisions := make([]*gormigrate.Migration, 0)

	for _, f := range revFuncs {
		revisions = append(revisions, f())
	}

	m := gormigrate.New(ms.DB.Gorm, migrationOptions, revisions)

	if err := m.RollbackTo(id); err != nil {
		log.Fatalf("Could not rollback: %v", err)
	}
}

var revFuncs = []func() *gormigrate.Migration{
	revision_0.Migration,
}
