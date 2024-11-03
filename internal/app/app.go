package app

import (
	"fmt"
	"koizumi55555/whisper-api/config"
	v1 "koizumi55555/whisper-api/internal/controller/v1"
	"koizumi55555/whisper-api/internal/usecase"
	"koizumi55555/whisper-api/pkg/logger"

	"github.com/labstack/echo/v4"
)

func Run(cfg *config.Config) error {

	handler := echo.New()
	l := logger.Logger{}
	// new
	apiUc := usecase.NewAPIUsecase(&l)

	// echo

	if err := v1.NewRouter(handler, apiUc, &l); err != nil {
		return fmt.Errorf("/v1 handler error: %w", err)
	}
	err := handler.Start(":8080")
	if err != nil {
		return err
	}
	return nil
}
