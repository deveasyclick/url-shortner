package config

import (
	"os"
	"strconv"
)

var (
	PORT        int
	ENV         string
	DB_HOST     string
	DB_NAME     string
	DB_USER     string
	DB_PORT     int
	DB_PASSWORD string
	APP_URL     string
	WEB_PORT    int
)

func init() {
	PORT, _ = strconv.Atoi(os.Getenv("PORT"))
	ENV = os.Getenv("ENV")
	APP_URL = os.Getenv("APP_URL")
	DB_HOST = os.Getenv("DB_HOST")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	WEB_PORT, _ = strconv.Atoi(os.Getenv("WEB_PORT"))
	DB_PORT, _ = strconv.Atoi(os.Getenv("DB_PORT"))
}
