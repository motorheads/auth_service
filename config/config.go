package config

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Database string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "postgres",
		Password: "password",
		Database: "catalog_service",
	}
	return &dbConfig
}
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Database,
	)
}
