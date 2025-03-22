package routes

import (
	"example.com/rest-api/middlewares"
	"example.com/rest-api/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine,
	patientService *service.PatientService,
	userService *service.UserService,
	campaignService *service.CampaignService) {

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.GET("/api/patient/:id", func(c *gin.Context) {
		getPatient(c, patientService)
	})
	authenticated.POST("/api/patients", func(c *gin.Context) {
		createPatient(c, patientService)
	})

	authenticated.POST("/api/admin/campaigns", func(c *gin.Context) {
		createCampign(c, campaignService)
	})

	authenticated.GET("/api/admin/campaigns", func(c *gin.Context) {
		listCampaigns(c, campaignService)
	})

	server.POST("/api/signup", func(c *gin.Context) {
		signup(c, userService)
	})
	server.POST("/api/login", func(c *gin.Context) {
		login(c, userService)
	})
}
