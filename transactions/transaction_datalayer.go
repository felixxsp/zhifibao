package transaction

import (
	"context"
	"fmt"
	"zhifubao/domain/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RealMongo struct {
	database   *mongo.Database
	collection *mongo.Collection
	person     *mongo.Collection
}

func NewRealMongo(db *mongo.Database) *RealMongo {
	return &RealMongo{
		database:   db,
		collection: db.Collection("transactions"),
		person:     db.Collection("person"),
	}
}

func (db *RealMongo) NewTransaction(ctx context.Context, item entity.Transaction) int {
	result, _ := db.collection.InsertOne(ctx, item)
	fmt.Println(result)
	return 200
}

func (db *RealMongo) ViewTransaction(ctx context.Context, item entity.Trc_req_one) entity.Transaction {
	var result entity.Transaction
	db.collection.FindOne(ctx, bson.M{"uuid": item.Transaction}).Decode(&result)
	return result
}

func (db *RealMongo) ViewMulti(ctx context.Context, item entity.Trc_req_multi) []entity.Transaction {
	var results []entity.Transaction
	cursor, _ := db.collection.Find(ctx, bson.M{"time": bson.M{"$gte": item.FilterStart, "$lte": item.FilterEnd}})
	for cursor.Next(ctx) {
		var result entity.Transaction
		cursor.Decode(&result)
		results = append(results, result)
	}
	return results
}
