package main

import (
	"fmt"
	"github.com/ShagaleevAlexey/go-clean/internal/config"
	"github.com/ShagaleevAlexey/go-clean/internal/domain/db"
	"github.com/ShagaleevAlexey/go-clean/internal/domain/queries"
	web2 "github.com/ShagaleevAlexey/go-clean/internal/platform/web"
	"github.com/ShagaleevAlexey/go-clean/internal/presenters/web"
	users2 "github.com/ShagaleevAlexey/go-clean/internal/services/users"
	"github.com/ShagaleevAlexey/go-clean/internal/usecases/users"
	"github.com/ShagaleevAlexey/go-clean/internal/utils/logging"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"time"
)

func webServe(webConf *web2.WebConfig, appConf *config.AppConfig, database *db.DB, queriesFactory *queries.QueriesFactory, c chan IMessage) {
	r := web2.NewRouter(chi.NewRouter())

	co := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(co.Handler)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", web2.HealthHandler)

	usersUseCase := makeUsersUsecase(database, queriesFactory, appConf)

	web.Mounts(r, usersUseCase)

	webApp, err := web2.NewWeb(webConf, r)
	if err != nil {
		c <- NewMessage(err)
		return
	}

	if err := webApp.ListenAndServer(); err != nil {
		c <- NewMessage(err)
		return
	}
}

func makeUsersUsecase(database *db.DB, queriesFactory *queries.QueriesFactory, appConf *config.AppConfig) *users.UsersUseCases {
	usersService := users2.NewUsersServices(database, appConf, queriesFactory)
	return users.NewUsersUseCases(database, appConf, queriesFactory, usersService)
}

func main() {
	var err error

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	appConf, err := config.NewAppConfigFromEnv()
	if err != nil {
		log.Fatalf("<backend> failed parse app config: %s", err)
	}

	webConf, err := web2.NewWebConfigFromEnv()
	if err != nil {
		log.Fatalf("<backend> failed parse web config: %s", err)
	}

	loggingConfig, err := logging.NewLoggingConfigFromEnv()
	if err != nil {
		log.Fatalf("<backend> failed parse logging config: %s", err)
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

	queriesFactory := queries.NewQueriesFactory()

	c := make(chan IMessage)
	go webServe(webConf, appConf, database, queriesFactory, c)

	select {
	case msg := <-c:
		log.Fatalf("<backend> failed: %v", msg)
	}
}

type IMessage interface {
	String() string
}

type Message struct {
	err error
}

func NewMessage(err error) *Message {
	return &Message{
		err: err,
	}
}

func (m *Message) String() string {
	return fmt.Sprintf("%v", m.err)
}
