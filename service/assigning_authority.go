package service

import (
	"context"

	"example.com/rest-api/models"
	"example.com/rest-api/repository"
)

// Service contains dependencies and implements methods for business logic
type AssigningAuthorityService struct {
	repo repository.Repository
}

// New creates new instance of Service
func NewAssigningAuthrityService(repo repository.Repository) *AssigningAuthorityService {
	return &AssigningAuthorityService{
		repo: repo,
	}
}

func (ps *AssigningAuthorityService) CreateAssigningAuthority(ctx *context.Context, campaign *models.AssigningAuthority) error {
	return ps.repo.SaveAssigningAuthority(ctx, campaign)
}
