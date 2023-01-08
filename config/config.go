package config

import (
  "log"
  "github.com/caarlos0/env/v6"
)

type Config struct {
  ENVIRONMENT string `env:"ENVIRONMENT" envDefault:"local"`
  WEB_PORT    int    `env:"WEB_PORT" envDefault:"8080"`
  // RDB
  DB_HOST     string `env:"DB_HOST" envDefault:"127.0.0.1"`
  DB_PORT     int    `env:"DB_PORT" envDefault:"3306"`
  DB_USER     string `env:"DB_USER" envDefault:"docker"`
  DB_PASSWORD string `env:"DB_PASSWORD" envDefault:"docker"`
  DB_NAME     string `env:"DB_NAME" envDefault:"birthday_reminder"`
  DB_RDBMS    string `env:"DB_RDBMS" envDefault:"postgresql"`
  // Validation
  VALIDATE_REQUEST_BODY_MESSAGE string `env:"VALIDATE_REQUEST_BODY_MESSAGE" envDefault:"Validation Error. Property:%v Value:%v"`
  VALIDATE_REQUEST_BODY_PROPERTY string `env:"VALIDATE_REQUEST_BODY_PROPERTY" envDefault:"%v"`
}

var err error

// construct
func NewConfig() (config *Config) {
  config, err = parseConfig()
  if err != nil {
    log.Fatalf("failed load config. %v", err)
  }
  return
}

func parseConfig() (*Config, error) {
  config := &Config{}
  if err = env.Parse(config); err != nil {
    return nil, err
  }
  return config, nil
}
