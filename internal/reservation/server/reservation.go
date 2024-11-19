package server

import (
	"errors"
	"github.com/AskaryanKarine/bmstu-ds-2/pkg/models"
	"github.com/AskaryanKarine/bmstu-ds-2/pkg/validation"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func (s *Server) createReservation(c echo.Context) error {
	var body models.CreateReservationRequest

	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
	}

	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
	}

	return nil
}

func (s *Server) getReservationByUid(c echo.Context) error {
	username, ok := c.Get("username").(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "failed to get username"})
	}

	uid := c.Param("uid")
	err := c.Validate(struct {
		Uid string `json:"uid" validate:"uuid"`
	}{uid})
	if err != nil {
		return c.JSON(http.StatusBadRequest, validation.ConvertToError(err, "failed to validate uid in path"))
	}

	res, err := s.rs.GetReservationByUUID(c.Request().Context(), uid, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, models.ErrorResponse{Message: err.Error()})
		}
		if errors.Is(err, models.WrongUsernameError) {
			return c.JSON(http.StatusForbidden, models.ErrorResponse{Message: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (s *Server) getAllReservationsByUser(c echo.Context) error {
	username, ok := c.Get("username").(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "failed to get username"})
	}

	res, err := s.rs.GetAllReservationByUsername(c.Request().Context(), username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, models.ErrorResponse{Message: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func (s *Server) canceledReservations(c echo.Context) error {
	username, ok := c.Get("username").(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "failed to get username"})
	}

	uid := c.Param("uid")
	err := c.Validate(struct {
		Uid string `json:"uid" validate:"uuid"`
	}{uid})
	if err != nil {
		return c.JSON(http.StatusBadRequest, validation.ConvertToError(err, "failed to validate uid in path"))
	}

	_, err = s.rs.GetReservationByUUID(c.Request().Context(), uid, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, models.ErrorResponse{Message: err.Error()})
		}
		if errors.Is(err, models.WrongUsernameError) {
			return c.JSON(http.StatusForbidden, models.ErrorResponse{Message: err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}

	err = s.rs.Delete(c.Request().Context(), uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusNoContent, echo.Map{})
}
