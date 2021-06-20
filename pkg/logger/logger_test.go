package logger_test

import (
	"github.com/chensienyong/stocky/pkg/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(suites *testing.T) {

	suites.Parallel()
	config := logger.Configuration{
		EnableConsole:     true,
		ConsoleJSONFormat: true,
		ConsoleLevel:      logger.Info,
	}
	suites.Run("Should return an error when invalid level is passed", func(t *testing.T) {
		invalidConfig := config
		invalidConfig.ConsoleLevel = "invalid"
		err := logger.NewLogger(invalidConfig)
		assert.Error(t, err)
		assert.Equal(t, err, logger.ErrInvalidLogLevel)
	})

	suites.Run("Should create a logger when config is proper", func(t *testing.T) {
		err := logger.NewLogger(config)
		logger.WithFields(logger.Fields{ "methods": "GET" })
		assert.NoError(t, err)
		assert.NotNil(t, logger.Log)
	})
}
