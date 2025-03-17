package service

import (
	"example.com/rest-api/models"
)

// Service contains dependencies and implements methods for business logic
type PatientService struct {
	repo models.Repository
}

// New creates new instance of Service
func NewPatientService(repo models.Repository) *PatientService {
	return &PatientService{
		repo: repo,
	}
}

func (ps *PatientService) GetPatientByUhid(uhid string) (*models.Patient, error) {
	return ps.repo.GetPatientByUhid(uhid)
}

func (ps *PatientService) Save(patient *models.Patient) error {
	return ps.repo.SavePatient(patient)
}

type UserService struct {
	repo models.Repository
}

func (ps *UserService) ValidateCredentials(user *models.User) (string, error) {
	return ps.repo.ValidateCredentials(user)
}

func (us *UserService) Save(user *models.User) error {
	return us.repo.SaveUser(user)
}

// New creates new instance of Service
func NewUserService(repo models.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}
