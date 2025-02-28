package config

import (
	"os"
	"strconv"
)

var (
	PORT        int
	APP_ENV     string
	DB_HOST     string
	DB_PORT     string = "5432"
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
)

func init() {
	PORT, _ = strconv.Atoi(os.Getenv("PORT"))
	APP_ENV = os.Getenv("APP_ENV")
	DB_HOST = os.Getenv("DB_HOST")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
}
