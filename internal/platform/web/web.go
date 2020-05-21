package web

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/go-chi/chi"
)

type Web struct {
	config *WebConfig
	http   *http.Server
}

func NewWeb(httpConf *WebConfig, mux http.Handler) (*Web, error) {
	h := &http.Server{
		Addr:         httpConf.Addr(),
		Handler:      mux,
		ReadTimeout:  httpConf.ReadTimeout,
		WriteTimeout: httpConf.WriteTimeout,
	}

	web := &Web{
		config: httpConf,
		http:   h,
	}

	return web, nil
}

func (a *Web) NewRouter() chi.Router {
	return chi.NewRouter()
}

func (a *Web) ListenAndServer() error {
	log.Infof("web: Listen %s", a.config.Addr())
	if err := a.http.ListenAndServe(); err != nil {
		return nil
	}

	return nil
}

func (a *Web) Close() error {
	err := a.http.Close()

	if err != nil {
		return err
	}

	return nil
}

func (a *Web) Shutdown(ctx context.Context) error {
	err := a.http.Shutdown(ctx)

	if err != nil {
		return err
	}

	return nil
}
