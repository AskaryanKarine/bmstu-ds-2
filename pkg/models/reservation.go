package models

type ReservationResponse struct {
	ReservationUid string            `json:"reservationUid"`
	Hotel          HotelInfo         `json:"hotel"`
	StartDate      string            `json:"startDate"`
	EndDate        string            `json:"endDate"`
	Status         PaymentStatusType `json:"status"`
	Payment        PaymentInfo       `json:"payment"`
}

type CreateReservationRequest struct {
	HotelUid  string `json:"hotelUid"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type CreateReservationResponse struct {
	ReservationUid string            `json:"reservationUid"`
	Discount       string            `json:"discount"`
	Status         PaymentStatusType `json:"status"`
	Payment        PaymentInfo       `json:"payment"`
	CreateReservationRequest
}
