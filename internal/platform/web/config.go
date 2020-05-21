package web

import (
	"fmt"
	"time"

	"github.com/caarlos0/env"
)

type IWebConfig interface {
	Addr() string
}

type WebConfig struct {
	Host         string        `json:"host,omitempty" env:"HTTP_HOST" envDefault:"0.0.0.0"`
	Port         int           `json:"port,omitempty" env:"HTTP_PORT" envDefault:"49876"`
	ReadTimeout  time.Duration `json:"read_timeout,omitempty" env:"HTTP_READ_TIMEOUT" envDefault:"60s"`
	WriteTimeout time.Duration `json:"write_timeout,omitempty" env:"HTTP_WRITE_TIMEOUT" envDefault:"60s"`
}

func NewWebConfigFromEnv() (*WebConfig, error) {
	c := &WebConfig{}
	if err := c.Import(); err != nil {
		return nil, err
	}

	return c, nil
}

func (h *WebConfig) Import() error {
	if err := env.Parse(h); err != nil {
		return err
	}

	return nil
}

func (h *WebConfig) Addr() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}
