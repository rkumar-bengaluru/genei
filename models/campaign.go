package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID       uuid.UUID      `json:"uuid"`
	Name     sql.NullString `json:"name"`
	UserId   sql.NullString `json:"userId"`
	UserType sql.NullString `json:"userType"`
}

type Program struct {
	Status       sql.NullString `json:"status"`
	DistrictName sql.NullString `json:"districtName"`
	Number       sql.NullString `json:"programNumber"`
	ShortCode    sql.NullString `json:"programShortCode"`
}

type ArogyaCampaign struct {
	ID           uuid.UUID        `json:"uuid"`
	Name         sql.NullString   `binding:"required" json:"name"`
	DistrictName sql.NullString   `json:"districtName"`
	VillageName  sql.NullString   `json:"villageName"`
	Taluk        sql.NullString   `binding:"required" json:"talukaName"`
	PincodeId    sql.NullString   `json:"villagePinCode"`
	CampId       sql.NullString   `json:"externalCampId"`
	WorkOrderId  sql.NullString   `json:"work_order_id"`
	Visibility   []sql.NullString `json:"visibility"`
	Status       sql.NullString   `json:"status"`
	CreatedBy    AuditLog         `json:"createdBy"`
	CreatedAt    time.Time        `json:"createdAt"`
	UpdatedAt    time.Time        `json:"updatedAt"`
	StateName    sql.NullString   `json:"stateName"`
	Description  sql.NullString   `json:"description"`

	EstimatedTargetScreening sql.NullInt64   `binding:"required" json:"estimatedNumberOfScreenings"`
	LabourInspectorName      sql.NullString  `json:"labour_inspector_name"`
	UnionName                sql.NullString  `json:"union_name"`
	UnionLeaderName          sql.NullString  `json:"union_leader_name"`
	Latitude                 sql.NullFloat64 `binding:"required" json:"latitude"`
	Longitude                sql.NullFloat64 `binding:"required" json:"longitude"`

	ScreeningStartDate time.Time      `binding:"required" json:"screeningStartDate"`
	ScreeningStartTime sql.NullString `binding:"required" json:"screeningStartTime"`

	ApplicationAccessId  sql.NullString `json:"application_access_id"`
	AssigningAuthorityId sql.NullString `json:"assigning_authority_id"`
	StoreId              sql.NullString `json:"store_id"`

	Program Program `json:"programId"`
}

// String implements the fmt.Stringer interface.
func (c ArogyaCampaign) String() string {
	msg := `
	 Name-%s,District-%s,Village-%s, Taluk-%s, PincodeId-%s, CampId-%s,WorkOrderId-%s, Visibility-%s, Status-%s, CreatedBy-%s, StateName-%s, CreatedAt-%s
	`
	return fmt.Sprintf(msg, c.Name, c.Program.DistrictName, c.VillageName, c.Taluk, c.PincodeId, c.CampId,
		c.Program.Number, c.Visibility, c.Program.Status, c.CreatedBy, c.StateName, c.CreatedAt)
}

type Campaign struct {
	ID                       uuid.UUID      `json:"uuid"`
	EstimatedTargetScreening int            `binding:"required" json:"estimated_target_screening"`
	LabourInspectorName      sql.NullString `binding:"required" json:"labour_inspector_name"`
	UnionName                sql.NullString `binding:"required" json:"union_name"`
	UnionLeaderName          sql.NullString `binding:"required" json:"union_leader_name"`
	Latitude                 sql.NullString `binding:"required" json:"latitude"`
	Longitude                sql.NullString `binding:"required" json:"longitude"`
	Taluk                    sql.NullString `binding:"required" json:"taluk"`
	CampName                 sql.NullString `binding:"required" json:"camp_name"`
	Description              sql.NullString `binding:"required" json:"description"`
	ScreeningStartDate       time.Time      `binding:"required" json:"screening_start_date"`
	ScreeningStartTime       string         `binding:"required" json:"screening_start_time"`
	ApplicationAccessId      uuid.UUID      `binding:"required" json:"application_access_id"`
	DistrictId               uuid.UUID      `binding:"required" json:"district_id"`
	StateId                  uuid.UUID      `binding:"required" json:"state_id"`
	AssigningAuthorityId     uuid.UUID      `binding:"required" json:"assigning_authority_id"`
	StoreId                  uuid.UUID      `binding:"required" json:"store_id"`
	WorkOrderId              uuid.UUID      `binding:"required" json:"work_order_id"`
	PincodeId                uuid.UUID      `binding:"required" json:"pin_code_id"`
}
