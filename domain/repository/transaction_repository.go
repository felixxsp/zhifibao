package repository

import (
	"context"
	"zhifubao/domain/entity"
)

type Transaction_Database interface {
	NewTransaction(context.Context, entity.Transaction) int
	ViewTransaction(context.Context, entity.Trc_req_one) entity.Transaction
	ViewMulti(context.Context, entity.Trc_req_multi) []entity.Transaction
}
