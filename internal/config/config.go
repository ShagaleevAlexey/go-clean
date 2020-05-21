package config

import (
	"reflect"
	"github.com/caarlos0/env"
)

type AppConfig struct {
	JwtSignature string `json:"host,omitempty" env:"JWT_SIG"`
}

func NewAppConfigFromEnv() (*AppConfig, error) {
	appConf := &AppConfig{}
	if err := appConf.Import(); err != nil {
		return nil, err
	}

	return appConf, nil
}

func (ac *AppConfig) Import() error {
	funcs := map[reflect.Type]env.ParserFunc{}

	if err := env.ParseWithFuncs(ac, funcs); err != nil {
		return err
	}

	return nil
}
