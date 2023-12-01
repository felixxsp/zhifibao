package usecases

import (
	"context"
	"zhifubao/domain/entity"
)

type Transaction_Usecase interface {
	NewTransaction(context.Context, entity.Transaction) (int, error)
	ViewTransaction(context.Context, entity.Trc_req_one) (entity.Transaction, error)
	ViewMulti(context.Context, entity.Trc_req_multi) ([]entity.Transaction, error)
}
