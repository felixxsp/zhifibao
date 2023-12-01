package main

import (
	"context"
	"zhifubao/domain/repository"
	"zhifubao/domain/usecases"
	"zhifubao/person"
	"zhifubao/transaction"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mainCtx := context.Background()

	router := gin.Default()

	client, _ := mongo.Connect(mainCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	database := client.Database("zhifubao")

	var personData repository.Person_Database = person.NewRealMongo(database)
	//var fakeData repository.Person_Database = person.InitFake()
	var personUC usecases.Person_Usecase = person.NewUsecase(personData)
	var personHandler person.Handler = person.NewHandler(personUC, router)

	var trcData repository.Transaction_Database = transaction.NewRealMongo(database)
	var trcUC usecases.Transaction_Usecase = transaction.NewUsecase(personData, trcData)
	var trcHanlder transaction.Handler = transaction.NewHandler(trcUC, router)

	personHandler.Standby(mainCtx)
	trcHanlder.Standby(mainCtx)

	router.Run(":7480")
}
