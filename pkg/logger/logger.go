package logger

import (
	"errors"
)

// Configuration stores the config for the logger
type Configuration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         string
	FileLocation      string
}

const (
	//Debug has verbose message
	Debug = "debug"
	//Info is default log level
	Info = "info"
	//Warn is for logging messages about possible issues
	Warn = "warn"
	//Error is for logging errors
	Error = "error"
	//Fatal is for logging fatal messages. The system shuts down after logging the message.
	Fatal = "fatal"
)

var ErrInvalidLogLevel = errors.New("Invalid log level")

//Logger is our contract for the logger
type Logger interface {
	Debugf(format string, args ...interface{})

	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})

	Panicf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger
}

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

var Log Logger

//NewLogger returns an instance of logger
func NewLogger(config Configuration) error {
	logger, err := newZapLogger(config)
	if err != nil {
		return err
	}
	Log = logger
	return nil
}

func Debugf(format string, args ...interface{}) {
	Log.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	Log.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	Log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	Log.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	Log.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	Log.Panicf(format, args...)
}

func WithFields(keyValues Fields) Logger {
	return Log.WithFields(keyValues)
}

func init() {
	defaultConfig := Configuration{
		EnableConsole:     true,
		ConsoleLevel:      Info,
		ConsoleJSONFormat: true,
	}
	NewLogger(defaultConfig)
}
