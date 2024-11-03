package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/rs/zerolog"
)

type (
	Interface interface {
		Debug(message interface{}, args ...interface{})
		Info(message string, args ...interface{})
		Warn(message string, args ...interface{})
		Error(message interface{}, args ...interface{})
		Fatal(message interface{}, args ...interface{})
	}

	Logger struct {
		zeroLogger *zerolog.Logger
	}
)

// init zerolog
func (l *Logger) NewHandlerZerolog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// リクエストの開始時間
		start := time.Now()

		// リクエスト処理を実行
		err := next(c)

		// リクエストに関する情報をログに出力
		l.zeroLogger.Info().
			Str("method", c.Request().Method).
			Str("path", c.Request().URL.Path).
			Int("status", c.Response().Status).
			Str("remote_ip", c.RealIP()).
			Dur("latency", time.Since(start)).
			Msg("handled request")

		return err
	}
}

// New -.
func New(level string) *Logger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)

	skipFrameCount := 3
	logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &Logger{
		zeroLogger: &logger,
	}
}

// Debug -.
func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.msg("debug", message, args...)
}

// Info -.
func (l *Logger) Info(message string, args ...interface{}) {
	l.log(message, args...)
}

// Warn -.
func (l *Logger) Warn(args ...interface{}) {}

// Error -.
func (l *Logger) Error(message interface{}, args ...interface{}) {
	if l.zeroLogger.GetLevel() == zerolog.DebugLevel {
		l.Debug(message, args...)
	}

	l.msg("error", message, args...)
}

// Errorf -.
func (l *Logger) Errorf(tmp string, args ...interface{}) {}

// Fatal -.
func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.msg("fatal", message, args...)

	os.Exit(1)
}

func (l *Logger) log(message string, args ...interface{}) {
	if len(args) == 0 {
		l.zeroLogger.Info().Msg(message)
	} else {
		l.zeroLogger.Info().Msgf(message, args...)
	}
}

func (l *Logger) msg(level string, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.log(msg.Error(), args...)
	case string:
		l.log(msg, args...)
	default:
		l.log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}
