package repository

import (
	"context"
	"github.com/pkg/errors"

	"github.com/hoangtk0100/dc-go-23/ex_07/pkg/model"
)

type paymentRepo struct {
	db *DB
}

func NewPaymentRepository(db *DB) *paymentRepo {
	return &paymentRepo{db}
}

func (p *paymentRepo) Create(ctx context.Context, data *model.Payment) (*model.Payment, error) {
	tx := p.db.Begin()
	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, errors.WithStack(err)
	}

	return data, nil
}
