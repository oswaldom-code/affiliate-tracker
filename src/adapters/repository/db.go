package repository

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/oswaldom-code/affiliate-tracker/pkg/config"
	"github.com/oswaldom-code/affiliate-tracker/pkg/log"
	"github.com/oswaldom-code/affiliate-tracker/src/services/ports"
)

// repository handles the database context
type repository struct {
	db *gorm.DB
}

var dbStore *repository

// New returns a new instance of a repository
func New(dsn config.DBConfig) ports.Repository {
	var dsnStrConnection string
	switch dsn.Engine {
	case "postgre":
		dsnStrConnection = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%v sslmode=%s TimeZone=America/Lima",
			dsn.Host, dsn.User, dsn.Password, dsn.Database, dsn.Port, dsn.SSLMode)
	case "mysql":
		dsnStrConnection = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
			dsn.User, dsn.Password, dsn.Host, dsn.Port, dsn.Database)
	default:
		log.ErrorWithFields("Invalid database engine", log.Fields{"engine": dsn.Engine})
	}

	// configure connection
	config := &gorm.Config{
		// SkipDefaultTransaction: (default false) - skip default transaction for each request
		// (useful for performance) 30 % faster but you need to handle transactions manually (begin, commit, rollback)
		SkipDefaultTransaction: true,
		FullSaveAssociations:   false, // default is true
	}
	db, err := gorm.Open(postgres.Open(dsnStrConnection), config)
	if err != nil {
		log.ErrorWithFields("error connecting to db ", log.Fields{
			"engine":   dsn.Engine,
			"host":     dsn.Host,
			"port":     dsn.Port,
			"database": dsn.Database,
			"username": dsn.User,
			"err":      err,
		})
		os.Exit(1)
	}
	dbStore = &repository{db: db.Set("gorm:auto_preload", true)}

	return dbStore
}

func NewRepository() ports.Repository {
	log.DebugWithFields("Creating new database connection",
		log.Fields{"dsn": config.GetDBConfig()})
	if dbStore == nil {
		return New(config.GetDBConfig())
	}
	return dbStore
}
