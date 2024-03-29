package locations

import (
	"net/http"

	"github.com/labstack/echo/v4"

	app "openapi/internal/app/stock/location"
	domain "openapi/internal/domain/stock/location"
	"openapi/internal/infra/database"
	infra "openapi/internal/infra/repository/sqlboiler/stock/location"

	openapi_types "github.com/oapi-codegen/runtime/types"

	oapicodegen "openapi/internal/infra/oapicodegen/stock"
)

// Put is a function that handles the HTTP PUT request for updating an existing stock location.
func PutStockLocation(ctx echo.Context, stockLocationId openapi_types.UUID) error {
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

	// Binding
	req := &oapicodegen.PutStockLocationJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Precondition
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

	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Main Process
	reqDto := &app.UpdateRequestDto{
		Id:   stockLocationId,
		Name: req.Name,
	}
	err = app.Update(reqDto, repository)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Postprocess
	return ctx.JSON(http.StatusOK, nil)
}
