package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/gommon/log"
)

type Config struct {
	AppEnv             string `env:"APP_ENV" env-default:"test"`
	Port               int    `env:"PORT" env-default:"8080"`
	LoyaltyService     string `env:"LOYALTY_SERVICE"`
	PaymentService     string `env:"PAYMENT_SERVICE"`
	ReservationService string `env:"RESERVATION_SERVICE"`
}

func NewConfig() (Config, error) {
	localPath := "./configs/gateway.env"
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return Config{}, err
	}

	log.Info(cfg)

	if cfg.AppEnv != "test" {
		return Config{}, nil
	}

	err = cleanenv.ReadConfig(localPath, &cfg)
	if err != nil {
		return Config{}, err
	}

	log.Info(cfg)

	return cfg, nil
}
