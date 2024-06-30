package logger

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	log.Logger
}

func New(logLevel string, outputFile string) (*Logger, error) {

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return nil, fmt.Errorf("invalid log level: %s", level)
	}

	src := os.Stderr

	if len(outputFile) != 0 {
		src, err = os.Create(outputFile)
		if err != nil {
			return nil, fmt.Errorf("%w: error opening log file", err)
		}
	}

	return &Logger{
		Logger: log.Logger{
			Level: level,
			Out:   src,
			Formatter: &log.JSONFormatter{
				PrettyPrint: true,
			},
		},
	}, nil
}
