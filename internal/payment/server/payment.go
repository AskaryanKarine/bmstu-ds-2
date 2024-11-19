package server

import (
	"errors"
	"fmt"
	"github.com/AskaryanKarine/bmstu-ds-2/pkg/models"
	"github.com/AskaryanKarine/bmstu-ds-2/pkg/validation"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func (s *Server) getPaymentInfo(c echo.Context) error {
	uid := c.Param("uid")
	err := c.Validate(struct {
		Uid string `json:"uid" validate:"uuid"`
	}{uid})
	if err != nil {
		return c.JSON(http.StatusBadRequest, validation.ConvertToError(err, "failed to validate uid in path"))
	}

	res, err := s.ps.GetPaymentInfoByUUID(c.Request().Context(), uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, models.ErrorResponse{Message: fmt.Sprintf("payment info with %s uid not found", uid)})
		}
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func (s *Server) setCanceledStatus(c echo.Context) error {
	uid := c.Param("uid")
	err := c.Validate(struct {
		Uid string `json:"uid" validate:"uuid"`
	}{uid})
	if err != nil {
		return c.JSON(http.StatusBadRequest, validation.ConvertToError(err, "failed to validate uid in path"))
	}
	err = s.ps.Delete(c.Request().Context(), uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusNoContent, echo.Map{})
}
