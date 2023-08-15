package repository

import (
	"context"

	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/constant"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/db"
	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/model"
)

type paymentRepo struct {
	db *db.DB
}

func NewPaymentRepository(db *db.DB) *paymentRepo {
	return &paymentRepo{db}
}

func (p *paymentRepo) Create(ctx context.Context, data *model.Payment) (*model.Payment, error) {
	data.Status = string(constant.PaymentStatusPaid)

	return data, nil
}
