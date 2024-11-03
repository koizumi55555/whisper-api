package usecase

import (
	"github.com/github.com/koizumi55555/whisper-api/internal/entitiy"
	"github.com/labstack/echo"
)

type (
	APIUseCase interface {
		GetUsers(ctx echo.Context) ([]entitiy.User, error)
	}
	MasterRepository interface {
		GetUsers(ctx echo.Context) ([]entitiy.User, error)
	}

	AuthUseCase interface {
		GetBasicAuthUsers(ctx echo.Context, userName string) (entitiy.BasicAuthUser, error)
	}
	AuthRepository interface {
		GetBasicAuthUser(ctx echo.Context, userName string) (entitiy.BasicAuthUser, error)
	}

	KmsRepository interface{}
)
