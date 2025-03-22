package models

import (
	"time"

	"github.com/google/uuid"
)

type Campaign struct {
	ID                       uuid.UUID `json:"uuid"`
	EstimatedTargetScreening int       `binding:"required" json:"estimated_target_screening"`
	LabourInspectorName      string    `binding:"required" json:"labour_inspector_name"`
	UnionName                string    `binding:"required" json:"union_name"`
	UnionLeaderName          string    `binding:"required" json:"union_leader_name"`
	Latitude                 string    `binding:"required" json:"latitude"`
	Longitude                string    `binding:"required" json:"longitude"`
	Taluk                    string    `binding:"required" json:"taluk"`
	CampName                 string    `binding:"required" json:"camp_name"`
	Description              string    `binding:"required" json:"description"`
	ScreeningStartDate       time.Time `binding:"required" json:"screening_start_date"`
	ScreeningStartTime       string    `binding:"required" json:"screening_start_time"`
	ApplicationAccessId      uuid.UUID `binding:"required" json:"application_access_id"`
	DistrictId               uuid.UUID `binding:"required" json:"district_id"`
	StateId                  uuid.UUID `binding:"required" json:"state_id"`
	AssigningAuthorityId     uuid.UUID `binding:"required" json:"assigning_authority_id"`
	StoreId                  uuid.UUID `binding:"required" json:"store_id"`
	WorkOrderId              uuid.UUID `binding:"required" json:"work_order_id"`
	PincodeId                uuid.UUID `binding:"required" json:"pin_code_id"`
}
