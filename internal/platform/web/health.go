package web

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Health"))

	if err != nil {
		log.Error(err)
	}

	//w.WriteHeader(http.StatusOK)
}
