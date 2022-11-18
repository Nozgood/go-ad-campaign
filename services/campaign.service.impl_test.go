package services

import (
	"context"
	"go-ad-campaign/models"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	campaignCollection *mongo.Collection
	ctx				   context.Context	
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
	campaignCollection = mongoClient.Database("campaigndb").Collection("campaigns")
}

/*
Create a mock Object and a mock Server to test the functions
*/
var test = models.Campaign{
	Name: "test",
	StartDate: models.Date {
		Day: 12,
		Month: 07,
		Year: 2022,
	},
	EndDate: models.Date {
		Day: 14,
		Month: 07,
		Year: 2022,
	},
	Price: 400,
	Objective: 1000,
	PricePerDisplay: 1.2,
}


func TestCreateCampaign(t *testing.T) {
	testService := NewCampaignService(campaignCollection, ctx)
	result := testService.CreateCampaign(&test)
	if result == nil {
		t.Logf("createCampaign PASSED. Expected %v, got %v \n", nil, result)
	} else {
		t.Errorf("createCampaign FAILED. Expected %v, got %v \n", nil, result)
	}
}

func TestGetAllCampaign(t *testing.T) {
	testService := NewCampaignService(campaignCollection, ctx)
	 _, err := testService.GetAll()
	if err != nil {
		t.Errorf("getAllCampaigns FAILED. Expected %v, got %v \n", 
		nil, err)
	} else {
		t.Logf("getAllCampaigns PASSED \n")
	}
}

func TestGetCampaignByName(t *testing.T) {
	testService := NewCampaignService(campaignCollection, ctx)
	name := "test"
	_, err := testService.GetByName(&name)
	if err != nil {
		t.Errorf("getCampaignByName FAILED. Expected %v, got %v \n", 
		nil, err)
	} else {
		t.Logf("getCampaignByName PASSED \n")
	}
}

func TestGetCampaignByNameDoesntExist(t *testing.T) {
	testService := NewCampaignService(campaignCollection, ctx)
	name := "nexistepas"
	_, err := testService.GetByName(&name)
	if err == nil {
		t.Errorf("getCampaignByName FAILED. Expected %v, got %v \n", 
		nil, err)
	} else {
		t.Logf("getCampaignByName PASSED \n")
	}
}

func TestUpdateCampaign(t *testing.T) {
	testService := NewCampaignService(campaignCollection, ctx)
	name := test.Name
	err := testService.UpdateCampaign(name, &test)

	if err != nil {
		t.Errorf("UpdateCampaign FAILED. Expected %v, got %v \n", 
		nil, err)
	} else {
		t.Logf("UpdateCampaign PASSED \n")
	}
}

func TestUpdateCampaignDoesntExist(t *testing.T) {
	testService := NewCampaignService(campaignCollection, ctx)
	name := "nexistepas"
	err := testService.UpdateCampaign(name, &test)

	if err == nil {
		t.Errorf("UpdateCampaign FAILED. Expected not nil value, got %v \n",  err)
	} else {
		t.Logf("UpdateCampaign PASSED \n")
	}
}

func TestDeleteCampaign(t *testing.T) {
	testService := NewCampaignService(campaignCollection, ctx)
	err := testService.DeleteCampaign(&test.Name)

	if err != nil {
		t.Errorf("DeleteCampaign FAILED. Expected %v, got %v \n", 
		nil, err)
	} else {
		t.Logf("UpdateCampaign PASSED \n")
	}
}

func TestDeleteCampaignDoesntExis(t *testing.T) {
	testService := NewCampaignService(campaignCollection, ctx)
	name := "nexistepas"
	err := testService.DeleteCampaign(&name)

	if err == nil {
		t.Errorf("UpdateCampaign FAILED. Expected not nil value, got %v \n",  err)
	} else {
		t.Logf("UpdateCampaign PASSED \n")
	}
}