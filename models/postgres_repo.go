package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"example.com/rest-api/logger"
	"example.com/rest-api/utils"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const (
	schema = "mydb"
)

// Repository contains db dependencies and implements methods to interact with db
type postgresRepository struct {
	DB *sql.DB
}

// GetDB returns the db
func (r *postgresRepository) GetDB() *sql.DB {
	return r.DB
}

func (r *postgresRepository) GetPatientByUhid(uhid string) (*Patient, error) {
	ctx := context.Background()
	log := logger.Get(ctx).With(
		zap.String("uhid", uhid),
		zap.String("method", "GetPatientByUhid"))

	p := Patient{}
	var inputArgs []interface{}
	sqlQuery := `SELECT uhid, barcode, name, labour_id, age,
	                 gender, mobile, district, taluk, camp FROM patients WHERE uhid = ?`

	sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)
	inputArgs = append(inputArgs, uhid)
	log.Info(fmt.Sprintf("query : %v", sqlQuery))

	row := DB.QueryRow(sqlQuery, inputArgs...)

	err := row.Scan(&p.Uhid, &p.Barcode, &p.Name, &p.LabourId, &p.Age,
		&p.Gender, &p.Mobile, &p.District, &p.Taluk, &p.Camp)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &p, nil
}

func (r *postgresRepository) SavePatient(p *Patient) error {
	ctx := context.Background()
	log := logger.Get(ctx).With(
		zap.String("barcode", p.Barcode),
		zap.String("method", "SavePatient"))
	var inputArgs []interface{}
	var userid int
	sqlQuery := `INSERT INTO patients(uhid, barcode, name, labour_id, age,
	                     gender, mobile, district, taluk, camp) 
					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id `

	inputArgs = append(inputArgs, p.Uhid, p.Barcode, p.Name, p.LabourId, p.Age,
		p.Gender, p.Mobile, p.District, p.Taluk, p.Camp)
	sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)

	log.Info(fmt.Sprintf("query : %v", sqlQuery))
	err := DB.QueryRow(sqlQuery, inputArgs...).Scan(&userid)

	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("beneficiary added with id : %v", userid))
	return nil
}

func (r *postgresRepository) SaveUser(u *User) error {
	ctx := context.Background()
	log := logger.Get(ctx).With(
		zap.String("username", u.Email),
		zap.String("role", u.Role),
		zap.String("method", "SaveUser"))

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	var inputArgs []interface{}
	sqlQuery := `INSERT INTO users(email, password,role) VALUES (?, ?, ?)`

	inputArgs = append(inputArgs, u.Email, hashedPassword, u.Role)
	sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)
	log.Info(fmt.Sprintf("query : %v", sqlQuery))
	result, err := DB.ExecContext(ctx, sqlQuery, inputArgs...)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		logger.Get(ctx).Error(err.Error())
		return err
	}

	log.Info("User created successfully. ")
	return err
}

func (r *postgresRepository) ValidateCredentials(u *User) (string, error) {
	ctx := context.Background()
	var inputArgs []interface{}
	log := logger.Get(ctx).With(zap.String("username", u.Email),
		zap.String("role", u.Role),
		zap.String("method", "ValidateCredentials"))

	sqlQuery := "SELECT id, password,role FROM users WHERE email = ?"

	sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)
	inputArgs = append(inputArgs, u.Email)
	log.Info(fmt.Sprintf("query : %v", sqlQuery))

	row := DB.QueryRow(sqlQuery, inputArgs...)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword, &u.Role)

	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	log.Info("reading user from db successful")
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return "", errors.New("credentials invalid")
	}

	return u.Role, nil
}
