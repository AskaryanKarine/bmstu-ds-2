package server

import (
	"context"
	inner_models "github.com/AskaryanKarine/bmstu-ds-2/internal/payment/models"
	"github.com/AskaryanKarine/bmstu-ds-2/pkg/app"
	"github.com/AskaryanKarine/bmstu-ds-2/pkg/models"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
	ps   paymentStorage
}

type paymentStorage interface {
	GetPaymentInfoByUUID(ctx context.Context, uuid string) (models.PaymentInfo, error)
	Delete(ctx context.Context, uuid string) error
	Create(ctx context.Context, payment inner_models.Payment) (string, error)
}

func NewServer(ps paymentStorage) *Server {
	e := echo.New()
	s := &Server{
		echo: e,
		ps:   ps,
	}

	app.SetStandardSetting(e)
	app.AddHealthCheck(e)

	api := s.echo.Group("/api/v1")

	api.GET("/payments/:uid", s.getPaymentInfo)
	api.POST("/payments", s.CreatePayment)

	api.DELETE("/reservations/:uid", s.setCanceledStatus)

	return s
}

func (s *Server) Run(port int) {
	app.Run(s.echo, port)
}
