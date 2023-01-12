package config

import (
  "log"
  "github.com/caarlos0/env/v6"
)

// Inject environment variables with CI/CD!
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
  VALIDATE_REQUEST_BODY_MESSAGE  string `env:"VALIDATE_REQUEST_BODY_MESSAGE" envDefault:"Validation Error. Property:%v Value:%v"`
  VALIDATE_REQUEST_BODY_PROPERTY string `env:"VALIDATE_REQUEST_BODY_PROPERTY" envDefault:"%v"`
  // JWT Require! Inject environment variables with CI/CD!
  JWT_SECRET_KEY string `env:"JWT_SECRET_KEY" envDefault:"-----BEGIN RSA PRIVATE KEY-----
  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
  -----END RSA PRIVATE KEY-----"`
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
