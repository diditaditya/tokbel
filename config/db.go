package config

import (
	"fmt"
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
}

func GetDBConfig() *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

func (conf *DBConfig) GetDBUrl() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.DBName, conf.Password)
}
