package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Importing the PostgreSQL driver for database connectivity.
	"go.uber.org/zap"

	"github.com/RedrikShuhartRed/TaskManager/config"
)

// Storage struct holds the database connection pool.
type Storage struct {
	DB *sqlx.DB
}

// ConnectDB establishes a connection to the PostgreSQL database.
// It checks if the 'verbatasks' database exists, and if not, it creates it.
func ConnectDB(cfg *config.Config) (*Storage, error) {

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.DbPort, cfg.Sslmode)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		zap.S().Errorf("error DB connection : %s", err)
		return nil, err
	}

	rows, err := db.Query("SELECT 1 FROM pg_database WHERE datname = 'verbatasks'")
	if err != nil {
		zap.S().Errorf("error checking database existence: %s", err)
		return nil, err
	}

	var result int
	for rows.Next() {

		err := rows.Scan(&result)
		if err != nil {
			zap.S().Errorf("error reading query result in pg_database: %s", err)
			return nil, err
		}
	}

	if result != 1 {
		_, err = db.Exec("CREATE DATABASE verbatasks")
		if err != nil {
			zap.S().Errorf("Error creating database:%s", err)
			return nil, err
		}

	}
	connStr = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=verbatasks sslmode=%s", cfg.User, cfg.Password, cfg.Host, cfg.DbPort, cfg.Sslmode)
	db, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		zap.S().Errorf("Error connecting to verbatasks database: %s", err)
		return nil, err
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS tasks (
		id BIGSERIAL NOT NULL PRIMARY KEY,
		title VARCHAR(40) NOT NULL,
		description VARCHAR(100) ,
		due_date TIMESTAMP NOT NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP)`,
	)

	if err != nil {
		zap.S().Errorf("Error creating tasks table:%s", err)
		return nil, err
	}

	return &Storage{DB: db}, nil
}

// CloseDB closes the database connection.
func (s *Storage) CloseDB() {
	err := s.DB.Close()
	if err != nil {
		zap.S().Errorf("error close connecting to database:%v", err)
	}
}
