package repository

import (
	"context"
	"fmt"

	"example.com/rest-api/logger"
	"example.com/rest-api/models"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func (r *postgresRepository) CreateCampaign(ctx *context.Context, c *models.Campaign) error {
	log := logger.Get(*ctx).With(
		zap.String("camp", c.CampName),
		zap.String("method", "CreateCampaign"))
	var inputArgs []interface{}
	var campId string

	sqlQuery := `INSERT INTO campaigns(district_id, state_id, estimated_target_screening, 
		                               labour_inspector_name, union_name, union_leader_name,
	                                   latitude, longitude, pin_code_id, 
									   taluk_name, application_access_id,camp_name, 
									   description, screening_start_date, screening_start_time,
									   assigning_authority_id, store_id,work_order_id)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14,
							$15, $16, $17, $18) RETURNING id `

	inputArgs = append(inputArgs, c.DistrictId, c.StateId, c.EstimatedTargetScreening,
		c.LabourInspectorName, c.UnionName, c.UnionLeaderName,
		c.Latitude, c.Longitude, c.PincodeId, c.Taluk, c.ApplicationAccessId, c.CampName,
		c.Description, c.ScreeningStartDate, c.ScreeningStartTime, c.AssigningAuthorityId, c.StoreId, c.WorkOrderId,
	)
	sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)

	log.Debug("Executing query", zap.String("query", sqlQuery))
	err := DB.QueryRow(sqlQuery, inputArgs...).Scan(&campId)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	log.Info(fmt.Sprintf("campaign added with id : %v", campId))
	return nil
}
