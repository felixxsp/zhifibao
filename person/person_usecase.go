package person

import (
	"context"
	"只服宝/entity"

	"github.com/google/uuid"
)

type RealUsecase struct {
	datalayer Database
}

type Usecase interface {
	NewPerson(context.Context, entity.Person) int
	Login(context.Context, entity.DTO_login) int
	Logout(context.Context, entity.DTO_login) int
}

func NewUsecase(datalayer Database) *RealUsecase {
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
	return uc.datalayer.NewPerson(ctx, person)
}

func (uc *RealUsecase) Login(ctx context.Context, info entity.DTO_login) int {
	if !uc.datalayer.IsActive(ctx, info) {
		return 500
	}
	check := uc.datalayer.GetPerson(ctx, info)
	if info.Password == check.Password {
		return uc.datalayer.Login(ctx, info)
	}
	return 400
}

func (uc *RealUsecase) Logout(ctx context.Context, info entity.DTO_login) int {
	if !uc.datalayer.IsActive(ctx, info) {
		return uc.datalayer.Logout(ctx, info)
	}
	return 400
}
