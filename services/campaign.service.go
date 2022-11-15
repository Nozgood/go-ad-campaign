package services

import "go-ad-campaign/models"

type CampaignService interface {
	CreateCampaign(*models.Campaign) error
	GetByName(*string) (*models.Campaign, error)
	GetAll() ([]*models.Campaign, error)
	UpdateCampaign(*models.Campaign) error
	DeleteCampaign(*string) error 
}