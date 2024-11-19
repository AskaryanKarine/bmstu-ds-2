package server

import (
	innermodels "github.com/AskaryanKarine/bmstu-ds-2/internal/reservation/models"
	"github.com/AskaryanKarine/bmstu-ds-2/pkg/models"
	"github.com/AskaryanKarine/bmstu-ds-2/pkg/validation"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Server) getAllHotels(c echo.Context) error {
	var qParams innermodels.PaginationParams

	if err := c.Bind(&qParams); err != nil {
		return c.JSON(http.StatusBadRequest,
			models.ErrorResponse{Message: err.Error()},
		)
	}

	if err := c.Validate(qParams); err != nil {
		return c.JSON(http.StatusBadRequest, validation.ConvertToError(err, "failed to validate query params"))
	}

	result, count, err := s.hs.GetAllHotels(c.Request().Context(), qParams)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, models.PaginationResponse{
		Page:          qParams.Page,
		PageSize:      qParams.Size,
		TotalElements: count,
		Items:         result,
	})
}