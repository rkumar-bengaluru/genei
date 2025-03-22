package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/service"
	"example.com/rest-api/zcontext"
	"github.com/gin-gonic/gin"
)

func createCampign(context *gin.Context, svc *service.CampaignService) {
	traceCtx := zcontext.BackgroundContext()
	var camp models.Campaign
	err := context.ShouldBindJSON(&camp)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = svc.CreateCampaign(&traceCtx, &camp)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save camp."})
		return
	}

	context.JSON(http.StatusOK, camp)
}
