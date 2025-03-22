package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                   int64
	UserName             string    `json:"user_name"`
	Password             string    `binding:"required" json:"password"`
	Mobile               string    `json:"mobile"`
	EmailId              string    `binding:"required" json:"email_id"`
	Gender               string    `json:"gender"`
	FirstName            string    `json:"first_name"`
	MiddleName           string    `json:" middle_name"`
	LastName             string    `json:"last_name"`
	CTC                  string    `json:"ctc"`
	DateOfBirth          time.Time `json:"date_of_birth"`
	DateOfJoining        time.Time `json:"date_of_joining"`
	LastWorkingDay       time.Time `json:"last_working_day"`
	CampaignId           uuid.UUID `json:"campaign_id"`
	AssigningAuthorityId uuid.UUID `json:"assigning_authority_id"`
	RoleId               uuid.UUID `json:"role_id"`
	DepartmentId         uuid.UUID `json:"department_id"`
	ApplicationAccessId  uuid.UUID `json:"application_access_id"`
}
