package repository

import (
	"context"
	"fmt"

	"example.com/rest-api/logger"
	"example.com/rest-api/models"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func (r *postgresRepository) GetPatientByUhid(ctx *context.Context, uhid string) (*models.Registration, error) {
	log := logger.Get(*ctx).With(
		zap.String("uhid", uhid),
		zap.String("method", "GetPatientByUhid"))

	p := models.Registration{}
	var inputArgs []interface{}
	sqlQuery := `SELECT registration_date, uhid, barcode, name, labour_id, age,
                gender, mobile, district, taluk, camp FROM patients WHERE uhid = $1`
	log.Debug("Executing query", zap.String("query", sqlQuery), zap.String("uhid", uhid)) // Use Debug level
	sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)
	inputArgs = append(inputArgs, uhid)
	log.Debug("Executing query", zap.String("query", sqlQuery))

	row := DB.QueryRow(sqlQuery, inputArgs...)

	err := row.Scan(&p.RegistrationDate, &p.Uhid, &p.Barcode, &p.Name, &p.LabourId, &p.Age,
		&p.Gender, &p.Mobile, &p.DistrictId, &p.Taluk, &p.CampaignId)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &p, nil
}

func (r *postgresRepository) SaveRegistration(ctx *context.Context, p *models.Registration) error {
	log := logger.Get(*ctx).With(
		zap.String("barcode", p.Barcode),
		zap.String("method", "SavePatient"))
	var inputArgs []interface{}
	var registrationId string
	sqlQuery := `INSERT INTO registrations(registration_date, uhid, barcode, name, labour_id, age,
	                     gender, mobile, district_id, taluk, campaign_id)
					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id `

	inputArgs = append(inputArgs, p.RegistrationDate, p.Uhid, p.Barcode, p.Name, p.LabourId, p.Age,
		p.Gender, p.Mobile, p.DistrictId, p.Taluk, p.CampaignId)
	sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)

	log.Debug("Executing query", zap.String("query", sqlQuery))
	err := DB.QueryRow(sqlQuery, inputArgs...).Scan(&registrationId)

	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("beneficiary added with id : %v", registrationId))
	return nil
}
