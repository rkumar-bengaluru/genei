package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"example.com/rest-api/logger"
	"example.com/rest-api/utils"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

const (
	defaultMaxConns     = 10
	defaultIdleConns    = 5
	defaultMaxLifeTime  = 30
	defaultIdleLifeTime = 5
	minIdleConns        = 2
	minConnIdleTime     = 1
	defaultFactor       = 2
)

type Config struct {
	Host, Port, User, Password, DBName, SSLMode string
	MaxOpenConnections                          int
	MaxIdleConnections                          int
	MaxConnectionLifeTime                       int64
	MaxConnectionIdleTime                       int64
}

// POSTGRES_USER=myuser  -e POSTGRES_PASSWORD=mypassword -e POSTGRES_DB=mydb
func InitConfig(ctx context.Context) *Config {
	host := utils.ReadStr(ctx, "POSTGRES_HOST")
	port := utils.ReadStr(ctx, "POSTGRES_PORT")
	user := utils.ReadStr(ctx, "POSTGRES_USER")
	password := utils.ReadStr(ctx, "POSTGRES_PWD")
	dbname := utils.ReadStr(ctx, "POSTGRES_DBNAME")
	sslMode := utils.ReadStr(ctx, "POSTGRES_SSLMODE")
	logger.Get(ctx).Info(fmt.Sprintf("user %v", user))
	return &Config{
		Host:                  host,
		Port:                  port,
		User:                  user,
		Password:              password,
		DBName:                dbname,
		SSLMode:               sslMode,
		MaxOpenConnections:    utils.ReadIntWithDefault(ctx, "POSTGRES_MAX_OPEN_CONNECTION", defaultMaxConns),
		MaxIdleConnections:    utils.ReadIntWithDefault(ctx, "POSTGRES_MAX_IDLE_CONNECTION", defaultIdleConns),
		MaxConnectionLifeTime: utils.ReadInt64WithDefault(ctx, "POSTGRES_MAX_CONNECTION_LIFE_TIME", defaultMaxLifeTime),
		MaxConnectionIdleTime: utils.ReadInt64WithDefault(ctx, "POSTGRES_MAX_CONNECTION_IDLE_TIME", defaultIdleLifeTime),
	}
}

func CreateDB(ctx context.Context, serviceName string) *sql.DB {
	return CreateDBWithConfig(ctx, serviceName, InitConfig(ctx))
}

func (c *Config) normalize() {
	if c.MaxOpenConnections == 0 {
		c.MaxOpenConnections = defaultMaxConns
	}
	if c.MaxIdleConnections == 0 {
		c.MaxIdleConnections = defaultIdleConns
	}
	if c.MaxConnectionLifeTime == 0 {
		c.MaxConnectionLifeTime = defaultMaxLifeTime
	}
	if c.MaxConnectionIdleTime == 0 {
		c.MaxConnectionIdleTime = defaultIdleLifeTime
	}
	if c.MaxOpenConnections < c.MaxIdleConnections {
		c.MaxOpenConnections = c.MaxIdleConnections
		c.MaxIdleConnections /= defaultFactor
		if c.MaxIdleConnections < minIdleConns {
			c.MaxIdleConnections = minIdleConns
			c.MaxOpenConnections = defaultFactor * minIdleConns
		}
	}
	if c.MaxConnectionLifeTime < c.MaxConnectionIdleTime {
		c.MaxConnectionLifeTime = c.MaxConnectionIdleTime
		if c.MaxConnectionLifeTime < minConnIdleTime {
			c.MaxConnectionIdleTime = minConnIdleTime
			c.MaxConnectionLifeTime = defaultFactor * minConnIdleTime
		}
	}
}

func CreateDBWithConfig(ctx context.Context, serviceName string, postgresConfiguration *Config) *sql.DB {
	postgresConfiguration.normalize()
	dsn := "host=" + postgresConfiguration.Host + " port=" + postgresConfiguration.Port + " user='" +
		postgresConfiguration.User + "' password='" + postgresConfiguration.Password + "' dbname='" +
		postgresConfiguration.DBName + "' sslmode=" + postgresConfiguration.SSLMode
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Get(ctx).With(zap.Error(err)).Panic("connection open pg failed")
	}
	db.SetMaxOpenConns(postgresConfiguration.MaxOpenConnections)
	db.SetMaxIdleConns(postgresConfiguration.MaxIdleConnections)
	db.SetConnMaxLifetime(time.Duration(postgresConfiguration.MaxConnectionLifeTime) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(postgresConfiguration.MaxConnectionIdleTime) * time.Minute)

	for i := 3; i >= 0; i-- {
		res, err := db.Exec("SELECT 1")
		if err == nil {
			createTables(db)
			return db
		}
		logger.Get(ctx).With(zap.Error(err)).Error(fmt.Sprintf("res %v", res))
	}
	panic("DB connection failed")
}

func createTables(DB *sql.DB) {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY ,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		role TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table.")
	}

	createPatientsTable := `
	CREATE TABLE IF NOT EXISTS patients (
		id SERIAL PRIMARY KEY ,
		uhid TEXT NOT NULL,
		barcode TEXT NOT NULL,
		name TEXT NOT NULL,
		labour_id TEXT NOT NULL,
		age INTEGER NOT NULL,
		gender TEXT NOT NULL,
		mobile TEXT NOT NULL,
		district TEXT NOT NULL,
		taluk TEXT NOT NULL,
		camp TEXT NOT NULL,
		lab_test_status INTEGER ,
		report_url TEXT 
	)
	`

	_, err = DB.Exec(createPatientsTable)

	if err != nil {
		panic("Could not create patients table.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY ,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATE NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id SERIAL PRIMARY KEY ,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table.")
	}
}
