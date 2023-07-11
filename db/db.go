package db

import (
  "fmt"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"

  "emr/config"
)

// DB handles database connections
type DB struct {
  *gorm.DB
}

// Connect initializes database connection
func Connect(config *config.Config) (*DB, error) {
  dbURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
    config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword)
  
  db, err := gorm.Open("postgres", dbURL)
  if err != nil {
    return nil, err
  }

  return &DB{db}, nil
}
