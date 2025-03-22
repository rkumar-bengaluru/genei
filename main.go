package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/logger"
	"example.com/rest-api/repository"
	"example.com/rest-api/routes"
	"example.com/rest-api/service"
	"example.com/rest-api/zcontext"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func main() {

	ctx := zcontext.BackgroundContext()
	log := logger.Get(ctx).With(zap.String("methodName", "main"))
	conn := db.CreateDB(ctx, "genei-server")
	defer conn.Close()
	// initialize repository
	repo := repository.NewRepository(conn)
	// initialize services
	patientService := service.NewPatientService(repo)
	userService := service.NewUserService(repo)
	campaignService := service.NewCampaignService(repo)
	//seed.SeedCampaignData(campaignService)
	server := gin.Default()

	routes.RegisterRoutes(server, patientService, userService, campaignService)
	log.Info("Server Started Successful...")

	server.Run(":8080") // localhost:8080

}
