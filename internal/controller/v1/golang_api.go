package v1

import (
	"net/http"

	"github.com/koizumi55555/whisper-api/pkg/logger"

	"github.com/github.com/koizumi55555/whisper-api/internal/usecase"
	"github.com/labstack/echo"
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
	handler.GET("/t", r.GetUsers)
}

func (r *apiRoutes) GetUsers(c echo.Context) error {

	users, err := r.apiUC.GetUsers(c)
	if err != nil {
		return err
	}

	m := []model.User{
		{
			Id:       users[0].UserName,
			Username: users[0].UserName,
			Email:    users[0].UserName,
		},
	}

	c.JSON(http.StatusOK, m)
	return nil
}
