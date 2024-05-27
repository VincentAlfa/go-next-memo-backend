package database

import (
	"os"
	"github.com/joho/godotenv"
)

type DataSourceName struct {
	User string
	Password string 
	Host string
	Port string
	Database string
}

func DbSourceName() (dsn *DataSourceName) {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	return &DataSourceName{
		User : os.Getenv("DB_USER"),
		Password : os.Getenv("DB_PASSWORD"),
		Host : os.Getenv("DB_HOST"),
		Port : os.Getenv("DB_PORT"),
		Database : os.Getenv("DB_NAME"),
	}
}