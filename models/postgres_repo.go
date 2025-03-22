package models

// import (
// 	"context"
// 	"database/sql"
// 	"errors"
// 	"fmt"

// 	"example.com/rest-api/logger"
// 	"example.com/rest-api/utils"
// 	"github.com/jmoiron/sqlx"
// 	"go.uber.org/zap"
// )

// func (r *postgresRepository) GetPatientByUhid(uhid string) (*Registration, error) {
// 	ctx := context.Background()
// 	log := logger.Get(ctx).With(
// 		zap.String("uhid", uhid),
// 		zap.String("method", "GetPatientByUhid"))

// 	p := Registration{}
// 	var inputArgs []interface{}
// 	sqlQuery := `SELECT registration_date, uhid, barcode, name, labour_id, age,
// 	                 gender, mobile, district, taluk, camp FROM patients WHERE uhid = ?`

// 	sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)
// 	inputArgs = append(inputArgs, uhid)
// 	log.Info(fmt.Sprintf("query : %v", sqlQuery))

// 	row := DB.QueryRow(sqlQuery, inputArgs...)

// 	err := row.Scan(&p.RegistrationDate, &p.Uhid, &p.Barcode, &p.Name, &p.LabourId, &p.Age,
// 		&p.Gender, &p.Mobile, &p.DistrictId, &p.Taluk, &p.CampaignId)

// 	if err != nil {
// 		log.Error(err.Error())
// 		return nil, err
// 	}

// 	return &p, nil
// }

// func (r *postgresRepository) SaveRegistration(p *Registration) error {
// 	ctx := context.Background()
// 	log := logger.Get(ctx).With(
// 		zap.String("barcode", p.Barcode),
// 		zap.String("method", "SavePatient"))
// 	var inputArgs []interface{}
// 	var userid int
// 	sqlQuery := `INSERT INTO registrations(registration_date, uhid, barcode, name, labour_id, age,
// 	                     gender, mobile, district_id, taluk, campaign_id)
// 					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id `

// 	inputArgs = append(inputArgs, p.RegistrationDate, p.Uhid, p.Barcode, p.Name, p.LabourId, p.Age,
// 		p.Gender, p.Mobile, p.DistrictId, p.Taluk, p.CampaignId)
// 	sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)

// 	log.Info(fmt.Sprintf("query : %v", sqlQuery))
// 	err := DB.QueryRow(sqlQuery, inputArgs...).Scan(&userid)

// 	if err != nil {
// 		return err
// 	}
// 	log.Info(fmt.Sprintf("beneficiary added with id : %v", userid))
// 	return nil
// }

// func (r *postgresRepository) SaveUser(u *User) error {
// 	ctx := context.Background()
// 	log := logger.Get(ctx).With(
// 		zap.String("username", u.EmailId),
// 		zap.String("role", u.RoleId.String()),
// 		zap.String("method", "SaveUser"))

// 	hashedPassword, err := utils.HashPassword(u.Password)

// 	if err != nil {
// 		log.Error(err.Error())
// 		return err
// 	}

// 	var inputArgs []interface{}
// 	sqlQuery := `INSERT INTO users(user_name, password,mobile,email_id,assigning_authority_id,
// 				role_id,department_id,application_access_id,gender,first_name,
// 				last_name,middle_name,date_of_birth,date_of_joining,last_working_day,
// 				monthly_ctc,campaign_id) VALUES (?, ?, ?)`

// 	inputArgs = append(inputArgs, u.UserName, hashedPassword, u.Mobile, u.EmailId, u.AssigningAuthorityId,
// 		u.RoleId, u.DepartmentId, u.ApplicationAccessId, u.Gender, u.FirstName,
// 		u.LastName, u.MiddleName, u.DateOfBirth, u.DateOfJoining, u.LastWorkingDay,
// 		u.CTC, u.CampaignId)
// 	sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)
// 	log.Info(fmt.Sprintf("query : %v", sqlQuery))
// 	result, err := DB.ExecContext(ctx, sqlQuery, inputArgs...)

// 	if err != nil {
// 		log.Error(err.Error())
// 		return err
// 	}

// 	_, err = result.RowsAffected()

// 	if err != nil {
// 		logger.Get(ctx).Error(err.Error())
// 		return err
// 	}

// 	log.Info("User created successfully. ")
// 	return err
// }

// func (r *postgresRepository) ValidateCredentials(u *User) (string, error) {
// 	ctx := context.Background()
// 	var inputArgs []interface{}
// 	log := logger.Get(ctx).With(zap.String("username", u.EmailId),
// 		zap.String("role", u.RoleId.String()),
// 		zap.String("method", "ValidateCredentials"))

// 	sqlQuery := "SELECT id, password,role_id FROM users WHERE user_name = ?"

// 	sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)
// 	inputArgs = append(inputArgs, u.EmailId)
// 	log.Info(fmt.Sprintf("query : %v", sqlQuery))

// 	row := DB.QueryRow(sqlQuery, inputArgs...)

// 	var retrievedPassword string
// 	var roleId string
// 	err := row.Scan(&u.ID, &retrievedPassword, &roleId)

// 	if err != nil {
// 		log.Error(err.Error())
// 		return "", err
// 	}

// 	log.Info("reading user from db successful")
// 	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

// 	if !passwordIsValid {
// 		return "", errors.New("credentials invalid")
// 	}

// 	// fetch role
// 	sqlQueryFetchRole := "SELECT name FROM roles WHERE id = ?"
// 	var inputArgsRole []interface{}
// 	sqlQueryFetchRole = sqlx.Rebind(sqlx.DOLLAR, sqlQueryFetchRole)
// 	inputArgsRole = append(inputArgsRole, roleId)
// 	log.Info(fmt.Sprintf("query : %v", sqlQueryFetchRole))

// 	row = DB.QueryRow(sqlQueryFetchRole, inputArgsRole...)
// 	var roleName string
// 	err = row.Scan(&roleName)

// 	if err != nil {
// 		log.Error(err.Error())
// 		return "", err
// 	}

// 	return roleName, nil
// }
