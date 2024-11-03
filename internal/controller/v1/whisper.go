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
		return c.JSON(http.StatusBadRequest, model.InlineResponse400{
			Error: err.Error(),
		})
	}

	// request
	text, err := r.apiUC.TranscribeAudio(filepath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.InlineResponse500{
			Error: err.Error(),
		})
	}

	// replay 200
	return c.JSON(http.StatusOK, model.InlineResponse200{
		Text: text,
	})

}
