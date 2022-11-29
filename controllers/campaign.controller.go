package controllers

import (
	"go-ad-campaign/models"
	"go-ad-campaign/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CampaignController struct {
	CampaignService services.CampaignService
}

// constructor 
func New(campaignService services.CampaignService) CampaignController {
	return CampaignController{
		CampaignService: campaignService,
	}
}

// Create a Campaign
func (cc *CampaignController) CreateCampaign(ctx *gin.Context) { // gin context hold all the req informations we send
	var campaign models.Campaign
	if err := ctx.ShouldBindJSON(&campaign); err != nil { // error handling during create campaign variable
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 
	}
	err := cc.CampaignService.CreateCampaign(&campaign)
	if err != nil { // error handling during interaction with MongoDB
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "campaign created"}) // response
}

// Get all Campaigns
func (cc *CampaignController) GetAll(ctx *gin.Context)  {
	campaigns, err := cc.CampaignService.GetAll()

	if err != nil { // error handling during interaction with MongoDB
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK, campaigns)
}

// Get a Campaign by name
func (cc *CampaignController) GetByName(ctx *gin.Context)  {
	campaignName := ctx.Param("name")
	campaign, err := cc.CampaignService.GetByName(&campaignName)
	if err != nil { // error handling during interaction with MongoDB
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK, campaign) // response
}

// Update a Campaign
func (cc *CampaignController) UpdateCampaign(ctx *gin.Context) {
	campaignName := ctx.Param("name")
	var campaign models.Campaign
	if err := ctx.ShouldBindJSON(&campaign); err != nil { // error handling during create campaign variable
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return 
	}
	err := cc.CampaignService.UpdateCampaign(campaignName, &campaign)
	if err != nil { // error handling during interaction with MongoDB
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}

	ctx.JSON(http.StatusOK, campaign) // response
}

// delete a Campaign 
func (cc *CampaignController) DeleteCampaign(ctx *gin.Context)  {
	campaignName := ctx.Param("name")
	err := cc.CampaignService.DeleteCampaign(&campaignName)
	if err != nil { // error handling during interaction with MongoDB
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return 
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "campaign deleted"}) // response
} 

// Routing 
func (cc *CampaignController) RegisterCampaignRoutes(rg *gin.RouterGroup) {
	campaignRoute := rg.Group("/campaign")
	campaignRoute.POST("/", cc.CreateCampaign) // final route is http://localhost:PORT/campaign/create for example
	campaignRoute.GET("/", cc.GetAll)
	campaignRoute.GET("/:name", cc.GetByName)
	campaignRoute.PATCH("/:name", cc.UpdateCampaign)
	campaignRoute.DELETE("/:name", cc.DeleteCampaign)
}