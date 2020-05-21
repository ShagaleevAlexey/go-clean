package logging

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

type LogFormat uint32

const (
	PlainFormat LogFormat = 0
	JsonFormat  LogFormat = 1
)

type LoggingConfig struct {
	LogLevel  logrus.Level `json:"log_level,omitempty" env:"LOG_LEVEL" envDefault:"debug"`
	LogFormat LogFormat    `json:"log_format,omitempty" env:"LOG_FORMAT" envDefault:"plain"`
}

func NewLoggingConfigFromEnv() (*LoggingConfig, error) {
	loggingConf := &LoggingConfig{}

	funcs := map[reflect.Type]env.ParserFunc{}
	type Level uint32

	funcs[reflect.TypeOf(logrus.DebugLevel)] = parseLogLevel
	funcs[reflect.TypeOf(PlainFormat)] = parseLogFormat

	if err := env.ParseWithFuncs(loggingConf, funcs); err != nil {
		return nil, err
	}

	return loggingConf, nil
}

func parseLogLevel(in string) (interface{}, error) {
	return logrus.ParseLevel(in)
}

func parseLogFormat(fm string) (interface{}, error) {
	switch strings.ToLower(fm) {
	case "plain":
		return PlainFormat, nil
	case "fatal":
		return JsonFormat, nil
	}

	var f LogFormat
	return f, fmt.Errorf("not a valid Format: %q", fm)
}
