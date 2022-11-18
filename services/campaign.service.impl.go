package services

import (
	"context"
	"errors"
	"go-ad-campaign/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CampaignServiceImpl struct {
	campaignCollection 	*mongo.Collection
	ctx 				context.Context
}

//initialisation
func NewCampaignService(campaignCollection *mongo.Collection, ctx context.Context) CampaignService {
	return &CampaignServiceImpl {
		campaignCollection: campaignCollection,
		ctx: ctx,
	}
}

// Create a Campaign
func (c *CampaignServiceImpl) CreateCampaign(campaign *models.Campaign) error {
	_, err := c.campaignCollection.InsertOne(c.ctx, campaign)
	return err 
}

// Get all Campaigns
func (c *CampaignServiceImpl) GetAll() ([]*models.Campaign, error) {
	var campaigns []*models.Campaign
	cursor, err := c.campaignCollection.Find(c.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(c.ctx) {
		var campaign models.Campaign 
		err := cursor.Decode(&campaign)
		if err != nil {
			return nil, err
		}
		campaigns = append(campaigns, &campaign)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(c.ctx)

	if len(campaigns) == 0 {
		return nil, errors.New("no campaign to get")
	}

	return campaigns, nil
}

// Get a Campaign by name
func (c *CampaignServiceImpl) GetByName(name *string) (*models.Campaign, error) {
	var campaign *models.Campaign
	query := bson.D{bson.E{Key: "campaign_name", Value: name}} // db.collection.find({name: ""})
	err := c.campaignCollection.FindOne(c.ctx, query).Decode(&campaign)
	return campaign, err
}

// Update a Campaign
func (c *CampaignServiceImpl) UpdateCampaign(campaignName string, campaign *models.Campaign) error {
	filter := bson.D{bson.E{Key: "campaign_name", Value: campaignName}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "campaign_name", Value: campaign.Name}, bson.E{Key: "campaign_start", Value: campaign.StartDate}, bson.E{Key: "campaign_end", Value: campaign.EndDate}, bson.E{Key: "campaign_price", Value: campaign.Price}, bson.E{Key: "campaign_objective", Value: campaign.Objective}, bson.E{Key: "campaign_pricePerDisplay", Value: campaign.PricePerDisplay}}}}
	result, _ := c.campaignCollection.UpdateOne(c.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched campaign found for update")
	}
	return nil
}

// delete a Campaign 
func (c *CampaignServiceImpl) DeleteCampaign(name *string) error {
	filter := bson.D{bson.E{Key: "campaign_name", Value: name}}
	result, _ := c.campaignCollection.DeleteOne(c.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched campaign found for update")
	}
	return nil
} 
