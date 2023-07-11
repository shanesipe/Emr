package config

import (
  "os"

  "github.com/spf13/viper"
)

// Config stores application configuration
type Config struct {
  DBHost string
  DBPort int 
  DBUser string
  DBPassword string
  DBName string

  SMTPServer string
  SMTPPort int
  SMTPEmail string
  SMTPPassword string

  JWTSecret string
}

// Load loads configuration from files and environment variables
func Load() Config {

  viper.SetConfigName("config")
  viper.SetConfigType("yaml")
  viper.AddConfigPath(".")
  viper.AutomaticEnv()

  err := viper.ReadInConfig()
  if err != nil {
    panic(err)
  }

  var c Config
  
  err = viper.Unmarshal(&c)
  if err != nil {
    panic(err)
  }

  return c
}

// DBURL returns PostgreSQL database URL 
func (c *Config) DBURL() string {
  return "postgres://" + c.DBUser + ":" + c.DBPassword + "@" + c.DBHost + "/" + c.DBName
}

// GetEnv returns environment variable value
func GetEnv(key string) string {
  return os.Getenv(key)
}
