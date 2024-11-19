package server

import (
	"github.com/AskaryanKarine/bmstu-ds-2/pkg/models"
)

type loyaltyClient interface {
	GetLoyaltyByUser(username string) (models.LoyaltyInfoResponse, error)
	DecreaseLoyalty(username string) error
}

type reservationClient interface {
	GetHotels(page, size int) (models.PaginationResponse, error)
	GetReservationByUUID(username, uuid string) (models.ExtendedReservationResponse, error)
	GetReservationsByUser(username string) ([]models.ExtendedReservationResponse, error)
	CancelReservation(username, uuid string) error
}

type paymentClient interface {
	GetByUUID(uuid string) (models.PaymentInfo, error)
	Cancel(uuid string) error
}
