package repository

import (
	"context"

	"example.com/rest-api/logger"
	"example.com/rest-api/models"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func (r *postgresRepository) SaveAssigningAuthority(ctx *context.Context, u *models.AssigningAuthority) error {
	log := logger.Get(*ctx).With(
		zap.String("name", u.Name),
		zap.String("method", "SaveAssigningAuthority"))

	var inputArgs []interface{}
	sqlQuery := `INSERT INTO assigning_authority(name) 
				 VALUES (?)`

	inputArgs = append(inputArgs, u.Name)
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

	log.Info("models.AssigningAuthority created successfully. ")
	return err
}
