package database

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host       string
	Port       int
	DBName     string
	DBUser     string
	DBPassword string
}

// BuildDBConfig use for building DB config
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:       "postgres",
		Port:       5432,
		DBName:     "car_game",
		DBUser:     "root",
		DBPassword: "password",
	}
	return &dbConfig
}

// DbURL use for create DB connection URL
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s port=%d sslmode=disable TimeZone=Asia/Bangkok",
		dbConfig.Host,
		dbConfig.DBName,
		dbConfig.DBUser,
		dbConfig.DBPassword,
		dbConfig.Port,
	)
}
