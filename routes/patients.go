package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getPatient(context *gin.Context) {
	var patient models.Patient
	uhid := context.Param("id")
	if uhid == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse uhid."})
		return
	}

	p, err := patient.GetPatientByUhid(uhid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusOK, p)
}

func createPatient(context *gin.Context) {
	var patient models.Patient
	err := context.ShouldBindJSON(&patient)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse models.Patient data."})
		return
	}

	err = patient.Save()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create models.Patient. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "models.Patient created!", "patient": patient})
}
