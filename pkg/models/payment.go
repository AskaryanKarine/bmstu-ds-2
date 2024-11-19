package models

type PaymentStatusType string

const (
	PAID     PaymentStatusType = "PAID"
	CANCELED PaymentStatusType = "CANCELED"
)

type PaymentInfo struct {
	// Status - статус операции оплаты
	Status PaymentStatusType `json:"status"`
	// Price - сумма операции
	Price int `json:"price"`
}
