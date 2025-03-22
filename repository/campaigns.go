package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"example.com/rest-api/logger"
	"example.com/rest-api/models"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func (r *postgresRepository) ListCampaigns(ctx *context.Context, pageSize, offset int) ([]*models.ArogyaCampaign, error) {
	log := logger.Get(*ctx).With(
		zap.String("pageSize", strconv.Itoa(pageSize)),
		zap.String("Offset", strconv.Itoa(offset)),
		zap.String("method", "CreateCampaign"))

	log.Info("Executing Query for Campaign Pagination ")

	sqlQuery := `SELECT id, name, distict, village, taluk_name, pin_code, camp_id,
				work_order, visibility, status, created_by, created_at,updated_at,
				state_name, estimated_target_screening,
				latitude, longitude,screening_start_date,screening_start_time
				FROM campaigns LIMIT $1 OFFSET $2`
	log.Debug("Executing query", zap.String("query", sqlQuery))
	rows, err := DB.Query(sqlQuery, pageSize, offset)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	// prepare the response
	campaigns := []*models.ArogyaCampaign{}
	for rows.Next() {
		var campaign models.ArogyaCampaign
		var visibility string
		var createdBy string
		if err := rows.Scan(&campaign.ID, &campaign.Name, &campaign.DistrictName, &campaign.VillageName,
			&campaign.Taluk, &campaign.PincodeId, &campaign.CampId, &campaign.WorkOrderId,
			&visibility, &campaign.Status, &createdBy, &campaign.CreatedAt,
			&campaign.UpdatedAt, &campaign.StateName, &campaign.EstimatedTargetScreening,
			&campaign.Latitude, &campaign.Longitude, &campaign.ScreeningStartDate,
			&campaign.ScreeningStartTime); err != nil {
			log.Error(err.Error())
			return nil, err
		}
		campaign.Visibility = append(campaign.Visibility, sql.NullString{
			String: visibility,
			Valid:  true,
		})
		campaign.CreatedBy = models.AuditLog{
			Name: sql.NullString{
				String: createdBy,
				Valid:  true,
			},
		}
		campaigns = append(campaigns, &campaign)
	}

	return campaigns, nil
}

func (r *postgresRepository) CreateCampaign(ctx *context.Context, c *models.ArogyaCampaign) error {
	log := logger.Get(*ctx).With(
		zap.String("camp", c.Name.String),
		zap.String("method", "CreateCampaign"))
	var inputArgs []interface{}
	var campId string

	sqlQuery := `INSERT INTO campaigns(name, distict, village, 
		                               taluk_name, pin_code, camp_id,
	                                   work_order, visibility, status, 
									   created_by, created_at,updated_at,
									   state_name, description, estimated_target_screening,
									   labour_inspector_name, union_name,union_leader_name,
									   latitude, longitude,screening_start_date,
									   screening_start_time)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14,
							$15, $16, $17, $18, $19, $20, $21, $22) RETURNING id `

	inputArgs = append(inputArgs, c.Name, c.Program.DistrictName, c.VillageName,
		c.Taluk, c.PincodeId, c.CampId,
		c.Program.Number, c.Visibility[0], c.Program.Status,
		c.CreatedBy.Name, c.CreatedAt, c.UpdatedAt,
		c.StateName, c.Description, c.EstimatedTargetScreening,
		c.LabourInspectorName, c.UnionName, c.UnionLeaderName,
		c.Latitude, c.Longitude, c.ScreeningStartDate, c.ScreeningStartTime)
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
