package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/service"
	"github.com/gin-gonic/gin"
)

func getPatient(context *gin.Context, svc *service.PatientService) {
	uhid := context.Param("id")
	if uhid == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse uhid."})
		return
	}

	patient, err := svc.GetPatientByUhid(uhid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusOK, patient)
}

func createPatient(context *gin.Context, svc *service.PatientService) {
	var patient models.Patient
	err := context.ShouldBindJSON(&patient)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse models.Patient data."})
		return
	}

	err = svc.Save(&patient)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create models.Patient. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "models.Patient created!", "patient": patient})
}
