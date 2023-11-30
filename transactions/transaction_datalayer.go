package transaction

import (
	"context"
	"只服宝/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RealMongo struct {
	database   *mongo.Database
	collection *mongo.Collection
	person     *mongo.Collection
}

type Database interface {
	NewTransaction(context.Context, entity.Transaction) int
	ViewTransaction(context.Context, entity.DTO_trc) entity.Transaction
	ViewAll(context.Context, entity.DTO_trc) []entity.Transaction
}

func NewRealMongo(db *mongo.Database) *RealMongo {
	return &RealMongo{
		database:   db,
		collection: db.Collection("transactions"),
		person:     db.Collection("person"),
	}
}

func (db *RealMongo) NewTransaction(ctx context.Context, item entity.Transaction) int {
	db.collection.InsertOne(ctx, item)
	var person entity.Person

	db.person.FindOne(ctx, bson.M{"uuid": item.Person}).Decode(&person)
	person.Transactions = append(person.Transactions, item)
	db.collection.UpdateOne(ctx, bson.M{"uuid": person.UUID}, bson.M{"$set": person})
	return 200
}

func (db *RealMongo) ViewTransaction(ctx context.Context, item entity.DTO_trc) entity.Transaction {
	var result entity.Transaction
	db.collection.FindOne(ctx, bson.M{"uuid": item.Transaction}).Decode(&result)
	return result
}

func (db *RealMongo) ViewAll(ctx context.Context, item entity.DTO_trc) []entity.Transaction {
	var result []entity.Transaction
	return result
}
