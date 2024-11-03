package v1

import (
	"github.com/koizumi55555/whisper-api/pkg/logger"

	"github.com/github.com/koizumi55555/whisper-api/internal/usecase"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(
	handler *echo.Echo,
	apiUc usecase.APIUseCase,
	l *logger.Logger,
) error {
	// Zerolog を使ったリクエストのロギングをミドルウェアとして追加
	handler.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return l.NewHandlerZerolog(next)
	})
	handler.Use(middleware.Recover())

	v1h := handler.Group("/v1")

	NewAPIRoutes(v1h, apiUc, l)
	return nil
}
