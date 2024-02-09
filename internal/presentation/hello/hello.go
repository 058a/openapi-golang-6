package hello

import (
	oapicodegen "openapi/internal/infrastructure/oapicodegen/hello"

	"github.com/labstack/echo/v4"
)

type Api struct {}

func New() oapicodegen.ServerInterface {
	return &Api{}
}

func RegisterHandlers(e *echo.Echo, si oapicodegen.ServerInterface) {
	oapicodegen.RegisterHandlers(e, si)
}

func (a *Api) GetHello(ctx echo.Context) error {
	return GetHello(ctx)
}