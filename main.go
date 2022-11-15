package main

import (
	"context"
	"fmt"
	"go-ad-campaign/controllers"
	"go-ad-campaign/services"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server 				*gin.Engine
	campaignService 	services.CampaignService
	CampaignController 	controllers.CampaignController
	ctx					context.Context
	campaignCollection 	*mongo.Collection
	mongoClient 		*mongo.Client
	err 				error
)

func init() {
	ctx = context.TODO()
	mongoConn := options.Client().ApplyURI("mongodb+srv://noz:okgoogle@projet7.tkmrk.mongodb.net/?retryWrites=true&w=majority")
	mongoClient, err := mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo connection done")

	campaignCollection = mongoClient.Database("campaigndb").Collection("campaigns")
	campaignService = services.NewCampaignService(campaignCollection, ctx)
	CampaignController = controllers.New(campaignService)
	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(ctx)

	basepath := server.Group("/v1")
	CampaignController.RegisterCampaignRoutes(basepath)

	log.Fatal(server.Run(":8080"))
}