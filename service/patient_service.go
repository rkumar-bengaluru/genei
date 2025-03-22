package service

import (
	"context"

	"example.com/rest-api/models"
	"example.com/rest-api/repository"
)

// Service contains dependencies and implements methods for business logic
type PatientService struct {
	repo repository.Repository
}

// New creates new instance of Service
func NewPatientService(repo repository.Repository) *PatientService {
	return &PatientService{
		repo: repo,
	}
}

func (ps *PatientService) GetPatientByUhid(ctx *context.Context, uhid string) (*models.Registration, error) {
	return ps.repo.GetPatientByUhid(ctx, uhid)
}

func (ps *PatientService) Save(ctx *context.Context, patient *models.Registration) error {
	return ps.repo.SaveRegistration(ctx, patient)
}
