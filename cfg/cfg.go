package cfg

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	AppConfig AppConfig
	PgConfig  PostgresqlConfig
}

type (
	PostgresqlConfig struct {
		Host     string
		Port     int
		User     string
		Password string
		DbName   string
		SslMode  string
	}
	AppConfig struct {
		Host string
		Port int
	}
)

func createAppConfig() (AppConfig, error) {
	var appCfg AppConfig

	host, ok := os.LookupEnv("APP_HOST")
	if !ok {
		return appCfg, errors.New("env APP_HOST not set")
	}

	port, ok := os.LookupEnv("APP_PORT")
	if !ok {
		return appCfg, errors.New("env APP_PORT not set")
	}

	convPort, err := strconv.Atoi(port)
	if err != nil {
		return appCfg, errors.New("env APP_PORT not able to convert to int")
	}

	appCfg.Host = host
	appCfg.Port = convPort

	return appCfg, nil
}

func createPostgresConfig() (PostgresqlConfig, error) {
	var pgCfg PostgresqlConfig

	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return pgCfg, errors.New("env DB_PORT not set")
	}

	convPort, err := strconv.Atoi(port)
	if err != nil {
		return pgCfg, errors.New("env DB_PORT not able to convert to int")
	}

	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return pgCfg, errors.New("env DB_HOST not set")
	}

	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		return pgCfg, errors.New("env DB_USER not set")
	}

	password, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return pgCfg, errors.New("env DB_PASSWORD not set")
	}

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return pgCfg, errors.New("env DB_NAME not set")
	}

	sslMode, ok := os.LookupEnv("DB_SSL_MODE")
	if !ok {
		return pgCfg, errors.New("env DB_SSL_MODE not set")
	}

	pgCfg.Port = convPort
	pgCfg.Host = host
	pgCfg.User = user
	pgCfg.Password = password
	pgCfg.DbName = dbName
	pgCfg.SslMode = sslMode

	return pgCfg, nil
}

func CreateConfig() (Config, error) {
	var cfg Config

	appConfig, err := createAppConfig()
	if err != nil {
		return cfg, err
	}

	pgConfig, err := createPostgresConfig()
	if err != nil {
		return cfg, err
	}

	cfg.AppConfig = appConfig
	cfg.PgConfig = pgConfig

	return cfg, nil
}
