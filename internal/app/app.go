package app

import (
	"fmt"

	"github.com/koizumi55555/whisper-api/config"
	v1 "github.com/koizumi55555/whisper-api/internal/controller/v1"
	"github.com/koizumi55555/whisper-api/internal/usecase"
	"github.com/koizumi55555/whisper-api/internal/usecase/master_repo"
	"github.com/koizumi55555/whisper-api/pkg/logger"

	"github.com/labstack/echo"
)

func Run(cfg *config.Config) error {

	handler := echo.New()
	l := logger.Logger{}
	// new
	mRepo := master_repo.New(&l)
	apiUc := usecase.NewAPIUsecase(mRepo, &l)

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
