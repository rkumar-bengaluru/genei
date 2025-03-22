package repository

import (
	"context"
	"errors"

	"example.com/rest-api/logger"
	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func (r *postgresRepository) SaveUser(ctx *context.Context, u *models.User) error {
	log := logger.Get(*ctx).With(
		zap.String("username", u.EmailId),
		zap.String("role", u.RoleId.String()),
		zap.String("method", "SaveUser"))

	var inputArgs []interface{}
	sqlQuery := `INSERT INTO users(user_name, password,mobile,email_id,role_id,gender,
				 first_name,last_name,middle_name,date_of_birth,date_of_joining,last_working_day,
				 monthly_ctc,department_id,application_access_id, assigning_authority_id) 
				 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	inputArgs = append(inputArgs, u.UserName, u.Password, u.Mobile, u.EmailId,
		u.RoleId, u.Gender, u.FirstName, u.LastName,
		u.MiddleName, u.DateOfBirth, u.DateOfJoining,
		u.LastWorkingDay, u.CTC, u.DepartmentId, u.ApplicationAccessId, u.AssigningAuthorityId)
	sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)
	log.Debug("Executing query", zap.String("query", sqlQuery))

	result, err := DB.ExecContext(*ctx, sqlQuery, inputArgs...)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		logger.Get(*ctx).Error(err.Error())
		return err
	}

	log.Info("User created successfully. ")
	return err
}

func (r *postgresRepository) GetPassword(ctx *context.Context, u *models.User) (string, error) {
	var inputArgs []interface{}
	log := logger.Get(*ctx).With(zap.String("username", u.EmailId),
		zap.String("role", u.RoleId.String()),
		zap.String("method", "GetPassword"))

	sqlQuery := "SELECT id, password,role_id FROM users WHERE user_name = ?"

	sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)
	inputArgs = append(inputArgs, u.EmailId)
	log.Debug("Executing query", zap.String("query", sqlQuery))

	row := DB.QueryRow(sqlQuery, inputArgs...)

	var retrievedPassword string
	var retrievedRoleId string
	err := row.Scan(&u.ID, &retrievedPassword, &retrievedRoleId)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}
	log.Info("retrived role id ", zap.String("roleId", retrievedRoleId))

	roleId, err := uuid.Parse(retrievedRoleId)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	u.RoleId = roleId

	return retrievedPassword, nil
}

func (r *postgresRepository) GetRole(ctx *context.Context, roleId uuid.UUID) (string, error) {
	log := logger.Get(*ctx).With(
		zap.String("method", "GetRole"))
	// fetch role
	sqlQueryFetchRole := "SELECT name FROM roles WHERE id = ?"
	var inputArgsRole []interface{}
	sqlQueryFetchRole = sqlx.Rebind(sqlx.DOLLAR, sqlQueryFetchRole)
	inputArgsRole = append(inputArgsRole, roleId.String())
	log.Debug("Executing query", zap.String("query", sqlQueryFetchRole))

	row := DB.QueryRow(sqlQueryFetchRole, inputArgsRole...)
	var roleName string
	err := row.Scan(&roleName)

	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	return roleName, nil
}

func (r *postgresRepository) ValidateCredentials(ctx *context.Context, u *models.User) (string, error) {
	var inputArgs []interface{}
	log := logger.Get(*ctx).With(zap.String("username", u.EmailId),
		zap.String("role", u.RoleId.String()),
		zap.String("method", "ValidateCredentials"))

	sqlQuery := "SELECT id, password,role_id FROM users WHERE user_name = ?"

	sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)
	inputArgs = append(inputArgs, u.EmailId)
	log.Debug("Executing query", zap.String("query", sqlQuery))

	row := DB.QueryRow(sqlQuery, inputArgs...)

	var retrievedPassword string
	var roleId string
	err := row.Scan(&u.ID, &retrievedPassword, &roleId)

	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	log.Info("reading user from db successful")
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return "", errors.New("credentials invalid")
	}

	// fetch role
	sqlQueryFetchRole := "SELECT name FROM roles WHERE id = ?"
	var inputArgsRole []interface{}
	sqlQueryFetchRole = sqlx.Rebind(sqlx.DOLLAR, sqlQueryFetchRole)
	inputArgsRole = append(inputArgsRole, roleId)
	log.Debug("Executing query", zap.String("query", sqlQuery))

	row = DB.QueryRow(sqlQueryFetchRole, inputArgsRole...)
	var roleName string
	err = row.Scan(&roleName)

	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	return roleName, nil
}
