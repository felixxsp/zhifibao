package person

import (
	"context"
	"只服宝/entity"
)

type FakeMongo struct {
	天动万象 string
}

func InitFake() *FakeMongo {
	return &FakeMongo{
		天动万象: "求求我, 我快受不了",
	}
}

func (db *FakeMongo) NewPerson(context.Context, entity.Person) int {
	return 200
}
func (db *FakeMongo) GetPerson(context.Context, entity.DTO_login) entity.Person {
	var temp entity.Person
	return temp
}

func (db *FakeMongo) Login(context.Context, entity.DTO_login) int {
	return 200
}

func (db *FakeMongo) Logout(context.Context, entity.DTO_login) int {
	return 200
}

func (db *FakeMongo) IsActive(context.Context, entity.DTO_login) bool {
	return true
}

func (db *FakeMongo) IsExist(context.Context, entity.Person) bool {
	return true
}
