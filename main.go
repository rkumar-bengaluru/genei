package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/logger"
	"example.com/rest-api/models"
	"example.com/rest-api/routes"
	"example.com/rest-api/service"
	"example.com/rest-api/zcontext"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	ctx := zcontext.BackgroundContext()
	log := logger.Get(ctx).With(zap.String("methodName", "main"))
	conn := db.CreateDB(ctx, "genei-server")
	defer conn.Close()
	// initialize repository
	repo := models.NewRepository(conn)
	// initialize services
	patientService := service.NewPatientService(repo)
	userService := service.NewUserService(repo)
	server := gin.Default()

	routes.RegisterRoutes(server, patientService, userService)
	log.Info("Server Started Successful...")

	server.Run(":8080") // localhost:8080

}
