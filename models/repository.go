package models

import (
	"database/sql"
	"sync"
)

var instance *postgresRepository
var once sync.Once
var DB *sql.DB

// NewRepository creates new instance of prismatosql Repository
func NewRepository(db *sql.DB) Repository {
	DB = db
	return getInstance(db)
}

func getInstance(db *sql.DB) *postgresRepository {
	once.Do(func() {
		instance = &postgresRepository{
			DB: db,
		}
	})
	return instance
}

//go:generate mockgen -source=repository.go -destination=./mocks/repository.go
type Repository interface {
	GetDB() *sql.DB
	GetPatientByUhid(uhid string) (*Patient, error)
	SavePatient(p *Patient) error
	SaveUser(p *User) error
	ValidateCredentials(u *User) (string, error)
}
