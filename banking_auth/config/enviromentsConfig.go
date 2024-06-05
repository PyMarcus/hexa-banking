package config

import (
	"errors"
	"os"

	"github.com/PyMarcus/banking_auth/logger"
	"github.com/joho/godotenv"
)

type IConfig interface {
	SanityCheckConfig()
}

var GlobalConfig *config

// NewGlobalConfigSet set the .env variables
func NewGlobalConfigSet() {
	err := godotenv.Load()
    if err != nil {
        panic(err)
    }
	GlobalConfig = &config{
		DBHost:     os.Getenv("DB_HOST"),
        DBPort:     os.Getenv("DB_PORT"),
        DBUser:     os.Getenv("DB_USER"),
        DBPassword: os.Getenv("DB_PASSWORD"),
		ServerIP:  os.Getenv("SERVER_IP"),
		ServerPort:  os.Getenv("SERVER_PORT"),
		DBDatabase:  os.Getenv("DB_DATABASE"),
		JWT_SECRET_KEY: os.Getenv("JWT_SECRET_KEY"),
	}
	GlobalConfig.SanityCheckConfig()
}

type config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBDatabase string
	ServerIP   string
	ServerPort string
	JWT_SECRET_KEY string 
}

func (c config) SanityCheckConfig() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error while load .env " + err.Error())
		panic(err)
	}
	if os.Getenv("DB_HOST") == "" || os.Getenv("DB_PORT") == "" || os.Getenv("DB_USER") == "" || os.Getenv("DB_PASSWORD") == "" || os.Getenv("SERVER_IP") == "" || os.Getenv("SERVER_PORT") == "" {
		logger.Error("Enviroment variables is not set ")
		panic(errors.New(".env is not set"))
	}
}

