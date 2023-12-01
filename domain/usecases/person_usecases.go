package usecases

import (
	"context"
	"zhifubao/domain/entity"
)

type Person_Usecase interface {
	NewPerson(context.Context, entity.Person) int
	Login(context.Context, entity.Login_req) int
	Logout(context.Context, entity.Login_req) int
	GetFriends(context.Context, entity.Login_req) []entity.Person
}
