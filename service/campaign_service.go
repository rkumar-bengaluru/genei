package service

import (
	"context"

	"example.com/rest-api/models"
	"example.com/rest-api/repository"
)

// Service contains dependencies and implements methods for business logic
type CampaignService struct {
	repo repository.Repository
}

// New creates new instance of Service
func NewCampaignService(repo repository.Repository) *CampaignService {
	return &CampaignService{
		repo: repo,
	}
}

func (ps *CampaignService) CreateCampaign(ctx *context.Context, campaign *models.ArogyaCampaign) error {
	return ps.repo.CreateCampaign(ctx, campaign)
}

func (ps *CampaignService) ListCampaign(ctx *context.Context, pageSize, offset int) ([]*models.ArogyaCampaign, error) {
	return ps.repo.ListCampaigns(ctx, pageSize, offset)
}
