package person

import (
	"context"
	"只服宝/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RealMongo struct {
	database   *mongo.Database
	collection *mongo.Collection
}

type Database interface {
	NewPerson(context.Context, entity.Person) int
	GetPerson(context.Context, entity.DTO_login) entity.Person
	Login(context.Context, entity.DTO_login) int
	Logout(context.Context, entity.DTO_login) int
	IsActive(context.Context, entity.DTO_login) bool
	IsExist(context.Context, entity.Person) bool
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
		return 500
	} else {
		return 201
	}
}

func (db *RealMongo) GetPerson(ctx context.Context, person entity.DTO_login) entity.Person {
	var update entity.Person
	db.collection.FindOne(ctx, bson.M{"uuid": person.PersonID}).Decode(&update)
	return update
}

func (db *RealMongo) Login(ctx context.Context, person entity.DTO_login) int {
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

func (db *RealMongo) Logout(ctx context.Context, person entity.DTO_login) int {
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

func (db *RealMongo) IsActive(ctx context.Context, person entity.DTO_login) bool {
	var check entity.Person
	db.collection.FindOne(ctx, bson.M{"uuid": person.PersonID}).Decode(&check)
	return check.Active
}

func (db *RealMongo) IsExist(ctx context.Context, person entity.Person) bool {
	check, _ := db.collection.CountDocuments(ctx, bson.M{"username": person.Username})
	return check != 0
}
