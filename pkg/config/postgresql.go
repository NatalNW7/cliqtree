package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func initPostgreSQL() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt: true,
	})

	if err != nil {
		logger.Errorf("PostgreSQL error: %v", err)
		return nil, err
	}

	return db, nil
}

type PostgreSQLInfo struct {
	DatabaseInfo
	OpenedConnections int `json:"opened_connections"`
}

func postgresInfo() PostgreSQLInfo {
	var version string
	var maxConn int
	var opConn int

	db.Raw("SHOW server_version;").Scan(&version)
	db.Raw("SHOW max_connections;").Scan(&maxConn)
	db.Raw("SELECT count(*)::int FROM pg_stat_activity WHERE datname = ?;", os.Getenv("POSTGRES_DB")).Scan(&opConn)

	return PostgreSQLInfo{
		DatabaseInfo: DatabaseInfo{
			Name: "PostgreSQL",
			Version: version,
			MaxConnections: maxConn,
		},
		OpenedConnections: opConn,
	}
}