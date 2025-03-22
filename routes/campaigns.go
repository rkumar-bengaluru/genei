package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"example.com/rest-api/service"
	"example.com/rest-api/zcontext"
	"github.com/gin-gonic/gin"
)

func createCampign(context *gin.Context, svc *service.CampaignService) {
	traceCtx := zcontext.BackgroundContext()
	var camp models.ArogyaCampaign
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

func listCampaigns(ctx *gin.Context, svc *service.CampaignService) {
	traceCtx := zcontext.BackgroundContext()

	// Parse query parameters for pagination
	pageStr := ctx.Query("page")
	pageSizeStr := ctx.Query("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1 // Default to page 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10 // Default page size of 10
	}

	offset := (page - 1) * pageSize

	resp, err := svc.ListCampaign(&traceCtx, pageSize, offset)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not list campaigns."})
		return
	}

	ctx.JSON(http.StatusOK, &resp)

}
