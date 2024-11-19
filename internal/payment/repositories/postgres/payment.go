package postgres

import (
	"context"
	"fmt"
	"github.com/AskaryanKarine/bmstu-ds-2/pkg/models"
	"gorm.io/gorm"
)

type paymentStorage struct {
	db *gorm.DB
}

func NewPaymentStorage(db *gorm.DB) *paymentStorage {
	return &paymentStorage{
		db: db,
	}
}

const (
	paymentTable = "payment p"
)

func (p *paymentStorage) GetPaymentInfoByUUID(ctx context.Context, uuid string) (models.PaymentInfo, error) {
	var paymentInfo models.PaymentInfo

	err := p.db.WithContext(ctx).Table(paymentTable).Where("p.payment_uid = ?", uuid).Take(&paymentInfo).Error
	if err != nil {
		return models.PaymentInfo{}, fmt.Errorf("failed to get reservation by uuid %s: %w", uuid, err)
	}

	return paymentInfo, nil
}

func (p *paymentStorage) Delete(ctx context.Context, uuid string) error {
	err := p.db.WithContext(ctx).Table(paymentTable).Where("p.payment_uid = ?", uuid).Update("status", models.CANCELED).Error
	if err != nil {
		return fmt.Errorf("failed deleting payment %s: %w", uuid, err)
	}
	return nil
}
