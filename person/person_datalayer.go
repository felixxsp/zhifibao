package person

import (
	"context"
	"zhifubao/domain"
	"zhifubao/domain/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RealMongo struct {
	database   *mongo.Database
	collection *mongo.Collection
}

func NewRealMongo(db *mongo.Database) *RealMongo {
	return &RealMongo{
		database:   db,
		collection: db.Collection("person"),
	}
}

func (db *RealMongo) NewPerson(ctx context.Context, new entity.Person) int {
	_, err := db.collection.InsertOne(ctx, new)
	if err != nil {
		print(err)
		return 500
	} else {
		return 201
	}
}

func (db *RealMongo) GetPerson(ctx context.Context, person entity.Login_req) entity.Person {
	var update entity.Person
	db.collection.FindOne(ctx, bson.M{"uuid": person.PersonID}).Decode(&update)
	return update
}

func (db *RealMongo) Login(ctx context.Context, person entity.Login_req) int {
	var update entity.Person
	db.collection.FindOne(ctx, bson.M{"uuid": person.PersonID}).Decode(&update)
	update.ActiveDevice = person.Device
	update.Active = true
	_, err := db.collection.UpdateOne(ctx, bson.M{"uuid": person.PersonID}, bson.M{"$set": update})
	if err != nil {
		return 500
	} else {
		return 200
	}
}

func (db *RealMongo) Logout(ctx context.Context, person entity.Login_req) int {
	var update entity.Person
	db.collection.FindOne(ctx, bson.M{"uuid": person.PersonID}).Decode(&update)
	update.Active = false
	_, err := db.collection.UpdateOne(ctx, bson.M{"uuid": person.PersonID}, bson.M{"$set": update})
	if err != nil {
		return 500
	} else {
		return 200
	}
}

func (db *RealMongo) IsActive(ctx context.Context, person entity.Login_req) bool {
	var check entity.Person
	db.collection.FindOne(ctx, bson.M{"uuid": person.PersonID}).Decode(&check)
	return check.Active
}

// jwt

func (db *RealMongo) IsExist(ctx context.Context, person entity.Person) bool {
	check, _ := db.collection.CountDocuments(ctx, bson.M{"username": person.Username})
	return check == 0
}

func (db *RealMongo) ErrTest(ctx context.Context) error {
	return nil
}

func (db *RealMongo) Autentication(ctx context.Context, subject entity.Trc_req_one) (bool, error) {
	var comparison entity.Person
	db.collection.FindOne(ctx, bson.M{"uuid": subject.PersonID}).Decode(comparison)
	if comparison.Name == "" {
		return false, domain.ErrNotFound
	} else if !comparison.Active {
		return false, domain.ErrLogout
	} else if comparison.ActiveDevice != subject.Device {
		return false, domain.ErrDeviceIncompatible
	}
	return true, nil
}

func (db *RealMongo) UpdateBalance(ctx context.Context, source entity.Transaction) bool {
	var update entity.Person
	db.collection.FindOne(ctx, bson.M{"uuid": source.Person}).Decode(&update)
	if !source.Type {
		update.Balance -= source.Amount
	} else {
		update.Balance += source.Amount
	}
	newFriend := source.Receiver
	if len(update.Friends) > 0 {
		for _, now := range update.Friends {
			if now.UUID != newFriend {
				var temp entity.Login_req
				temp.PersonID = newFriend
				update.Friends = append(update.Friends, db.GetPerson(ctx, temp))
			}
		}
	}

	_, err := db.collection.UpdateOne(ctx, bson.M{"uuid": update.UUID}, bson.M{"$set": update})
	if err != nil {
		return false
	} else {
		return true
	}
}

// embedding - who friends are
// add transfer uuid
