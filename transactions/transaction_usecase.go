package transaction

import (
	"context"
	"net/http"
	"time"
	"zhifubao/domain/entity"
	"zhifubao/domain/repository"

	"github.com/google/uuid"
)

type RealUsecase struct {
	transactionData repository.Transaction_Database
	personData      repository.Person_Database
}

func NewUsecase(datalayer repository.Person_Database, self repository.Transaction_Database) *RealUsecase {
	return &RealUsecase{
		personData:      datalayer,
		transactionData: self,
	}
}

func (uc *RealUsecase) NewTransaction(ctx context.Context, item entity.Transaction) (int, error) {
	var sent entity.Trc_req_one
	sent.Device = item.Device
	sent.PersonID = item.Person
	sent.Transaction = item.UUID
	_, err := uc.personData.Autentication(ctx, sent)
	if err != nil {
		return http.StatusConflict, err
	}
	uc.personData.UpdateBalance(ctx, item)

	var completion entity.Login_req
	completion.Device = item.Device
	completion.PersonID = item.Person
	person := uc.personData.GetPerson(ctx, completion)

	item.Balance = person.Balance
	item.UUID, _ = uuid.NewUUID()
	item.Time = time.Now().Unix()

	uc.transactionData.NewTransaction(ctx, item)
	return http.StatusOK, nil
}

func (uc *RealUsecase) ViewTransaction(ctx context.Context, item entity.Trc_req_one) (entity.Transaction, error) {
	var result entity.Transaction
	_, err := uc.personData.Autentication(ctx, item)
	if err != nil {
		return result, err
	}
	return uc.transactionData.ViewTransaction(ctx, item), nil
}

func (uc *RealUsecase) ViewMulti(ctx context.Context, item entity.Trc_req_multi) ([]entity.Transaction, error) {
	var result []entity.Transaction
	var sent entity.Trc_req_one
	sent.Device = item.Device
	sent.PersonID = item.PersonID
	_, err := uc.personData.Autentication(ctx, sent)
	if err != nil {
		return result, err
	}
	return uc.transactionData.ViewMulti(ctx, item), nil
}
