package stock

import (
	oapicodegen "openapi/internal/infrastructure/oapicodegen/stock"
	"openapi/internal/presentation/stock/items"
	"openapi/internal/presentation/stock/locations"

	openapi_types "github.com/oapi-codegen/runtime/types"

	"github.com/labstack/echo/v4"
)

type Api struct {}

func New() oapicodegen.ServerInterface {
	return &Api{}
}

func RegisterHandlers(e *echo.Echo, si oapicodegen.ServerInterface) {
	oapicodegen.RegisterHandlers(e, si)
}

func (a *Api) PostStockLocation(ctx echo.Context) error {
	return locations.PostStockLocation(ctx)
}

func (a *Api) PutStockLocation(ctx echo.Context, stockLocationId openapi_types.UUID) error {
	return locations.PutStockLocation(ctx, stockLocationId)
}

func (a *Api) DeleteStockLocation(ctx echo.Context, stockLocationId openapi_types.UUID) error {
	return locations.DeleteStockLocation(ctx, stockLocationId)
}

func (a *Api) PostStockItem(ctx echo.Context) error {
	return items.PostStockItem(ctx)
}

func (a *Api) PutStockItem(ctx echo.Context, stockItemId openapi_types.UUID) error {
	return items.PutStockItem(ctx, stockItemId)
}

func (a *Api) DeleteStockItem(ctx echo.Context, stockItemId openapi_types.UUID) error {
	return items.DeleteStockItem(ctx, stockItemId)
}