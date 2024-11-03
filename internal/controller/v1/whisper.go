package v1

import (
	"koizumi55555/whisper-api/internal/controller/model"
	"koizumi55555/whisper-api/internal/usecase"
	"koizumi55555/whisper-api/pkg/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

type apiRoutes struct {
	apiUC usecase.APIUseCase
	l     *logger.Logger
}

func NewAPIRoutes(
	handler *echo.Group, apiUC usecase.APIUseCase, l *logger.Logger,
) {
	r := &apiRoutes{
		apiUC: apiUC,
		l:     l,
	}
	handler.GET("/transcribe", r.Transcribe)
}

func (r *apiRoutes) Transcribe(c echo.Context) error {
	// validate
	filepath, err := model.ValidateTranscribe(c)
	if err != nil {
		return err
	}

	// request
	text, err := r.apiUC.TranscribeAudio(filepath)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, text)

}
