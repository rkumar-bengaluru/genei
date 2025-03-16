package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/logger"
	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	log := logger.Get(context).With(zap.String("username", user.Email))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	log.Info("Validating User...")
	role, err := user.ValidateCredentials(context)

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
