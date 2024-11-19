package models

import "github.com/AskaryanKarine/bmstu-ds-2/pkg/models"

type ExpandedLoyalty struct {
	models.LoyaltyInfoResponse
	Discount int
}
