package logger

import (
	"dtam-fund-cms-backend/config"
	"dtam-fund-cms-backend/domain/ports"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type ZeroLogger struct {
	log zerolog.Logger
}

func EstablishZeroLogger(
	cfg config.App,
) *ZeroLogger {

	level := parseLevel(cfg.LogLevel)
	zerolog.SetGlobalLevel(level)

	var writer io.Writer

	if cfg.Stage == "development" {
		writer = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	} else {
		writer = os.Stdout
	}

	logger := zerolog.New(writer).
		With().
		Timestamp().
		Str("env", cfg.Stage).
		Logger()

	return &ZeroLogger{log: logger}
}

func NewZeroLogger(zaplogger *ZeroLogger) ports.Logger {
	return zaplogger
}

func (zl *ZeroLogger) Info(msg string, field map[string]interface{}) {
	zl.log.Info().Fields(field).Msg(msg)
}

func (zl *ZeroLogger) Debug(msg string, field map[string]interface{}) {
	zl.log.Debug().Fields(field).Msg(msg)
}

func (zl *ZeroLogger) Warn(msg string, field map[string]interface{}) {
	zl.log.Warn().Fields(field).Msg(msg)
}

func (zl *ZeroLogger) Error(msg string, err error) {
	zl.log.Error().Err(err).Msg(msg)
}

func (zl *ZeroLogger) ErrorF(msg string, err error, fields map[string]interface{}) {
	zl.log.Error().Err(err).Fields(fields).Msg(msg)
}

func (zl *ZeroLogger) APILogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next() // process request

		latency := time.Since(start)

		fields := map[string]interface{}{
			"method":  c.Method(),
			"path":    c.Path(),
			"status":  c.Response().StatusCode(),
			"latency": latency.String(),
			"ip":      c.IP(),
		}

		if err != nil {
			zl.log.Error().Err(err).Fields(fields).Msg("API error")
			return err
		}

		zl.log.Debug().Fields(fields).Msg("API request")

		return nil
	}
}

func parseLevel(l string) zerolog.Level {
	switch strings.ToLower(l) {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	}
	return zerolog.InfoLevel
}
