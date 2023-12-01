package repository

import (
	"context"
	"zhifubao/domain/entity"
)

type Person_Database interface {
	NewPerson(context.Context, entity.Person) int
	GetPerson(context.Context, entity.Login_req) entity.Person
	Login(context.Context, entity.Login_req) int
	Logout(context.Context, entity.Login_req) int
	IsActive(context.Context, entity.Login_req) bool
	IsExist(context.Context, entity.Person) bool
	ErrTest(context.Context) error
	Autentication(context.Context, entity.Trc_req_one) (bool, error)
	UpdateBalance(context.Context, entity.Transaction) bool
	GetFriends(context.Context, entity.Login_req) []entity.Person
}
