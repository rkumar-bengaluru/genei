package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/logger"
	"example.com/rest-api/routes"
	"example.com/rest-api/zcontext"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	ctx := zcontext.BackgroundContext()
	log := logger.Get(ctx).With(zap.String("methodName", "main"))

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)
	log.Info("Server Started Successful...")
	server.Run(":8080") // localhost:8080

}
