package locations

import (
	"net/http"

	"github.com/labstack/echo/v4"

	app "openapi/internal/app/stock/location"
	"openapi/internal/infra/database"
	oapicodegen "openapi/internal/infra/oapicodegen/stock"
	infra "openapi/internal/infra/repository/sqlboiler/stock/location"
)

// PostStockLocation is a function that handles the HTTP POST request for creating a new stock item.
func PostStockLocation(ctx echo.Context) error {
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
	req := &oapicodegen.PostStockLocationJSONRequestBody{}
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Precondition
	if err := ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Main Process
	reqDto := &app.CreateRequestDto{
		Name: req.Name,
	}
	resDto, err := app.Create(reqDto, repository)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Postprocess
	res := &oapicodegen.Created{Id: resDto.Id}

	// Postcondition
	if err := ctx.Validate(res); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, res)
}
