package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"example.com/rest-api/logger"
	"example.com/rest-api/utils"
	"github.com/jmoiron/sqlx"
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

func CreateDB(ctx context.Context, serviceName string) *sqlx.DB {
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

func CreateDBWithConfig(ctx context.Context, serviceName string, postgresConfiguration *Config) *sqlx.DB {
	postgresConfiguration.normalize()
	dsn := "host=" + postgresConfiguration.Host + " port=" + postgresConfiguration.Port + " user='" +
		postgresConfiguration.User + "' password='" + postgresConfiguration.Password + "' dbname='" +
		postgresConfiguration.DBName + "' sslmode=" + postgresConfiguration.SSLMode
	db, err := sqlx.Connect("postgres", dsn)
	//db, err := sql.Open("postgres", dsn)
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
			//createTables(db)
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

	//

	createWorkOrderTable := `
	CREATE TABLE IF NOT EXISTS work_orders (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name TEXT NOT NULL
	)
	`

	_, err = DB.Exec(createWorkOrderTable)

	if err != nil {
		panic("Could not create work_orders table.")
	}

	createStateTable := `
	CREATE TABLE IF NOT EXISTS states (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name TEXT NOT NULL
	)
	`

	_, err = DB.Exec(createStateTable)

	if err != nil {
		panic("Could not create states table.")
	}

	createPincodesTable := `
	CREATE TABLE IF NOT EXISTS pin_codes (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		pin_code TEXT NOT NULL,
		state_id UUID,
		district_id UUID,
		FOREIGN KEY(state_id) REFERENCES states(id),
		FOREIGN KEY(district_id) REFERENCES districts(id)
	)
	`

	_, err = DB.Exec(createPincodesTable)

	if err != nil {
		panic("Could not create states table.")
	}

	createDistrictTable := `
	CREATE TABLE IF NOT EXISTS districts (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name TEXT NOT NULL,
		state_id UUID,
		FOREIGN KEY(state_id) REFERENCES states(id)
	)
	`

	_, err = DB.Exec(createDistrictTable)

	if err != nil {
		panic("Could not create campaigns table.")
	}

	createRolesTable := `
	CREATE TABLE IF NOT EXISTS roles (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name TEXT NOT NULL
	)
	`

	_, err = DB.Exec(createRolesTable)

	if err != nil {
		panic("Could not create campaigns table.")
	}

	createAssigningAuthorityTable := `
	CREATE TABLE IF NOT EXISTS assigning_authority (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name TEXT NOT NULL
	)
	`

	_, err = DB.Exec(createAssigningAuthorityTable)

	if err != nil {
		panic("Could not create campaigns table.")
	}

	createStoreTable := `
	CREATE TABLE IF NOT EXISTS stores (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name TEXT NOT NULL
	)
	`

	_, err = DB.Exec(createStoreTable)

	if err != nil {
		panic("Could not create campaigns table.")
	}

	createCampaignTable := `
	CREATE TABLE IF NOT EXISTS campaigns (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		district_id UUID,
		state_id UUID,
		estimated_target_screening INTEGER NOT NULL,
		labour_inspector_name TEXT NOT NULL,
		union_name TEXT NOT NULL,
		union_leader_name TEXT NOT NULL,
		latitude TEXT NOT NULL,
		longitude TEXT NOT NULL,
		pin_code_id UUID,
		taluk_name TEXT NOT NULL,
		visibility_access UUID,
		camp_name TEXT NOT NULL,
		description TEXT NOT NULL,
		screening_start_date DATE NOT NULL,
		screening_start_time TEXT NOT NULL,
		assigning_authority_id UUID,
		store_id UUID,
		FOREIGN KEY(district_id) REFERENCES districts(id),
		FOREIGN KEY(state_id) REFERENCES states(id),
		FOREIGN KEY(pin_code_id) REFERENCES pin_codes(id),
		FOREIGN KEY(visibility_access) REFERENCES roles(id),
		FOREIGN KEY(assigning_authority_id) REFERENCES assigning_authority(id),
		FOREIGN KEY(store_id) REFERENCES stores(id)
	)
	`

	_, err = DB.Exec(createCampaignTable)

	if err != nil {
		panic("Could not create campaigns table.")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		uhid TEXT NOT NULL,
		barcode TEXT NOT NULL,
		name TEXT NOT NULL,
		labour_id TEXT NOT NULL,
		age INTEGER NOT NULL,
		gender TEXT NOT NULL,
		mobile TEXT NOT NULL,
		taluk TEXT NOT NULL,
		lab_test_status INTEGER ,
		report_url TEXT ,
		campaign_id UUID ,
		district_id UUID,
		FOREIGN KEY(campaign_id) REFERENCES campaigns(id),
		FOREIGN KEY(district_id) REFERENCES districts(id)
	)
	`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table.")
	}

	createReportsTable := `
	CREATE TABLE IF NOT EXISTS reports (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

		blood_group_id UUID ,
		parameter_group_id UUID,
		registration_id UUID ,
		campaign_id UUID ,
		FOREIGN KEY(registration_id) REFERENCES registrations(id),
		FOREIGN KEY(campaign_id) REFERENCES campaigns(id),
		FOREIGN KEY(parameter_group_id) REFERENCES parameter_groups(id)
	)
	`

	_, err = DB.Exec(createReportsTable)

	if err != nil {
		panic("Could not create registrations table.")
	}

	createParameterGroupTable := `
	CREATE TABLE IF NOT EXISTS parameter_groups (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		group_name TEXT NOT NULL,
		report_sequence INTEGER NOT NULL
	)
	`

	_, err = DB.Exec(createParameterGroupTable)

	if err != nil {
		panic("Could not create registrations table.")
	}

	createParametersTable := `
	CREATE TABLE IF NOT EXISTS parameters (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		name TEXT NOT NULL,
		result_type_id UUID,
		result TEXT NOT NULL,
		unit TEXT NOT NULL,
		interpretation TEXT NOT NULL,
		bio_ref_interval json NOT NULL,
		report_sequence INTEGER NOT NULL,
		parameter_group_id UUID,
		parameter_rule_id UUID,
		comments TEXT NOT NULL,
		FOREIGN KEY(parameter_group_id) REFERENCES parameter_groups(id),
		FOREIGN KEY(parameter_rule_id) REFERENCES parameter_rules(id)
	)
	`

	_, err = DB.Exec(createParametersTable)

	if err != nil {
		panic("Could not create registrations table.")
	}

	createResultTypeTable := `
	CREATE TABLE IF NOT EXISTS result_type (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		type INTEGER NOT NULL,
		drop_down_table_name TEXT NOT NULL
	)
	`

	_, err = DB.Exec(createResultTypeTable)

	if err != nil {
		panic("Could not create registrations table.")
	}

	createParameterRuleTable := `
	CREATE TABLE IF NOT EXISTS parameter_rules (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		expression TEXT NOT NULL
	)
	`

	_, err = DB.Exec(createParameterRuleTable)

	if err != nil {
		panic("Could not create registrations table.")
	}

}
