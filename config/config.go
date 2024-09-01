package config

import (
	"os"

	"go.uber.org/zap"
)

const (
	// Default values for configuration settings
	defaultPort     = "8080"
	defaultUser     = "postgres"
	defaultPassword = "root"
	defaultHost     = "127.0.0.1"
	defaultDbPort   = "5432"
	defaultSslmode  = "disable"
)

// Config holds the configuration settings for the application.
type Config struct {
	Port     string
	User     string
	Password string
	Host     string
	DbPort   string
	Sslmode  string
}

// NewConfig creates a new Config instance by reading environment variables.
func NewConfig() *Config {

	port := os.Getenv("TASK_PORT")
	if port == "" {
		port = defaultPort
		zap.S().Infof("TASK_PORT environment variable required but not set, used default")
	}

	user := os.Getenv("TASK_USER")
	if user == "" {
		user = defaultUser
		zap.S().Infof("TASK_USER environment variable required but not set, used default")
	}

	password := os.Getenv("TASK_PASSWORD")
	if password == "" {
		password = defaultPassword
		zap.S().Infof("TASK_PASSWORD environment variable required but not set, used default")
	}

	host := os.Getenv("TASK_HOST")
	if host == "" {
		host = defaultHost
		zap.S().Infof("TASK_HOST environment variable required but not set, used defaultt")
	}

	dbPort := os.Getenv("TASK_DBPORT")
	if dbPort == "" {
		dbPort = defaultDbPort
		zap.S().Infof("TASK_DBPORT environment variable required but not set, used default")
	}

	sslmode := os.Getenv("TASK_SSLMODE")
	if sslmode == "" {
		sslmode = defaultSslmode
		zap.S().Infof("TASK_SSLMODE environment variable required but not set, used default")
	}

	return &Config{
		Port:     port,
		User:     user,
		Password: password,
		Host:     host,
		DbPort:   dbPort,
		Sslmode:  sslmode,
	}
}
