package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Level specifies what level the log should output.
type Level int

// This holds all of the log levels.
var (
	DebugLevel Level = 1
	InfoLevel  Level = 2
	WarnLevel  Level = 3
	ErrorLevel Level = 4
	FatalLevel Level = 5
)

// log is the global logger.
var log = logrus.New()

// GetLogger returns the global logger.
func GetLogger() *logrus.Logger {
	return log
}

// SetLevel will set the
func SetLevel(level Level) *logrus.Logger {
	switch level {
	case DebugLevel:
		log.SetLevel(logrus.DebugLevel)
	case InfoLevel:
		log.SetLevel(logrus.InfoLevel)
	case WarnLevel:
		log.SetLevel(logrus.WarnLevel)
	case ErrorLevel:
		log.SetLevel(logrus.ErrorLevel)
	case FatalLevel:
		log.SetLevel(logrus.FatalLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	return log
}

// SetLogPath will make the logger log to a file.
func SetLogPath(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	log.SetOutput(f)
	return nil
}
