package main

import (
	"github.com/ShagaleevAlexey/go-clean/internal/domain/db"
	"github.com/ShagaleevAlexey/go-clean/internal/domain/migrations"
	"github.com/ShagaleevAlexey/go-clean/internal/utils/logging"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func upgrade(args []string) {
	var err error

	loggingConfig, err := logging.NewLoggingConfigFromEnv()
	if err != nil {
		log.Fatalf("<rule35 server> failed parse logging config: %s", err)
	}

	logger := logging.NewLogging(loggingConfig)

	dbConf, err := db.NewDBSQLiteConfig()
	if err != nil {
		log.Fatalf("<backend> failed parse db config: %s", err)
	}

	database, err := db.NewDB(dbConf, logger)
	if err != nil {
		log.Fatalf("<backend> failed db connection: %s", err)
	}
	defer database.Close()

	ms := migrations.InitMigrationService(database)
	ms.UpgradeProcess()
}

func rollback(args []string) {
	revId := args[0]
	log.Infof("Rollback to '%s' revision", revId)
	var err error

	loggingConfig, err := logging.NewLoggingConfigFromEnv()
	if err != nil {
		log.Fatalf("<rule35 server> failed parse logging config: %s", err)
	}

	logger := logging.NewLogging(loggingConfig)

	dbConf, err := db.NewDBSQLiteConfig()
	if err != nil {
		log.Fatalf("<backend> failed parse db config: %s", err)
	}

	database, err := db.NewDB(dbConf, logger)
	if err != nil {
		log.Fatalf("<backend> failed db connection: %s", err)
	}
	defer database.Close()

	ms := migrations.InitMigrationService(database)
	ms.RollbackProcess(revId)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	upgrade(nil)
}