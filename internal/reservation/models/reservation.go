package models

import "github.com/AskaryanKarine/bmstu-ds-2/pkg/models"

type ReservationResponse struct {
	ReservationUid string
	StartDate      string
	EndDate        string
	Status         models.PaymentStatusType
	PaymentUID     string
}

func (r *ReservationResponse) ToResponse(info models.HotelInfo) models.ExtendedReservationResponse {
	return models.ExtendedReservationResponse{
		ReservationResponse: models.ReservationResponse{
			ReservationUid: r.ReservationUid,
			StartDate:      r.StartDate,
			EndDate:        r.EndDate,
			Status:         r.Status,
			Hotel:          info,
		},
		PaymentUID: r.PaymentUID,
	}
}

type ReservationTable struct {
	ReservationUid string
	Username       string
	PaymentUid     string
	HotelID        int
	Status         models.PaymentStatusType
	StartDate      string
	EndDate        string
}

func ToReservationTable(model models.ExtendedCreateReservationResponse) ReservationTable {
	return ReservationTable{
		PaymentUid: model.PaymentUid,
		StartDate:  model.StartDate,
		EndDate:    model.EndDate,
	}
}
