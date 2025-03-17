package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/logger"
	"example.com/rest-api/models"
	"example.com/rest-api/service"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func signup(context *gin.Context, us *service.UserService) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = us.Save(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func login(context *gin.Context, us *service.UserService) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	log := logger.Get(context).With(zap.String("username", user.Email),
		zap.String("method", "login"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	log.Info("Validating User...")
	role, err := us.ValidateCredentials(&user)

	if err != nil {
		log.Error(err.Error())
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}
	log.Info("Generating Token...")
	token, err := utils.GenerateToken(user.Email, user.ID, role)

	if err != nil {
		log.Error(err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	fmt.Println(context.Get("userId"))
	fmt.Println(role)

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
