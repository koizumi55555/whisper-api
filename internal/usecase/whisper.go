package usecase

import (
	"github.com/github.com/koizumi55555/whisper-api/internal/entitiy"
	"github.com/github.com/koizumi55555/whisper-api/pkg/logger"
	"github.com/labstack/echo"
)

type apiUseCase struct {
	masterRepo MasterRepository
	l          *logger.Logger
}

func NewAPIUsecase(mRepo MasterRepository, l *logger.Logger) *apiUseCase {
	return &apiUseCase{
		masterRepo: mRepo,
		l:          l,
	}
}

func (u *apiUseCase) GetUsers(
	ctx echo.Context,
) ([]entitiy.User, error) {
	users, err := u.masterRepo.GetUsers(ctx)
	if err != nil {
		return []entitiy.User{}, err
	}
	return users, nil

}
