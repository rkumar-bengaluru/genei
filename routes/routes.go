package routes

import (
	"example.com/rest-api/middlewares"
	"example.com/rest-api/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine,
	patientService *service.PatientService,
	userService *service.UserService) {
	server.GET("/events", getEvents)    // GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent) // /events/1, /events/5

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	authenticated.GET("/api/patient/:id", func(c *gin.Context) {
		getPatient(c, patientService)
	})
	authenticated.POST("/api/patients", func(c *gin.Context) {
		createPatient(c, patientService)
	})

	server.POST("/api/signup", func(c *gin.Context) {
		signup(c, userService)
	})
	server.POST("/api/login", func(c *gin.Context) {
		login(c, userService)
	})
}
