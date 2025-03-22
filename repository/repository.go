package repository

import (
	"context"
	"sync"

	"example.com/rest-api/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Repository contains db dependencies and implements methods to interact with db
type postgresRepository struct {
	DB *sqlx.DB
}

// GetDB returns the db
func (r *postgresRepository) GetDB() *sqlx.DB {
	return r.DB
}

var instance *postgresRepository
var once sync.Once
var DB *sqlx.DB

// NewRepository creates new instance of prismatosql Repository
func NewRepository(db *sqlx.DB) Repository {
	DB = db
	return getInstance(db)
}

func getInstance(db *sqlx.DB) *postgresRepository {
	once.Do(func() {
		instance = &postgresRepository{
			DB: db,
		}
	})
	return instance
}

//go:generate mockgen -source=repository.go -destination=./mocks/repository.go
type Repository interface {
	GetDB() *sqlx.DB
	GetPatientByUhid(ctx *context.Context, uhid string) (*models.Registration, error)
	SaveRegistration(ctx *context.Context, p *models.Registration) error
	SaveUser(ctx *context.Context, p *models.User) error
	ValidateCredentials(ctx *context.Context, u *models.User) (string, error)
	GetPassword(ctx *context.Context, u *models.User) (string, error)
	GetRole(ctx *context.Context, roleId uuid.UUID) (string, error)
	CreateCampaign(ctx *context.Context, c *models.ArogyaCampaign) error
	ListCampaigns(ctx *context.Context, pageSize, offset int) ([]*models.ArogyaCampaign, error)
	SaveAssigningAuthority(ctx *context.Context, u *models.AssigningAuthority) error
}
