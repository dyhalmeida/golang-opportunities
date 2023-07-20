package config

import "errors"

var (
	logger *Logger
)

func Init() error {
	return errors.New("fake error")
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
