package person

import (
	"context"
	"zhifubao/domain"
	"zhifubao/domain/entity"
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
func (db *FakeMongo) GetPerson(context.Context, entity.Login_req) entity.Person {
	var temp entity.Person
	return temp
}

func (db *FakeMongo) Login(context.Context, entity.Login_req) int {
	return 200
}

func (db *FakeMongo) Logout(context.Context, entity.Login_req) int {
	return 200
}

func (db *FakeMongo) IsActive(context.Context, entity.Login_req) bool {
	return true
}

func (db *FakeMongo) IsExist(context.Context, entity.Person) bool {
	return true
}

func (db *FakeMongo) ErrTest(ctx context.Context) error {
	return domain.ErrNotFound
}

func (db *FakeMongo) Autentication(ctx context.Context, test entity.Trc_req_one) (bool, error) {
	return true, nil
}

func (db *FakeMongo) UpdateBalance(ctx context.Context, source entity.Transaction) bool {
	return true
}

// func TestAutentication(t *testing.T) {
// 	ctx := context.Background()
// 	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
// 	database := client.Database("zhifubao")
// 	db := database.Collection("person")
// 	var comparison entity.Person
// 	db.FindOne(ctx, bson.M{"uuid": subject.PersonID}).Decode(comparison)
// 	if comparison.Name == "" {
// 		return false, domain.ErrNotFound
// 	} else if !comparison.Active {
// 		return false, domain.ErrLogout
// 	} else if comparison.ActiveDevice != subject.Device {
// 		return false, domain.ErrDeviceIncompatible
// 	}
// 	return true, nil
// }
