package person

import (
	"context"
	"fmt"
	"net/http"
	"zhifubao/domain/entity"
	"zhifubao/domain/repository"

	"github.com/google/uuid"
)

type RealUsecase struct {
	datalayer repository.Person_Database
}

func NewUsecase(datalayer repository.Person_Database) *RealUsecase {
	return &RealUsecase{
		datalayer: datalayer,
	}
}

func (uc *RealUsecase) NewPerson(ctx context.Context, person entity.Person) int {
	if !uc.datalayer.IsExist(ctx, person) {
		return 500
	}
	person.UUID, _ = uuid.NewUUID()
	person.Balance = 0.0
	person.Active = false
	fmt.Println(person)
	return uc.datalayer.NewPerson(ctx, person)
}

func (uc *RealUsecase) Login(ctx context.Context, info entity.Login_req) int {
	if uc.datalayer.IsActive(ctx, info) {
		return http.StatusConflict
	}
	check := uc.datalayer.GetPerson(ctx, info)
	if info.Password == check.Password {
		return uc.datalayer.Login(ctx, info)
	}
	return http.StatusBadRequest
}

func (uc *RealUsecase) Logout(ctx context.Context, info entity.Login_req) int {
	if uc.datalayer.IsActive(ctx, info) {
		return uc.datalayer.Logout(ctx, info)
	}
	return 400
}
