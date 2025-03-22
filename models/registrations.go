package models

import (
	"time"

	"github.com/google/uuid"
)

type Registration struct {
	ID               uuid.UUID `json:"uuid"`
	RegistrationDate time.Time `binding:"required" json:"registration_date"`
	Uhid             string    `binding:"required" json:"uhid"`
	Barcode          string    `binding:"required" json:"barcode"`
	Name             string    `binding:"required" json:"name"`
	LabourId         string    `binding:"required" json:"labour_id"`
	Age              int       `binding:"required" json:"age"`
	Gender           string    `binding:"required" json:"gender"`
	Mobile           string    `binding:"required" json:"mobile"`
	Taluk            string    `binding:"required" json:"taluk"`
	LabTestStatus    int       `json:"lab_test_status"`
	ReportUrl        string    `json:"report_url"`
	CampaignId       uuid.UUID `binding:"required" json:"campaign_id"`
	DistrictId       uuid.UUID `binding:"required" json:"district_id"`
}
