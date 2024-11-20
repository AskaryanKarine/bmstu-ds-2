package models

import "github.com/AskaryanKarine/bmstu-ds-2/pkg/models"

type Payment struct {
	PaymentUid string
	Status     models.PaymentStatusType
	Price      int
}
