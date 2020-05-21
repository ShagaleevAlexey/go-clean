package logging

import (
	"os"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

func NewLogging(ca *LoggingConfig) *log.Logger {
	switch ca.LogFormat {
	case PlainFormat:
		log.SetFormatter(&log.TextFormatter{
			ForceColors:      true,
			CallerPrettyfier: callerPrettyfier,
		})
	case JsonFormat:
		log.SetFormatter(&log.JSONFormatter{
			CallerPrettyfier: callerPrettyfier,
		})
	}

	log.SetLevel(ca.LogLevel)
	log.SetReportCaller(true)

	return log.StandardLogger()
}

func callerPrettyfier(f *runtime.Frame) (string, string) {
	fn := strings.Split(f.Function, " ")[0]
	dir, err := os.Getwd()

	if err != nil {
		dir = f.File
	} else {
		dir = strings.Replace(f.File, dir, "", 1)
	}

	return fn, dir
}
