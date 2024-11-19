package server

import (
	"context"
	"github.com/AskaryanKarine/bmstu-ds-2/pkg/app"
	"github.com/AskaryanKarine/bmstu-ds-2/pkg/models"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
	lr   loyaltyRepository
}

type loyaltyRepository interface {
	GetByUser(ctx context.Context, username string) (models.LoyaltyInfoResponse, error)
	UpdateByUser(ctx context.Context, username string, usersLoyalty models.LoyaltyInfoResponse) error
}

func NewServer(lr loyaltyRepository) *Server {
	e := echo.New()
	s := &Server{
		echo: e,
		lr:   lr,
	}

	app.SetStandardSetting(e)
	app.AddHealthCheck(e)

	api := s.echo.Group("/api/v1")

	api.GET("/loyalty", s.getLoyaltyByUser, app.GetUsernameMW())

	reservations := api.Group("/reservations")
	//reservations.POST("", s.createReservation, settings.GetUsernameMW())
	reservations.DELETE("/decrease", s.decreaseCounter, app.GetUsernameMW())

	return s
}

func (s *Server) Run(port int) {
	app.Run(s.echo, port)
}
