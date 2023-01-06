package config

import (
  "github.com/caarlos0/env/v6"
)

type Config struct {
  ENVIRONMENT string `env:"ENVIRONMENT" envDefault:"local"`
  WEB_PORT    int    `env:"WEB_PORT" envDefault:"8080"`
  DB_HOST     string `env:"DB_HOST" envDefault:"127.0.0.1"`
  DB_PORT     int    `env:"DB_PORT" envDefault:"3306"`
  DB_USER     string `env:"DB_USER" envDefault:"docker"`
  DB_PASSWORD string `env:"DB_PASSWORD" envDefault:"docker"`
  DB_NAME     string `env:"DB_NAME" envDefault:"birthday_reminder"`
  DB_RDBMS    string `env:"DB_RDBMS" envDefault:"postgresql"`
}

func NewConfig() (*Config, error) {
  config := &Config{}
  if err := env.Parse(config); err != nil {
    return nil, err
  }
  return config, nil
}
