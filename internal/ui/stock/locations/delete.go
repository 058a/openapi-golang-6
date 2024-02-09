package locations

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	app "openapi/internal/app/stock/location"
	domain "openapi/internal/domain/stock/location"
	"openapi/internal/infra/database"
	infra "openapi/internal/infra/repository/sqlboiler/stock/location"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Delete is a function that handles the HTTP DELETE request for deleting an existing stock item.
func DeleteStockLocation(ctx echo.Context, stockLocationId openapi_types.UUID) error {
	// Preprocess
	db, err := database.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer db.Close()

	repository, err := infra.NewRepository(db)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Prepcondition
	if stockLocationId == uuid.Nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid stock location id")
	}

	id, err := domain.NewId(stockLocationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	found, err := repository.Find(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if !found {
		return echo.NewHTTPError(http.StatusNotFound, "stock location not found")
	}

	// Main Process
	reqDto := &app.DeleteRequestDto{
		Id: stockLocationId,
	}
	if err := app.Delete(reqDto, repository); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Postprocess
	return ctx.JSON(http.StatusOK, nil)
}
