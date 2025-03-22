package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/service"
	"example.com/rest-api/zcontext"
	"github.com/gin-gonic/gin"
)

func getPatient(context *gin.Context, svc *service.PatientService) {
	traceCtx := zcontext.BackgroundContext()
	uhid := context.Param("id")
	if uhid == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse uhid."})
		return
	}

	patient, err := svc.GetPatientByUhid(&traceCtx, uhid)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusOK, patient)
}

func createPatient(context *gin.Context, svc *service.PatientService) {
	traceCtx := zcontext.BackgroundContext()
	var patient models.Registration
	err := context.ShouldBindJSON(&patient)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse models.Patient data."})
		return
	}

	err = svc.Save(&traceCtx, &patient)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create models.Patient. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "models.Patient created!", "patient": patient})
}
